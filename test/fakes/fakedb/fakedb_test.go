package fakedb

import (
	"reflect"
	"testing"

	pb "github.com/marcosantonastasi/arex_challenge/api/arex/v1"
)

func TestFakeDb_GetAllInvestors(t *testing.T) {
	tests := []struct {
		name string
		db   *FakeDb
		want []*pb.Investor
	}{
		{
			name: "GetAllInvestors() returns exactly the data json file",
			db:   &FakeDb{},
			want: FakeAllInvestorsList,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.db.GetAllInvestors(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FakeDb.GetAllInvestors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFakeDb_GetAllIssuers(t *testing.T) {
	tests := []struct {
		name string
		db   *FakeDb
		want []*pb.Issuer
	}{
		{
			name: "GetAllIssuers() returns exactly the data json file",
			db:   &FakeDb{},
			want: FakeAllIssuersList,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.db.GetAllIssuers(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FakeDb.GetAllIssuers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFakeDb_GetAllInvoices(t *testing.T) {
	tests := []struct {
		name string
		db   *FakeDb
		want []*pb.Invoice
	}{
		{
			name: "GetAllInvoices() returns exactly the data json file",
			db:   &FakeDb{},
			want: FakeAllInvoicesList,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.db.GetAllInvoices(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FakeDb.GetAllInvoices() = %v, want %v", got, tt.want)
			}
		})
	}
}
