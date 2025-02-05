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

// BotsAnswerWebhookJSONQueryRequest represents TL type `bots.answerWebhookJSONQuery#e6213f4d`.
// Answers a custom query; for bots only
//
// See https://core.telegram.org/method/bots.answerWebhookJSONQuery for reference.
type BotsAnswerWebhookJSONQueryRequest struct {
	// Identifier of a custom query
	QueryID int64
	// JSON-serialized answer to the query
	Data DataJSON
}

// BotsAnswerWebhookJSONQueryRequestTypeID is TL type id of BotsAnswerWebhookJSONQueryRequest.
const BotsAnswerWebhookJSONQueryRequestTypeID = 0xe6213f4d

// String implements fmt.Stringer.
func (a *BotsAnswerWebhookJSONQueryRequest) String() string {
	if a == nil {
		return "BotsAnswerWebhookJSONQueryRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("BotsAnswerWebhookJSONQueryRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tQueryID: ")
	sb.WriteString(fmt.Sprint(a.QueryID))
	sb.WriteString(",\n")
	sb.WriteString("\tData: ")
	sb.WriteString(fmt.Sprint(a.Data))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (a *BotsAnswerWebhookJSONQueryRequest) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode bots.answerWebhookJSONQuery#e6213f4d as nil")
	}
	b.PutID(BotsAnswerWebhookJSONQueryRequestTypeID)
	b.PutLong(a.QueryID)
	if err := a.Data.Encode(b); err != nil {
		return fmt.Errorf("unable to encode bots.answerWebhookJSONQuery#e6213f4d: field data: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *BotsAnswerWebhookJSONQueryRequest) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode bots.answerWebhookJSONQuery#e6213f4d to nil")
	}
	if err := b.ConsumeID(BotsAnswerWebhookJSONQueryRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode bots.answerWebhookJSONQuery#e6213f4d: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode bots.answerWebhookJSONQuery#e6213f4d: field query_id: %w", err)
		}
		a.QueryID = value
	}
	{
		if err := a.Data.Decode(b); err != nil {
			return fmt.Errorf("unable to decode bots.answerWebhookJSONQuery#e6213f4d: field data: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for BotsAnswerWebhookJSONQueryRequest.
var (
	_ bin.Encoder = &BotsAnswerWebhookJSONQueryRequest{}
	_ bin.Decoder = &BotsAnswerWebhookJSONQueryRequest{}
)

// BotsAnswerWebhookJSONQuery invokes method bots.answerWebhookJSONQuery#e6213f4d returning error if any.
// Answers a custom query; for bots only
//
// Possible errors:
//  400 QUERY_ID_INVALID: The query ID is invalid
//  400 USER_BOT_INVALID: This method can only be called by a bot
//
// See https://core.telegram.org/method/bots.answerWebhookJSONQuery for reference.
// Can be used by bots.
func (c *Client) BotsAnswerWebhookJSONQuery(ctx context.Context, request *BotsAnswerWebhookJSONQueryRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
