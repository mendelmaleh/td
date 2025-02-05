// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"fmt"
	"strings"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}

// UpdatesGetChannelDifferenceRequest represents TL type `updates.getChannelDifference#3173d78`.
// Returns the difference between the current state of updates of a certain channel and transmitted.
//
// See https://core.telegram.org/method/updates.getChannelDifference for reference.
type UpdatesGetChannelDifferenceRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Set to true to skip some possibly unneeded updates and reduce server-side load
	Force bool
	// The channel
	Channel InputChannelClass
	// Messsage filter
	Filter ChannelMessagesFilterClass
	// Persistent timestamp (see updates¹)
	//
	// Links:
	//  1) https://core.telegram.org/api/updates
	Pts int
	// How many updates to fetch, max 100000Ordinary (non-bot) users are supposed to pass 10-100
	Limit int
}

// UpdatesGetChannelDifferenceRequestTypeID is TL type id of UpdatesGetChannelDifferenceRequest.
const UpdatesGetChannelDifferenceRequestTypeID = 0x3173d78

// String implements fmt.Stringer.
func (g *UpdatesGetChannelDifferenceRequest) String() string {
	if g == nil {
		return "UpdatesGetChannelDifferenceRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("UpdatesGetChannelDifferenceRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(fmt.Sprint(g.Flags))
	sb.WriteString(",\n")
	sb.WriteString("\tChannel: ")
	sb.WriteString(fmt.Sprint(g.Channel))
	sb.WriteString(",\n")
	sb.WriteString("\tFilter: ")
	sb.WriteString(fmt.Sprint(g.Filter))
	sb.WriteString(",\n")
	sb.WriteString("\tPts: ")
	sb.WriteString(fmt.Sprint(g.Pts))
	sb.WriteString(",\n")
	sb.WriteString("\tLimit: ")
	sb.WriteString(fmt.Sprint(g.Limit))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

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
		g.Force = true
	} else {
		g.Flags.Unset(0)
		g.Force = false
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
// Possible errors:
//  400 CHANNEL_INVALID: The provided channel is invalid
//  400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup
//  403 CHANNEL_PUBLIC_GROUP_NA: channel/supergroup not available
//  400 FROM_MESSAGE_BOT_DISABLED: Bots can't use fromMessage min constructors
//  400 MSG_ID_INVALID: Invalid message ID provided
//  400 PERSISTENT_TIMESTAMP_EMPTY: Persistent timestamp empty
//  400 PERSISTENT_TIMESTAMP_INVALID: Persistent timestamp invalid
//  400 PINNED_DIALOGS_TOO_MUCH: Too many pinned dialogs
//  400 RANGES_INVALID: Invalid range provided
//
// See https://core.telegram.org/method/updates.getChannelDifference for reference.
// Can be used by bots.
func (c *Client) UpdatesGetChannelDifference(ctx context.Context, request *UpdatesGetChannelDifferenceRequest) (UpdatesChannelDifferenceClass, error) {
	var result UpdatesChannelDifferenceBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.ChannelDifference, nil
}
