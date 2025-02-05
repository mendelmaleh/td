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

// MessagesAcceptUrlAuthRequest represents TL type `messages.acceptUrlAuth#f729ea98`.
// Use this to accept a Seamless Telegram Login authorization request, for more info click here »¹
//
// Links:
//  1) https://core.telegram.org/api/url-authorization
//
// See https://core.telegram.org/method/messages.acceptUrlAuth for reference.
type MessagesAcceptUrlAuthRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Set this flag to allow the bot to send messages to you (if requested)
	WriteAllowed bool
	// The location of the message
	Peer InputPeerClass
	// Message ID of the message with the login button
	MsgID int
	// ID of the login button
	ButtonID int
}

// MessagesAcceptUrlAuthRequestTypeID is TL type id of MessagesAcceptUrlAuthRequest.
const MessagesAcceptUrlAuthRequestTypeID = 0xf729ea98

// String implements fmt.Stringer.
func (a *MessagesAcceptUrlAuthRequest) String() string {
	if a == nil {
		return "MessagesAcceptUrlAuthRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesAcceptUrlAuthRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(fmt.Sprint(a.Flags))
	sb.WriteString(",\n")
	sb.WriteString("\tPeer: ")
	sb.WriteString(fmt.Sprint(a.Peer))
	sb.WriteString(",\n")
	sb.WriteString("\tMsgID: ")
	sb.WriteString(fmt.Sprint(a.MsgID))
	sb.WriteString(",\n")
	sb.WriteString("\tButtonID: ")
	sb.WriteString(fmt.Sprint(a.ButtonID))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (a *MessagesAcceptUrlAuthRequest) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode messages.acceptUrlAuth#f729ea98 as nil")
	}
	b.PutID(MessagesAcceptUrlAuthRequestTypeID)
	if err := a.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.acceptUrlAuth#f729ea98: field flags: %w", err)
	}
	if a.Peer == nil {
		return fmt.Errorf("unable to encode messages.acceptUrlAuth#f729ea98: field peer is nil")
	}
	if err := a.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.acceptUrlAuth#f729ea98: field peer: %w", err)
	}
	b.PutInt(a.MsgID)
	b.PutInt(a.ButtonID)
	return nil
}

// SetWriteAllowed sets value of WriteAllowed conditional field.
func (a *MessagesAcceptUrlAuthRequest) SetWriteAllowed(value bool) {
	if value {
		a.Flags.Set(0)
		a.WriteAllowed = true
	} else {
		a.Flags.Unset(0)
		a.WriteAllowed = false
	}
}

// Decode implements bin.Decoder.
func (a *MessagesAcceptUrlAuthRequest) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode messages.acceptUrlAuth#f729ea98 to nil")
	}
	if err := b.ConsumeID(MessagesAcceptUrlAuthRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.acceptUrlAuth#f729ea98: %w", err)
	}
	{
		if err := a.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.acceptUrlAuth#f729ea98: field flags: %w", err)
		}
	}
	a.WriteAllowed = a.Flags.Has(0)
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.acceptUrlAuth#f729ea98: field peer: %w", err)
		}
		a.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.acceptUrlAuth#f729ea98: field msg_id: %w", err)
		}
		a.MsgID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.acceptUrlAuth#f729ea98: field button_id: %w", err)
		}
		a.ButtonID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesAcceptUrlAuthRequest.
var (
	_ bin.Encoder = &MessagesAcceptUrlAuthRequest{}
	_ bin.Decoder = &MessagesAcceptUrlAuthRequest{}
)

// MessagesAcceptUrlAuth invokes method messages.acceptUrlAuth#f729ea98 returning error if any.
// Use this to accept a Seamless Telegram Login authorization request, for more info click here »¹
//
// Links:
//  1) https://core.telegram.org/api/url-authorization
//
// See https://core.telegram.org/method/messages.acceptUrlAuth for reference.
func (c *Client) MessagesAcceptUrlAuth(ctx context.Context, request *MessagesAcceptUrlAuthRequest) (UrlAuthResultClass, error) {
	var result UrlAuthResultBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.UrlAuthResult, nil
}
