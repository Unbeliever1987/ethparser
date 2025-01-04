package repository

import (
	"context"
	"reflect"
	"testing"

	"github.com/Unbeliever1987/ethparser/internal/model"
)

var mockBlock1 = model.Block{
	Number: 1,
	Transactions: []model.Transaction{
		{
			From:   "0x111111",
			To:     "0x222222",
			Amount: "0x1234",
		},
		{
			From:   "0x333333",
			To:     "ox444444",
			Amount: "0x5678",
		},
	},
}

var mockBlock2 = model.Block{
	Number: 2,
	Transactions: []model.Transaction{
		{
			From:   "0x222222",
			To:     "0x111111",
			Amount: "0x8888",
		},
		{
			From:   "0x555555",
			To:     "ox666666",
			Amount: "0x9999",
		},
	},
}

func Test_impl_RetrieveLatestBlock(t *testing.T) {
	tests := []struct {
		name    string
		want    model.Block
		wantErr bool
	}{
		{
			name: "success",
			want: mockBlock2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := New("db_connection")
			// Mock block insertion.
			if err := r.InsertBlockAndTransactions(context.Background(), mockBlock1); err != nil {
				t.Errorf("failed to insert mock data, err: %v", err)
				return
			}

			if err := r.InsertBlockAndTransactions(context.Background(), mockBlock2); err != nil {
				t.Errorf("failed to insert mock data, err: %v", err)
				return
			}

			got, err := r.RetrieveLatestBlock(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("impl.RetrieveLatestBlock() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("impl.RetrieveLatestBlock() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_impl_RetrieveBlockByNumber(t *testing.T) {
	type args struct {
		number uint64
	}
	tests := []struct {
		name    string
		args    args
		want    model.Block
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				number: 2,
			},
			want: mockBlock2,
		},
		{
			name: "err: no block found",
			args: args{
				number: 999,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := New("db_connection")
			// Mock block insertion.
			if err := r.InsertBlockAndTransactions(context.Background(), mockBlock1); err != nil {
				t.Errorf("failed to insert mock data, err: %v", err)
				return
			}

			if err := r.InsertBlockAndTransactions(context.Background(), mockBlock2); err != nil {
				t.Errorf("failed to insert mock data, err: %v", err)
				return
			}

			got, err := r.RetrieveBlockByNumber(context.Background(), tt.args.number)
			if (err != nil) != tt.wantErr {
				t.Errorf("impl.RetrieveBlockByNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("impl.RetrieveBlockByNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
