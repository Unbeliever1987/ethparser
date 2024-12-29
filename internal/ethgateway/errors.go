package ethgateway

import (
	"errors"
)

var (
	errEmptyBlockNumber         = errors.New("empty block number returned")
	errUnexpectedHTTPStatusCode = errors.New("unexpected status code")
	errBlockNumberFormat        = errors.New("block number should start with 0x")
	errTxValueFormat            = errors.New("transaction value should start with 0x")
)
