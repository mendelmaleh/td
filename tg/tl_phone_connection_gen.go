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

// PhoneConnection represents TL type `phoneConnection#9d4c17c0`.
// Identifies an endpoint that can be used to connect to the other user in a phone call
//
// See https://core.telegram.org/constructor/phoneConnection for reference.
type PhoneConnection struct {
	// Endpoint ID
	ID int64
	// IP address of endpoint
	IP string
	// IPv6 address of endpoint
	Ipv6 string
	// Port ID
	Port int
	// Our peer tag
	PeerTag []byte
}

// PhoneConnectionTypeID is TL type id of PhoneConnection.
const PhoneConnectionTypeID = 0x9d4c17c0

// String implements fmt.Stringer.
func (p *PhoneConnection) String() string {
	if p == nil {
		return "PhoneConnection(nil)"
	}
	var sb strings.Builder
	sb.WriteString("PhoneConnection")
	sb.WriteString("{\n")
	sb.WriteString("\tID: ")
	sb.WriteString(fmt.Sprint(p.ID))
	sb.WriteString(",\n")
	sb.WriteString("\tIP: ")
	sb.WriteString(fmt.Sprint(p.IP))
	sb.WriteString(",\n")
	sb.WriteString("\tIpv6: ")
	sb.WriteString(fmt.Sprint(p.Ipv6))
	sb.WriteString(",\n")
	sb.WriteString("\tPort: ")
	sb.WriteString(fmt.Sprint(p.Port))
	sb.WriteString(",\n")
	sb.WriteString("\tPeerTag: ")
	sb.WriteString(fmt.Sprint(p.PeerTag))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (p *PhoneConnection) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode phoneConnection#9d4c17c0 as nil")
	}
	b.PutID(PhoneConnectionTypeID)
	b.PutLong(p.ID)
	b.PutString(p.IP)
	b.PutString(p.Ipv6)
	b.PutInt(p.Port)
	b.PutBytes(p.PeerTag)
	return nil
}

// Decode implements bin.Decoder.
func (p *PhoneConnection) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode phoneConnection#9d4c17c0 to nil")
	}
	if err := b.ConsumeID(PhoneConnectionTypeID); err != nil {
		return fmt.Errorf("unable to decode phoneConnection#9d4c17c0: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode phoneConnection#9d4c17c0: field id: %w", err)
		}
		p.ID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode phoneConnection#9d4c17c0: field ip: %w", err)
		}
		p.IP = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode phoneConnection#9d4c17c0: field ipv6: %w", err)
		}
		p.Ipv6 = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode phoneConnection#9d4c17c0: field port: %w", err)
		}
		p.Port = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode phoneConnection#9d4c17c0: field peer_tag: %w", err)
		}
		p.PeerTag = value
	}
	return nil
}

// construct implements constructor of PhoneConnectionClass.
func (p PhoneConnection) construct() PhoneConnectionClass { return &p }

// Ensuring interfaces in compile-time for PhoneConnection.
var (
	_ bin.Encoder = &PhoneConnection{}
	_ bin.Decoder = &PhoneConnection{}

	_ PhoneConnectionClass = &PhoneConnection{}
)

// PhoneConnectionWebrtc represents TL type `phoneConnectionWebrtc#635fe375`.
// WebRTC connection parameters
//
// See https://core.telegram.org/constructor/phoneConnectionWebrtc for reference.
type PhoneConnectionWebrtc struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether this is a TURN endpoint
	Turn bool
	// Whether this is a STUN endpoint
	Stun bool
	// Endpoint ID
	ID int64
	// IP address
	IP string
	// IPv6 address
	Ipv6 string
	// Port
	Port int
	// Username
	Username string
	// Password
	Password string
}

// PhoneConnectionWebrtcTypeID is TL type id of PhoneConnectionWebrtc.
const PhoneConnectionWebrtcTypeID = 0x635fe375

// String implements fmt.Stringer.
func (p *PhoneConnectionWebrtc) String() string {
	if p == nil {
		return "PhoneConnectionWebrtc(nil)"
	}
	var sb strings.Builder
	sb.WriteString("PhoneConnectionWebrtc")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(fmt.Sprint(p.Flags))
	sb.WriteString(",\n")
	sb.WriteString("\tID: ")
	sb.WriteString(fmt.Sprint(p.ID))
	sb.WriteString(",\n")
	sb.WriteString("\tIP: ")
	sb.WriteString(fmt.Sprint(p.IP))
	sb.WriteString(",\n")
	sb.WriteString("\tIpv6: ")
	sb.WriteString(fmt.Sprint(p.Ipv6))
	sb.WriteString(",\n")
	sb.WriteString("\tPort: ")
	sb.WriteString(fmt.Sprint(p.Port))
	sb.WriteString(",\n")
	sb.WriteString("\tUsername: ")
	sb.WriteString(fmt.Sprint(p.Username))
	sb.WriteString(",\n")
	sb.WriteString("\tPassword: ")
	sb.WriteString(fmt.Sprint(p.Password))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (p *PhoneConnectionWebrtc) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode phoneConnectionWebrtc#635fe375 as nil")
	}
	b.PutID(PhoneConnectionWebrtcTypeID)
	if err := p.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phoneConnectionWebrtc#635fe375: field flags: %w", err)
	}
	b.PutLong(p.ID)
	b.PutString(p.IP)
	b.PutString(p.Ipv6)
	b.PutInt(p.Port)
	b.PutString(p.Username)
	b.PutString(p.Password)
	return nil
}

