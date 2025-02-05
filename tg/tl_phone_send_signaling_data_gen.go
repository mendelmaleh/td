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

// PhoneSendSignalingDataRequest represents TL type `phone.sendSignalingData#ff7a9383`.
// Send VoIP signaling data
//
// See https://core.telegram.org/method/phone.sendSignalingData for reference.
type PhoneSendSignalingDataRequest struct {
	// Phone call
	Peer InputPhoneCall
	// Signaling payload
	Data []byte
}

// PhoneSendSignalingDataRequestTypeID is TL type id of PhoneSendSignalingDataRequest.
const PhoneSendSignalingDataRequestTypeID = 0xff7a9383

// String implements fmt.Stringer.
func (s *PhoneSendSignalingDataRequest) String() string {
	if s == nil {
		return "PhoneSendSignalingDataRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("PhoneSendSignalingDataRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tPeer: ")
	sb.WriteString(fmt.Sprint(s.Peer))
	sb.WriteString(",\n")
	sb.WriteString("\tData: ")
	sb.WriteString(fmt.Sprint(s.Data))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (s *PhoneSendSignalingDataRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode phone.sendSignalingData#ff7a9383 as nil")
	}
	b.PutID(PhoneSendSignalingDataRequestTypeID)
	if err := s.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.sendSignalingData#ff7a9383: field peer: %w", err)
	}
	b.PutBytes(s.Data)
	return nil
}

// Decode implements bin.Decoder.
func (s *PhoneSendSignalingDataRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode phone.sendSignalingData#ff7a9383 to nil")
	}
	if err := b.ConsumeID(PhoneSendSignalingDataRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode phone.sendSignalingData#ff7a9383: %w", err)
	}
	{
		if err := s.Peer.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.sendSignalingData#ff7a9383: field peer: %w", err)
		}
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode phone.sendSignalingData#ff7a9383: field data: %w", err)
		}
		s.Data = value
	}
	return nil
}

// Ensuring interfaces in compile-time for PhoneSendSignalingDataRequest.
var (
	_ bin.Encoder = &PhoneSendSignalingDataRequest{}
	_ bin.Decoder = &PhoneSendSignalingDataRequest{}
)

// PhoneSendSignalingData invokes method phone.sendSignalingData#ff7a9383 returning error if any.
// Send VoIP signaling data
//
// See https://core.telegram.org/method/phone.sendSignalingData for reference.
// Can be used by bots.
func (c *Client) PhoneSendSignalingData(ctx context.Context, request *PhoneSendSignalingDataRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
