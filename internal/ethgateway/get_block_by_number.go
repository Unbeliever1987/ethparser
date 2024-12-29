package ethgateway

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/Unbeliever1987/ethparser/internal/model"
)

type getBlockByNumberResponse struct {
	Result struct {
		Number       string `json:"number"`
		Transactions []struct {
			From  string `json:"from"`
			To    string `json:"to"`
			Value string `json:"value"`
		} `json:"transactions"`
	} `json:"result"`
}

func (i impl) GetBlockByNumber(ctx context.Context, number string) (model.Block, error) {
	reqBody, err := json.Marshal(map[string]any{
		"jsonrpc": "2.0",
		"method":  "eth_getBlockByNumber",
		"params": []any{
			number,
			true,
		},
		"id": 1,
	})
	if err != nil {
		return model.Block{}, err
	}

	req, err := http.NewRequest(http.MethodPost, i.baseURL, bytes.NewBuffer(reqBody))
	if err != nil {
		return model.Block{}, err
	}

	resp, err := i.httpClient.Do(req)
	if err != nil {
		return model.Block{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.Block{}, fmt.Errorf("%w, status code: %d", errUnexpectedHTTPStatusCode, resp.StatusCode)
	}

	respBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.Block{}, err
	}

	var respBody getBlockByNumberResponse
	if err = json.Unmarshal(respBodyBytes, &respBody); err != nil {
		return model.Block{}, err
	}

	block, err := respBody.toBlock()
	if err != nil {
		return model.Block{}, err
	}

	return block, nil
}

func (resp getBlockByNumberResponse) toBlock() (model.Block, error) {
	if resp.Result.Number == "" {
		return model.Block{}, errEmptyBlockNumber
	}

	if !strings.HasPrefix(resp.Result.Number, "0x") {
		return model.Block{}, errBlockNumberFormat
	}

	number, err := strconv.ParseUint(resp.Result.Number[2:], 16, 64)
	if err != nil {
		return model.Block{}, err
	}

	var transactions []model.Transaction

	for _, t := range resp.Result.Transactions {
		transactions = append(transactions, model.Transaction{
			From:   t.From,
			To:     t.To,
			Amount: t.Value,
		})
	}

	return model.Block{
		Number:       number,
		Transactions: transactions,
	}, nil
}
