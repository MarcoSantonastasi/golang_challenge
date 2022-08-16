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
	pgHostName := os.Getenv("POSTGRES_HOSTNAME")
	pgDbName := os.Getenv("POSTGRES_STUB_DB")

	testDb = db.NewPgDb(pgUser, pgPwd, pgHostName, pgDbName)

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

func TestDb_GetAllBids(t *testing.T) {
	tests := []struct {
		name string
		db   db.IDb
		want *[]*pb.Bid
	}{
		{
			name: "gets the list of all Bids from the bids table",
			db:   testDb,
			want: data.SeededAllBidsList,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.db.GetAllBids(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Got:\n \t%v\n expected:\n \t%v\n", got, tt.want)
			}
		})
	}
}

func TestDb_NewBid(t *testing.T) {
	tests := []struct {
		name string
		db   db.IDb
		want *pb.Bid
	}{
		{
			name: "crates a new bid invoking the db function",
			db:   testDb,
			want: data.NewBidData,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.db.NewBid(
				&pb.NewBidRequest{
					InvoiceId:       data.NewBidData.InvoiceId,
					BidderAccountId: data.NewBidData.BidderAccountId,
					Offer:           data.NewBidData.Offer,
				})

			if err != nil {
				t.Errorf("Got an error form the database:\t%+v", err)
			}

			// Trick to pass the test withou killing myself with json values
			if got != nil {
				tt.want.Id = got.Id
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Got:\n%v\nexpected:\n%v", got, tt.want)
			}
			// WARNING: should delete the invoice that has been createed
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
			got, err := tt.db.NewInvoice(
				&pb.NewInvoiceRequest{
					IssuerAccountId: data.NewInvoiceData.IssuerAccountId,
					Reference:       data.NewInvoiceData.Reference,
					Denom:           data.NewInvoiceData.Denom,
					Amount:          data.NewInvoiceData.Amount,
					Asking:          data.NewInvoiceData.Asking})

			if err != nil {
				t.Errorf("Got an error\t%v", err)
			}
			tt.want.Id = got.Id
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Got:\n \t%v\n expected:\n \t%v\n", got, tt.want)
			}
			// WARNING: should delete the invoice that has been createed
		})
	}
}
