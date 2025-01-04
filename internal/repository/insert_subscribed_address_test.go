package repository

import (
	"context"
	"reflect"
	"testing"
)

func Test_impl_InsertSubscribedAddress(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    []string
	}{
		{
			name: "success",
			args: args{
				address: "0x111111",
			},
			want: []string{"0x111111"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := New("db_connection")
			if err := r.InsertSubscribedAddress(context.Background(), tt.args.address); (err != nil) != tt.wantErr {
				t.Errorf("impl.InsertSubscribedAddress() error = %v, wantErr %v", err, tt.wantErr)
			}

			got, err := r.RetrieveAllSubcribedAddresses(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("impl.RetrieveAllSubcribedAddresses() error = %v", err)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("impl.RetrieveAllSubcribedAddresses() = %v, want %v", got, tt.want)
			}
		})
	}
}
