package db_test

import (
	"log"
	"os"
	"reflect"
	"testing"

	db "github.com/marcosantonastasi/arex_challenge/internal/db"
	data "github.com/marcosantonastasi/arex_challenge/internal/fixtures/data"
)

var testDb db.IDb

func TestMain(m *testing.M) {
	pgUser := os.Getenv("POSTGRES_USER")
	pgPwd := os.Getenv("POSTGRES_PASSWORD")
	pgHostName := os.Getenv("POSTGRES_HOSTNAME")
	pgDbName := os.Getenv("POSTGRES_STUB_DB")

	if pgUser == "" || pgPwd == "" || pgHostName == "" || pgDbName == "" {
		log.Fatal("stubdb unit test is missing .env variables")
	}

	testDb = db.NewPgDb(pgUser, pgPwd, pgHostName, pgDbName)

	testDb.Connect()

	testDbErr := testDb.Connect()
	if testDbErr != nil {
		log.Fatalf("stubdb unit test conneciton to PostgresDB failed %+v", testDbErr)
	}
	log.Printf("stubdb unit test conneciton to PostgresDB %+v", testDb)

	defer testDb.Close()

	os.Exit(m.Run())
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
				t.Errorf("testDb.GetAllInvestors() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("testDb.GetAllInvestors() = %+v, want %+v", got, tt.want)
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
				t.Errorf("testDb.GetAllIssuers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("testDb.GetAllIssuers() = %v, want %v", got, tt.want)
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
				t.Errorf("testDb.GetAllInvoices() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("testDb.GetAllInvoices() = %v, want %v", got, tt.want)
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
				t.Errorf("testDb.GetInvoiceById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("testDb.GetInvoiceById() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPgDb_NewInvoice(t *testing.T) {
	type args struct {
		newInvoiceData db.Invoice
	}
	tests := []struct {
		name    string
		db      db.IDb
		args    args
		want    *db.Invoice
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.db.NewInvoice(tt.args.newInvoiceData)
			if (err != nil) != tt.wantErr {
				t.Errorf("testDb.NewInvoice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("testDb.NewInvoice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPgDb_GetAllBids(t *testing.T) {
	tests := []struct {
		name    string
		db      db.IDb
		want    *[]*db.Bid
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.db.GetAllBids()
			if (err != nil) != tt.wantErr {
				t.Errorf("testDb.GetAllBids() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("testDb.GetAllBids() = %v, want %v", got, tt.want)
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
		db      db.IDb
		args    args
		want    *db.Bid
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.db.GetBidById(tt.args.bidId)
			if (err != nil) != tt.wantErr {
				t.Errorf("testDb.GetBidById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("testDb.GetBidById() = %v, want %v", got, tt.want)
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
		db      db.IDb
		args    args
		want    *[]*db.Bid
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.db.GetBidsByInvoiceId(tt.args.invoiceId)
			if (err != nil) != tt.wantErr {
				t.Errorf("testDb.GetBidsByInvoiceId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("testDb.GetBidsByInvoiceId() = %v, want %v", got, tt.want)
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
		db      db.IDb
		args    args
		want    *[]*db.Bid
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.db.GetBidsByInvestorId(tt.args.investorId)
			if (err != nil) != tt.wantErr {
				t.Errorf("testDb.GetBidsByInvestorId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("testDb.GetBidsByInvestorId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPgDb_NewBid(t *testing.T) {
	type args struct {
		newBidData db.Bid
	}
	tests := []struct {
		name    string
		db      db.IDb
		args    args
		want    *db.Bid
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.db.NewBid(tt.args.newBidData)
			if (err != nil) != tt.wantErr {
				t.Errorf("testDb.NewBid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("testDb.NewBid() = %v, want %v", got, tt.want)
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
		db      db.IDb
		args    args
		want    *[]*db.Bid
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.db.GetFulfillingBids(tt.args.invoiceId)
			if (err != nil) != tt.wantErr {
				t.Errorf("testDb.GetFulfillingBids() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("testDb.GetFulfillingBids() = %v, want %v", got, tt.want)
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
		db      db.IDb
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
				t.Errorf("testDb.AdjudicateBid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("testDb.AdjudicateBid() = %v, want %v", got, tt.want)
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
		db      db.IDb
		args    args
		want    *[]*db.Bid
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.db.AllRunningBidsToLost(tt.args.invoiceId)
			if (err != nil) != tt.wantErr {
				t.Errorf("testDb.AllRunningBidsToLost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("testDb.AllRunningBidsToLost() = %v, want %v", got, tt.want)
			}
		})
	}
}
