package tgtest

import (
	"errors"

	"golang.org/x/xerrors"

	"github.com/gotd/td/internal/mt"
	"github.com/gotd/td/tg"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/internal/crypto"
	"github.com/gotd/td/internal/proto"
)

func (s *Server) Send(k Session, t proto.MessageType, encoder bin.Encoder) error {
	conn, ok := s.users.getConnection(k.AuthKey)
	if !ok {
		return errors.New("invalid key: connection not found")
	}

	var b bin.Buffer
	if err := encoder.Encode(&b); err != nil {
		return xerrors.Errorf("failed to encode data: %w", err)
	}

	data := crypto.EncryptedMessageData{
		SessionID:              k.SessionID,
		MessageDataLen:         int32(b.Len()),
		MessageDataWithPadding: b.Copy(),
		MessageID:              s.msgID.New(t),
	}

	err := s.cipher.Encrypt(k.AuthKey, data, &b)
	if err != nil {
		return xerrors.Errorf("failed to encrypt message: %w", err)
	}

	return conn.Send(s.ctx, &b)
}

func (s *Server) SendResult(k Session, id int64, msg bin.Encoder) error {
	var buf bin.Buffer

	if err := msg.Encode(&buf); err != nil {
		return xerrors.Errorf("failed to encode result data: %w", err)
	}

	return s.Send(k, proto.MessageServerResponse, &proto.Result{
		RequestMessageID: id,
		Result:           buf.Raw(),
	})
}

func (s *Server) sendSessionCreated(k Session, serverSalt int64) error {
	return s.Send(k, proto.MessageFromServer, &mt.NewSessionCreated{
		ServerSalt: serverSalt,
	})
}

func (s *Server) SendConfig(k Session, id int64) error {
	s.log.Debug("SendConfig")
	return s.SendResult(k, id, &tg.Config{})
}

func (s *Server) SendUpdates(k Session, updates ...tg.UpdateClass) error {
	if len(updates) == 0 {
		return nil
	}

	return s.Send(k, proto.MessageFromServer, &tg.Updates{
		Updates: updates,
		Date:    int(s.clock.Now().Unix()),
	})
}

func (s *Server) SendAck(k Session, ids ...int64) error {
	return s.Send(k, proto.MessageFromServer, &mt.MsgsAck{MsgIds: ids})
}

func (s *Server) SendPong(k Session, msgID, pingID int64) error {
	return s.Send(k, proto.MessageServerResponse, &mt.Pong{
		MsgID:  msgID,
		PingID: pingID,
	})
}

func (s *Server) ForceDisconnect(k Session) {
	s.users.deleteConnection(k.AuthKey)
}
