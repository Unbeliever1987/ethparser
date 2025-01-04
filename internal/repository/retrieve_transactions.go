package repository

import (
	"context"

	"github.com/Unbeliever1987/ethparser/internal/model"
)

// RetrieveTransactionsByFromAddress implements Repository.
func (i *impl) RetrieveTransactionsByFromAddress(ctx context.Context, address string) ([]model.Transaction, error) {
	i.mu.Lock()
	defer i.mu.Unlock()

	var transactions []model.Transaction
	for _, transaction := range i.transactions {
		if transaction.from != address {
			continue
		}
		transactions = append(transactions, model.Transaction{
			From:   transaction.from,
			To:     transaction.to,
			Amount: transaction.amount,
		})
	}
	return transactions, nil
}

// RetrieveTransactionsByToAddress implements Repository.
func (r *impl) RetrieveTransactionsByToAddress(ctx context.Context, address string) ([]model.Transaction, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var transactions []model.Transaction
	for _, transaction := range r.transactions {
		if transaction.to != address {
			continue
		}
		transactions = append(transactions, model.Transaction{
			From:   transaction.from,
			To:     transaction.to,
			Amount: transaction.amount,
		})
	}
	return transactions, nil
}
