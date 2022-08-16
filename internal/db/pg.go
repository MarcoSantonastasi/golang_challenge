package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgtype"
	pg "github.com/jackc/pgx/v4"
)

func NewPgDb(
	pgUser string,
	pgPwd string,
	pgHostname string,
	pgDbname string,
) *PgDb {
	return &PgDb{
		pgUser:     pgUser,
		pgPwd:      pgPwd,
		pgHostname: pgHostname,
		pgDbname:   pgDbname,
	}
}

type PgDb struct {
	pgUser     string
	pgPwd      string
	pgHostname string
	pgDbname   string
	conn       *pg.Conn
}

func (db *PgDb) Connect() error {
	conn, err := pg.Connect(context.Background(), fmt.Sprintf("postgres://%s:%s@%s:5432/%s", db.pgUser, db.pgPwd, db.pgHostname, db.pgDbname))
	if err != nil {
		return fmt.Errorf("unable to connect to database: %v", err)
	}
	db.conn = conn
	return nil
}

func (db *PgDb) Close() error {
	err := db.conn.Close(context.Background())
	if err != nil {
		return fmt.Errorf("unable to close the database: %v", err)
	}
	db.conn = nil
	return nil
}

func (db *PgDb) GetAllInvestors() (*[]*Account, error) {
	data := new([]*Account)
	rows, err := db.conn.Query(context.Background(), "select id, name, type, balance from investors")
	if err != nil {
		return nil, fmt.Errorf("failed query for GetAllInvestors: %+v", err)
	}
	for rows.Next() {
		row := new(Account)
		if err := rows.Scan(&row.Id, &row.Name, &row.Type, &row.Balance); err != nil {
			return nil, fmt.Errorf("error scanning query response for GetAllInvestors: %+v", err)
		}
		*data = append(*data, row)
	}
	//defer rows.Close()
	rows.Close()
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error in some rows of query response for GetAllInvestors: %+v", err)
	}
	return data, nil
}

func (db *PgDb) GetAllIssuers() (*[]*Account, error) {
	data := new([]*Account)
	rows, err := db.conn.Query(
		context.Background(),
		"select id, name, type, balance from issuers",
	)
	if err != nil {
		return nil, fmt.Errorf("failed query for GetAllIssuers: %+v", err)
	}
	for rows.Next() {
		row := new(Account)
		if err := rows.Scan(&row.Id, &row.Name, &row.Type, &row.Balance); err != nil {
			return nil, fmt.Errorf("error scanning query response for GetAllIssuers: %+v", err)
		}
		*data = append(*data, row)
	}
	// defer rows.Close()
	rows.Close()
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error in some rows of query response for GetAllIssuers: %+v", err)
	}

	return data, nil
}

func (db *PgDb) GetAllInvoices() (*[]*Invoice, error) {
	data := new([]*Invoice)
	rows, err := db.conn.Query(
		context.Background(),
		`select
			id,
			issuer_account_id,
			reference,
			denom,
			amount,
			asking,
			state
		from
			invoices`,
	)
	if err != nil {
		return nil, fmt.Errorf("failed query for GetAllInvoices: %+v", err)
	}
	for rows.Next() {
		row := new(Invoice)
		if err := rows.Scan(&row.Id, &row.IssuerAccountId, &row.Reference, &row.Denom, &row.Amount, &row.Asking, &row.State); err != nil {
			return nil, fmt.Errorf("error scanning query response for GetAllInvoices: %+v", err)
		}
		*data = append(*data, row)
	}
	// defer rows.Close()
	rows.Close()
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error in some rows of query response for GetAllIssuers: %+v", err)
	}

	return data, nil
}

func (db *PgDb) GetInvoiceById(invoiceId string) (*Invoice, error) {
	data := new(Invoice)

	row := db.conn.QueryRow(
		context.Background(),
		"select * from invoices where id = $1",
		invoiceId,
	)
	if err := row.Scan(&data); err != nil {
		return nil, fmt.Errorf("error scanning db response record field for GetInvoiceById: %+v", err)
	}
	return data, nil
}

