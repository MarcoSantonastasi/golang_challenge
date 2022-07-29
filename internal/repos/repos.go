package repos

import (
	pb "github.com/marcosantonastasi/arex_challenge/api/arex/v1"
	db "github.com/marcosantonastasi/arex_challenge/internal/db"
)

type IInvestorsRepository interface {
	GetAllInvestors() ([]*pb.Investor, error)
}

type InvestorsRepository struct {
	Db db.IDb
}

func (repo *InvestorsRepository) GetAllInvestors() ([]*pb.Investor, error) {
	return nil, nil
}

type IIssuersRepository interface {
	GetAllIssuers() ([]*pb.Issuer, error)
}

type IssuersRepository struct {
	Db db.IDb
}

func (repo *IssuersRepository) GetAllIssuers() ([]*pb.Issuer, error) {
	return nil, nil
}

type IInvoicesRepository interface {
	GetAllInvoices() ([]*pb.Invoice, error)
}

type InvoicesRepository struct {
	Db db.IDb
}

func (repo *InvoicesRepository) GetAllInvoices() ([]*pb.Invoice, error) {
	return nil, nil
}
