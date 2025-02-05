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

// MessagesEditInlineBotMessageRequest represents TL type `messages.editInlineBotMessage#83557dba`.
// Edit an inline bot message
//
// See https://core.telegram.org/method/messages.editInlineBotMessage for reference.
type MessagesEditInlineBotMessageRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Disable webpage preview
	NoWebpage bool
	// Sent inline message ID
	ID InputBotInlineMessageID
	// Message
	//
	// Use SetMessage and GetMessage helpers.
	Message string
	// Media
	//
	// Use SetMedia and GetMedia helpers.
	Media InputMediaClass
	// Reply markup for inline keyboards
	//
	// Use SetReplyMarkup and GetReplyMarkup helpers.
	ReplyMarkup ReplyMarkupClass
	// Message entities for styled text¹
	//
	// Links:
	//  1) https://core.telegram.org/api/entities
	//
	// Use SetEntities and GetEntities helpers.
	Entities []MessageEntityClass
}

// MessagesEditInlineBotMessageRequestTypeID is TL type id of MessagesEditInlineBotMessageRequest.
const MessagesEditInlineBotMessageRequestTypeID = 0x83557dba

// String implements fmt.Stringer.
func (e *MessagesEditInlineBotMessageRequest) String() string {
	if e == nil {
		return "MessagesEditInlineBotMessageRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesEditInlineBotMessageRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(fmt.Sprint(e.Flags))
	sb.WriteString(",\n")
	sb.WriteString("\tID: ")
	sb.WriteString(fmt.Sprint(e.ID))
	sb.WriteString(",\n")
	if e.Flags.Has(11) {
		sb.WriteString("\tMessage: ")
		sb.WriteString(fmt.Sprint(e.Message))
		sb.WriteString(",\n")
	}
	if e.Flags.Has(14) {
		sb.WriteString("\tMedia: ")
		sb.WriteString(fmt.Sprint(e.Media))
		sb.WriteString(",\n")
	}
	if e.Flags.Has(2) {
		sb.WriteString("\tReplyMarkup: ")
		sb.WriteString(fmt.Sprint(e.ReplyMarkup))
		sb.WriteString(",\n")
	}
	if e.Flags.Has(3) {
		sb.WriteByte('[')
		for _, v := range e.Entities {
			sb.WriteString(fmt.Sprint(v))
		}
		sb.WriteByte(']')
	}
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (e *MessagesEditInlineBotMessageRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode messages.editInlineBotMessage#83557dba as nil")
	}
	b.PutID(MessagesEditInlineBotMessageRequestTypeID)
	if err := e.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.editInlineBotMessage#83557dba: field flags: %w", err)
	}
	if err := e.ID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.editInlineBotMessage#83557dba: field id: %w", err)
	}
	if e.Flags.Has(11) {
		b.PutString(e.Message)
	}
	if e.Flags.Has(14) {
		if e.Media == nil {
			return fmt.Errorf("unable to encode messages.editInlineBotMessage#83557dba: field media is nil")
		}
		if err := e.Media.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.editInlineBotMessage#83557dba: field media: %w", err)
		}
	}
	if e.Flags.Has(2) {
		if e.ReplyMarkup == nil {
			return fmt.Errorf("unable to encode messages.editInlineBotMessage#83557dba: field reply_markup is nil")
		}
		if err := e.ReplyMarkup.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.editInlineBotMessage#83557dba: field reply_markup: %w", err)
		}
	}
	if e.Flags.Has(3) {
		b.PutVectorHeader(len(e.Entities))
		for idx, v := range e.Entities {
			if v == nil {
				return fmt.Errorf("unable to encode messages.editInlineBotMessage#83557dba: field entities element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode messages.editInlineBotMessage#83557dba: field entities element with index %d: %w", idx, err)
			}
		}
	}
	return nil
}

// SetNoWebpage sets value of NoWebpage conditional field.
func (e *MessagesEditInlineBotMessageRequest) SetNoWebpage(value bool) {
	if value {
		e.Flags.Set(1)
		e.NoWebpage = true
	} else {
		e.Flags.Unset(1)
		e.NoWebpage = false
	}
}

