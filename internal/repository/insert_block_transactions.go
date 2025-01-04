package repository

import (
	"context"

	"github.com/Unbeliever1987/ethparser/internal/model"
)

// InsertBlockAndTransactions implements Repository.
func (r *repository) InsertBlockAndTransactions(ctx context.Context, block model.Block) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	var transactionRefs []*ormTransaction
	for _, transaction := range block.Transactions {
		ormTransaction := ormTransaction{
			from:   transaction.From,
			to:     transaction.To,
			amount: transaction.Amount,
		}

		r.transactions = append(r.transactions, ormTransaction)
		transactionRefs = append(transactionRefs, &ormTransaction)
	}

	r.blocks = append(r.blocks, ormBlock{
		number:       block.Number,
		transactions: transactionRefs,
	})

	return nil
}
