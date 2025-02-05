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

// PhoneJoinGroupCallRequest represents TL type `phone.joinGroupCall#5f9c8e62`.
//
// See https://core.telegram.org/method/phone.joinGroupCall for reference.
type PhoneJoinGroupCallRequest struct {
	// Flags field of PhoneJoinGroupCallRequest.
	Flags bin.Fields
	// Muted field of PhoneJoinGroupCallRequest.
	Muted bool
	// Call field of PhoneJoinGroupCallRequest.
	Call InputGroupCall
	// Params field of PhoneJoinGroupCallRequest.
	Params DataJSON
}

// PhoneJoinGroupCallRequestTypeID is TL type id of PhoneJoinGroupCallRequest.
const PhoneJoinGroupCallRequestTypeID = 0x5f9c8e62

// String implements fmt.Stringer.
func (j *PhoneJoinGroupCallRequest) String() string {
	if j == nil {
		return "PhoneJoinGroupCallRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("PhoneJoinGroupCallRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(fmt.Sprint(j.Flags))
	sb.WriteString(",\n")
	sb.WriteString("\tCall: ")
	sb.WriteString(fmt.Sprint(j.Call))
	sb.WriteString(",\n")
	sb.WriteString("\tParams: ")
	sb.WriteString(fmt.Sprint(j.Params))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (j *PhoneJoinGroupCallRequest) Encode(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't encode phone.joinGroupCall#5f9c8e62 as nil")
	}
	b.PutID(PhoneJoinGroupCallRequestTypeID)
	if err := j.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.joinGroupCall#5f9c8e62: field flags: %w", err)
	}
	if err := j.Call.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.joinGroupCall#5f9c8e62: field call: %w", err)
	}
	if err := j.Params.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.joinGroupCall#5f9c8e62: field params: %w", err)
	}
	return nil
}

// SetMuted sets value of Muted conditional field.
func (j *PhoneJoinGroupCallRequest) SetMuted(value bool) {
	if value {
		j.Flags.Set(0)
		j.Muted = true
	} else {
		j.Flags.Unset(0)
		j.Muted = false
	}
}

// Decode implements bin.Decoder.
func (j *PhoneJoinGroupCallRequest) Decode(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't decode phone.joinGroupCall#5f9c8e62 to nil")
	}
	if err := b.ConsumeID(PhoneJoinGroupCallRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode phone.joinGroupCall#5f9c8e62: %w", err)
	}
	{
		if err := j.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.joinGroupCall#5f9c8e62: field flags: %w", err)
		}
	}
	j.Muted = j.Flags.Has(0)
	{
		if err := j.Call.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.joinGroupCall#5f9c8e62: field call: %w", err)
		}
	}
	{
		if err := j.Params.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.joinGroupCall#5f9c8e62: field params: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for PhoneJoinGroupCallRequest.
var (
	_ bin.Encoder = &PhoneJoinGroupCallRequest{}
	_ bin.Decoder = &PhoneJoinGroupCallRequest{}
)

// PhoneJoinGroupCall invokes method phone.joinGroupCall#5f9c8e62 returning error if any.
//
// See https://core.telegram.org/method/phone.joinGroupCall for reference.
func (c *Client) PhoneJoinGroupCall(ctx context.Context, request *PhoneJoinGroupCallRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
