package ethparser

import (
	"context"
	"log"

	"github.com/Unbeliever1987/ethparser/internal/model"
)

// GetTransactions implements Parser.
func (i *impl) GetTransactions(ctx context.Context, address string) []model.Transaction {
	var transactions []model.Transaction

	fromAddresses, err := i.repository.RetrieveTransactionsByFromAddress(ctx, address)
	if err != nil {
		log.Printf("failed to retrieve transactions by from address, err: %v", err)
	}

	toAddresses, err := i.repository.RetrieveTransactionsByToAddress(ctx, address)
	if err != nil {
		log.Printf("failed to retrieve transactions by to address, err: %v", err)
	}

	transactions = append(transactions, fromAddresses...)
	transactions = append(transactions, toAddresses...)

	return transactions
}
