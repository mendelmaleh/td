package telegram

import (
	"context"
	"errors"
	"fmt"
	"io"
	"runtime/debug"
	"strings"
	"sync"

	"github.com/cenkalti/backoff/v4"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"golang.org/x/xerrors"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/clock"
	"github.com/gotd/td/internal/crypto"
	"github.com/gotd/td/internal/mt"
	"github.com/gotd/td/internal/proto"
	"github.com/gotd/td/internal/tmap"
	"github.com/gotd/td/mtproto"
	"github.com/gotd/td/session"
	"github.com/gotd/td/tg"
)

// UpdateHandler will be called on received updates from Telegram.
type UpdateHandler func(ctx context.Context, u *tg.Updates) error

// Available MTProto default server addresses.
//
// See https://my.telegram.org/apps.
const (
	AddrProduction = "149.154.167.50:443"
	AddrTest       = "149.154.167.40:443"
)

// Test-only credentials. Can be used with AddrTest and TestAuth to
// test authentication.
//
// Reference:
//	* https://github.com/telegramdesktop/tdesktop/blob/5f665b8ecb48802cd13cfb48ec834b946459274a/docs/api_credentials.md
const (
	TestAppID   = 17349
	TestAppHash = "344583e45741c457fe1862106095a5eb"
)

type clientStorage interface {
	Load(ctx context.Context) (*session.Data, error)
	Save(ctx context.Context, data *session.Data) error
}

type clientConn interface {
	Run(ctx context.Context) error
	InvokeRaw(ctx context.Context, input bin.Encoder, output bin.Decoder) error
}

// Client represents a MTProto client to Telegram.
type Client struct {
	// tg provides RPC calls via Client.
	tg *tg.Client

	connMux  sync.Mutex
	connAddr string
	connOpt  mtproto.Options
	conn     clientConn
	cfg      tg.Config
	restart  chan struct{}

	// Wrappers for external world, like logs or PRNG.
	// Should be immutable.
	rand  io.Reader
	log   *zap.Logger
	clock clock.Clock

	ctx    context.Context
	cancel context.CancelFunc

	appID   int    // immutable
	appHash string // immutable
	storage clientStorage

	ready     chan struct{}
	readyOnce sync.Once

	updateHandler UpdateHandler // immutable
}

func (c *Client) onMessage(b *bin.Buffer) error {
	return c.handleUpdates(b)
}

// getVersion optimistically gets current client version.
//
// Does not handle replace directives.
func getVersion() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return ""
	}
	// Hard-coded package name. Probably we can generate this via parsing
	// the go.mod file.
	const pkg = "github.com/gotd/td"
	for _, d := range info.Deps {
		if strings.HasPrefix(d.Path, pkg) {
			return d.Version
		}
	}
	return ""
}

// Port is default port used by telegram.
const Port = 443

// NewClient creates new unstarted client.
func NewClient(appID int, appHash string, opt Options) *Client {
	// Set default values, if user does not set.
	opt.setDefaults()

	clientCtx, clientCancel := context.WithCancel(context.Background())
	client := &Client{
		rand:          opt.Random,
		log:           opt.Logger,
		ctx:           clientCtx,
		cancel:        clientCancel,
		appID:         appID,
		appHash:       appHash,
		updateHandler: opt.UpdateHandler,
		connAddr:      opt.Addr,
		clock:         opt.Clock,
	}

	// Including version into client logger to help with debugging.
	if v := getVersion(); v != "" {
		client.log = client.log.With(zap.String("v", v))
	}

	if opt.SessionStorage != nil {
		client.storage = &session.Loader{
			Storage: opt.SessionStorage,
		}
	}

	client.connOpt = mtproto.Options{
		PublicKeys:    opt.PublicKeys,
		Transport:     opt.Transport,
		Network:       opt.Network,
		Random:        opt.Random,
		Logger:        opt.Logger,
		AckBatchSize:  opt.AckBatchSize,
		AckInterval:   opt.AckInterval,
		RetryInterval: opt.RetryInterval,
		MaxRetries:    opt.MaxRetries,
		MessageID:     opt.MessageID,
		Clock:         opt.Clock,

		Types: tmap.New(
			tg.TypesMap(),
			mt.TypesMap(),
			proto.TypesMap(),
		),
	}
	client.conn = client.createConn(connModeUpdates)

	// Initializing internal RPC caller.
	client.tg = tg.NewClient(client)

	return client
}