func (db *PgDb) NewInvoice(newInvoiceData Invoice) (*Invoice, error) {
	data := new(Invoice)
	row := db.conn.QueryRow(
		context.Background(),
		`insert into invoices (
			issuer_account_id,
			reference,
			denom,
			amount,
			asking
		)
		values($1, $2, $3, $4, $5)
		returning
		    id,
			issuer_account_id,
			reference,
			denom,
			amount,
			asking`,
		newInvoiceData.IssuerAccountId,
		newInvoiceData.Reference,
		newInvoiceData.Denom,
		newInvoiceData.Amount,
		newInvoiceData.Asking,
	)

	if err := row.Scan(
		&data.Id,
		&data.IssuerAccountId,
		&data.Reference,
		&data.Denom,
		&data.Amount,
		&data.Asking); err != nil {

		return nil, fmt.Errorf("error scanning query response for NewInvoice: %+v", err)
	}

	return data, nil
}

func (db *PgDb) GetAllBids() (*[]*Bid, error) {
	data := new([]*Bid)
	rows, err := db.conn.Query(
		context.Background(),
		`select
			id,
			invoice_id,
			bidder_account_id,
			offer,
			state
		from
		    bids`,
	)
	if err != nil {
		return nil, fmt.Errorf("failed query for GetAllBids: %+v", err)
	}
	for rows.Next() {
		row := new(Bid)
		if err := rows.Scan(
			&row.Id,
			&row.InvoiceId,
			&row.BidderAccountId,
			&row.Offer,
			&row.State); err != nil {
			return nil, fmt.Errorf("error scanning query response for GetAllBids: %+v", err)
		}
		*data = append(*data, row)
	}
	//defer rows.Close()
	rows.Close()
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error in some rows of query response for GetAllBids: %+v", err)
	}
	return data, nil
}

func (db *PgDb) GetBidById(bidId int64) (*Bid, error) {
	data := new(Bid)

	row := db.conn.QueryRow(
		context.Background(),
		"select * from bids_with_invoice where id = $1",
		bidId,
	)
	if err := row.Scan(&data); err != nil {
		return nil, fmt.Errorf("error scanning db response record field for GetBidById: %+v", err)
	}
	return data, nil
}

func (db *PgDb) GetBidsByInvoiceId(invoiceId string) (*[]*Bid, error) {
	data := new([]*Bid)
	rows, err := db.conn.Query(
		context.Background(),
		`select
			id,
			invoice_id,
			bidder_account_id,
			offer,
			state
		from
		    bids
		where
			invoice_id = $1`,
		invoiceId,
	)
	if err != nil {
		return nil, fmt.Errorf("failed query for GetBidsByInvoiceId: %+v", err)
	}
	for rows.Next() {
		row := new(Bid)
		if err := rows.Scan(
			&row.Id,
			&row.InvoiceId,
			&row.BidderAccountId,
			&row.Offer,
			&row.State); err != nil {
			return nil, fmt.Errorf("error scanning query response for GetBidsByInvoiceId: %+v", err)
		}
		*data = append(*data, row)
	}
	//defer rows.Close()
	rows.Close()
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error in some rows of query response for GetBidsByInvoiceId: %+v", err)
	}
	return data, nil
}

func (db *PgDb) GetBidsByInvestorId(investorId string) (*[]*Bid, error) {
	data := new([]*Bid)
	rows, err := db.conn.Query(
		context.Background(),
		`select
			id,
			invoice_id,
			bidder_account_id,
			offer,
			state,
			invoice_issuer_account_id,
			invoice_reference,
			invoice_denom,
			invoice_amount,
			invoice_asking,
			invoice_state 
		from
		    bids_with_invoice
		where
			bidder_account_id = $1`,
		investorId,
	)
	if err != nil {
		return nil, fmt.Errorf("failed query for GetBidsByInvestorId: %+v", err)
	}
	for rows.Next() {
		row := new(Bid)
		if err := rows.Scan(
			&row.Id,
			&row.InvoiceId,
			&row.BidderAccountId,
			&row.Offer,
			&row.State); err != nil {
			return nil, fmt.Errorf("error scanning query response for GetBidsByInvestorId: %+v", err)
		}
		*data = append(*data, row)
	}
	//defer rows.Close()
	rows.Close()
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error in some rows of query response for GetBidsByInvestorId: %+v", err)
	}
	return data, nil
}

