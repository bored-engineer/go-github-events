//go:build !goexperiment.jsonv2

package github

import (
	"encoding/json"
	"fmt"
)

// Casts a RawEvent to a specific EventPayload type.
func Cast[T EventPayload](in Event[RawEvent]) (out Event[T], _ error) {
	out = Event[T]{
		ID:           in.ID,
		Type:         in.Type,
		Actor:        in.Actor,
		Repository:   in.Repository,
		Public:       in.Public,
		CreatedAt:    in.CreatedAt,
		Organization: in.Organization,
	}
	if err := json.Unmarshal(*in.Payload, &out.Payload); err != nil {
		return out, fmt.Errorf("json.Unmarshal of payload into %T failed: %w", out, err)
	}
	return out, nil
}
