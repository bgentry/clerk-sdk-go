// Code generated by "gen"; DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.
// Last generated at 2024-02-07 11:36:32.720909014 +0000 UTC
package actortoken

import (
	"context"

	"github.com/clerk/clerk-sdk-go/v2"
)

// Create creates a new actor token.
func Create(ctx context.Context, params *CreateParams) (*clerk.ActorToken, error) {
	return getClient().Create(ctx, params)
}

// Revoke revokes a pending actor token.
func Revoke(ctx context.Context, id string) (*clerk.ActorToken, error) {
	return getClient().Revoke(ctx, id)
}

func getClient() *Client {
	return &Client{
		Backend: clerk.GetBackend(),
	}
}
