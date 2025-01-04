package repository

import (
	"context"

	"github.com/Unbeliever1987/ethparser/internal/model"
)

// RetrieveLatestBlock implements Repository.
func (r *repository) RetrieveLatestBlock(ctx context.Context) (model.Block, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	latestBlock := r.blocks[len(r.blocks)-1]

	var transactions []model.Transaction
	for _, transaction := range latestBlock.transactions {
		transactions = append(transactions, model.Transaction{
			From:   transaction.from,
			To:     transaction.to,
			Amount: transaction.amount,
		})
	}

	return model.Block{
		Number:       latestBlock.number,
		Transactions: transactions,
	}, nil
}
