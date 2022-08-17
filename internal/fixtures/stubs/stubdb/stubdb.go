package stubdb

import (
	"fmt"

	db "github.com/marcosantonastasi/arex_challenge/internal/db"
	data "github.com/marcosantonastasi/arex_challenge/internal/fixtures/data"
)

type StubDb struct {
}

func (db *StubDb) Connect() error {
	return nil
}

func (db *StubDb) Close() error {
	return nil
}

func (db *StubDb) GetAllInvestors() (*[]*db.Account, error) {
	return data.FakeAllInvestorsList, nil
}

func (db *StubDb) GetAllIssuers() (*[]*db.Account, error) {
	return data.FakeAllIssuersList, nil
}

func (db *StubDb) GetAllBids() (*[]*db.Bid, error) {
	return data.FakeAllBidsList, nil
}

func (db *StubDb) GetBidById(bidId int64) (*db.Bid, error) {
	data := data.FakeAllBidsList
	if bidId >= 0 && int(bidId) < len(*data) {
		return (*data)[bidId], nil
	} else {
		return nil, fmt.Errorf("bidId does not exist on the stubDb")
	}
}

func (db *StubDb) GetBidsByInvoiceId(invoiceId string) (*[]*db.Bid, error) {
	return data.FakeAllBidsList, nil
}

func (db *StubDb) GetBidsByInvestorId(invoiceId string) (*[]*db.Bid, error) {
	return data.FakeAllBidsList, nil
}

func (db *StubDb) NewBid(db.Bid) (*db.Bid, error) {
	return data.NewBidData, nil
}

func (db *StubDb) GetFulfillingBids(invoiceId string) (*[]*db.Bid, error) {
	return data.FakeAllBidsList, nil
}

func (db *StubDb) AdjudicateBid(bidId int64) (*int64, error) {
	var paidAmount int64 = 300000
	return &paidAmount, nil
}

func (db *StubDb) AllRunningBidsToLost(invoiceId string) (*[]*db.Bid, error) {
	return data.FakeAllBidsList, nil
}

func (db *StubDb) GetAllInvoices() (*[]*db.Invoice, error) {
	return data.FakeAllInvoicesList, nil
}

func (db *StubDb) GetInvoiceById(invoiceId string) (*db.Invoice, error) {
	data := data.FakeAllInvoicesList
	return (*data)[1], nil
}

func (db *StubDb) NewInvoice(db.Invoice) (*db.Invoice, error) {
	return data.NewInvoiceData, nil
}

var TestStubDb = new(StubDb)
