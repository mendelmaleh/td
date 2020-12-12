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

// UpdatesGetChannelDifferenceRequest represents TL type `updates.getChannelDifference#3173d78`.
// Returns the difference between the current state of updates of a certain channel and transmitted.
//
// See https://core.telegram.org/method/updates.getChannelDifference for reference.
type UpdatesGetChannelDifferenceRequest struct {
	// Flags, see TL conditional fields
	Flags bin.Fields
	// Set to true to skip some possibly unneeded updates and reduce server-side load
	Force bool
	// The channel
	Channel InputChannelClass
	// Messsage filter
	Filter ChannelMessagesFilterClass
	// Persistent timestamp (see updates)
	Pts int
	// How many updates to fetch, max 100000Ordinary (non-bot) users are supposed to pass 10-100
	Limit int
}

// UpdatesGetChannelDifferenceRequestTypeID is TL type id of UpdatesGetChannelDifferenceRequest.
const UpdatesGetChannelDifferenceRequestTypeID = 0x3173d78

// Encode implements bin.Encoder.
func (g *UpdatesGetChannelDifferenceRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode updates.getChannelDifference#3173d78 as nil")
	}
	b.PutID(UpdatesGetChannelDifferenceRequestTypeID)
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode updates.getChannelDifference#3173d78: field flags: %w", err)
	}
	if g.Channel == nil {
		return fmt.Errorf("unable to encode updates.getChannelDifference#3173d78: field channel is nil")
	}
	if err := g.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode updates.getChannelDifference#3173d78: field channel: %w", err)
	}
	if g.Filter == nil {
		return fmt.Errorf("unable to encode updates.getChannelDifference#3173d78: field filter is nil")
	}
	if err := g.Filter.Encode(b); err != nil {
		return fmt.Errorf("unable to encode updates.getChannelDifference#3173d78: field filter: %w", err)
	}
	b.PutInt(g.Pts)
	b.PutInt(g.Limit)
	return nil
}

// SetForce sets value of Force conditional field.
func (g *UpdatesGetChannelDifferenceRequest) SetForce(value bool) {
	if value {
		g.Flags.Set(0)
	} else {
		g.Flags.Unset(0)
	}
}

// Decode implements bin.Decoder.
func (g *UpdatesGetChannelDifferenceRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode updates.getChannelDifference#3173d78 to nil")
	}
	if err := b.ConsumeID(UpdatesGetChannelDifferenceRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode updates.getChannelDifference#3173d78: %w", err)
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode updates.getChannelDifference#3173d78: field flags: %w", err)
		}
	}
	g.Force = g.Flags.Has(0)
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode updates.getChannelDifference#3173d78: field channel: %w", err)
		}
		g.Channel = value
	}
	{
		value, err := DecodeChannelMessagesFilter(b)
		if err != nil {
			return fmt.Errorf("unable to decode updates.getChannelDifference#3173d78: field filter: %w", err)
		}
		g.Filter = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode updates.getChannelDifference#3173d78: field pts: %w", err)
		}
		g.Pts = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode updates.getChannelDifference#3173d78: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// Ensuring interfaces in compile-time for UpdatesGetChannelDifferenceRequest.
var (
	_ bin.Encoder = &UpdatesGetChannelDifferenceRequest{}
	_ bin.Decoder = &UpdatesGetChannelDifferenceRequest{}
)

// UpdatesGetChannelDifference invokes method updates.getChannelDifference#3173d78 returning error if any.
// Returns the difference between the current state of updates of a certain channel and transmitted.
//
// See https://core.telegram.org/method/updates.getChannelDifference for reference.
func (c *Client) UpdatesGetChannelDifference(ctx context.Context, request *UpdatesGetChannelDifferenceRequest) (UpdatesChannelDifferenceClass, error) {
	var result UpdatesChannelDifferenceBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.ChannelDifference, nil
}
