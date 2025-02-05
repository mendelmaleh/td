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

// PhoneEditGroupCallMemberRequest represents TL type `phone.editGroupCallMember#63146ae4`.
//
// See https://core.telegram.org/method/phone.editGroupCallMember for reference.
type PhoneEditGroupCallMemberRequest struct {
	// Flags field of PhoneEditGroupCallMemberRequest.
	Flags bin.Fields
	// Muted field of PhoneEditGroupCallMemberRequest.
	Muted bool
	// Call field of PhoneEditGroupCallMemberRequest.
	Call InputGroupCall
	// UserID field of PhoneEditGroupCallMemberRequest.
	UserID InputUserClass
}

// PhoneEditGroupCallMemberRequestTypeID is TL type id of PhoneEditGroupCallMemberRequest.
const PhoneEditGroupCallMemberRequestTypeID = 0x63146ae4

// String implements fmt.Stringer.
func (e *PhoneEditGroupCallMemberRequest) String() string {
	if e == nil {
		return "PhoneEditGroupCallMemberRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("PhoneEditGroupCallMemberRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(fmt.Sprint(e.Flags))
	sb.WriteString(",\n")
	sb.WriteString("\tCall: ")
	sb.WriteString(fmt.Sprint(e.Call))
	sb.WriteString(",\n")
	sb.WriteString("\tUserID: ")
	sb.WriteString(fmt.Sprint(e.UserID))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (e *PhoneEditGroupCallMemberRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode phone.editGroupCallMember#63146ae4 as nil")
	}
	b.PutID(PhoneEditGroupCallMemberRequestTypeID)
	if err := e.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.editGroupCallMember#63146ae4: field flags: %w", err)
	}
	if err := e.Call.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.editGroupCallMember#63146ae4: field call: %w", err)
	}
	if e.UserID == nil {
		return fmt.Errorf("unable to encode phone.editGroupCallMember#63146ae4: field user_id is nil")
	}
	if err := e.UserID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.editGroupCallMember#63146ae4: field user_id: %w", err)
	}
	return nil
}

// SetMuted sets value of Muted conditional field.
func (e *PhoneEditGroupCallMemberRequest) SetMuted(value bool) {
	if value {
		e.Flags.Set(0)
		e.Muted = true
	} else {
		e.Flags.Unset(0)
		e.Muted = false
	}
}

// Decode implements bin.Decoder.
func (e *PhoneEditGroupCallMemberRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode phone.editGroupCallMember#63146ae4 to nil")
	}
	if err := b.ConsumeID(PhoneEditGroupCallMemberRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode phone.editGroupCallMember#63146ae4: %w", err)
	}
	{
		if err := e.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.editGroupCallMember#63146ae4: field flags: %w", err)
		}
	}
	e.Muted = e.Flags.Has(0)
	{
		if err := e.Call.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.editGroupCallMember#63146ae4: field call: %w", err)
		}
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode phone.editGroupCallMember#63146ae4: field user_id: %w", err)
		}
		e.UserID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for PhoneEditGroupCallMemberRequest.
var (
	_ bin.Encoder = &PhoneEditGroupCallMemberRequest{}
	_ bin.Decoder = &PhoneEditGroupCallMemberRequest{}
)

// PhoneEditGroupCallMember invokes method phone.editGroupCallMember#63146ae4 returning error if any.
//
// See https://core.telegram.org/method/phone.editGroupCallMember for reference.
func (c *Client) PhoneEditGroupCallMember(ctx context.Context, request *PhoneEditGroupCallMemberRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
