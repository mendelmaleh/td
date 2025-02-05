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

// MessagesSearchRequest represents TL type `messages.search#c352eec`.
// Gets back found messages
//
// See https://core.telegram.org/method/messages.search for reference.
type MessagesSearchRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// User or chat, histories with which are searched, or (inputPeerEmpty)¹ constructor for global search
	//
	// Links:
	//  1) https://core.telegram.org/constructor/inputPeerEmpty
	Peer InputPeerClass
	// Text search request
	Q string
	// Only return messages sent by the specified user ID
	//
	// Use SetFromID and GetFromID helpers.
	FromID InputPeerClass
	// Thread ID¹
	//
	// Links:
	//  1) https://core.telegram.org/api/threads
	//
	// Use SetTopMsgID and GetTopMsgID helpers.
	TopMsgID int
	// Filter to return only specified message types
	Filter MessagesFilterClass
	// If a positive value was transferred, only messages with a sending date bigger than the transferred one will be returned
	MinDate int
	// If a positive value was transferred, only messages with a sending date smaller than the transferred one will be returned
	MaxDate int
	// Only return messages starting from the specified message ID
	OffsetID int
	// Additional offset¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	AddOffset int
	// Number of results to return¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	Limit int
	// Maximum message ID to return¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	MaxID int
	// Minimum message ID to return¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	MinID int
	// Hash¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	Hash int
}

// MessagesSearchRequestTypeID is TL type id of MessagesSearchRequest.
const MessagesSearchRequestTypeID = 0xc352eec

// String implements fmt.Stringer.
func (s *MessagesSearchRequest) String() string {
	if s == nil {
		return "MessagesSearchRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesSearchRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(fmt.Sprint(s.Flags))
	sb.WriteString(",\n")
	sb.WriteString("\tPeer: ")
	sb.WriteString(fmt.Sprint(s.Peer))
	sb.WriteString(",\n")
	sb.WriteString("\tQ: ")
	sb.WriteString(fmt.Sprint(s.Q))
	sb.WriteString(",\n")
	if s.Flags.Has(0) {
		sb.WriteString("\tFromID: ")
		sb.WriteString(fmt.Sprint(s.FromID))
		sb.WriteString(",\n")
	}
	if s.Flags.Has(1) {
		sb.WriteString("\tTopMsgID: ")
		sb.WriteString(fmt.Sprint(s.TopMsgID))
		sb.WriteString(",\n")
	}
	sb.WriteString("\tFilter: ")
	sb.WriteString(fmt.Sprint(s.Filter))
	sb.WriteString(",\n")
	sb.WriteString("\tMinDate: ")
	sb.WriteString(fmt.Sprint(s.MinDate))
	sb.WriteString(",\n")
	sb.WriteString("\tMaxDate: ")
	sb.WriteString(fmt.Sprint(s.MaxDate))
	sb.WriteString(",\n")
	sb.WriteString("\tOffsetID: ")
	sb.WriteString(fmt.Sprint(s.OffsetID))
	sb.WriteString(",\n")
	sb.WriteString("\tAddOffset: ")
	sb.WriteString(fmt.Sprint(s.AddOffset))
	sb.WriteString(",\n")
	sb.WriteString("\tLimit: ")
	sb.WriteString(fmt.Sprint(s.Limit))
	sb.WriteString(",\n")
	sb.WriteString("\tMaxID: ")
	sb.WriteString(fmt.Sprint(s.MaxID))
	sb.WriteString(",\n")
	sb.WriteString("\tMinID: ")
	sb.WriteString(fmt.Sprint(s.MinID))
	sb.WriteString(",\n")
	sb.WriteString("\tHash: ")
	sb.WriteString(fmt.Sprint(s.Hash))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (s *MessagesSearchRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.search#c352eec as nil")
	}
	b.PutID(MessagesSearchRequestTypeID)
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.search#c352eec: field flags: %w", err)
	}
	if s.Peer == nil {
		return fmt.Errorf("unable to encode messages.search#c352eec: field peer is nil")
	}
	if err := s.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.search#c352eec: field peer: %w", err)
	}
	b.PutString(s.Q)
	if s.Flags.Has(0) {
		if s.FromID == nil {
			return fmt.Errorf("unable to encode messages.search#c352eec: field from_id is nil")
		}
		if err := s.FromID.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.search#c352eec: field from_id: %w", err)
		}
	}
	if s.Flags.Has(1) {
		b.PutInt(s.TopMsgID)
	}
	if s.Filter == nil {
		return fmt.Errorf("unable to encode messages.search#c352eec: field filter is nil")
	}
	if err := s.Filter.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.search#c352eec: field filter: %w", err)
	}
	b.PutInt(s.MinDate)
	b.PutInt(s.MaxDate)
	b.PutInt(s.OffsetID)
	b.PutInt(s.AddOffset)
	b.PutInt(s.Limit)
	b.PutInt(s.MaxID)
	b.PutInt(s.MinID)
	b.PutInt(s.Hash)
	return nil
}

