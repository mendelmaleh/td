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

// ChatAdminRights represents TL type `chatAdminRights#5fb224d5`.
// Represents the rights of an admin in a channel/supergroup¹.
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// See https://core.telegram.org/constructor/chatAdminRights for reference.
type ChatAdminRights struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// If set, allows the admin to modify the description of the channel/supergroup¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	ChangeInfo bool
	// If set, allows the admin to post messages in the channel¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	PostMessages bool
	// If set, allows the admin to also edit messages from other admins in the channel¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	EditMessages bool
	// If set, allows the admin to also delete messages from other admins in the channel¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	DeleteMessages bool
	// If set, allows the admin to ban users from the channel/supergroup¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	BanUsers bool
	// If set, allows the admin to invite users in the channel/supergroup¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	InviteUsers bool
	// If set, allows the admin to pin messages in the channel/supergroup¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	PinMessages bool
	// If set, allows the admin to add other admins with the same (or more limited) permissions in the channel/supergroup¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	AddAdmins bool
	// Whether this admin is anonymous
	Anonymous bool
	// ManageCall field of ChatAdminRights.
	ManageCall bool
}

// ChatAdminRightsTypeID is TL type id of ChatAdminRights.
const ChatAdminRightsTypeID = 0x5fb224d5

// String implements fmt.Stringer.
func (c *ChatAdminRights) String() string {
	if c == nil {
		return "ChatAdminRights(nil)"
	}
	var sb strings.Builder
	sb.WriteString("ChatAdminRights")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(fmt.Sprint(c.Flags))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (c *ChatAdminRights) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatAdminRights#5fb224d5 as nil")
	}
	b.PutID(ChatAdminRightsTypeID)
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chatAdminRights#5fb224d5: field flags: %w", err)
	}
	return nil
}

// SetChangeInfo sets value of ChangeInfo conditional field.
func (c *ChatAdminRights) SetChangeInfo(value bool) {
	if value {
		c.Flags.Set(0)
		c.ChangeInfo = true
	} else {
		c.Flags.Unset(0)
		c.ChangeInfo = false
	}
}

// SetPostMessages sets value of PostMessages conditional field.
func (c *ChatAdminRights) SetPostMessages(value bool) {
	if value {
		c.Flags.Set(1)
		c.PostMessages = true
	} else {
		c.Flags.Unset(1)
		c.PostMessages = false
	}
}

// SetEditMessages sets value of EditMessages conditional field.
func (c *ChatAdminRights) SetEditMessages(value bool) {
	if value {
		c.Flags.Set(2)
		c.EditMessages = true
	} else {
		c.Flags.Unset(2)
		c.EditMessages = false
	}
}

// SetDeleteMessages sets value of DeleteMessages conditional field.
func (c *ChatAdminRights) SetDeleteMessages(value bool) {
	if value {
		c.Flags.Set(3)
		c.DeleteMessages = true
	} else {
		c.Flags.Unset(3)
		c.DeleteMessages = false
	}
}

// SetBanUsers sets value of BanUsers conditional field.
func (c *ChatAdminRights) SetBanUsers(value bool) {
	if value {
		c.Flags.Set(4)
		c.BanUsers = true
	} else {
		c.Flags.Unset(4)
		c.BanUsers = false
	}
}

// SetInviteUsers sets value of InviteUsers conditional field.
func (c *ChatAdminRights) SetInviteUsers(value bool) {
	if value {
		c.Flags.Set(5)
		c.InviteUsers = true
	} else {
		c.Flags.Unset(5)
		c.InviteUsers = false
	}
}

// SetPinMessages sets value of PinMessages conditional field.
func (c *ChatAdminRights) SetPinMessages(value bool) {
	if value {
		c.Flags.Set(7)
		c.PinMessages = true
	} else {
		c.Flags.Unset(7)
		c.PinMessages = false
	}
}

// SetAddAdmins sets value of AddAdmins conditional field.
func (c *ChatAdminRights) SetAddAdmins(value bool) {
	if value {
		c.Flags.Set(9)
		c.AddAdmins = true
	} else {
		c.Flags.Unset(9)
		c.AddAdmins = false
	}
}

// SetAnonymous sets value of Anonymous conditional field.
func (c *ChatAdminRights) SetAnonymous(value bool) {
	if value {
		c.Flags.Set(10)
		c.Anonymous = true
	} else {
		c.Flags.Unset(10)
		c.Anonymous = false
	}
}

// SetManageCall sets value of ManageCall conditional field.
func (c *ChatAdminRights) SetManageCall(value bool) {
	if value {
		c.Flags.Set(11)
		c.ManageCall = true
	} else {
		c.Flags.Unset(11)
		c.ManageCall = false
	}
}

// Decode implements bin.Decoder.
func (c *ChatAdminRights) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatAdminRights#5fb224d5 to nil")
	}
	if err := b.ConsumeID(ChatAdminRightsTypeID); err != nil {
		return fmt.Errorf("unable to decode chatAdminRights#5fb224d5: %w", err)
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode chatAdminRights#5fb224d5: field flags: %w", err)
		}
	}
	c.ChangeInfo = c.Flags.Has(0)
	c.PostMessages = c.Flags.Has(1)
	c.EditMessages = c.Flags.Has(2)
	c.DeleteMessages = c.Flags.Has(3)
	c.BanUsers = c.Flags.Has(4)
	c.InviteUsers = c.Flags.Has(5)
	c.PinMessages = c.Flags.Has(7)
	c.AddAdmins = c.Flags.Has(9)
	c.Anonymous = c.Flags.Has(10)
	c.ManageCall = c.Flags.Has(11)
	return nil
}

// Ensuring interfaces in compile-time for ChatAdminRights.
var (
	_ bin.Encoder = &ChatAdminRights{}
	_ bin.Decoder = &ChatAdminRights{}
)
