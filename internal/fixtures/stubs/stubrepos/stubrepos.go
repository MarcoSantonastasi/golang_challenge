package stubrepos

import (
	pb "github.com/marcosantonastasi/arex_challenge/api/arex/v1"
	data "github.com/marcosantonastasi/arex_challenge/internal/fixtures/data"
)

type FakeInvestorsRepository struct {
}

func (repo *FakeInvestorsRepository) GetAllInvestors() (*[]*pb.Investor, error) {
	return data.FakeAllInvestorsList, nil
}

type FakeIssuersRepository struct {
}

func (repo *FakeIssuersRepository) GetAllIssuers() (*[]*pb.Issuer, error) {
	return data.FakeAllIssuersList, nil
}

type FakeInvoicesRepository struct {
}

func (repo *FakeInvoicesRepository) GetAllInvoices() (*[]*pb.Invoice, error) {
	return data.FakeAllInvoicesList, nil
}
