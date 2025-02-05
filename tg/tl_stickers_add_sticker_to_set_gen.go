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

// StickersAddStickerToSetRequest represents TL type `stickers.addStickerToSet#8653febe`.
// Add a sticker to a stickerset, bots only. The sticker set must have been created by the bot.
//
// See https://core.telegram.org/method/stickers.addStickerToSet for reference.
type StickersAddStickerToSetRequest struct {
	// The stickerset
	Stickerset InputStickerSetClass
	// The sticker
	Sticker InputStickerSetItem
}

// StickersAddStickerToSetRequestTypeID is TL type id of StickersAddStickerToSetRequest.
const StickersAddStickerToSetRequestTypeID = 0x8653febe

// String implements fmt.Stringer.
func (a *StickersAddStickerToSetRequest) String() string {
	if a == nil {
		return "StickersAddStickerToSetRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("StickersAddStickerToSetRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tStickerset: ")
	sb.WriteString(fmt.Sprint(a.Stickerset))
	sb.WriteString(",\n")
	sb.WriteString("\tSticker: ")
	sb.WriteString(fmt.Sprint(a.Sticker))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (a *StickersAddStickerToSetRequest) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode stickers.addStickerToSet#8653febe as nil")
	}
	b.PutID(StickersAddStickerToSetRequestTypeID)
	if a.Stickerset == nil {
		return fmt.Errorf("unable to encode stickers.addStickerToSet#8653febe: field stickerset is nil")
	}
	if err := a.Stickerset.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stickers.addStickerToSet#8653febe: field stickerset: %w", err)
	}
	if err := a.Sticker.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stickers.addStickerToSet#8653febe: field sticker: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *StickersAddStickerToSetRequest) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode stickers.addStickerToSet#8653febe to nil")
	}
	if err := b.ConsumeID(StickersAddStickerToSetRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode stickers.addStickerToSet#8653febe: %w", err)
	}
	{
		value, err := DecodeInputStickerSet(b)
		if err != nil {
			return fmt.Errorf("unable to decode stickers.addStickerToSet#8653febe: field stickerset: %w", err)
		}
		a.Stickerset = value
	}
	{
		if err := a.Sticker.Decode(b); err != nil {
			return fmt.Errorf("unable to decode stickers.addStickerToSet#8653febe: field sticker: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for StickersAddStickerToSetRequest.
var (
	_ bin.Encoder = &StickersAddStickerToSetRequest{}
	_ bin.Decoder = &StickersAddStickerToSetRequest{}
)

// StickersAddStickerToSet invokes method stickers.addStickerToSet#8653febe returning error if any.
// Add a sticker to a stickerset, bots only. The sticker set must have been created by the bot.
//
// Possible errors:
//  400 BOT_MISSING: This method can only be run by a bot
//  400 STICKERSET_INVALID: The provided sticker set is invalid
//
// See https://core.telegram.org/method/stickers.addStickerToSet for reference.
// Can be used by bots.
func (c *Client) StickersAddStickerToSet(ctx context.Context, request *StickersAddStickerToSetRequest) (*MessagesStickerSet, error) {
	var result MessagesStickerSet

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
