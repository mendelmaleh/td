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

// PhoneAcceptCallRequest represents TL type `phone.acceptCall#3bd2b4a0`.
// Accept incoming call
//
// See https://core.telegram.org/method/phone.acceptCall for reference.
type PhoneAcceptCallRequest struct {
	// The call to accept
	Peer InputPhoneCall
	// Parameter for E2E encryption key exchange »
	GB []byte
	// Phone call settings
	Protocol PhoneCallProtocol
}

// PhoneAcceptCallRequestTypeID is TL type id of PhoneAcceptCallRequest.
const PhoneAcceptCallRequestTypeID = 0x3bd2b4a0

// Encode implements bin.Encoder.
func (a *PhoneAcceptCallRequest) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode phone.acceptCall#3bd2b4a0 as nil")
	}
	b.PutID(PhoneAcceptCallRequestTypeID)
	if err := a.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.acceptCall#3bd2b4a0: field peer: %w", err)
	}
	b.PutBytes(a.GB)
	if err := a.Protocol.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.acceptCall#3bd2b4a0: field protocol: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *PhoneAcceptCallRequest) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode phone.acceptCall#3bd2b4a0 to nil")
	}
	if err := b.ConsumeID(PhoneAcceptCallRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode phone.acceptCall#3bd2b4a0: %w", err)
	}
	{
		if err := a.Peer.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.acceptCall#3bd2b4a0: field peer: %w", err)
		}
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode phone.acceptCall#3bd2b4a0: field g_b: %w", err)
		}
		a.GB = value
	}
	{
		if err := a.Protocol.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.acceptCall#3bd2b4a0: field protocol: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for PhoneAcceptCallRequest.
var (
	_ bin.Encoder = &PhoneAcceptCallRequest{}
	_ bin.Decoder = &PhoneAcceptCallRequest{}
)

// PhoneAcceptCall invokes method phone.acceptCall#3bd2b4a0 returning error if any.
// Accept incoming call
//
// See https://core.telegram.org/method/phone.acceptCall for reference.
func (c *Client) PhoneAcceptCall(ctx context.Context, request *PhoneAcceptCallRequest) (*PhonePhoneCall, error) {
	var result PhonePhoneCall

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
