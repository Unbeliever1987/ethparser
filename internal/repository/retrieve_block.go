package repository

import (
	"context"

	"github.com/Unbeliever1987/ethparser/internal/model"
)

// RetrieveLatestBlock implements Repository.
func (i *impl) RetrieveLatestBlock(ctx context.Context) (model.Block, error) {
	i.mu.Lock()
	defer i.mu.Unlock()

	latestBlock := i.blocks[len(i.blocks)-1]

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

// RetrieveBlockByNumber implements Repository.
func (i *impl) RetrieveBlockByNumber(ctx context.Context, number uint64) (model.Block, error) {
	i.mu.Lock()
	defer i.mu.Unlock()

	for _, block := range i.blocks {
		if block.number == number {
			var transactions []model.Transaction
			for _, txn := range block.transactions {
				transactions = append(transactions, model.Transaction{
					From:   txn.from,
					To:     txn.to,
					Amount: txn.amount,
				})
			}
			return model.Block{
				Number:       block.number,
				Transactions: transactions,
			}, nil
		}
	}

	return model.Block{}, ErrBlockNotFound
}
