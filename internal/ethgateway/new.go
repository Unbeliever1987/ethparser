package ethgateway

import (
	"context"
	"net/http"

	"github.com/Unbeliever1987/ethparser/internal/model"
)

type Gateway interface {
	GetBlockByNumber(ctx context.Context, number string) (model.Block, error)
}

type impl struct {
	baseURL    string
	httpClient *http.Client
}

func New(host string, httpClient *http.Client) Gateway {
	return &impl{
		baseURL:    host,
		httpClient: httpClient,
	}
}
