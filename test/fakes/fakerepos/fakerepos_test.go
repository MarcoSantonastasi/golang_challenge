package fakerepos

import (
	"reflect"
	"testing"

	pb "github.com/marcosantonastasi/arex_challenge/api/arex/v1"
)

func TestFakeInvestorsRepository_GetAllInvestors(t *testing.T) {
	tests := []struct {
		name    string
		repo    *FakeInvestorsRepository
		want    []*pb.Investor
		wantErr bool
	}{
		{
			name:    "GetAllInvestors () returns exactly the data json file",
			repo:    &FakeInvestorsRepository{},
			want:    FakeAllInvestorsList,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.repo.GetAllInvestors()
			if (err != nil) != tt.wantErr {
				t.Errorf("FakeInvestorsRepository.GetAllInvestors() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FakeInvestorsRepository.GetAllInvestors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFakeIssuersRepository_GetAllIssuers(t *testing.T) {
	tests := []struct {
		name    string
		repo    *FakeIssuersRepository
		want    []*pb.Issuer
		wantErr bool
	}{
		{
			name:    "GetAllIssuers () returns exactly the data json file",
			repo:    &FakeIssuersRepository{},
			want:    FakeAllIssuersList,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.repo.GetAllIssuers()
			if (err != nil) != tt.wantErr {
				t.Errorf("FakeIssuersRepository.GetAllIssuers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FakeIssuersRepository.GetAllIssuers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFakeInvoicesRepository_GetAllInvoices(t *testing.T) {
	tests := []struct {
		name    string
		repo    *FakeInvoicesRepository
		want    []*pb.Invoice
		wantErr bool
	}{
		{
			name:    "GetAllInvoices () returns exactly the data json file",
			repo:    &FakeInvoicesRepository{},
			want:    FakeAllInvoicesList,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.repo.GetAllInvoices()
			if (err != nil) != tt.wantErr {
				t.Errorf("FakeInvoiceRepository.GetAllInvoices() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FakeInvoiceRepository.GetAllInvoices() = %v, want %v", got, tt.want)
			}
		})
	}
}