func (db *PgDb) NewBid(newBidData Bid) (*Bid, error) {
	var buff pgtype.Record
	data := new(Bid)

	row := db.conn.QueryRow(
		context.Background(),
		"select bid($1, $2, $3)",
		newBidData.InvoiceId,
		newBidData.BidderAccountId,
		newBidData.Offer,
	)
	if err := row.Scan(&buff); err != nil {
		return nil, fmt.Errorf("error scanning db response record field for NewBid: %+v", err)
	}

	if err := buff.Fields[0].AssignTo(&data.Id); err != nil {
		return nil, fmt.Errorf("error scanning db response record field for NewBid: %+v", err)
	}
	if err := buff.Fields[1].AssignTo(&data.InvoiceId); err != nil {
		return nil, fmt.Errorf("error scanning db response record field for NewBid: %+v", err)
	}
	if err := buff.Fields[2].AssignTo(&data.BidderAccountId); err != nil {
		return nil, fmt.Errorf("error scanning db response record field for NewBid: %+v", err)
	}
	if err := buff.Fields[3].AssignTo(&data.Offer); err != nil {
		return nil, fmt.Errorf("error scanning db response record field for NewBid: %+v", err)
	}
	if err := buff.Fields[4].AssignTo(&data.State); err != nil {
		return nil, fmt.Errorf("error scanning db response record field for NewBid: %+v", err)
	}

	return data, nil
}

func (db *PgDb) GetFulfillingBids(invoiceId string) (*[]*Bid, error) {
	data := new([]*Bid)
	rows, err := db.conn.Query(
		context.Background(),
		`select
			id,
			invoice_id,
			bidder_account_id,
			offer,
			state,
			invoice_issuer_account_id,
			invoice_reference,
			invoice_denom,
			invoice_amount,
			invoice_asking,
			invoice_state 
		from
		    bids_with_invoice
		where
			invoice_id = $1 AND
			offer >= invoice_asking`,
		invoiceId,
	)
	if err != nil {
		return nil, fmt.Errorf("failed query for GetFulfillingBids: %+v", err)
	}
	for rows.Next() {
		row := new(Bid)
		if err := rows.Scan(
			&row.Id,
			&row.InvoiceId,
			&row.BidderAccountId,
			&row.Offer,
			&row.State); err != nil {
			return nil, fmt.Errorf("error scanning query response for GetFulfillingBids: %+v", err)
		}
		*data = append(*data, row)
	}
	//defer rows.Close()
	rows.Close()
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error in some rows of query response for GetFulfillingBids: %+v", err)
	}
	return data, nil
}

func (db *PgDb) AdjudicateBid(bidId int64) (*int64, error) {
	saleAmount := new(int64)
	row := db.conn.QueryRow(
		context.Background(),
		"select adjudicate_bid($1)",
		bidId,
	)
	if err := row.Scan(&saleAmount); err != nil {
		return nil, fmt.Errorf("error scanning db response record field for AdjudicateInvoice: %+v", err)
	}

	return saleAmount, nil
}

func (db *PgDb) AllRunningBidsToLost(invoiceId string) (*[]*Bid, error) {
	data := new([]*Bid)
	rows, err := db.conn.Query(
		context.Background(),
		"select all_running_bids_to_lost($1)",
		invoiceId,
	)
	if err != nil {
		return nil, fmt.Errorf("failed query for AllRunningBidsToLost: %+v", err)
	}
	for rows.Next() {
		row := new(Bid)
		if err := rows.Scan(
			&row.Id,
			&row.InvoiceId,
			&row.BidderAccountId,
			&row.Offer,
			&row.State,
		); err != nil {
			return nil, fmt.Errorf("error scanning query response for AllRunningBidsToLost: %+v", err)
		}
		*data = append(*data, row)
	}
	// defer rows.Close()
	rows.Close()
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error in some rows of query response for AllRunningBidsToLost: %+v", err)
	}
	return data, nil
}
