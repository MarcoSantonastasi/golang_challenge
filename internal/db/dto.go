package db

type Account struct {
	Id      string
	Name    string
	Type    string
	Balance int64
}

type Invoice struct {
	Id              string
	IssuerAccountId string
	Reference       string
	Denom           string
	Amount          int64
	Asking          int64
	State           string
}

type Bid struct {
	Id              int64
	InvoiceId       string
	BidderAccountId string
	Offer           int64
	State           string
}

type BidWithInvoice struct {
	Id                     int64
	InvoiceId              string
	BidderAccountId        string
	Offer                  int64
	State                  string
	InvoiceIssuerAccountId string
	InvoiceReference       string
	InvoiceDenom           string
	InvoiceAmount          int64
	InvoiceAsking          int64
	InvoiceState           string
}
