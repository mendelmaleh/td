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

// ThemeSettings represents TL type `themeSettings#9c14984a`.
// Theme settings
//
// See https://core.telegram.org/constructor/themeSettings for reference.
type ThemeSettings struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Base theme
	BaseTheme BaseThemeClass
	// Accent color, RGB24 format
	AccentColor int
	// Message gradient color (top), RGB24 format
	//
	// Use SetMessageTopColor and GetMessageTopColor helpers.
	MessageTopColor int
	// Message gradient color (bottom), RGB24 format
	//
	// Use SetMessageBottomColor and GetMessageBottomColor helpers.
	MessageBottomColor int
	// Wallpaper
	//
	// Use SetWallpaper and GetWallpaper helpers.
	Wallpaper WallPaperClass
}

// ThemeSettingsTypeID is TL type id of ThemeSettings.
const ThemeSettingsTypeID = 0x9c14984a

// String implements fmt.Stringer.
func (t *ThemeSettings) String() string {
	if t == nil {
		return "ThemeSettings(nil)"
	}
	var sb strings.Builder
	sb.WriteString("ThemeSettings")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(fmt.Sprint(t.Flags))
	sb.WriteString(",\n")
	sb.WriteString("\tBaseTheme: ")
	sb.WriteString(fmt.Sprint(t.BaseTheme))
	sb.WriteString(",\n")
	sb.WriteString("\tAccentColor: ")
	sb.WriteString(fmt.Sprint(t.AccentColor))
	sb.WriteString(",\n")
	if t.Flags.Has(0) {
		sb.WriteString("\tMessageTopColor: ")
		sb.WriteString(fmt.Sprint(t.MessageTopColor))
		sb.WriteString(",\n")
	}
	if t.Flags.Has(0) {
		sb.WriteString("\tMessageBottomColor: ")
		sb.WriteString(fmt.Sprint(t.MessageBottomColor))
		sb.WriteString(",\n")
	}
	if t.Flags.Has(1) {
		sb.WriteString("\tWallpaper: ")
		sb.WriteString(fmt.Sprint(t.Wallpaper))
		sb.WriteString(",\n")
	}
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (t *ThemeSettings) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode themeSettings#9c14984a as nil")
	}
	b.PutID(ThemeSettingsTypeID)
	if err := t.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode themeSettings#9c14984a: field flags: %w", err)
	}
	if t.BaseTheme == nil {
		return fmt.Errorf("unable to encode themeSettings#9c14984a: field base_theme is nil")
	}
	if err := t.BaseTheme.Encode(b); err != nil {
		return fmt.Errorf("unable to encode themeSettings#9c14984a: field base_theme: %w", err)
	}
	b.PutInt(t.AccentColor)
	if t.Flags.Has(0) {
		b.PutInt(t.MessageTopColor)
	}
	if t.Flags.Has(0) {
		b.PutInt(t.MessageBottomColor)
	}
	if t.Flags.Has(1) {
		if t.Wallpaper == nil {
			return fmt.Errorf("unable to encode themeSettings#9c14984a: field wallpaper is nil")
		}
		if err := t.Wallpaper.Encode(b); err != nil {
			return fmt.Errorf("unable to encode themeSettings#9c14984a: field wallpaper: %w", err)
		}
	}
	return nil
}

// SetMessageTopColor sets value of MessageTopColor conditional field.
func (t *ThemeSettings) SetMessageTopColor(value int) {
	t.Flags.Set(0)
	t.MessageTopColor = value
}

// GetMessageTopColor returns value of MessageTopColor conditional field and
// boolean which is true if field was set.
func (t *ThemeSettings) GetMessageTopColor() (value int, ok bool) {
	if !t.Flags.Has(0) {
		return value, false
	}
	return t.MessageTopColor, true
}

// SetMessageBottomColor sets value of MessageBottomColor conditional field.
func (t *ThemeSettings) SetMessageBottomColor(value int) {
	t.Flags.Set(0)
	t.MessageBottomColor = value
}

// GetMessageBottomColor returns value of MessageBottomColor conditional field and
// boolean which is true if field was set.
func (t *ThemeSettings) GetMessageBottomColor() (value int, ok bool) {
	if !t.Flags.Has(0) {
		return value, false
	}
	return t.MessageBottomColor, true
}

// SetWallpaper sets value of Wallpaper conditional field.
func (t *ThemeSettings) SetWallpaper(value WallPaperClass) {
	t.Flags.Set(1)
	t.Wallpaper = value
}

// GetWallpaper returns value of Wallpaper conditional field and
// boolean which is true if field was set.
func (t *ThemeSettings) GetWallpaper() (value WallPaperClass, ok bool) {
	if !t.Flags.Has(1) {
		return value, false
	}
	return t.Wallpaper, true
}

// Decode implements bin.Decoder.
func (t *ThemeSettings) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode themeSettings#9c14984a to nil")
	}
	if err := b.ConsumeID(ThemeSettingsTypeID); err != nil {
		return fmt.Errorf("unable to decode themeSettings#9c14984a: %w", err)
	}
	{
		if err := t.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode themeSettings#9c14984a: field flags: %w", err)
		}
	}
	{
		value, err := DecodeBaseTheme(b)
		if err != nil {
			return fmt.Errorf("unable to decode themeSettings#9c14984a: field base_theme: %w", err)
		}
		t.BaseTheme = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode themeSettings#9c14984a: field accent_color: %w", err)
		}
		t.AccentColor = value
	}
	if t.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode themeSettings#9c14984a: field message_top_color: %w", err)
		}
		t.MessageTopColor = value
	}
	if t.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode themeSettings#9c14984a: field message_bottom_color: %w", err)
		}
		t.MessageBottomColor = value
	}
	if t.Flags.Has(1) {
		value, err := DecodeWallPaper(b)
		if err != nil {
			return fmt.Errorf("unable to decode themeSettings#9c14984a: field wallpaper: %w", err)
		}
		t.Wallpaper = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ThemeSettings.
var (
	_ bin.Encoder = &ThemeSettings{}
	_ bin.Decoder = &ThemeSettings{}
)
