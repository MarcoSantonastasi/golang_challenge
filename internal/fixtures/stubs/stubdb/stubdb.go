package stubdb

import (
	pb "github.com/marcosantonastasi/arex_challenge/api/arex/v1"
	data "github.com/marcosantonastasi/arex_challenge/internal/fixtures/data"
)

type StubDb struct {
}

func (db *StubDb) Connect() {}

func (db *StubDb) Close() {}

func (db *StubDb) GetAllInvestors() *[]*pb.Investor {
	return data.FakeAllInvestorsList
}

func (db *StubDb) GetAllIssuers() *[]*pb.Issuer {
	return data.FakeAllIssuersList
}

func (db *StubDb) GetAllBids() *[]*pb.Bid {
	return data.FakeAllBidsList
}

func (db *StubDb) NewBid(*pb.NewBidRequest) (*pb.Bid, error) {
	return data.NewBidData, nil
}

func (db *StubDb) GetAllInvoices() *[]*pb.Invoice {
	return data.FakeAllInvoicesList
}

func (db *StubDb) NewInvoice(*pb.NewInvoiceRequest) (*pb.Invoice, error) {
	return data.NewInvoiceData, nil
}

func (db *StubDb) Adjudicate(invoiceId string) any {
	return struct {
		invoiceId       string
		bidderAccountId string
		amount          int64
	}{
		invoiceId:       "af80d0ea-78b9-45b1-a7b0-d1ddd0fbd6fe",
		bidderAccountId: "991842fe-2e97-4481-a560-8d985a82ae74",
		amount:          420000,
	}
}

func (db *StubDb) AllRunningBidsToLost(invoiceId string) any {
	return struct {
		id                int64
		invoice_id        string
		bidder_account_id string
		offer             int64
	}{
		id:                4,
		invoice_id:        "ceeaece4-ca5c-4d31-9fd6-90a90854fed9",
		bidder_account_id: "991842fe-2e97-4481-a560-8d985a82ae74",
		offer:             200001,
	}
}

var TestStubDb = new(StubDb)
