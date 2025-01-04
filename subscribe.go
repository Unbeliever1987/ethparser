package ethparser

import (
	"context"
	"log"
)

// Subscribe implements Parser.
func (i *impl) Subscribe(ctx context.Context, address string) bool {
	if err := i.repository.InsertSubscribedAddress(ctx, address); err != nil {
		log.Printf("failed to insert subscribed address, err: %v", err)
		return false
	}

	return true
}
