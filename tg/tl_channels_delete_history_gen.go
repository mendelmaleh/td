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

// ChannelsDeleteHistoryRequest represents TL type `channels.deleteHistory#af369d42`.
// Delete the history of a supergroup
//
// See https://core.telegram.org/method/channels.deleteHistory for reference.
type ChannelsDeleteHistoryRequest struct {
	// Supergroup whose history must be deleted
	Channel InputChannelClass
	// ID of message up to which the history must be deleted
	MaxID int
}

// ChannelsDeleteHistoryRequestTypeID is TL type id of ChannelsDeleteHistoryRequest.
const ChannelsDeleteHistoryRequestTypeID = 0xaf369d42

// Encode implements bin.Encoder.
func (d *ChannelsDeleteHistoryRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode channels.deleteHistory#af369d42 as nil")
	}
	b.PutID(ChannelsDeleteHistoryRequestTypeID)
	if d.Channel == nil {
		return fmt.Errorf("unable to encode channels.deleteHistory#af369d42: field channel is nil")
	}
	if err := d.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.deleteHistory#af369d42: field channel: %w", err)
	}
	b.PutInt(d.MaxID)
	return nil
}

// Decode implements bin.Decoder.
func (d *ChannelsDeleteHistoryRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode channels.deleteHistory#af369d42 to nil")
	}
	if err := b.ConsumeID(ChannelsDeleteHistoryRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.deleteHistory#af369d42: %w", err)
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.deleteHistory#af369d42: field channel: %w", err)
		}
		d.Channel = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channels.deleteHistory#af369d42: field max_id: %w", err)
		}
		d.MaxID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelsDeleteHistoryRequest.
var (
	_ bin.Encoder = &ChannelsDeleteHistoryRequest{}
	_ bin.Decoder = &ChannelsDeleteHistoryRequest{}
)

// ChannelsDeleteHistory invokes method channels.deleteHistory#af369d42 returning error if any.
// Delete the history of a supergroup
//
// See https://core.telegram.org/method/channels.deleteHistory for reference.
func (c *Client) ChannelsDeleteHistory(ctx context.Context, request *ChannelsDeleteHistoryRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
