// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"fmt"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)

// HelpGetConfigRequest represents TL type `help.getConfig#c4f9186b`.
// Returns current configuration, including data center configuration.
//
// See https://core.telegram.org/method/help.getConfig for reference.
type HelpGetConfigRequest struct {
}

// HelpGetConfigRequestTypeID is TL type id of HelpGetConfigRequest.
const HelpGetConfigRequestTypeID = 0xc4f9186b

// Encode implements bin.Encoder.
func (g *HelpGetConfigRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode help.getConfig#c4f9186b as nil")
	}
	b.PutID(HelpGetConfigRequestTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (g *HelpGetConfigRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode help.getConfig#c4f9186b to nil")
	}
	if err := b.ConsumeID(HelpGetConfigRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode help.getConfig#c4f9186b: %w", err)
	}
	return nil
}

// Ensuring interfaces in compile-time for HelpGetConfigRequest.
var (
	_ bin.Encoder = &HelpGetConfigRequest{}
	_ bin.Decoder = &HelpGetConfigRequest{}
)

// HelpGetConfig invokes method help.getConfig#c4f9186b returning error if any.
// Returns current configuration, including data center configuration.
//
// See https://core.telegram.org/method/help.getConfig for reference.
func (c *Client) HelpGetConfig(ctx context.Context) (*Config, error) {
	var result Config

	request := &HelpGetConfigRequest{}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