// SetMessage sets value of Message conditional field.
func (e *MessagesEditInlineBotMessageRequest) SetMessage(value string) {
	e.Flags.Set(11)
	e.Message = value
}

// GetMessage returns value of Message conditional field and
// boolean which is true if field was set.
func (e *MessagesEditInlineBotMessageRequest) GetMessage() (value string, ok bool) {
	if !e.Flags.Has(11) {
		return value, false
	}
	return e.Message, true
}

// SetMedia sets value of Media conditional field.
func (e *MessagesEditInlineBotMessageRequest) SetMedia(value InputMediaClass) {
	e.Flags.Set(14)
	e.Media = value
}

// GetMedia returns value of Media conditional field and
// boolean which is true if field was set.
func (e *MessagesEditInlineBotMessageRequest) GetMedia() (value InputMediaClass, ok bool) {
	if !e.Flags.Has(14) {
		return value, false
	}
	return e.Media, true
}

// SetReplyMarkup sets value of ReplyMarkup conditional field.
func (e *MessagesEditInlineBotMessageRequest) SetReplyMarkup(value ReplyMarkupClass) {
	e.Flags.Set(2)
	e.ReplyMarkup = value
}

// GetReplyMarkup returns value of ReplyMarkup conditional field and
// boolean which is true if field was set.
func (e *MessagesEditInlineBotMessageRequest) GetReplyMarkup() (value ReplyMarkupClass, ok bool) {
	if !e.Flags.Has(2) {
		return value, false
	}
	return e.ReplyMarkup, true
}

// SetEntities sets value of Entities conditional field.
func (e *MessagesEditInlineBotMessageRequest) SetEntities(value []MessageEntityClass) {
	e.Flags.Set(3)
	e.Entities = value
}

// GetEntities returns value of Entities conditional field and
// boolean which is true if field was set.
func (e *MessagesEditInlineBotMessageRequest) GetEntities() (value []MessageEntityClass, ok bool) {
	if !e.Flags.Has(3) {
		return value, false
	}
	return e.Entities, true
}

// Decode implements bin.Decoder.
func (e *MessagesEditInlineBotMessageRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode messages.editInlineBotMessage#83557dba to nil")
	}
	if err := b.ConsumeID(MessagesEditInlineBotMessageRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.editInlineBotMessage#83557dba: %w", err)
	}
	{
		if err := e.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.editInlineBotMessage#83557dba: field flags: %w", err)
		}
	}
	e.NoWebpage = e.Flags.Has(1)
	{
		if err := e.ID.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.editInlineBotMessage#83557dba: field id: %w", err)
		}
	}
	if e.Flags.Has(11) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.editInlineBotMessage#83557dba: field message: %w", err)
		}
		e.Message = value
	}
	if e.Flags.Has(14) {
		value, err := DecodeInputMedia(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.editInlineBotMessage#83557dba: field media: %w", err)
		}
		e.Media = value
	}
	if e.Flags.Has(2) {
		value, err := DecodeReplyMarkup(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.editInlineBotMessage#83557dba: field reply_markup: %w", err)
		}
		e.ReplyMarkup = value
	}
	if e.Flags.Has(3) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.editInlineBotMessage#83557dba: field entities: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessageEntity(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.editInlineBotMessage#83557dba: field entities: %w", err)
			}
			e.Entities = append(e.Entities, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesEditInlineBotMessageRequest.
var (
	_ bin.Encoder = &MessagesEditInlineBotMessageRequest{}
	_ bin.Decoder = &MessagesEditInlineBotMessageRequest{}
)

// MessagesEditInlineBotMessage invokes method messages.editInlineBotMessage#83557dba returning error if any.
// Edit an inline bot message
//
// Possible errors:
//  400 MESSAGE_ID_INVALID: The provided message id is invalid
//  400 MESSAGE_NOT_MODIFIED: The message text has not changed
//
// See https://core.telegram.org/method/messages.editInlineBotMessage for reference.
// Can be used by bots.
func (c *Client) MessagesEditInlineBotMessage(ctx context.Context, request *MessagesEditInlineBotMessageRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
