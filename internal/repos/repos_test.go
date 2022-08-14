package repos

import (
	"reflect"
	"testing"

	pb "github.com/marcosantonastasi/arex_challenge/api/arex/v1"
	data "github.com/marcosantonastasi/arex_challenge/internal/fixtures/data"
	stubdb "github.com/marcosantonastasi/arex_challenge/internal/fixtures/stubs/stubdb"
)

func TestInvestorsRepository_GetAllInvestors(t *testing.T) {
	tests := []struct {
		name    string
		repo    *InvestorsRepository
		want    *[]*pb.Investor
		wantErr bool
	}{
		{
			name:    "gets all Investors on the database (3 for newly seeded db)",
			repo:    &InvestorsRepository{Db: stubdb.TestStubDb},
			want:    data.FakeAllInvestorsList,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.repo.GetAllInvestors()
			if (err != nil) != tt.wantErr {
				t.Errorf("InvestorsRepository.GetAllInvestors() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InvestorsRepository.GetAllInvestors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIssuersRepository_GetAllIssuers(t *testing.T) {
	tests := []struct {
		name    string
		repo    *IssuersRepository
		want    *[]*pb.Issuer
		wantErr bool
	}{
		{
			name:    "gets all Issuers on the database (3 for newly seeded db)",
			repo:    &IssuersRepository{Db: stubdb.TestStubDb},
			want:    data.FakeAllIssuersList,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.repo.GetAllIssuers()
			if (err != nil) != tt.wantErr {
				t.Errorf("IssuersRepository.GetAllIssuers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IssuersRepository.GetAllIssuers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBidsRepository_GetAllBids(t *testing.T) {
	tests := []struct {
		name    string
		repo    *BidsRepository
		want    *[]*pb.Bid
		wantErr bool
	}{
		{
			name:    "gets all Bids on the database (3 for newly seeded db)",
			repo:    &BidsRepository{Db: stubdb.TestStubDb},
			want:    data.FakeAllBidsList,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.repo.GetAllBids()
			if (err != nil) != tt.wantErr {
				t.Errorf("BidsRepository.GetAllBids() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BidsRepository.GetAllBids() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInvoicesRepository_GetAllInvoices(t *testing.T) {
	tests := []struct {
		name    string
		repo    *InvoicesRepository
		want    *[]*pb.Invoice
		wantErr bool
	}{
		{
			name:    "gets all Invoices on the database (3 for newly seeded db)",
			repo:    &InvoicesRepository{Db: stubdb.TestStubDb},
			want:    data.FakeAllInvoicesList,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.repo.GetAllInvoices()
			if (err != nil) != tt.wantErr {
				t.Errorf("InvoicesRepository.GetAllInvoices() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InvoicesRepository.GetAllInvoices() = %v, want %v", got, tt.want)
			}
		})
	}
}
