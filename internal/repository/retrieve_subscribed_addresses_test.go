package repository

import (
	"context"
	"reflect"
	"testing"
)

func Test_impl_RetrieveAllSubcribedAddresses(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "success",
			want: []string{"0x111111"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := New("db_connection")
			if err := r.InsertSubscribedAddress(context.Background(), "0x111111"); err != nil {
				t.Errorf("impl.InsertSubscribedAddress() error = %v", err)
				return
			}

			got, err := r.RetrieveAllSubcribedAddresses(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("impl.RetrieveAllSubcribedAddresses() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("impl.RetrieveAllSubcribedAddresses() = %v, want %v", got, tt.want)
			}
		})
	}
}
