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

// AccountFinishTakeoutSessionRequest represents TL type `account.finishTakeoutSession#1d2652ee`.
// Finish account takeout session
//
// See https://core.telegram.org/method/account.finishTakeoutSession for reference.
type AccountFinishTakeoutSessionRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Data exported successfully
	Success bool
}

// AccountFinishTakeoutSessionRequestTypeID is TL type id of AccountFinishTakeoutSessionRequest.
const AccountFinishTakeoutSessionRequestTypeID = 0x1d2652ee

// String implements fmt.Stringer.
func (f *AccountFinishTakeoutSessionRequest) String() string {
	if f == nil {
		return "AccountFinishTakeoutSessionRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("AccountFinishTakeoutSessionRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(fmt.Sprint(f.Flags))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (f *AccountFinishTakeoutSessionRequest) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode account.finishTakeoutSession#1d2652ee as nil")
	}
	b.PutID(AccountFinishTakeoutSessionRequestTypeID)
	if err := f.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.finishTakeoutSession#1d2652ee: field flags: %w", err)
	}
	return nil
}

// SetSuccess sets value of Success conditional field.
func (f *AccountFinishTakeoutSessionRequest) SetSuccess(value bool) {
	if value {
		f.Flags.Set(0)
		f.Success = true
	} else {
		f.Flags.Unset(0)
		f.Success = false
	}
}

// Decode implements bin.Decoder.
func (f *AccountFinishTakeoutSessionRequest) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode account.finishTakeoutSession#1d2652ee to nil")
	}
	if err := b.ConsumeID(AccountFinishTakeoutSessionRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.finishTakeoutSession#1d2652ee: %w", err)
	}
	{
		if err := f.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.finishTakeoutSession#1d2652ee: field flags: %w", err)
		}
	}
	f.Success = f.Flags.Has(0)
	return nil
}

// Ensuring interfaces in compile-time for AccountFinishTakeoutSessionRequest.
var (
	_ bin.Encoder = &AccountFinishTakeoutSessionRequest{}
	_ bin.Decoder = &AccountFinishTakeoutSessionRequest{}
)

// AccountFinishTakeoutSession invokes method account.finishTakeoutSession#1d2652ee returning error if any.
// Finish account takeout session
//
// Possible errors:
//  403 TAKEOUT_REQUIRED: A takeout session has to be initialized, first
//
// See https://core.telegram.org/method/account.finishTakeoutSession for reference.
func (c *Client) AccountFinishTakeoutSession(ctx context.Context, request *AccountFinishTakeoutSessionRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
