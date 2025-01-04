package ethblocksyncer

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Unbeliever1987/ethparser/internal/model"
)

func (s Runner) Run(ctx context.Context, blockNumber string) {
	go func() {
		for {
			time.Sleep(500 * time.Millisecond)
			err := s.sync(ctx, blockNumber)
			if err != nil {
				log.Fatalf("failed to sync block, err :%v", err)
			}
		}
	}()
}

func (s Runner) sync(ctx context.Context, blockNumber string) error {
	subscribedAddresses, err := s.repository.RetrieveAllSubcribedAddresses(ctx)
	if err != nil {
		return fmt.Errorf("failed to retrieve subscribed addresses, err: %v", err)
	}

	block, err := s.ethGateway.GetBlockByNumber(ctx, blockNumber)
	if err != nil {
		return fmt.Errorf("failed to get block by number, err: %v", err)
	}

	var filteredBlock model.Block
	filteredBlock.Number = block.Number

	var filteredTransactions []model.Transaction
	for _, address := range subscribedAddresses {
		for _, transaction := range block.Transactions {
			if transaction.From == address || transaction.To == address {
				filteredTransactions = append(filteredTransactions, transaction)
			}
		}
	}

	filteredBlock.Transactions = filteredTransactions

	err = s.repository.InsertBlockAndTransactions(ctx, filteredBlock)
	if err != nil {
		return fmt.Errorf("failed to insert block and transactions, err : %v", err)
	}

	return nil
}
