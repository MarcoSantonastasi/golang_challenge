package repos

import (
	pb "github.com/marcosantonastasi/arex_challenge/api/arex/v1"
	db "github.com/marcosantonastasi/arex_challenge/internal/db"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type IInvestorsRepository interface {
	GetAllInvestors() ([]*pb.Investor, error)
}

type InvestorsRepository struct {
	Db db.IDb
}

func (repo *InvestorsRepository) GetAllInvestors() ([]*pb.Investor, error) {
	if repo.Db == nil {
		return nil, status.Error(codes.Internal, "no database found for Investors")
	}
	return nil, nil
}

type IIssuersRepository interface {
	GetAllIssuers() ([]*pb.Issuer, error)
}

type IssuersRepository struct {
	Db db.IDb
}

func (repo *IssuersRepository) GetAllIssuers() ([]*pb.Issuer, error) {
	if repo.Db == nil {
		return nil, status.Error(codes.Internal, "no database found for Issuers")
	}
	return nil, nil
}

type IInvoicesRepository interface {
	GetAllInvoices() ([]*pb.Invoice, error)
}

type InvoicesRepository struct {
	Db db.IDb
}

func (repo *InvoicesRepository) GetAllInvoices() ([]*pb.Invoice, error) {
	if repo.Db == nil {
		return nil, status.Error(codes.Internal, "no database found for Invoices")
	}
	return nil, nil
}