// SetTurn sets value of Turn conditional field.
func (p *PhoneConnectionWebrtc) SetTurn(value bool) {
	if value {
		p.Flags.Set(0)
		p.Turn = true
	} else {
		p.Flags.Unset(0)
		p.Turn = false
	}
}

// SetStun sets value of Stun conditional field.
func (p *PhoneConnectionWebrtc) SetStun(value bool) {
	if value {
		p.Flags.Set(1)
		p.Stun = true
	} else {
		p.Flags.Unset(1)
		p.Stun = false
	}
}

// Decode implements bin.Decoder.
func (p *PhoneConnectionWebrtc) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode phoneConnectionWebrtc#635fe375 to nil")
	}
	if err := b.ConsumeID(PhoneConnectionWebrtcTypeID); err != nil {
		return fmt.Errorf("unable to decode phoneConnectionWebrtc#635fe375: %w", err)
	}
	{
		if err := p.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phoneConnectionWebrtc#635fe375: field flags: %w", err)
		}
	}
	p.Turn = p.Flags.Has(0)
	p.Stun = p.Flags.Has(1)
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode phoneConnectionWebrtc#635fe375: field id: %w", err)
		}
		p.ID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode phoneConnectionWebrtc#635fe375: field ip: %w", err)
		}
		p.IP = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode phoneConnectionWebrtc#635fe375: field ipv6: %w", err)
		}
		p.Ipv6 = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode phoneConnectionWebrtc#635fe375: field port: %w", err)
		}
		p.Port = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode phoneConnectionWebrtc#635fe375: field username: %w", err)
		}
		p.Username = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode phoneConnectionWebrtc#635fe375: field password: %w", err)
		}
		p.Password = value
	}
	return nil
}

// construct implements constructor of PhoneConnectionClass.
func (p PhoneConnectionWebrtc) construct() PhoneConnectionClass { return &p }

// Ensuring interfaces in compile-time for PhoneConnectionWebrtc.
var (
	_ bin.Encoder = &PhoneConnectionWebrtc{}
	_ bin.Decoder = &PhoneConnectionWebrtc{}

	_ PhoneConnectionClass = &PhoneConnectionWebrtc{}
)

// PhoneConnectionClass represents PhoneConnection generic type.
//
// See https://core.telegram.org/type/PhoneConnection for reference.
//
// Example:
//  g, err := DecodePhoneConnection(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *PhoneConnection: // phoneConnection#9d4c17c0
//  case *PhoneConnectionWebrtc: // phoneConnectionWebrtc#635fe375
//  default: panic(v)
//  }
type PhoneConnectionClass interface {
	bin.Encoder
	bin.Decoder
	construct() PhoneConnectionClass
	fmt.Stringer
}

// DecodePhoneConnection implements binary de-serialization for PhoneConnectionClass.
func DecodePhoneConnection(buf *bin.Buffer) (PhoneConnectionClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case PhoneConnectionTypeID:
		// Decoding phoneConnection#9d4c17c0.
		v := PhoneConnection{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PhoneConnectionClass: %w", err)
		}
		return &v, nil
	case PhoneConnectionWebrtcTypeID:
		// Decoding phoneConnectionWebrtc#635fe375.
		v := PhoneConnectionWebrtc{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PhoneConnectionClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode PhoneConnectionClass: %w", bin.NewUnexpectedID(id))
	}
}

// PhoneConnection boxes the PhoneConnectionClass providing a helper.
type PhoneConnectionBox struct {
	PhoneConnection PhoneConnectionClass
}

// Decode implements bin.Decoder for PhoneConnectionBox.
func (b *PhoneConnectionBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode PhoneConnectionBox to nil")
	}
	v, err := DecodePhoneConnection(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.PhoneConnection = v
	return nil
}

// Encode implements bin.Encode for PhoneConnectionBox.
func (b *PhoneConnectionBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.PhoneConnection == nil {
		return fmt.Errorf("unable to encode PhoneConnectionClass as nil")
	}
	return b.PhoneConnection.Encode(buf)
}
