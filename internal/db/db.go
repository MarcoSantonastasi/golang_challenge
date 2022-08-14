package db

import (
	"context"
	"fmt"
	"os"

	pg "github.com/jackc/pgx/v4"
	pb "github.com/marcosantonastasi/arex_challenge/api/arex/v1"
)

type IDb interface {
	Connect()
	Close()
	GetAllInvestors() *[]*pb.Investor
	GetAllIssuers() *[]*pb.Issuer
	GetAllBids() *[]*pb.Bid
	GetAllInvoices() *[]*pb.Invoice
	NewInvoice(*pb.Invoice) *pb.Invoice
	Bid(invoiceId string) any
	Adjudicate(invoiceId string) any
	AllRunningBidsToLost(invoiceId string) any
}

type PgDb struct {
	pgUser     string
	pgPwd      string
	pgHostname string
	pgDbname   string
	conn       *pg.Conn
}

func (db *PgDb) Connect() {
	conn, err := pg.Connect(context.Background(), fmt.Sprintf("postgres://%s:%s@%s:5432/%s", db.pgUser, db.pgPwd, db.pgHostname, db.pgDbname))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}
	db.conn = conn
}

func (db *PgDb) Close() {
	err := db.conn.Close(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to close the database: %v\n", err)
	}
	db.conn = nil
}

func (db *PgDb) GetAllInvestors() *[]*pb.Investor {
	data := new([]*pb.Investor)
	rows, err := db.conn.Query(context.Background(), "select id::varchar, name, balance from investors")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
	}
	defer rows.Close()
	for rows.Next() {
		row := new(pb.Investor)
		if err := rows.Scan(&row.Id, &row.Name, &row.Balance); err != nil {
			fmt.Printf("%v", err)
		}
		*data = append(*data, row)
	}
	if err := rows.Err(); err != nil {
		fmt.Printf("%v", err)
	}
	return data
}

func (db *PgDb) GetAllIssuers() *[]*pb.Issuer {
	data := new([]*pb.Issuer)
	rows, err := db.conn.Query(context.Background(), "select id::varchar, name, balance from issuers")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
	}
	defer rows.Close()
	for rows.Next() {
		row := new(pb.Issuer)
		if err := rows.Scan(&row.Id, &row.Name, &row.Balance); err != nil {
			fmt.Printf("%v", err)
		}
		*data = append(*data, row)
	}
	if err := rows.Err(); err != nil {
		fmt.Printf("%v", err)
	}
	return data
}

func (db *PgDb) GetAllBids() *[]*pb.Bid {
	data := new([]*pb.Bid)
	rows, err := db.conn.Query(
		context.Background(),
		`select
			id,
			invoice_id,
			bidder_account_id,
			offer,
			state
		from bids`,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
	}
	defer rows.Close()
	for rows.Next() {
		row := new(pb.Bid)
		if err := rows.Scan(&row.Id, &row.InvoiceId, &row.BidderAccountId, &row.Offer, &row.State); err != nil {
			fmt.Printf("%v", err)
		}
		*data = append(*data, row)
	}
	if err := rows.Err(); err != nil {
		fmt.Printf("%v", err)
	}
	return data
}

func (db *PgDb) GetAllInvoices() *[]*pb.Invoice {
	data := new([]*pb.Invoice)
	rows, err := db.conn.Query(
		context.Background(),
		`select
			id::varchar,
			issuer_account_id,
			reference,
			denom,
			amount,
			asking,
			state
		from invoices`,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
	}
	defer rows.Close()
	for rows.Next() {
		row := new(pb.Invoice)
		if err := rows.Scan(&row.Id, &row.IssuerAccountId, &row.Reference, &row.Denom, &row.Amount, &row.Asking, &row.State); err != nil {
			fmt.Printf("%v", err)
		}
		*data = append(*data, row)
	}
	if err := rows.Err(); err != nil {
		fmt.Printf("%v", err)
	}
	return data
}

func (db *PgDb) NewInvoice(newInvoiceData *pb.Invoice) *pb.Invoice {
	data := new(pb.Invoice)
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

	if err := row.Scan(&data.Id, &data.IssuerAccountId, &data.Reference, &data.Denom, &data.Amount, &data.Asking); err != nil {
		fmt.Printf("%+v", err)
	}
	return data
}

func (db *PgDb) Bid(invoiceId string) any {

	data := new(struct {
		invoiceId       string
		bidderAccountId string
		offer           int64
	})

	row := db.conn.QueryRow(
		context.Background(),
		"select bid($1)",
		invoiceId,
	)

	if err := row.Scan(&data.invoiceId, &data.bidderAccountId, &data.offer); err != nil {
		fmt.Printf("%+v", err)
	}
	return data
}

func (db *PgDb) Adjudicate(invoiceId string) any {
	data := new(struct {
		invoiceId       string
		bidderAccountId string
		amount          int64
	})

	row := db.conn.QueryRow(
		context.Background(),
		"select adjudicate($1)",
		invoiceId,
	)

	if err := row.Scan(&data.invoiceId, &data.bidderAccountId, &data.amount); err != nil {
		fmt.Printf("%+v", err)
	}

	return data
}

func (db *PgDb) AllRunningBidsToLost(invoiceId string) any {
	data := new(struct {
		invoiceId string
		bidId     string
	})

	row := db.conn.QueryRow(
		context.Background(),
		"select adjudicate($1)",
		invoiceId,
	)

	if err := row.Scan(&data.invoiceId, &data.bidId); err != nil {
		fmt.Printf("%+v", err)
	}
	return data
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
