package db_test

import (
	"os"
	"reflect"
	"testing"

	pb "github.com/marcosantonastasi/arex_challenge/api/arex/v1"
	db "github.com/marcosantonastasi/arex_challenge/internal/db"
	data "github.com/marcosantonastasi/arex_challenge/internal/fixtures/data"
)

var testDb db.IDb

func TestMain(m *testing.M) {
	pgUser := os.Getenv("POSTGRES_USER")
	pgPwd := os.Getenv("POSTGRES_PASSWORD")
	pgHostname := os.Getenv("POSTGRES_HOSTNAME")
	pgDbname := os.Getenv("POSTGRES_DB")

	testDb = db.NewPgDb(pgUser, pgPwd, pgHostname, pgDbname)

	testDb.Connect()

	exitCode := m.Run()

	testDb.Close()

	os.Exit(exitCode)
}

func TestDb_GetAllInvestors(t *testing.T) {
	tests := []struct {
		name string
		db   db.IDb
		want *[]*pb.Investor
	}{
		{
			name: "gets the list of all Investors from the investors view",
			db:   testDb,
			want: data.SeededAllInvestorsList,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.db.GetAllInvestors(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Got:\n \t%v\n expected:\n \t%v\n", got, tt.want)
			}
		})
	}
}

func TestDb_GetAllIssuers(t *testing.T) {
	tests := []struct {
		name string
		db   db.IDb
		want *[]*pb.Issuer
	}{
		{
			name: "gets the list of all Issuers from the issuers view",
			db:   testDb,
			want: data.SeededAllIssuersList,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.db.GetAllIssuers(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Got:\n \t%v\n expected:\n \t%v\n", got, tt.want)
			}
		})
	}
}

func TestDb_GetAllInvoices(t *testing.T) {
	tests := []struct {
		name string
		db   db.IDb
		want *[]*pb.Invoice
	}{
		{
			name: "gets the list of all Invoices from the invoices table",
			db:   testDb,
			want: data.SeededAllInvoicesList,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.db.GetAllInvoices(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Got:\n \t%v\n expected:\n \t%v\n", got, tt.want)
			}
		})
	}
}

func TestDb_NewInvoice(t *testing.T) {
	tests := []struct {
		name string
		db   db.IDb
		want *pb.Invoice
	}{
		{
			name: "crates a new invoice using a insert query",
			db:   testDb,
			want: data.NewInvoiceData,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.db.NewInvoice(tt.want)
			got.Id = ""
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Got:\n \t%v\n expected:\n \t%v\n", got, tt.want)
			}
			// WARNING: should delete the invoice that has been createed
		})
	}
}
