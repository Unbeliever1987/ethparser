package ethparser

import "github.com/Unbeliever1987/ethparser/internal/model"

type Parser interface {
	// last parsed block
	GetCurrentBlock() int
	// add address to observer
	Subscribe(address string) bool
	// list of inbound or outbound transactions for an address
	GetTransactions(address string) []model.Transaction
}
