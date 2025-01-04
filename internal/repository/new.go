package repository

import (
	"context"
	"sync"

	"github.com/Unbeliever1987/ethparser/internal/model"
)

type Repository interface {
	InsertBlockAndTransactions(ctx context.Context, block model.Block) error
	RetrieveBlockByNumber(ctx context.Context, number uint64) (model.Block, error)
	RetrieveLatestBlock(ctx context.Context) (model.Block, error)
	RetrieveTransactionsByFromAddress(ctx context.Context, address string) ([]model.Transaction, error)
	RetrieveTransactionsByToAddress(ctx context.Context, address string) ([]model.Transaction, error)
}

type ormBlock struct {
	number       uint64
	transactions []*ormTransaction
}

type ormTransaction struct {
	from   string
	to     string
	amount string
}

type impl struct {
	mu           sync.Mutex
	blocks       []ormBlock
	transactions []ormTransaction
}

func New() Repository {
	return &impl{}
}