// SetFromID sets value of FromID conditional field.
func (s *MessagesSearchRequest) SetFromID(value InputPeerClass) {
	s.Flags.Set(0)
	s.FromID = value
}

// GetFromID returns value of FromID conditional field and
// boolean which is true if field was set.
func (s *MessagesSearchRequest) GetFromID() (value InputPeerClass, ok bool) {
	if !s.Flags.Has(0) {
		return value, false
	}
	return s.FromID, true
}

// SetTopMsgID sets value of TopMsgID conditional field.
func (s *MessagesSearchRequest) SetTopMsgID(value int) {
	s.Flags.Set(1)
	s.TopMsgID = value
}

// GetTopMsgID returns value of TopMsgID conditional field and
// boolean which is true if field was set.
func (s *MessagesSearchRequest) GetTopMsgID() (value int, ok bool) {
	if !s.Flags.Has(1) {
		return value, false
	}
	return s.TopMsgID, true
}

// Decode implements bin.Decoder.
func (s *MessagesSearchRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.search#c352eec to nil")
	}
	if err := b.ConsumeID(MessagesSearchRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.search#c352eec: %w", err)
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.search#c352eec: field flags: %w", err)
		}
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.search#c352eec: field peer: %w", err)
		}
		s.Peer = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.search#c352eec: field q: %w", err)
		}
		s.Q = value
	}
	if s.Flags.Has(0) {
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.search#c352eec: field from_id: %w", err)
		}
		s.FromID = value
	}
	if s.Flags.Has(1) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.search#c352eec: field top_msg_id: %w", err)
		}
		s.TopMsgID = value
	}
	{
		value, err := DecodeMessagesFilter(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.search#c352eec: field filter: %w", err)
		}
		s.Filter = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.search#c352eec: field min_date: %w", err)
		}
		s.MinDate = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.search#c352eec: field max_date: %w", err)
		}
		s.MaxDate = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.search#c352eec: field offset_id: %w", err)
		}
		s.OffsetID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.search#c352eec: field add_offset: %w", err)
		}
		s.AddOffset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.search#c352eec: field limit: %w", err)
		}
		s.Limit = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.search#c352eec: field max_id: %w", err)
		}
		s.MaxID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.search#c352eec: field min_id: %w", err)
		}
		s.MinID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.search#c352eec: field hash: %w", err)
		}
		s.Hash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesSearchRequest.
var (
	_ bin.Encoder = &MessagesSearchRequest{}
	_ bin.Decoder = &MessagesSearchRequest{}
)

// MessagesSearch invokes method messages.search#c352eec returning error if any.
// Gets back found messages
//
// Possible errors:
//  400 CHANNEL_INVALID: The provided channel is invalid
//  400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup
//  400 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this
//  400 INPUT_CONSTRUCTOR_INVALID: The provided constructor is invalid
//  400 INPUT_USER_DEACTIVATED: The specified user was deleted
//  400 MSG_ID_INVALID: Invalid message ID provided
//  400 PEER_ID_INVALID: The provided peer id is invalid
//  400 PEER_ID_NOT_SUPPORTED: The provided peer ID is not supported
//  400 SEARCH_QUERY_EMPTY: The search query is empty
//  400 USER_ID_INVALID: The provided user ID is invalid
//
// See https://core.telegram.org/method/messages.search for reference.
func (c *Client) MessagesSearch(ctx context.Context, request *MessagesSearchRequest) (MessagesMessagesClass, error) {
	var result MessagesMessagesBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Messages, nil
}
