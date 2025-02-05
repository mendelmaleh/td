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

// ChannelsUpdateUsernameRequest represents TL type `channels.updateUsername#3514b3de`.
// Change the username of a supergroup/channel
//
// See https://core.telegram.org/method/channels.updateUsername for reference.
type ChannelsUpdateUsernameRequest struct {
	// Channel
	Channel InputChannelClass
	// New username
	Username string
}

// ChannelsUpdateUsernameRequestTypeID is TL type id of ChannelsUpdateUsernameRequest.
const ChannelsUpdateUsernameRequestTypeID = 0x3514b3de

// String implements fmt.Stringer.
func (u *ChannelsUpdateUsernameRequest) String() string {
	if u == nil {
		return "ChannelsUpdateUsernameRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("ChannelsUpdateUsernameRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tChannel: ")
	sb.WriteString(fmt.Sprint(u.Channel))
	sb.WriteString(",\n")
	sb.WriteString("\tUsername: ")
	sb.WriteString(fmt.Sprint(u.Username))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (u *ChannelsUpdateUsernameRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode channels.updateUsername#3514b3de as nil")
	}
	b.PutID(ChannelsUpdateUsernameRequestTypeID)
	if u.Channel == nil {
		return fmt.Errorf("unable to encode channels.updateUsername#3514b3de: field channel is nil")
	}
	if err := u.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.updateUsername#3514b3de: field channel: %w", err)
	}
	b.PutString(u.Username)
	return nil
}

// Decode implements bin.Decoder.
func (u *ChannelsUpdateUsernameRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode channels.updateUsername#3514b3de to nil")
	}
	if err := b.ConsumeID(ChannelsUpdateUsernameRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.updateUsername#3514b3de: %w", err)
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.updateUsername#3514b3de: field channel: %w", err)
		}
		u.Channel = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode channels.updateUsername#3514b3de: field username: %w", err)
		}
		u.Username = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelsUpdateUsernameRequest.
var (
	_ bin.Encoder = &ChannelsUpdateUsernameRequest{}
	_ bin.Decoder = &ChannelsUpdateUsernameRequest{}
)

// ChannelsUpdateUsername invokes method channels.updateUsername#3514b3de returning error if any.
// Change the username of a supergroup/channel
//
// Possible errors:
//  400 CHANNELS_ADMIN_PUBLIC_TOO_MUCH: You're admin of too many public channels, make some channels private to change the username of this channel
//  400 CHANNEL_INVALID: The provided channel is invalid
//  400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup
//  400 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this
//  400 CHAT_NOT_MODIFIED: The pinned message wasn't modified
//  403 CHAT_WRITE_FORBIDDEN: You can't write in this chat
//  400 USERNAME_INVALID: The provided username is not valid
//  400 USERNAME_NOT_MODIFIED: The username was not modified
//  400 USERNAME_OCCUPIED: The provided username is already occupied
//
// See https://core.telegram.org/method/channels.updateUsername for reference.
func (c *Client) ChannelsUpdateUsername(ctx context.Context, request *ChannelsUpdateUsernameRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
