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

// AccountAcceptAuthorizationRequest represents TL type `account.acceptAuthorization#e7027c94`.
// Sends a Telegram Passport authorization form, effectively sharing data with the service
//
// See https://core.telegram.org/method/account.acceptAuthorization for reference.
type AccountAcceptAuthorizationRequest struct {
	// Bot ID
	BotID int
	// Telegram Passport element types requested by the service
	Scope string
	// Service's public key
	PublicKey string
	// Types of values sent and their hashes
	ValueHashes []SecureValueHash
	// Encrypted values
	Credentials SecureCredentialsEncrypted
}

// AccountAcceptAuthorizationRequestTypeID is TL type id of AccountAcceptAuthorizationRequest.
const AccountAcceptAuthorizationRequestTypeID = 0xe7027c94

// Encode implements bin.Encoder.
func (a *AccountAcceptAuthorizationRequest) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode account.acceptAuthorization#e7027c94 as nil")
	}
	b.PutID(AccountAcceptAuthorizationRequestTypeID)
	b.PutInt(a.BotID)
	b.PutString(a.Scope)
	b.PutString(a.PublicKey)
	b.PutVectorHeader(len(a.ValueHashes))
	for idx, v := range a.ValueHashes {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode account.acceptAuthorization#e7027c94: field value_hashes element with index %d: %w", idx, err)
		}
	}
	if err := a.Credentials.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.acceptAuthorization#e7027c94: field credentials: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *AccountAcceptAuthorizationRequest) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode account.acceptAuthorization#e7027c94 to nil")
	}
	if err := b.ConsumeID(AccountAcceptAuthorizationRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.acceptAuthorization#e7027c94: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode account.acceptAuthorization#e7027c94: field bot_id: %w", err)
		}
		a.BotID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.acceptAuthorization#e7027c94: field scope: %w", err)
		}
		a.Scope = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.acceptAuthorization#e7027c94: field public_key: %w", err)
		}
		a.PublicKey = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode account.acceptAuthorization#e7027c94: field value_hashes: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value SecureValueHash
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode account.acceptAuthorization#e7027c94: field value_hashes: %w", err)
			}
			a.ValueHashes = append(a.ValueHashes, value)
		}
	}
	{
		if err := a.Credentials.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.acceptAuthorization#e7027c94: field credentials: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountAcceptAuthorizationRequest.
var (
	_ bin.Encoder = &AccountAcceptAuthorizationRequest{}
	_ bin.Decoder = &AccountAcceptAuthorizationRequest{}
)

// AccountAcceptAuthorization invokes method account.acceptAuthorization#e7027c94 returning error if any.
// Sends a Telegram Passport authorization form, effectively sharing data with the service
//
// See https://core.telegram.org/method/account.acceptAuthorization for reference.
func (c *Client) AccountAcceptAuthorization(ctx context.Context, request *AccountAcceptAuthorizationRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
