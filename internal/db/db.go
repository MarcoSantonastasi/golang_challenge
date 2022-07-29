package db

import (
	pb "github.com/marcosantonastasi/arex_challenge/api/arex/v1"
)

type IDb interface {
	GetAllInvestors() []*pb.Investor
	GetAllIssuers() []*pb.Issuer
	GetAllInvoices() []*pb.Invoice
}

type Db struct {
}

func (db *Db) GetAllInvestors() []*pb.Investor {
	return nil
}

func (db *Db) GetAllIssuers() []*pb.Issuer {
	return nil
}
func (db *Db) GetAllInvoices() []*pb.Invoice {
	return nil
}
