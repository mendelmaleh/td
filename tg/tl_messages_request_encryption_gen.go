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

// MessagesRequestEncryptionRequest represents TL type `messages.requestEncryption#f64daf43`.
// Sends a request to start a secret chat to the user.
//
// See https://core.telegram.org/method/messages.requestEncryption for reference.
type MessagesRequestEncryptionRequest struct {
	// User ID
	UserID InputUserClass
	// Unique client request ID required to prevent resending. This also doubles as the chat ID.
	RandomID int
	// A = g ^ a mod p, see Wikipedia¹
	//
	// Links:
	//  1) https://en.wikipedia.org/wiki/Diffie%E2%80%93Hellman_key_exchange
	GA []byte
}

// MessagesRequestEncryptionRequestTypeID is TL type id of MessagesRequestEncryptionRequest.
const MessagesRequestEncryptionRequestTypeID = 0xf64daf43

// String implements fmt.Stringer.
func (r *MessagesRequestEncryptionRequest) String() string {
	if r == nil {
		return "MessagesRequestEncryptionRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesRequestEncryptionRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tUserID: ")
	sb.WriteString(fmt.Sprint(r.UserID))
	sb.WriteString(",\n")
	sb.WriteString("\tRandomID: ")
	sb.WriteString(fmt.Sprint(r.RandomID))
	sb.WriteString(",\n")
	sb.WriteString("\tGA: ")
	sb.WriteString(fmt.Sprint(r.GA))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (r *MessagesRequestEncryptionRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.requestEncryption#f64daf43 as nil")
	}
	b.PutID(MessagesRequestEncryptionRequestTypeID)
	if r.UserID == nil {
		return fmt.Errorf("unable to encode messages.requestEncryption#f64daf43: field user_id is nil")
	}
	if err := r.UserID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.requestEncryption#f64daf43: field user_id: %w", err)
	}
	b.PutInt(r.RandomID)
	b.PutBytes(r.GA)
	return nil
}

// Decode implements bin.Decoder.
func (r *MessagesRequestEncryptionRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.requestEncryption#f64daf43 to nil")
	}
	if err := b.ConsumeID(MessagesRequestEncryptionRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.requestEncryption#f64daf43: %w", err)
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.requestEncryption#f64daf43: field user_id: %w", err)
		}
		r.UserID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.requestEncryption#f64daf43: field random_id: %w", err)
		}
		r.RandomID = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode messages.requestEncryption#f64daf43: field g_a: %w", err)
		}
		r.GA = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesRequestEncryptionRequest.
var (
	_ bin.Encoder = &MessagesRequestEncryptionRequest{}
	_ bin.Decoder = &MessagesRequestEncryptionRequest{}
)

// MessagesRequestEncryption invokes method messages.requestEncryption#f64daf43 returning error if any.
// Sends a request to start a secret chat to the user.
//
// Possible errors:
//  400 DH_G_A_INVALID: g_a invalid
//  400 USER_ID_INVALID: The provided user ID is invalid
//
// See https://core.telegram.org/method/messages.requestEncryption for reference.
func (c *Client) MessagesRequestEncryption(ctx context.Context, request *MessagesRequestEncryptionRequest) (EncryptedChatClass, error) {
	var result EncryptedChatBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.EncryptedChat, nil
}
