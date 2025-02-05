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

// MessagesSetInlineGameScoreRequest represents TL type `messages.setInlineGameScore#15ad9f64`.
// Use this method to set the score of the specified user in a game sent as an inline message (bots only).
//
// See https://core.telegram.org/method/messages.setInlineGameScore for reference.
type MessagesSetInlineGameScoreRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Set this flag if the game message should be automatically edited to include the current scoreboard
	EditMessage bool
	// Set this flag if the high score is allowed to decrease. This can be useful when fixing mistakes or banning cheaters
	Force bool
	// ID of the inline message
	ID InputBotInlineMessageID
	// User identifier
	UserID InputUserClass
	// New score
	Score int
}

// MessagesSetInlineGameScoreRequestTypeID is TL type id of MessagesSetInlineGameScoreRequest.
const MessagesSetInlineGameScoreRequestTypeID = 0x15ad9f64

// String implements fmt.Stringer.
func (s *MessagesSetInlineGameScoreRequest) String() string {
	if s == nil {
		return "MessagesSetInlineGameScoreRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesSetInlineGameScoreRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(fmt.Sprint(s.Flags))
	sb.WriteString(",\n")
	sb.WriteString("\tID: ")
	sb.WriteString(fmt.Sprint(s.ID))
	sb.WriteString(",\n")
	sb.WriteString("\tUserID: ")
	sb.WriteString(fmt.Sprint(s.UserID))
	sb.WriteString(",\n")
	sb.WriteString("\tScore: ")
	sb.WriteString(fmt.Sprint(s.Score))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (s *MessagesSetInlineGameScoreRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.setInlineGameScore#15ad9f64 as nil")
	}
	b.PutID(MessagesSetInlineGameScoreRequestTypeID)
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.setInlineGameScore#15ad9f64: field flags: %w", err)
	}
	if err := s.ID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.setInlineGameScore#15ad9f64: field id: %w", err)
	}
	if s.UserID == nil {
		return fmt.Errorf("unable to encode messages.setInlineGameScore#15ad9f64: field user_id is nil")
	}
	if err := s.UserID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.setInlineGameScore#15ad9f64: field user_id: %w", err)
	}
	b.PutInt(s.Score)
	return nil
}

// SetEditMessage sets value of EditMessage conditional field.
func (s *MessagesSetInlineGameScoreRequest) SetEditMessage(value bool) {
	if value {
		s.Flags.Set(0)
		s.EditMessage = true
	} else {
		s.Flags.Unset(0)
		s.EditMessage = false
	}
}

// SetForce sets value of Force conditional field.
func (s *MessagesSetInlineGameScoreRequest) SetForce(value bool) {
	if value {
		s.Flags.Set(1)
		s.Force = true
	} else {
		s.Flags.Unset(1)
		s.Force = false
	}
}

// Decode implements bin.Decoder.
func (s *MessagesSetInlineGameScoreRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.setInlineGameScore#15ad9f64 to nil")
	}
	if err := b.ConsumeID(MessagesSetInlineGameScoreRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.setInlineGameScore#15ad9f64: %w", err)
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.setInlineGameScore#15ad9f64: field flags: %w", err)
		}
	}
	s.EditMessage = s.Flags.Has(0)
	s.Force = s.Flags.Has(1)
	{
		if err := s.ID.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.setInlineGameScore#15ad9f64: field id: %w", err)
		}
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.setInlineGameScore#15ad9f64: field user_id: %w", err)
		}
		s.UserID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.setInlineGameScore#15ad9f64: field score: %w", err)
		}
		s.Score = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesSetInlineGameScoreRequest.
var (
	_ bin.Encoder = &MessagesSetInlineGameScoreRequest{}
	_ bin.Decoder = &MessagesSetInlineGameScoreRequest{}
)

// MessagesSetInlineGameScore invokes method messages.setInlineGameScore#15ad9f64 returning error if any.
// Use this method to set the score of the specified user in a game sent as an inline message (bots only).
//
// Possible errors:
//  400 MESSAGE_ID_INVALID: The provided message id is invalid
//  400 USER_BOT_REQUIRED: This method can only be called by a bot
//
// See https://core.telegram.org/method/messages.setInlineGameScore for reference.
// Can be used by bots.
func (c *Client) MessagesSetInlineGameScore(ctx context.Context, request *MessagesSetInlineGameScoreRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
