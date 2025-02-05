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

// ChannelsEditAdminRequest represents TL type `channels.editAdmin#d33c8902`.
// Modify the admin rights of a user in a supergroup/channel¹.
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// See https://core.telegram.org/method/channels.editAdmin for reference.
type ChannelsEditAdminRequest struct {
	// The supergroup/channel¹.
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	Channel InputChannelClass
	// The ID of the user whose admin rights should be modified
	UserID InputUserClass
	// The admin rights
	AdminRights ChatAdminRights
	// Indicates the role (rank) of the admin in the group: just an arbitrary string
	Rank string
}

// ChannelsEditAdminRequestTypeID is TL type id of ChannelsEditAdminRequest.
const ChannelsEditAdminRequestTypeID = 0xd33c8902

// String implements fmt.Stringer.
func (e *ChannelsEditAdminRequest) String() string {
	if e == nil {
		return "ChannelsEditAdminRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("ChannelsEditAdminRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tChannel: ")
	sb.WriteString(fmt.Sprint(e.Channel))
	sb.WriteString(",\n")
	sb.WriteString("\tUserID: ")
	sb.WriteString(fmt.Sprint(e.UserID))
	sb.WriteString(",\n")
	sb.WriteString("\tAdminRights: ")
	sb.WriteString(fmt.Sprint(e.AdminRights))
	sb.WriteString(",\n")
	sb.WriteString("\tRank: ")
	sb.WriteString(fmt.Sprint(e.Rank))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (e *ChannelsEditAdminRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode channels.editAdmin#d33c8902 as nil")
	}
	b.PutID(ChannelsEditAdminRequestTypeID)
	if e.Channel == nil {
		return fmt.Errorf("unable to encode channels.editAdmin#d33c8902: field channel is nil")
	}
	if err := e.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.editAdmin#d33c8902: field channel: %w", err)
	}
	if e.UserID == nil {
		return fmt.Errorf("unable to encode channels.editAdmin#d33c8902: field user_id is nil")
	}
	if err := e.UserID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.editAdmin#d33c8902: field user_id: %w", err)
	}
	if err := e.AdminRights.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.editAdmin#d33c8902: field admin_rights: %w", err)
	}
	b.PutString(e.Rank)
	return nil
}

// Decode implements bin.Decoder.
func (e *ChannelsEditAdminRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode channels.editAdmin#d33c8902 to nil")
	}
	if err := b.ConsumeID(ChannelsEditAdminRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.editAdmin#d33c8902: %w", err)
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.editAdmin#d33c8902: field channel: %w", err)
		}
		e.Channel = value
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.editAdmin#d33c8902: field user_id: %w", err)
		}
		e.UserID = value
	}
	{
		if err := e.AdminRights.Decode(b); err != nil {
			return fmt.Errorf("unable to decode channels.editAdmin#d33c8902: field admin_rights: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode channels.editAdmin#d33c8902: field rank: %w", err)
		}
		e.Rank = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelsEditAdminRequest.
var (
	_ bin.Encoder = &ChannelsEditAdminRequest{}
	_ bin.Decoder = &ChannelsEditAdminRequest{}
)

// ChannelsEditAdmin invokes method channels.editAdmin#d33c8902 returning error if any.
// Modify the admin rights of a user in a supergroup/channel¹.
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// Possible errors:
//  400 ADMINS_TOO_MUCH: There are too many admins
//  400 BOTS_TOO_MUCH: There are too many bots in this chat/channel
//  400 BOT_CHANNELS_NA: Bots can't edit admin privileges
//  400 BOT_GROUPS_BLOCKED: This bot can't be added to groups
//  400 CHANNEL_INVALID: The provided channel is invalid
//  400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup
//  403 CHAT_ADMIN_INVITE_REQUIRED: You do not have the rights to do this
//  400 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this
//  403 CHAT_WRITE_FORBIDDEN: You can't write in this chat
//  406 FRESH_CHANGE_ADMINS_FORBIDDEN: You were just elected admin, you can't add or modify other admins yet
//  400 INPUT_USER_DEACTIVATED: The specified user was deleted
//  400 PEER_ID_INVALID: The provided peer id is invalid
//  403 RIGHT_FORBIDDEN: Your admin rights do not allow you to do this
//  400 USERS_TOO_MUCH: The maximum number of users has been exceeded (to create a chat, for example)
//  400 USER_BLOCKED: User blocked
//  403 USER_CHANNELS_TOO_MUCH: One of the users you tried to add is already in too many channels/supergroups
//  400 USER_CREATOR: You can't leave this channel, because you're its creator
//  400 USER_ID_INVALID: The provided user ID is invalid
//  400 USER_NOT_MUTUAL_CONTACT: The provided user is not a mutual contact
//  403 USER_PRIVACY_RESTRICTED: The user's privacy settings do not allow you to do this
//  403 USER_RESTRICTED: You're spamreported, you can't create channels or chats.
//
// See https://core.telegram.org/method/channels.editAdmin for reference.
// Can be used by bots.
func (c *Client) ChannelsEditAdmin(ctx context.Context, request *ChannelsEditAdminRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
