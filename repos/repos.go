package repos

import (
	pb "github.com/marcosantonastasi/arex_challenge/api/arex/v1"
)

type IInvestorsRepository interface {
	GetAllInvestors() ([]*pb.Investor, error)
}

type InvestorsRepository struct {
}

func (repo *InvestorsRepository) GetAllInvestors() ([]*pb.Investor, error) {
	return nil, nil
}

type IIssuersRepository interface {
}

type IssuersRepository struct {
}

func (repo *IssuersRepository) GetAllIssuers() ([]*pb.Issuer, error) {
	return nil, nil
}

type IInvoiceRepository interface {
}

type InvoiceRepository struct {
}

func (repo *InvoiceRepository) GetAllInvoices() ([]*pb.Invoice, error) {
	return nil, nil
}
