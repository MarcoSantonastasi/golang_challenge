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

func TestPgDb_GetAllInvestors(t *testing.T) {
	tests := []struct {
		name    string
		db      db.IDb
		want    *[]*db.Account
		wantErr bool
	}{
		{
			name:    "gets the list of all Investors from the investors view",
			db:      testDb,
			want:    data.SeededAllInvestorsList,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.db.GetAllInvestors()
			if (err != nil) != tt.wantErr {
				t.Errorf("PgDb.GetAllInvestors() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PgDb.GetAllInvestors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPgDb_GetAllIssuers(t *testing.T) {
	tests := []struct {
		name    string
		db      db.IDb
		want    *[]*db.Account
		wantErr bool
	}{
		{
			name:    "gets the list of all Issuers from the issuers view",
			db:      testDb,
			want:    data.SeededAllIssuersList,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.db.GetAllIssuers()
			if (err != nil) != tt.wantErr {
				t.Errorf("PgDb.GetAllIssuers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PgDb.GetAllIssuers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPgDb_GetAllInvoices(t *testing.T) {
	tests := []struct {
		name    string
		db      db.IDb
		want    *[]*db.Invoice
		wantErr bool
	}{
		{
			name:    "gets the list of all Invoices from the invoices table",
			db:      testDb,
			want:    data.SeededAllInvoicesList,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.db.GetAllInvoices()
			if (err != nil) != tt.wantErr {
				t.Errorf("PgDb.GetAllInvoices() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PgDb.GetAllInvoices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPgDb_GetInvoiceById(t *testing.T) {
	type args struct {
		invoiceId string
	}
	tests := []struct {
		name    string
		db      db.IDb
		args    args
		want    *db.Invoice
		wantErr bool
	}{
		{
			name:    "gets invoice by id from the invoices table",
			db:      testDb,
			args:    args{invoiceId: "af80d0ea-78b9-45b1-a7b0-d1ddd0fbd6fe"},
			want:    (*data.SeededAllInvoicesList)[1],
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.db.GetInvoiceById(tt.args.invoiceId)
			if (err != nil) != tt.wantErr {
				t.Errorf("PgDb.GetInvoiceById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PgDb.GetInvoiceById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPgDb_NewInvoice(t *testing.T) {
	type args struct {
		newInvoiceData Invoice
	}
	tests := []struct {
		name    string
		db      *PgDb
		args    args
		want    *Invoice
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.db.NewInvoice(tt.args.newInvoiceData)
			if (err != nil) != tt.wantErr {
				t.Errorf("PgDb.NewInvoice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PgDb.NewInvoice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPgDb_GetAllBids(t *testing.T) {
	tests := []struct {
		name    string
		db      *PgDb
		want    *[]*Bid
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.db.GetAllBids()
			if (err != nil) != tt.wantErr {
				t.Errorf("PgDb.GetAllBids() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PgDb.GetAllBids() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPgDb_GetBidById(t *testing.T) {
	type args struct {
		bidId int64
	}
	tests := []struct {
		name    string
		db      *PgDb
		args    args
		want    *Bid
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.db.GetBidById(tt.args.bidId)
			if (err != nil) != tt.wantErr {
				t.Errorf("PgDb.GetBidById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PgDb.GetBidById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPgDb_GetBidsByInvoiceId(t *testing.T) {
	type args struct {
		invoiceId string
	}
	tests := []struct {
		name    string
		db      *PgDb
		args    args
		want    *[]*Bid
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.db.GetBidsByInvoiceId(tt.args.invoiceId)
			if (err != nil) != tt.wantErr {
				t.Errorf("PgDb.GetBidsByInvoiceId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PgDb.GetBidsByInvoiceId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPgDb_GetBidsByInvestorId(t *testing.T) {
	type args struct {
		investorId string
	}
	tests := []struct {
		name    string
		db      *PgDb
		args    args
		want    *[]*Bid
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.db.GetBidsByInvestorId(tt.args.investorId)
			if (err != nil) != tt.wantErr {
				t.Errorf("PgDb.GetBidsByInvestorId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PgDb.GetBidsByInvestorId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPgDb_NewBid(t *testing.T) {
	type args struct {
		newBidData Bid
	}
	tests := []struct {
		name    string
		db      *PgDb
		args    args
		want    *Bid
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.db.NewBid(tt.args.newBidData)
			if (err != nil) != tt.wantErr {
				t.Errorf("PgDb.NewBid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PgDb.NewBid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPgDb_GetFulfillingBids(t *testing.T) {
	type args struct {
		invoiceId string
	}
	tests := []struct {
		name    string
		db      *PgDb
		args    args
		want    *[]*Bid
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.db.GetFulfillingBids(tt.args.invoiceId)
			if (err != nil) != tt.wantErr {
				t.Errorf("PgDb.GetFulfillingBids() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PgDb.GetFulfillingBids() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPgDb_AdjudicateBid(t *testing.T) {
	type args struct {
		bidId int64
	}
	tests := []struct {
		name    string
		db      *PgDb
		args    args
		want    *int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.db.AdjudicateBid(tt.args.bidId)
			if (err != nil) != tt.wantErr {
				t.Errorf("PgDb.AdjudicateBid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PgDb.AdjudicateBid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPgDb_AllRunningBidsToLost(t *testing.T) {
	type args struct {
		invoiceId string
	}
	tests := []struct {
		name    string
		db      *PgDb
		args    args
		want    *[]*Bid
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.db.AllRunningBidsToLost(tt.args.invoiceId)
			if (err != nil) != tt.wantErr {
				t.Errorf("PgDb.AllRunningBidsToLost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PgDb.AllRunningBidsToLost() = %v, want %v", got, tt.want)
			}
		})
	}
}
