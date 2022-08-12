//go:build unit_tests
// +build unit_tests

package stubdb

import (
	"reflect"
	"testing"

	pb "github.com/marcosantonastasi/arex_challenge/api/arex/v1"
	db "github.com/marcosantonastasi/arex_challenge/internal/db"
)

func TestStubDb_GetAllInvestors(t *testing.T) {
	tests := []struct {
		name string
		db   db.IDb
		want []*pb.Investor
	}{
		{
			name: "returns exactly the data json file",
			db:   &StubDb{},
			want: loadFakeInvestorsData(),
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

func TestStubDb_GetAllIssuers(t *testing.T) {
	tests := []struct {
		name string
		db   db.IDb
		want []*pb.Issuer
	}{
		{
			name: "returns exactly the data json file",
			db:   &StubDb{},
			want: loadFakeIssuersData(),
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

func TestStubDb_GetAllInvoices(t *testing.T) {
	tests := []struct {
		name string
		db   db.IDb
		want []*pb.Invoice
	}{
		{
			name: "returns exactly the data json file",
			db:   &StubDb{},
			want: loadFakeInvoicesData(),
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

func TestStubDb_NewInvoice(t *testing.T) {
	tests := []struct {
		name string
		db   db.IDb
		want *pb.Invoice
	}{
		{
			name: "returns exactly the data json file",
			db:   &StubDb{},
			want: loadFakeNewInvoiceData(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.db.NewInvoice(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FakeDb.NewInvoice() = %v, want %v", got, tt.want)
			}
		})
	}
}
