package stubrepos

import (
	pb "github.com/marcosantonastasi/arex_challenge/api/arex/v1"
	"github.com/marcosantonastasi/arex_challenge/internal/db"
	data "github.com/marcosantonastasi/arex_challenge/internal/fixtures/data"
)

type StubInvestorsRepository struct {
}

func (repo *StubInvestorsRepository) GetAllInvestors() (*[]*db.Account, error) {
	return data.SeededAllInvestorsList, nil
}

type StubIssuersRepository struct {
}

func (repo *StubIssuersRepository) GetAllIssuers() (*[]*db.Account, error) {
	return data.SeededAllIssuersList, nil
}

type StubInvoicesRepository struct {
}

func (repo *StubInvoicesRepository) GetAllInvoices() (*[]*db.Invoice, error) {
	return data.SeededAllInvoicesList, nil
}

func (repo *StubInvoicesRepository) GetInvoiceById(invoiceId string) (*db.Invoice, error) {
	return (*data.SeededAllInvoicesList)[1], nil
}

func (repo *StubInvoicesRepository) NewInvoice(*db.Invoice) (*db.Invoice, error) {
	return data.NewInvoiceData, nil
}

type StubBidsRepository struct {
}

func (repo *StubBidsRepository) GetAllBids() (*[]*db.Bid, error) {
	return data.SeededAllBidsList, nil
}

func (repo *StubBidsRepository) NewBid(*pb.NewBidRequest) (*db.Bid, error) {
	return data.NewBidData, nil
}
