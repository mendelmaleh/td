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

// HelpGetUserInfoRequest represents TL type `help.getUserInfo#38a08d3`.
// Internal use
//
// See https://core.telegram.org/method/help.getUserInfo for reference.
type HelpGetUserInfoRequest struct {
	// User ID
	UserID InputUserClass
}

// HelpGetUserInfoRequestTypeID is TL type id of HelpGetUserInfoRequest.
const HelpGetUserInfoRequestTypeID = 0x38a08d3

// String implements fmt.Stringer.
func (g *HelpGetUserInfoRequest) String() string {
	if g == nil {
		return "HelpGetUserInfoRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("HelpGetUserInfoRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tUserID: ")
	sb.WriteString(fmt.Sprint(g.UserID))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (g *HelpGetUserInfoRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode help.getUserInfo#38a08d3 as nil")
	}
	b.PutID(HelpGetUserInfoRequestTypeID)
	if g.UserID == nil {
		return fmt.Errorf("unable to encode help.getUserInfo#38a08d3: field user_id is nil")
	}
	if err := g.UserID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode help.getUserInfo#38a08d3: field user_id: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *HelpGetUserInfoRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode help.getUserInfo#38a08d3 to nil")
	}
	if err := b.ConsumeID(HelpGetUserInfoRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode help.getUserInfo#38a08d3: %w", err)
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode help.getUserInfo#38a08d3: field user_id: %w", err)
		}
		g.UserID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for HelpGetUserInfoRequest.
var (
	_ bin.Encoder = &HelpGetUserInfoRequest{}
	_ bin.Decoder = &HelpGetUserInfoRequest{}
)

// HelpGetUserInfo invokes method help.getUserInfo#38a08d3 returning error if any.
// Internal use
//
// Possible errors:
//  403 USER_INVALID: Invalid user provided
//
// See https://core.telegram.org/method/help.getUserInfo for reference.
func (c *Client) HelpGetUserInfo(ctx context.Context, userid InputUserClass) (HelpUserInfoClass, error) {
	var result HelpUserInfoBox

	request := &HelpGetUserInfoRequest{
		UserID: userid,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.UserInfo, nil
}
