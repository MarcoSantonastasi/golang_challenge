package db

type IDb interface {
	Connect() error
	Close() error

	GetAllInvestors() (*[]*Account, error)

	GetAllIssuers() (*[]*Account, error)

	GetAllBids() (*[]*Bid, error)
	GetBidById(bidId int64) (*Bid, error)
	GetBidWithInvoiceById(bidId int64) (*BidWithInvoice, error)
	GetBidsByInvoiceId(invoiceId string) (*[]*Bid, error)
	GetBidsByInvestorId(investorId string) (*[]*BidWithInvoice, error)
	NewBid(Bid) (*Bid, error)
	GetFulfillingBids(invoiceId string) (*[]*Bid, error)
	AdjudicateBid(bidId int64) (*struct {
		BidId      int64
		PaidAmount int64
	}, error)
	AllRunningBidsToLost(invoiceId string) (*[]*Bid, error)

	GetAllInvoices() (*[]*Invoice, error)
	GetInvoiceById(invoiceId string) (*Invoice, error)
	NewInvoice(Invoice) (*Invoice, error)
}
