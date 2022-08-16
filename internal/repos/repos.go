package repos

import (
	"fmt"

	db "github.com/marcosantonastasi/arex_challenge/internal/db"
)

type IInvestorsRepository interface {
	GetAllInvestors() (*[]*db.Account, error)
}

type InvestorsRepository struct {
	Db db.IDb
}

func (repo *InvestorsRepository) GetAllInvestors() (*[]*db.Account, error) {
	if repo.Db == nil {
		return nil, fmt.Errorf("no database found for Investors")
	}
	data, err := repo.Db.GetAllInvestors()
	if err != nil {
		return nil, fmt.Errorf("repository error for GetAllInvestors: %+v", err)
	}

	return data, nil
}

type IIssuersRepository interface {
	GetAllIssuers() (*[]*db.Account, error)
}

type IssuersRepository struct {
	Db db.IDb
}

func (repo *IssuersRepository) GetAllIssuers() (*[]*db.Account, error) {
	if repo.Db == nil {
		return nil, fmt.Errorf("no database found for Issuers")
	}
	data, err := repo.Db.GetAllIssuers()
	if err != nil {
		return nil, fmt.Errorf("repository error for GetAllIssuers: %+v", err)
	}

	return data, nil
}

type IInvoicesRepository interface {
	GetAllInvoices() (*[]*db.Invoice, error)
	GetInvoiceById(invoiceId string) (*db.Invoice, error)
	NewInvoice(db.Invoice) (*db.Invoice, error)
}

type InvoicesRepository struct {
	Db db.IDb
}

func (repo *InvoicesRepository) GetAllInvoices() (*[]*db.Invoice, error) {
	if repo.Db == nil {
		return nil, fmt.Errorf("no database found for Invoices")
	}
	data, err := repo.Db.GetAllInvoices()
	if err != nil {
		return nil, fmt.Errorf("repository error for GetAllInvoices: %+v", err)
	}

	return data, nil
}

func (repo *InvoicesRepository) GetInvoiceById(invoiceId string) (*db.Invoice, error) {
	if repo.Db == nil {
		return nil, fmt.Errorf("no database found for Invoices")
	}
	data, err := repo.Db.GetInvoiceById(invoiceId)
	if err != nil {
		return nil, fmt.Errorf("repository error for GetInvoiceById: %+v", err)
	}

	return data, nil
}

func (repo *InvoicesRepository) NewInvoice(newInvoiceData db.Invoice) (*db.Invoice, error) {
	if repo.Db == nil {
		return nil, fmt.Errorf("no database found for Invoices")
	}
	data, err := repo.Db.NewInvoice(newInvoiceData)
	if err != nil {
		return nil, fmt.Errorf("repository error for NewInvoice: %+v", err)
	}

	return data, nil
}

type IBidsRepository interface {
	GetAllBids() (*[]*db.Bid, error)
	GetBidById(bidId int64) (*db.Bid, error)
	GetBidsByInvoiceId(invoiceId string) (*[]*db.Bid, error)
	GetBidsByInvestorId(investorId string) (*[]*db.Bid, error)
	NewBid(db.Bid) (*db.Bid, error)
	GetFulfillingBids(invoiceId string) (*[]*db.Bid, error)
	AdjudicateBid(bidId int64) (*int64, error)
	AllRunningBidsToLost(invoiceId string) (*[]*db.Bid, error)
}

type BidsRepository struct {
	Db db.IDb
}

func (repo *BidsRepository) GetAllBids() (*[]*db.Bid, error) {
	if repo.Db == nil {
		return nil, fmt.Errorf("no database found for Bids")
	}
	data, err := repo.Db.GetAllBids()
	if err != nil {
		return nil, fmt.Errorf("repository error for GetAllBids: %+v", err)
	}

	return data, nil
}

func (repo *BidsRepository) GetBidById(bidId int64) (*db.Bid, error) {
	if repo.Db == nil {
		return nil, fmt.Errorf("no database found for Bids")
	}
	data, err := repo.Db.GetBidById(bidId)
	if err != nil {
		return nil, fmt.Errorf("repository error for GetBidById: %+v", err)
	}

	return data, nil
}

func (repo *BidsRepository) GetBidsByInvoiceId(invoiceId string) (*[]*db.Bid, error) {
	if repo.Db == nil {
		return nil, fmt.Errorf("no database found for Bids")
	}
	data, err := repo.Db.GetBidsByInvoiceId(invoiceId)
	if err != nil {
		return nil, fmt.Errorf("repository error for GetBidsByInvoiceId: %+v", err)
	}

	return data, nil
}

func (repo *BidsRepository) GetBidsByInvestorId(investorId string) (*[]*db.Bid, error) {
	if repo.Db == nil {
		return nil, fmt.Errorf("no database found for Bids")
	}
	data, err := repo.Db.GetBidsByInvestorId(investorId)
	if err != nil {
		return nil, fmt.Errorf("repository error for GetBidsByInvestorId: %+v", err)
	}

	return data, nil
}

func (repo *BidsRepository) NewBid(newBidData db.Bid) (*db.Bid, error) {
	if repo.Db == nil {
		return nil, fmt.Errorf("no database found for Bids")
	}
	data, err := repo.Db.NewBid(newBidData)

	if err != nil {
		return nil, fmt.Errorf("repository error for NewBid: %+v", err)
	}

	return data, nil
}

func (repo *BidsRepository) GetFulfillingBids(invoiceId string) (*[]*db.Bid, error) {
	if repo.Db == nil {
		return nil, fmt.Errorf("no database found for Bids")
	}
	data, err := repo.Db.GetFulfillingBids(invoiceId)

	if err != nil {
		return nil, fmt.Errorf("repository error for GetFulfillingBids: %+v", err)
	}

	return data, nil
}

func (repo *BidsRepository) AdjudicateBid(bidId int64) (*int64, error) {
	if repo.Db == nil {
		return nil, fmt.Errorf("no database found for Invoices")
	}
	data, err := repo.Db.AdjudicateBid(bidId)

	if err != nil {
		return nil, fmt.Errorf("repository error for AdjudicateInvoice: %+v", err)
	}

	return data, nil
}

func (repo *BidsRepository) AllRunningBidsToLost(invoiceId string) (*[]*db.Bid, error) {
	if repo.Db == nil {
		return nil, fmt.Errorf("no database found for Bids")
	}
	data, err := repo.Db.AllRunningBidsToLost(invoiceId)
	if err != nil {
		return nil, fmt.Errorf("repository error for AllRunningBidsToLost: %+v", err)
	}

	return data, nil
}
