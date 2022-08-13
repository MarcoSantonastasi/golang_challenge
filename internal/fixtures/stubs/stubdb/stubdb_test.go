package stubdb

import (
	"reflect"
	"testing"

	pb "github.com/marcosantonastasi/arex_challenge/api/arex/v1"
	db "github.com/marcosantonastasi/arex_challenge/internal/db"
	data "github.com/marcosantonastasi/arex_challenge/internal/fixtures/data"
)

func TestStubDb_GetAllInvestors(t *testing.T) {
	tests := []struct {
		name string
		db   db.IDb
		want *[]*pb.Investor
	}{
		{
			name: "returns exactly the data json file",
			db:   TestStubDb,
			want: data.FakeAllInvestorsList,
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
		want *[]*pb.Issuer
	}{
		{
			name: "returns exactly the data json file",
			db:   TestStubDb,
			want: data.FakeAllIssuersList,
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
		want *[]*pb.Invoice
	}{
		{
			name: "returns exactly the data json file",
			db:   TestStubDb,
			want: data.FakeAllInvoicesList,
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
			db:   TestStubDb,
			want: data.NewInvoiceData,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.db.NewInvoice(tt.want); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FakeDb.NewInvoice() = %v, want %v", got, tt.want)
			}
		})
	}
}
