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

// MessagesSaveGifRequest represents TL type `messages.saveGif#327a30cb`.
// Add GIF to saved gifs list
//
// See https://core.telegram.org/method/messages.saveGif for reference.
type MessagesSaveGifRequest struct {
	// GIF to save
	ID InputDocumentClass
	// Whether to remove GIF from saved gifs list
	Unsave bool
}

// MessagesSaveGifRequestTypeID is TL type id of MessagesSaveGifRequest.
const MessagesSaveGifRequestTypeID = 0x327a30cb

// String implements fmt.Stringer.
func (s *MessagesSaveGifRequest) String() string {
	if s == nil {
		return "MessagesSaveGifRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesSaveGifRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tID: ")
	sb.WriteString(fmt.Sprint(s.ID))
	sb.WriteString(",\n")
	sb.WriteString("\tUnsave: ")
	sb.WriteString(fmt.Sprint(s.Unsave))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (s *MessagesSaveGifRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.saveGif#327a30cb as nil")
	}
	b.PutID(MessagesSaveGifRequestTypeID)
	if s.ID == nil {
		return fmt.Errorf("unable to encode messages.saveGif#327a30cb: field id is nil")
	}
	if err := s.ID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.saveGif#327a30cb: field id: %w", err)
	}
	b.PutBool(s.Unsave)
	return nil
}

// Decode implements bin.Decoder.
func (s *MessagesSaveGifRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.saveGif#327a30cb to nil")
	}
	if err := b.ConsumeID(MessagesSaveGifRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.saveGif#327a30cb: %w", err)
	}
	{
		value, err := DecodeInputDocument(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.saveGif#327a30cb: field id: %w", err)
		}
		s.ID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode messages.saveGif#327a30cb: field unsave: %w", err)
		}
		s.Unsave = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesSaveGifRequest.
var (
	_ bin.Encoder = &MessagesSaveGifRequest{}
	_ bin.Decoder = &MessagesSaveGifRequest{}
)

// MessagesSaveGif invokes method messages.saveGif#327a30cb returning error if any.
// Add GIF to saved gifs list
//
// Possible errors:
//  400 GIF_ID_INVALID: The provided GIF ID is invalid
//
// See https://core.telegram.org/method/messages.saveGif for reference.
func (c *Client) MessagesSaveGif(ctx context.Context, request *MessagesSaveGifRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
