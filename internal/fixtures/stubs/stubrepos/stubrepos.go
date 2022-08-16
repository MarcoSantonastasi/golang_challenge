package stubrepos

import (
	pb "github.com/marcosantonastasi/arex_challenge/api/arex/v1"
	"github.com/marcosantonastasi/arex_challenge/internal/db"
	data "github.com/marcosantonastasi/arex_challenge/internal/fixtures/data"
)

type FakeInvestorsRepository struct {
}

func (repo *FakeInvestorsRepository) GetAllInvestors() (*[]*db.Account, error) {
	return data.FakeAllInvestorsList, nil
}

type FakeIssuersRepository struct {
}

func (repo *FakeIssuersRepository) GetAllIssuers() (*[]*db.Account, error) {
	return data.FakeAllIssuersList, nil
}

type FakeInvoicesRepository struct {
}

func (repo *FakeInvoicesRepository) GetAllInvoices() (*[]*db.Invoice, error) {
	return data.FakeAllInvoicesList, nil
}

func (repo *FakeInvoicesRepository) NewInvoice(*pb.NewInvoiceRequest) (*db.Invoice, error) {
	return data.NewInvoiceData, nil
}

type FakeBidsRepository struct {
}

func (repo *FakeBidsRepository) GetAllBids() (*[]*db.Bid, error) {
	return data.FakeAllBidsList, nil
}

func (repo *FakeBidsRepository) NewBid(*pb.NewBidRequest) (*db.Bid, error) {
	return data.NewBidData, nil
}
