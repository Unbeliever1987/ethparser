package ethparser

import (
	"context"
	"net/http"

	"github.com/Unbeliever1987/ethparser/internal/ethblocksyncer"
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
}

func New(ctx context.Context, dbconn, ethHost, startBlockNumber string) Parser {
	repository := repository.New(dbconn)
	ethGateway := ethgateway.New(ethHost, http.DefaultClient)
	syncer := ethblocksyncer.New(repository, ethGateway)
	syncer.Run(ctx, startBlockNumber)

	return &impl{
		repository: repository,
	}
}
