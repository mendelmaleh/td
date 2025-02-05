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

// MessagesUnpinAllMessagesRequest represents TL type `messages.unpinAllMessages#f025bc8b`.
// Unpin¹ all pinned messages
//
// Links:
//  1) https://core.telegram.org/api/pin
//
// See https://core.telegram.org/method/messages.unpinAllMessages for reference.
type MessagesUnpinAllMessagesRequest struct {
	// Chat where to unpin
	Peer InputPeerClass
}

// MessagesUnpinAllMessagesRequestTypeID is TL type id of MessagesUnpinAllMessagesRequest.
const MessagesUnpinAllMessagesRequestTypeID = 0xf025bc8b

// String implements fmt.Stringer.
func (u *MessagesUnpinAllMessagesRequest) String() string {
	if u == nil {
		return "MessagesUnpinAllMessagesRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesUnpinAllMessagesRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tPeer: ")
	sb.WriteString(fmt.Sprint(u.Peer))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (u *MessagesUnpinAllMessagesRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode messages.unpinAllMessages#f025bc8b as nil")
	}
	b.PutID(MessagesUnpinAllMessagesRequestTypeID)
	if u.Peer == nil {
		return fmt.Errorf("unable to encode messages.unpinAllMessages#f025bc8b: field peer is nil")
	}
	if err := u.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.unpinAllMessages#f025bc8b: field peer: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (u *MessagesUnpinAllMessagesRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode messages.unpinAllMessages#f025bc8b to nil")
	}
	if err := b.ConsumeID(MessagesUnpinAllMessagesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.unpinAllMessages#f025bc8b: %w", err)
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.unpinAllMessages#f025bc8b: field peer: %w", err)
		}
		u.Peer = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesUnpinAllMessagesRequest.
var (
	_ bin.Encoder = &MessagesUnpinAllMessagesRequest{}
	_ bin.Decoder = &MessagesUnpinAllMessagesRequest{}
)

// MessagesUnpinAllMessages invokes method messages.unpinAllMessages#f025bc8b returning error if any.
// Unpin¹ all pinned messages
//
// Links:
//  1) https://core.telegram.org/api/pin
//
// See https://core.telegram.org/method/messages.unpinAllMessages for reference.
// Can be used by bots.
func (c *Client) MessagesUnpinAllMessages(ctx context.Context, peer InputPeerClass) (*MessagesAffectedHistory, error) {
	var result MessagesAffectedHistory

	request := &MessagesUnpinAllMessagesRequest{
		Peer: peer,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
