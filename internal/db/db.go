package db

type IDb interface {
	Connect() error
	Close() error

	GetAllInvestors() (*[]*Account, error)

	GetAllIssuers() (*[]*Account, error)

	GetAllBids() (*[]*Bid, error)
	GetBidById(bidId int64) (*Bid, error)
	GetBidsByInvoiceId(invoiceId string) (*[]*Bid, error)
	GetBidsByInvestorId(investorId string) (*[]*Bid, error)
	NewBid(Bid) (*Bid, error)
	GetFulfillingBids(invoiceId string) (*[]*Bid, error)
	AdjudicateBid(bidId int64) (*int64, error)
	AllRunningBidsToLost(invoiceId string) (*[]*Bid, error)

	GetAllInvoices() (*[]*Invoice, error)
	GetInvoiceById(invoiceId string) (*Invoice, error)
	NewInvoice(Invoice) (*Invoice, error)
}
