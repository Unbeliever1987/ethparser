package repository

import (
	"context"
	"reflect"
	"testing"

	"github.com/Unbeliever1987/ethparser/internal/model"
)

func Test_impl_RetrieveTransactionsByFromAddress(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name    string
		args    args
		want    []model.Transaction
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				address: "0x111111",
			},
			want: []model.Transaction{
				{
					From:   "0x111111",
					To:     "0x222222",
					Amount: "0x1234",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := New()
			// Mock block insertion.
			if err := r.InsertBlockAndTransactions(context.Background(), mockBlock1); err != nil {
				t.Errorf("failed to insert mock data, err: %v", err)
				return
			}

			if err := r.InsertBlockAndTransactions(context.Background(), mockBlock2); err != nil {
				t.Errorf("failed to insert mock data, err: %v", err)
				return
			}

			got, err := r.RetrieveTransactionsByFromAddress(context.Background(), tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("impl.RetrieveTransactionsByFromAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("impl.RetrieveTransactionsByFromAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_impl_RetrieveTransactionsByToAddress(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name    string
		args    args
		want    []model.Transaction
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				address: "0x111111",
			},
			want: []model.Transaction{
				{
					From:   "0x222222",
					To:     "0x111111",
					Amount: "0x8888",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := New()
			// Mock block insertion.
			if err := r.InsertBlockAndTransactions(context.Background(), mockBlock1); err != nil {
				t.Errorf("failed to insert mock data, err: %v", err)
				return
			}

			if err := r.InsertBlockAndTransactions(context.Background(), mockBlock2); err != nil {
				t.Errorf("failed to insert mock data, err: %v", err)
				return
			}

			got, err := r.RetrieveTransactionsByToAddress(context.Background(), tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("impl.RetrieveTransactionsByToAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("impl.RetrieveTransactionsByToAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}
