package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgtype"
	pg "github.com/jackc/pgx/v4"
	pb "github.com/marcosantonastasi/arex_challenge/api/arex/v1"
)

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
		fmt.Fprintf(os.Stderr, "unable to connect to database: %v", err)
		return err
	}
	db.conn = conn
	return nil
}

func (db *PgDb) Close() error {
	err := db.conn.Close(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to close the database: %v\n", err)
		return err
	}
	db.conn = nil
	return nil
}

func (db *PgDb) GetAllInvestors() (*[]*Account, error) {
	data := new([]*Account)
	rows, err := db.conn.Query(context.Background(), "select id, name, type, balance from investors")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed query for GetAllInvestors: %+v", err)
		return nil, err
	}
	for rows.Next() {
		row := new(Account)
		if err := rows.Scan(&row.Id, &row.Name, &row.Type, &row.Balance); err != nil {
			fmt.Printf(os.Stderr, "error scanning query response for GetAllInvestors: %+v", err)
			return nil, err
		}
		*data = append(*data, row)
	}
	//defer rows.Close()
	rows.Close()
	if err := rows.Err(); err != nil {
			fmt.Printf(os.Stderr, "error in some rows of query response for GetAllInvestors: %+v", err)
		return nil, err
	}
	return data, nil
}

func (db *PgDb) GetAllIssuers() (*[]*Account, error) {
	data := new([]*Account)
	rows, err := db.conn.Query(context.Background(), "select id, name, type, balance from issuers")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed query for GetAllIssuers: %+v", err)
		return nil, err
	}
	for rows.Next() {
		row := new(Account)
		if err := rows.Scan(&row.Id, &row.Name, &row.Type, &row.Balance); err != nil {
			fmt.Printf(os.Stderr, "error scanning query response for GetAllIssuers: %+v", err)
			return nil, err
		}
		*data = append(*data, row)
	}
	// defer rows.Close()
	rows.Close()
	if err := rows.Err(); err != nil {
			fmt.Printf(os.Stderr, "error in some rows of query response for GetAllIssuers: %+v", err)
		return nil, err
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
		fmt.Fprintf(os.Stderr, "failed query for GetAllBids: %+v", err)
		return nil, err
	}
	for rows.Next() {
		row := new(pb.Bid)
		if err := rows.Scan(
			&row.Id,
			&row.InvoiceId,
			&row.BidderAccountId,
			&row.Offer,
			&row.State);
		err != nil {
			fmt.Printf(os.Stderr, "error scanning query response for GetAllBids: %+v", err)
			return nil, err
		}

		*data = append(*data, row)
	}
	//defer rows.Close()
	rows.Close()
	if err := rows.Err(); err != nil {
		fmt.Printf(os.Stderr, "error in some rows of query response for GetAllIssuers: %+v", err)
		return nil, err
	}
	return data, nil
}

func (db *PgDb) NewBid(newBid IBid) (*Bid, error) {
	var buff pgtype.Record
	data := new(Bid)

	row := db.conn.QueryRow(
		context.Background(),
		"select bid($1, $2, $3)",
		newBid.InvoiceId,
		newBid.BidderAccountId,
		newBid.Offer,
	)
	if err := row.Scan(&buff); err != nil {
		fmt.Printf("error scanning db response record field for NewBid: %+v", err)
		return nil, err
	}

	if err := buff.Fields[0].AssignTo(&data.Id); err != nil {
		fmt.Printf("error assigning db response record field for NewBid: %+v", err)
		return nil, err
	}
	if err := buff.Fields[1].AssignTo(&data.InvoiceId); err != nil {
		fmt.Printf("error assigning db response record field for NewBid: %+v", err)
		return nil, err
	}
	if err := buff.Fields[2].AssignTo(&data.BidderAccountId); err != nil {
		fmt.Printf("error assigning db response record field for NewBid: %+v", err)
		return nil, err
	}
	if err := buff.Fields[3].AssignTo(&data.Offer); err != nil {
		fmt.Printf("error assigning db response record field for NewBid: %+v", err)
		return nil, err
	}
	if err := buff.Fields[4].AssignTo(&data.State); err != nil {
		fmt.Printf("error assigning db response record field for NewBid: %+v", err)
		return nil, err
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
		fmt.Fprintf(os.Stderr, "failed query for GetAllInvoices: %+v", err)
		return nil, err
	}
	for rows.Next() {
		row := new(pb.Invoice)
		if err := rows.Scan(&row.Id, &row.IssuerAccountId, &row.Reference, &row.Denom, &row.Amount, &row.Asking, &row.State); err != nil {
			fmt.Printf(os.Stderr, "error scanning query response for GetAllInvoices: %+v", err)
			return nil, err
		}
		*data = append(*data, row)
	}
	// defer rows.Close()
	rows.Close()
	if err := rows.Err(); err != nil {
		fmt.Printf(os.Stderr, "error in some rows of query response for GetAllIssuers: %+v", err)
		return nil, err
	}
	return data, nil
}

func (db *PgDb) NewInvoice(newInvoiceData IInvoice) (*Invoice, error) {
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
		&data.Asking);
	err != nil {
		fmt.Printf(os.Stderr, "error scanning query response for NewInvoice: %+v", err)
		return nil, err
	}

	return data, nil
}

func (db *PgDb) AdjudicateInvoice(invoiceData IInvoice) (*Invoice, error) {
	data := new(Invoice)
	row := db.conn.QueryRow(
		context.Background(),
		"select adjudicate($1)",
		invoiceData.Id,
	)

	if err := row.Scan(&data.invoiceId, &data.bidderAccountId, &data.amount); err != nil {
		fmt.Printf("error scanning db response record field for AdjudicateInvoice: %+v", err)
		return nil, err
	}

	return data, nil
}

func (db *PgDb) AllRunningBidsToLost(invoiceData IInvoice) (*[]*Bid, error) {
	data := new(*Bid)

	row := db.conn.QueryRow(
		context.Background(),
		"select adjudicate($1)",
		invoiceData.Id,
	)

	if err := row.Scan(&data.invoiceId, &data.bidId); err != nil {
		fmt.Printf("error scanning db response record field for AllRunningBidsToLost: %+v", err)
		return nil, err
	}
	return data, nil
}

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
