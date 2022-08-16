package stubrepos

import (
	"reflect"
	"testing"

	db "github.com/marcosantonastasi/arex_challenge/internal/db"
	data "github.com/marcosantonastasi/arex_challenge/internal/fixtures/data"
)

func TestFakeInvestorsRepository_GetAllInvestors(t *testing.T) {
	tests := []struct {
		name    string
		repo    *StubInvestorsRepository
		want    *[]*db.Account
		wantErr bool
	}{
		{
			name:    "GetAllInvestors () returns exactly the data json file",
			repo:    &StubInvestorsRepository{},
			want:    data.SeededAllInvestorsList,
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
		repo    *StubIssuersRepository
		want    *[]*db.Account
		wantErr bool
	}{
		{
			name:    "GetAllIssuers () returns exactly the data json file",
			repo:    &StubIssuersRepository{},
			want:    data.SeededAllIssuersList,
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
		repo    *StubInvoicesRepository
		want    *[]*db.Invoice
		wantErr bool
	}{
		{
			name:    "GetAllInvoices () returns exactly the data json file",
			repo:    &StubInvoicesRepository{},
			want:    data.SeededAllInvoicesList,
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
