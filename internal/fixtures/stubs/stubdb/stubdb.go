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
	return data.SeededAllInvestorsList, nil
}

func (db *StubDb) GetAllIssuers() (*[]*db.Account, error) {
	return data.SeededAllIssuersList, nil
}

func (db *StubDb) GetAllBids() (*[]*db.Bid, error) {
	return data.SeededAllBidsList, nil
}

func (db *StubDb) GetBidById(bidId int64) (*db.Bid, error) {
	data := data.SeededAllBidsList
	if bidId >= 0 && int(bidId) < len(*data) {
		return (*data)[bidId], nil
	} else {
		return nil, fmt.Errorf("bidId does not exist on the stubDb")
	}
}

func (db *StubDb) GetBidWithInvoiceById(bidId int64) (*db.BidWithInvoice, error) {
	data := data.SeededAllBidsWithInvoiceList
	if bidId >= 0 && int(bidId) < len(*data) {
		return (*data)[bidId], nil
	} else {
		return nil, fmt.Errorf("bidId does not exist on the stubDb")
	}
}

func (db *StubDb) GetBidsByInvoiceId(invoiceId string) (*[]*db.Bid, error) {
	return data.SeededAllBidsList, nil
}

func (db *StubDb) GetBidsByInvestorId(invoiceId string) (*[]*db.BidWithInvoice, error) {
	return data.SeededAllBidsWithInvoiceList, nil
}

func (db *StubDb) NewBid(db.Bid) (*db.Bid, error) {
	return data.NewBidData, nil
}

func (db *StubDb) GetFulfillingBids(invoiceId string) (*[]*db.Bid, error) {
	return data.SeededAllBidsList, nil
}

func (db *StubDb) AdjudicateBid(bidId int64) (*struct {
	BidId      int64
	PaidAmount int64
}, error) {
	return &struct {
		BidId      int64
		PaidAmount int64
	}{BidId: 4, PaidAmount: 200000}, nil
}

func (db *StubDb) AllRunningBidsToLost(invoiceId string) (*[]*db.Bid, error) {
	return data.SeededAllBidsList, nil
}

func (db *StubDb) GetAllInvoices() (*[]*db.Invoice, error) {
	return data.SeededAllInvoicesList, nil
}

func (db *StubDb) GetInvoiceById(invoiceId string) (*db.Invoice, error) {
	return (*data.SeededAllInvoicesList)[1], nil
}

func (db *StubDb) NewInvoice(db.Invoice) (*db.Invoice, error) {
	return data.NewInvoiceData, nil
}

var TestStubDb = new(StubDb)
