package ethparser

import (
	"context"
	"net/http"

	"github.com/Unbeliever1987/ethparser/internal/ethgateway"
	"github.com/Unbeliever1987/ethparser/internal/model"
	"github.com/Unbeliever1987/ethparser/internal/repository"
)

type Parser interface {
	// last parsed block
	GetCurrentBlock(ctx context.Context) int
	// add address to observer
	Subscribe(ctx context.Context, address string) bool
	// list of inbound or outbound transactions for an address
	GetTransactions(ctx context.Context, address string) []model.Transaction
}

type impl struct {
	repository repository.Repository
	ethgateway ethgateway.Gateway
}

func New(dbconn, ethHost, startBlockNumber string) Parser {
	return &impl{
		repository: repository.New(dbconn),
		ethgateway: ethgateway.New(ethHost, http.DefaultClient),
	}
}
