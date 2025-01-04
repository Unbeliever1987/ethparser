package ethparser

import (
	"context"
	"log"
)

// GetCurrentBlock implements Parser.
func (i *impl) GetCurrentBlock(ctx context.Context) int {
	block, err := i.repository.RetrieveLatestBlock(ctx)
	if err != nil {
		log.Printf("failed to retrieve latest block, err: %v", err)
		return -1
	}

	return int(block.Number)
}
