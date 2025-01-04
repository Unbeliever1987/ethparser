package repository

import (
	"context"

	"github.com/Unbeliever1987/ethparser/internal/model"
)

// InsertBlockAndTransactions implements Repository.
func (i *impl) InsertBlockAndTransactions(ctx context.Context, block model.Block) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	var transactionRefs []*ormTransaction
	for _, transaction := range block.Transactions {
		ormTransaction := ormTransaction{
			from:   transaction.From,
			to:     transaction.To,
			amount: transaction.Amount,
		}

		i.transactions = append(i.transactions, ormTransaction)
		transactionRefs = append(transactionRefs, &ormTransaction)
	}

	i.blocks = append(i.blocks, ormBlock{
		number:       block.Number,
		transactions: transactionRefs,
	})

	return nil
}
