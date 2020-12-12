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

// ContactsBlockFromRepliesRequest represents TL type `contacts.blockFromReplies#29a8962c`.
// Stop getting notifications about thread replies of a certain user in @replies
//
// See https://core.telegram.org/method/contacts.blockFromReplies for reference.
type ContactsBlockFromRepliesRequest struct {
	// Flags, see TL conditional fields
	Flags bin.Fields
	// Whether to delete the specified message as well
	DeleteMessage bool
	// Whether to delete all @replies messages from this user as well
	DeleteHistory bool
	// Whether to also report this user for spam
	ReportSpam bool
	// ID of the message in the @replies chat
	MsgID int
}

// ContactsBlockFromRepliesRequestTypeID is TL type id of ContactsBlockFromRepliesRequest.
const ContactsBlockFromRepliesRequestTypeID = 0x29a8962c

// Encode implements bin.Encoder.
func (b *ContactsBlockFromRepliesRequest) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode contacts.blockFromReplies#29a8962c as nil")
	}
	buf.PutID(ContactsBlockFromRepliesRequestTypeID)
	if err := b.Flags.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode contacts.blockFromReplies#29a8962c: field flags: %w", err)
	}
	buf.PutInt(b.MsgID)
	return nil
}

// SetDeleteMessage sets value of DeleteMessage conditional field.
func (b *ContactsBlockFromRepliesRequest) SetDeleteMessage(value bool) {
	if value {
		b.Flags.Set(0)
	} else {
		b.Flags.Unset(0)
	}
}

// SetDeleteHistory sets value of DeleteHistory conditional field.
func (b *ContactsBlockFromRepliesRequest) SetDeleteHistory(value bool) {
	if value {
		b.Flags.Set(1)
	} else {
		b.Flags.Unset(1)
	}
}

// SetReportSpam sets value of ReportSpam conditional field.
func (b *ContactsBlockFromRepliesRequest) SetReportSpam(value bool) {
	if value {
		b.Flags.Set(2)
	} else {
		b.Flags.Unset(2)
	}
}

// Decode implements bin.Decoder.
func (b *ContactsBlockFromRepliesRequest) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode contacts.blockFromReplies#29a8962c to nil")
	}
	if err := buf.ConsumeID(ContactsBlockFromRepliesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.blockFromReplies#29a8962c: %w", err)
	}
	{
		if err := b.Flags.Decode(buf); err != nil {
			return fmt.Errorf("unable to decode contacts.blockFromReplies#29a8962c: field flags: %w", err)
		}
	}
	b.DeleteMessage = b.Flags.Has(0)
	b.DeleteHistory = b.Flags.Has(1)
	b.ReportSpam = b.Flags.Has(2)
	{
		value, err := buf.Int()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.blockFromReplies#29a8962c: field msg_id: %w", err)
		}
		b.MsgID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ContactsBlockFromRepliesRequest.
var (
	_ bin.Encoder = &ContactsBlockFromRepliesRequest{}
	_ bin.Decoder = &ContactsBlockFromRepliesRequest{}
)

// ContactsBlockFromReplies invokes method contacts.blockFromReplies#29a8962c returning error if any.
// Stop getting notifications about thread replies of a certain user in @replies
//
// See https://core.telegram.org/method/contacts.blockFromReplies for reference.
func (c *Client) ContactsBlockFromReplies(ctx context.Context, request *ContactsBlockFromRepliesRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
