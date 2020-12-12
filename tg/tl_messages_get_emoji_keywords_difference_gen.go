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

// MessagesGetEmojiKeywordsDifferenceRequest represents TL type `messages.getEmojiKeywordsDifference#1508b6af`.
// Get changed emoji keywords
//
// See https://core.telegram.org/method/messages.getEmojiKeywordsDifference for reference.
type MessagesGetEmojiKeywordsDifferenceRequest struct {
	// Language code
	LangCode string
	// Previous emoji keyword localization version
	FromVersion int
}

// MessagesGetEmojiKeywordsDifferenceRequestTypeID is TL type id of MessagesGetEmojiKeywordsDifferenceRequest.
const MessagesGetEmojiKeywordsDifferenceRequestTypeID = 0x1508b6af

// Encode implements bin.Encoder.
func (g *MessagesGetEmojiKeywordsDifferenceRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getEmojiKeywordsDifference#1508b6af as nil")
	}
	b.PutID(MessagesGetEmojiKeywordsDifferenceRequestTypeID)
	b.PutString(g.LangCode)
	b.PutInt(g.FromVersion)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetEmojiKeywordsDifferenceRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getEmojiKeywordsDifference#1508b6af to nil")
	}
	if err := b.ConsumeID(MessagesGetEmojiKeywordsDifferenceRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getEmojiKeywordsDifference#1508b6af: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getEmojiKeywordsDifference#1508b6af: field lang_code: %w", err)
		}
		g.LangCode = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getEmojiKeywordsDifference#1508b6af: field from_version: %w", err)
		}
		g.FromVersion = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesGetEmojiKeywordsDifferenceRequest.
var (
	_ bin.Encoder = &MessagesGetEmojiKeywordsDifferenceRequest{}
	_ bin.Decoder = &MessagesGetEmojiKeywordsDifferenceRequest{}
)

// MessagesGetEmojiKeywordsDifference invokes method messages.getEmojiKeywordsDifference#1508b6af returning error if any.
// Get changed emoji keywords
//
// See https://core.telegram.org/method/messages.getEmojiKeywordsDifference for reference.
func (c *Client) MessagesGetEmojiKeywordsDifference(ctx context.Context, request *MessagesGetEmojiKeywordsDifferenceRequest) (*EmojiKeywordsDifference, error) {
	var result EmojiKeywordsDifference

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
