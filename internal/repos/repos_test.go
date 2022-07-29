package repos

import (
	"reflect"
	"testing"

	pb "github.com/marcosantonastasi/arex_challenge/api/arex/v1"
)

func TestInvestorsRepository_GetAllInvestors(t *testing.T) {
	tests := []struct {
		name    string
		repo    *InvestorsRepository
		want    []*pb.Investor
		wantErr bool
	}{
		// TODO: Add test cases.
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
		want    []*pb.Issuer
		wantErr bool
	}{
		// TODO: Add test cases.
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

func TestInvoicesRepository_GetAllInvoices(t *testing.T) {
	tests := []struct {
		name    string
		repo    *InvoicesRepository
		want    []*pb.Invoice
		wantErr bool
	}{
		// TODO: Add test cases.
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
