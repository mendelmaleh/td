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

// PageListItemText represents TL type `pageListItemText#b92fb6cd`.
// List item
//
// See https://core.telegram.org/constructor/pageListItemText for reference.
type PageListItemText struct {
	// Text
	Text RichTextClass
}

// PageListItemTextTypeID is TL type id of PageListItemText.
const PageListItemTextTypeID = 0xb92fb6cd

// String implements fmt.Stringer.
func (p *PageListItemText) String() string {
	if p == nil {
		return "PageListItemText(nil)"
	}
	var sb strings.Builder
	sb.WriteString("PageListItemText")
	sb.WriteString("{\n")
	sb.WriteString("\tText: ")
	sb.WriteString(fmt.Sprint(p.Text))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (p *PageListItemText) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode pageListItemText#b92fb6cd as nil")
	}
	b.PutID(PageListItemTextTypeID)
	if p.Text == nil {
		return fmt.Errorf("unable to encode pageListItemText#b92fb6cd: field text is nil")
	}
	if err := p.Text.Encode(b); err != nil {
		return fmt.Errorf("unable to encode pageListItemText#b92fb6cd: field text: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *PageListItemText) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode pageListItemText#b92fb6cd to nil")
	}
	if err := b.ConsumeID(PageListItemTextTypeID); err != nil {
		return fmt.Errorf("unable to decode pageListItemText#b92fb6cd: %w", err)
	}
	{
		value, err := DecodeRichText(b)
		if err != nil {
			return fmt.Errorf("unable to decode pageListItemText#b92fb6cd: field text: %w", err)
		}
		p.Text = value
	}
	return nil
}

// construct implements constructor of PageListItemClass.
func (p PageListItemText) construct() PageListItemClass { return &p }

// Ensuring interfaces in compile-time for PageListItemText.
var (
	_ bin.Encoder = &PageListItemText{}
	_ bin.Decoder = &PageListItemText{}

	_ PageListItemClass = &PageListItemText{}
)

// PageListItemBlocks represents TL type `pageListItemBlocks#25e073fc`.
// List item
//
// See https://core.telegram.org/constructor/pageListItemBlocks for reference.
type PageListItemBlocks struct {
	// Blocks
	Blocks []PageBlockClass
}

// PageListItemBlocksTypeID is TL type id of PageListItemBlocks.
const PageListItemBlocksTypeID = 0x25e073fc

// String implements fmt.Stringer.
func (p *PageListItemBlocks) String() string {
	if p == nil {
		return "PageListItemBlocks(nil)"
	}
	var sb strings.Builder
	sb.WriteString("PageListItemBlocks")
	sb.WriteString("{\n")
	sb.WriteByte('[')
	for _, v := range p.Blocks {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (p *PageListItemBlocks) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode pageListItemBlocks#25e073fc as nil")
	}
	b.PutID(PageListItemBlocksTypeID)
	b.PutVectorHeader(len(p.Blocks))
	for idx, v := range p.Blocks {
		if v == nil {
			return fmt.Errorf("unable to encode pageListItemBlocks#25e073fc: field blocks element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode pageListItemBlocks#25e073fc: field blocks element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *PageListItemBlocks) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode pageListItemBlocks#25e073fc to nil")
	}
	if err := b.ConsumeID(PageListItemBlocksTypeID); err != nil {
		return fmt.Errorf("unable to decode pageListItemBlocks#25e073fc: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode pageListItemBlocks#25e073fc: field blocks: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodePageBlock(b)
			if err != nil {
				return fmt.Errorf("unable to decode pageListItemBlocks#25e073fc: field blocks: %w", err)
			}
			p.Blocks = append(p.Blocks, value)
		}
	}
	return nil
}

// construct implements constructor of PageListItemClass.
func (p PageListItemBlocks) construct() PageListItemClass { return &p }

// Ensuring interfaces in compile-time for PageListItemBlocks.
var (
	_ bin.Encoder = &PageListItemBlocks{}
	_ bin.Decoder = &PageListItemBlocks{}

	_ PageListItemClass = &PageListItemBlocks{}
)

// PageListItemClass represents PageListItem generic type.
//
// See https://core.telegram.org/type/PageListItem for reference.
//
// Example:
//  g, err := DecodePageListItem(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *PageListItemText: // pageListItemText#b92fb6cd
//  case *PageListItemBlocks: // pageListItemBlocks#25e073fc
//  default: panic(v)
//  }
type PageListItemClass interface {
	bin.Encoder
	bin.Decoder
	construct() PageListItemClass
	fmt.Stringer
}

// DecodePageListItem implements binary de-serialization for PageListItemClass.
func DecodePageListItem(buf *bin.Buffer) (PageListItemClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case PageListItemTextTypeID:
		// Decoding pageListItemText#b92fb6cd.
		v := PageListItemText{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PageListItemClass: %w", err)
		}
		return &v, nil
	case PageListItemBlocksTypeID:
		// Decoding pageListItemBlocks#25e073fc.
		v := PageListItemBlocks{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PageListItemClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode PageListItemClass: %w", bin.NewUnexpectedID(id))
	}
}

// PageListItem boxes the PageListItemClass providing a helper.
type PageListItemBox struct {
	PageListItem PageListItemClass
}

// Decode implements bin.Decoder for PageListItemBox.
func (b *PageListItemBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode PageListItemBox to nil")
	}
	v, err := DecodePageListItem(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.PageListItem = v
	return nil
}

// Encode implements bin.Encode for PageListItemBox.
func (b *PageListItemBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.PageListItem == nil {
		return fmt.Errorf("unable to encode PageListItemClass as nil")
	}
	return b.PageListItem.Encode(buf)
}
