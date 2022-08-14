package repos

import (
	pb "github.com/marcosantonastasi/arex_challenge/api/arex/v1"
	db "github.com/marcosantonastasi/arex_challenge/internal/db"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type IInvestorsRepository interface {
	GetAllInvestors() (*[]*pb.Investor, error)
}

type InvestorsRepository struct {
	Db db.IDb
}

func (repo *InvestorsRepository) GetAllInvestors() (*[]*pb.Investor, error) {
	if repo.Db == nil {
		return nil, status.Error(codes.Internal, "no database found for Investors")
	}
	data := repo.Db.GetAllInvestors()
	return data, nil
}

type IIssuersRepository interface {
	GetAllIssuers() (*[]*pb.Issuer, error)
}

type IssuersRepository struct {
	Db db.IDb
}

func (repo *IssuersRepository) GetAllIssuers() (*[]*pb.Issuer, error) {
	if repo.Db == nil {
		return nil, status.Error(codes.Internal, "no database found for Issuers")
	}
	data := repo.Db.GetAllIssuers()
	return data, nil
}

type IBidsRepository interface {
	GetAllBids() (*[]*pb.Bid, error)
	NewBid(*pb.NewBidRequest) (*pb.Bid, error)
}

type BidsRepository struct {
	Db db.IDb
}

func (repo *BidsRepository) GetAllBids() (*[]*pb.Bid, error) {
	if repo.Db == nil {
		return nil, status.Error(codes.Internal, "no database found for Bids")
	}
	data := repo.Db.GetAllBids()
	return data, nil
}

func (repo *BidsRepository) NewBid(newBid *pb.NewBidRequest) (*pb.Bid, error) {
	if repo.Db == nil {
		return nil, status.Error(codes.Internal, "no database found for Bids")
	}
	data := repo.Db.NewBid(newBid)
	return data, nil
}

type IInvoicesRepository interface {
	GetAllInvoices() (*[]*pb.Invoice, error)
	NewInvoice(*pb.NewInvoiceRequest) (*pb.Invoice, error)
}

type InvoicesRepository struct {
	Db db.IDb
}

func (repo *InvoicesRepository) GetAllInvoices() (*[]*pb.Invoice, error) {
	if repo.Db == nil {
		return nil, status.Error(codes.Internal, "no database found for Invoices")
	}
	data := repo.Db.GetAllInvoices()
	return data, nil
}

func (repo *InvoicesRepository) NewInvoice(newInvoice *pb.NewInvoiceRequest) (*pb.Invoice, error) {
	if repo.Db == nil {
		return nil, status.Error(codes.Internal, "no database found for Invoices")
	}
	data := repo.Db.NewInvoice(newInvoice)
	return data, nil
}
