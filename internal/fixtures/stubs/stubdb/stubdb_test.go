//go:build unit_tests
// +build unit_tests

package stubdb

import (
	"reflect"
	"testing"

	pb "github.com/marcosantonastasi/arex_challenge/api/arex/v1"
	db "github.com/marcosantonastasi/arex_challenge/internal/db"
)

func TestFakeDb_GetAllInvestors(t *testing.T) {
	tests := []struct {
		name string
		db   db.IDb
		want []*pb.Investor
	}{
		{
			name: "GetAllInvestors() returns exactly the data json file",
			db:   &StubDb{},
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
		db   *StubDb
		want []*pb.Issuer
	}{
		{
			name: "GetAllIssuers() returns exactly the data json file",
			db:   &StubDb{},
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
		db   *StubDb
		want []*pb.Invoice
	}{
		{
			name: "GetAllInvoices() returns exactly the data json file",
			db:   &StubDb{},
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
