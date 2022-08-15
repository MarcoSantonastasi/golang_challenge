package db

type IAccount interface {
    Id string
    Name string
    Type string
    balance int64
}

type Account struct {
    Id string
    Name string
    Type string
    Balance int64
}


type IInvoice interface {
	Id string
	IssuerAccountId string
	Reference string
	Denom string
    Amount int64
	Asking int64
    State string
}

type Invoice struct {
	Id string
	IssuerAccountId string
	Reference string
	Denom string
    Amount int64
	Asking int64
    State string
}

type IBid interface {
	Id int64
	InvoiceId string
	BidderAccountId string
	Offer int64
	State string
}

type Bid struct {
	Id int64
	InvoiceId string
	BidderAccountId string
	Offer int64
	State string
}

type IDb interface {
	Connect() error
	Close() error
	GetAllInvestors() (*[]*Account, error)
	GetAllIssuers() (*[]*Account, error)
	GetAllBids() (*[]*Bid, error)
	GetAllInvoices() (*[]*Invoice, error)
	NewInvoice(IInvoice) (*Invoice, error)
	NewBid(IBid) (*Bid, error)
	AdjudicateInvoice(IInvoice) (*Invoice, error)
	AllRunningBidsToLost(IInvoice) (*[]*Bid, error)
}