func (c *Client) restoreConnection(ctx context.Context) error {
	if c.storage == nil {
		return nil
	}
	data, err := c.storage.Load(ctx)
	if errors.Is(err, session.ErrNotFound) {
		return nil
	}
	if err != nil {
		return xerrors.Errorf("load: %w", err)
	}

	// Restoring persisted auth key.
	var key crypto.AuthKey
	copy(key.Value[:], data.AuthKey)
	copy(key.ID[:], data.AuthKeyID)

	if key.Value.ID() != key.ID {
		return xerrors.New("corrupted key")
	}

	// Re-initializing connection from persisted state.
	c.log.Info("Connection restored from state",
		zap.String("addr", data.Addr),
		zap.String("key_id", fmt.Sprintf("%x", data.AuthKeyID)),
	)

	c.connMux.Lock()
	c.connOpt.Key = key
	c.connOpt.Salt = data.Salt
	c.connAddr = data.Addr
	c.conn = c.createConn(connModeUpdates)
	c.connMux.Unlock()

	return nil
}

func (c *Client) runUntilRestart(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	g, gCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return c.conn.Run(ctx)
	})
	g.Go(func() error {
		select {
		case <-gCtx.Done():
			return gCtx.Err()
		case <-c.restart:
			c.log.Debug("Restart triggered")
			// Should call cancel() to cancel gCtx.
			cancel()

			return nil
		}
	})

	return g.Wait()
}

func (c *Client) reconnectUntilClosed(ctx context.Context) error {
	c.restart = make(chan struct{})

	// TODO: Make this configurable.
	// Note that we currently have no timeout on connection, so this is
	// potentially eternal.
	b := backoff.NewExponentialBackOff()
	b.Clock = c.clock
	b.MaxElapsedTime = 0

	for {
		err := c.runUntilRestart(ctx)
		if err == nil {
			return nil
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-c.clock.After(b.NextBackOff()):
			c.log.Info("Restarting connection", zap.Error(err))

			c.connMux.Lock()
			c.conn = c.createConn(connModeUpdates)
			c.connMux.Unlock()
		}
	}
}

func (c *Client) onReady() {
	c.log.Debug("Ready")
	c.readyOnce.Do(func() {
		close(c.ready)
	})
}

func (c *Client) resetReady() {
	c.ready = make(chan struct{})
	c.readyOnce = sync.Once{}
}

// Run starts client session and block until connection close.
// The f callback is called on successful session initialization and Run
// will return on f() result.
//
// Context of callback will be canceled if fatal error is detected.
func (c *Client) Run(ctx context.Context, f func(ctx context.Context) error) error {
	c.log.Info("Starting")
	defer c.log.Info("Closed")

	c.resetReady()
	if err := c.restoreConnection(ctx); err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	g, gCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return c.reconnectUntilClosed(gCtx)
	})
	g.Go(func() error {
		select {
		case <-gCtx.Done():
			return gCtx.Err()
		case <-c.ready:
			if err := f(gCtx); err != nil {
				return xerrors.Errorf("callback: %w", err)
			}
			// Should call cancel() to cancel gCtx.
			// This will terminate c.conn.Run().
			c.log.Debug("Callback returned, stopping")
			cancel()
			return nil
		}
	})
	if err := g.Wait(); !xerrors.Is(err, context.Canceled) {
		return err
	}

	return nil
}

func (c *Client) saveSession(addr string, cfg tg.Config, s mtproto.Session) error {
	if c.storage == nil {
		return nil
	}

	data, err := c.storage.Load(c.ctx)
	if errors.Is(err, session.ErrNotFound) {
		// Initializing new state.
		err = nil
		data = &session.Data{}
	}
	if err != nil {
		return xerrors.Errorf("load: %w", err)
	}

	// Updating previous data.
	data.Config = cfg
	data.AuthKey = s.Key.Value[:]
	data.AuthKeyID = s.Key.ID[:]
	data.DC = cfg.ThisDC
	data.Addr = addr
	data.Salt = s.Salt

	if err := c.storage.Save(c.ctx, data); err != nil {
		return xerrors.Errorf("save: %w", err)
	}

	c.log.Debug("Data saved",
		zap.String("key_id", fmt.Sprintf("%x", data.AuthKeyID)),
	)
	return nil
}

func (c *Client) onSession(addr string, cfg tg.Config, s mtproto.Session) error {
	if err := c.saveSession(addr, cfg, s); err != nil {
		return xerrors.Errorf("save: %w", err)
	}

	c.connMux.Lock()
	c.connAddr = addr
	c.cfg = cfg
	c.onReady()
	c.connMux.Unlock()

	return nil
}

func (c *Client) createConn(mode connMode) clientConn {
	return newConn(
		c,
		c.connAddr,
		c.appID,
		mode,
		c.connOpt,
	)
}
