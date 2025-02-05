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

// MessagesGetMessagesViewsRequest represents TL type `messages.getMessagesViews#5784d3e1`.
// Get and increase the view counter of a message sent or forwarded from a channel¹
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// See https://core.telegram.org/method/messages.getMessagesViews for reference.
type MessagesGetMessagesViewsRequest struct {
	// Peer where the message was found
	Peer InputPeerClass
	// ID of message
	ID []int
	// Whether to mark the message as viewed and increment the view counter
	Increment bool
}

// MessagesGetMessagesViewsRequestTypeID is TL type id of MessagesGetMessagesViewsRequest.
const MessagesGetMessagesViewsRequestTypeID = 0x5784d3e1

// String implements fmt.Stringer.
func (g *MessagesGetMessagesViewsRequest) String() string {
	if g == nil {
		return "MessagesGetMessagesViewsRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesGetMessagesViewsRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tPeer: ")
	sb.WriteString(fmt.Sprint(g.Peer))
	sb.WriteString(",\n")
	sb.WriteByte('[')
	for _, v := range g.ID {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteString("\tIncrement: ")
	sb.WriteString(fmt.Sprint(g.Increment))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (g *MessagesGetMessagesViewsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getMessagesViews#5784d3e1 as nil")
	}
	b.PutID(MessagesGetMessagesViewsRequestTypeID)
	if g.Peer == nil {
		return fmt.Errorf("unable to encode messages.getMessagesViews#5784d3e1: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getMessagesViews#5784d3e1: field peer: %w", err)
	}
	b.PutVectorHeader(len(g.ID))
	for _, v := range g.ID {
		b.PutInt(v)
	}
	b.PutBool(g.Increment)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetMessagesViewsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getMessagesViews#5784d3e1 to nil")
	}
	if err := b.ConsumeID(MessagesGetMessagesViewsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getMessagesViews#5784d3e1: %w", err)
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getMessagesViews#5784d3e1: field peer: %w", err)
		}
		g.Peer = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getMessagesViews#5784d3e1: field id: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode messages.getMessagesViews#5784d3e1: field id: %w", err)
			}
			g.ID = append(g.ID, value)
		}
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getMessagesViews#5784d3e1: field increment: %w", err)
		}
		g.Increment = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetMessagesViewsRequest.
var (
	_ bin.Encoder = &MessagesGetMessagesViewsRequest{}
	_ bin.Decoder = &MessagesGetMessagesViewsRequest{}
)

// MessagesGetMessagesViews invokes method messages.getMessagesViews#5784d3e1 returning error if any.
// Get and increase the view counter of a message sent or forwarded from a channel¹
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// Possible errors:
//  400 CHANNEL_INVALID: The provided channel is invalid
//  400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup
//  400 CHAT_ID_INVALID: The provided chat id is invalid
//  400 MSG_ID_INVALID: Invalid message ID provided
//  400 PEER_ID_INVALID: The provided peer id is invalid
//
// See https://core.telegram.org/method/messages.getMessagesViews for reference.
func (c *Client) MessagesGetMessagesViews(ctx context.Context, request *MessagesGetMessagesViewsRequest) (*MessagesMessageViews, error) {
	var result MessagesMessageViews

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
