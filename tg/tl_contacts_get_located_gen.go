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

// ContactsGetLocatedRequest represents TL type `contacts.getLocated#d348bc44`.
// Get contacts near you
//
// See https://core.telegram.org/method/contacts.getLocated for reference.
type ContactsGetLocatedRequest struct {
	// Flags, see TL conditional fields
	Flags bin.Fields
	// While the geolocation of the current user is public, clients should update it in the background every half-an-hour or so, while setting this flag. Do this only if the new location is more than 1 KM away from the previous one, or if the previous location is unknown.
	Background bool
	// Geolocation
	GeoPoint InputGeoPointClass
	// If set, the geolocation of the current user will be public for the specified number of seconds; pass 0x7fffffff to disable expiry, 0 to make the current geolocation private; if the flag isn't set, no changes will be applied.
	//
	// Use SetSelfExpires and GetSelfExpires helpers.
	SelfExpires int
}

// ContactsGetLocatedRequestTypeID is TL type id of ContactsGetLocatedRequest.
const ContactsGetLocatedRequestTypeID = 0xd348bc44

// Encode implements bin.Encoder.
func (g *ContactsGetLocatedRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode contacts.getLocated#d348bc44 as nil")
	}
	b.PutID(ContactsGetLocatedRequestTypeID)
	if err := g.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode contacts.getLocated#d348bc44: field flags: %w", err)
	}
	if g.GeoPoint == nil {
		return fmt.Errorf("unable to encode contacts.getLocated#d348bc44: field geo_point is nil")
	}
	if err := g.GeoPoint.Encode(b); err != nil {
		return fmt.Errorf("unable to encode contacts.getLocated#d348bc44: field geo_point: %w", err)
	}
	if g.Flags.Has(0) {
		b.PutInt(g.SelfExpires)
	}
	return nil
}

// SetBackground sets value of Background conditional field.
func (g *ContactsGetLocatedRequest) SetBackground(value bool) {
	if value {
		g.Flags.Set(1)
	} else {
		g.Flags.Unset(1)
	}
}

// SetSelfExpires sets value of SelfExpires conditional field.
func (g *ContactsGetLocatedRequest) SetSelfExpires(value int) {
	g.Flags.Set(0)
	g.SelfExpires = value
}

// GetSelfExpires returns value of SelfExpires conditional field and
// boolean which is true if field was set.
func (g *ContactsGetLocatedRequest) GetSelfExpires() (value int, ok bool) {
	if !g.Flags.Has(0) {
		return value, false
	}
	return g.SelfExpires, true
}

// Decode implements bin.Decoder.
func (g *ContactsGetLocatedRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode contacts.getLocated#d348bc44 to nil")
	}
	if err := b.ConsumeID(ContactsGetLocatedRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.getLocated#d348bc44: %w", err)
	}
	{
		if err := g.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode contacts.getLocated#d348bc44: field flags: %w", err)
		}
	}
	g.Background = g.Flags.Has(1)
	{
		value, err := DecodeInputGeoPoint(b)
		if err != nil {
			return fmt.Errorf("unable to decode contacts.getLocated#d348bc44: field geo_point: %w", err)
		}
		g.GeoPoint = value
	}
	if g.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.getLocated#d348bc44: field self_expires: %w", err)
		}
		g.SelfExpires = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ContactsGetLocatedRequest.
var (
	_ bin.Encoder = &ContactsGetLocatedRequest{}
	_ bin.Decoder = &ContactsGetLocatedRequest{}
)

// ContactsGetLocated invokes method contacts.getLocated#d348bc44 returning error if any.
// Get contacts near you
//
// See https://core.telegram.org/method/contacts.getLocated for reference.
func (c *Client) ContactsGetLocated(ctx context.Context, request *ContactsGetLocatedRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
