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
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("impl.GetBlockByNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
