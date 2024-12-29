package ethgateway

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/Unbeliever1987/ethparser/internal/ethgateway"
	"github.com/Unbeliever1987/ethparser/internal/model"
)

func Test_impl_GetBlockByNumber(t *testing.T) {

	type args struct {
		number string
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
				number: "0x1476e8c",
			},
			want: model.Block{
				Number: 21458572,
				Transactions: []model.Transaction{
					{
						From:   "0x042523db4f3effc33d2742022b2490258494f8b3",
						To:     "0xa69babef1ca67a37ffaf7a485dfff3382056e78c",
						Amount: "0xfcac00",
					},
					{
						From:   "0x91aae0aafd9d2d730111b395c6871f248d7bd728",
						To:     "0x98c3d3183c4b8a650614ad179a1a98be0a8d6b8e",
						Amount: "0x0",
					},
					{
						From:   "0x5e2b6c6b2240d582995537d3fafdb556e4a3822f",
						To:     "0x98c3d3183c4b8a650614ad179a1a98be0a8d6b8e",
						Amount: "0x0",
					},
					{
						From:   "0x3a67b184803641ba107c4c75276d747d123e49ca",
						To:     "0x51c72848c68a965f66fa7a88855f9f7784502a7f",
						Amount: "0x0",
					},
					{
						From:   "0x187bc65342c5c902e8c2084017c7aaf4ec53addf",
						To:     "0x80a64c6d7f12c47b7c66c5b4e20e72bc1fcd5d9e",
						Amount: "0xb1a2bc2ec50000",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: to change to a test host name.
			ethGwy := ethgateway.New("https://ethereum-rpc.publicnode.com", http.DefaultClient)

			got, err := ethGwy.GetBlockByNumber(context.Background(), tt.args.number)
			if (err != nil) != tt.wantErr {
				t.Errorf("impl.GetBlockByNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// To many transactions in one block, just check the first 5 trasnactions.
			got.Transactions = got.Transactions[:5]
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("impl.GetBlockByNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
