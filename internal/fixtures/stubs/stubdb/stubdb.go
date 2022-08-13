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

func (db *StubDb) GetAllInvoices() *[]*pb.Invoice {
	return data.FakeAllInvoicesList
}

func (db *StubDb) NewInvoice(*pb.Invoice) *pb.Invoice {
	return data.NewInvoiceData
}

func (db *StubDb) Bid(invoiceId string) any {
	return struct {
		invoiceId       string
		bidderAccountId string
		offer           int64
	}{
		invoiceId:       "",
		bidderAccountId: "",
		offer:           0,
	}
}

func (db *StubDb) Adjudicate(invoiceId string) any {
	return struct {
		invoiceId       string
		bidderAccountId string
		amount          int64
	}{
		invoiceId:       "",
		bidderAccountId: "",
		amount:          0,
	}
}

func (db *StubDb) AllRunningBidsToLost(invoiceId string) any {
	return struct {
		invoiceId string
		bidId     string
	}{
		invoiceId: "",
		bidId:     "",
	}
}

var TestStubDb = new(StubDb)
