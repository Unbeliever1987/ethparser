package repository

import (
	"context"
	"reflect"
	"testing"

	"github.com/Unbeliever1987/ethparser/internal/model"
)

func Test_impl_InsertBlockAndTransactions(t *testing.T) {
	type args struct {
		block model.Block
	}
	tests := []struct {
		name             string
		args             args
		wantBlock        model.Block
		wantTransactions []model.Transaction
		wantErr          bool
	}{
		{
			name: "success",
			args: args{
				block: model.Block{
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
				},
			},
			wantBlock: model.Block{
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
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := New()
			if err := r.InsertBlockAndTransactions(context.Background(), tt.args.block); (err != nil) != tt.wantErr {
				t.Errorf("impl.InsertBlockAndTransactions() error = %v, wantErr %v", err, tt.wantErr)
			}

			got, err := r.RetrieveBlockByNumber(context.Background(), tt.args.block.Number)
			if (err != nil) != tt.wantErr {
				t.Errorf("impl.InsertBlockAndTransactions() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(got, tt.wantBlock) {
				t.Errorf("impl.InsertBlockAndTransactions() = %v, want %v", got, tt.wantBlock)
			}
		})
	}
}
