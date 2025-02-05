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

// AccountUpdateNotifySettingsRequest represents TL type `account.updateNotifySettings#84be5b93`.
// Edits notification settings from a given user/group, from all users/all groups.
//
// See https://core.telegram.org/method/account.updateNotifySettings for reference.
type AccountUpdateNotifySettingsRequest struct {
	// Notification source
	Peer InputNotifyPeerClass
	// Notification settings
	Settings InputPeerNotifySettings
}

// AccountUpdateNotifySettingsRequestTypeID is TL type id of AccountUpdateNotifySettingsRequest.
const AccountUpdateNotifySettingsRequestTypeID = 0x84be5b93

// String implements fmt.Stringer.
func (u *AccountUpdateNotifySettingsRequest) String() string {
	if u == nil {
		return "AccountUpdateNotifySettingsRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("AccountUpdateNotifySettingsRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tPeer: ")
	sb.WriteString(fmt.Sprint(u.Peer))
	sb.WriteString(",\n")
	sb.WriteString("\tSettings: ")
	sb.WriteString(fmt.Sprint(u.Settings))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (u *AccountUpdateNotifySettingsRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode account.updateNotifySettings#84be5b93 as nil")
	}
	b.PutID(AccountUpdateNotifySettingsRequestTypeID)
	if u.Peer == nil {
		return fmt.Errorf("unable to encode account.updateNotifySettings#84be5b93: field peer is nil")
	}
	if err := u.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.updateNotifySettings#84be5b93: field peer: %w", err)
	}
	if err := u.Settings.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.updateNotifySettings#84be5b93: field settings: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (u *AccountUpdateNotifySettingsRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode account.updateNotifySettings#84be5b93 to nil")
	}
	if err := b.ConsumeID(AccountUpdateNotifySettingsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.updateNotifySettings#84be5b93: %w", err)
	}
	{
		value, err := DecodeInputNotifyPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.updateNotifySettings#84be5b93: field peer: %w", err)
		}
		u.Peer = value
	}
	{
		if err := u.Settings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.updateNotifySettings#84be5b93: field settings: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountUpdateNotifySettingsRequest.
var (
	_ bin.Encoder = &AccountUpdateNotifySettingsRequest{}
	_ bin.Decoder = &AccountUpdateNotifySettingsRequest{}
)

// AccountUpdateNotifySettings invokes method account.updateNotifySettings#84be5b93 returning error if any.
// Edits notification settings from a given user/group, from all users/all groups.
//
// Possible errors:
//  400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup
//  400 MSG_ID_INVALID: Invalid message ID provided
//  400 PEER_ID_INVALID: The provided peer id is invalid
//  400 SETTINGS_INVALID: Invalid settings were provided
//
// See https://core.telegram.org/method/account.updateNotifySettings for reference.
func (c *Client) AccountUpdateNotifySettings(ctx context.Context, request *AccountUpdateNotifySettingsRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
