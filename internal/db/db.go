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
	GetAllInvestors() []*pb.Investor
	GetAllIssuers() []*pb.Issuer
	GetAllInvoices() []*pb.Invoice
	NewInvoice(*pb.Invoice) *pb.Invoice
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
		os.Exit(1)
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

func (db *PgDb) GetAllInvestors() (data []*pb.Investor) {
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
		data = append(data, row)
	}
	if err := rows.Err(); err != nil {
		fmt.Printf("%v", err)
	}
	return
}

func (db *PgDb) GetAllIssuers() (data []*pb.Issuer) {
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
		data = append(data, row)
	}
	if err := rows.Err(); err != nil {
		fmt.Printf("%v", err)
	}
	return
}

func (db *PgDb) GetAllInvoices() (data []*pb.Invoice) {
	rows, err := db.conn.Query(
		context.Background(),
		`select
			id::varchar,
			issuer_id,
			reference,
			denom,
			amount,
			asking
		from invoices`,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
	}
	defer rows.Close()
	for rows.Next() {
		row := new(pb.Invoice)
		if err := rows.Scan(&row.Id, &row.IssuerId, &row.Reference, &row.Denom, &row.Amount, &row.Asking); err != nil {
			fmt.Printf("%v", err)
		}
		data = append(data, row)
	}
	if err := rows.Err(); err != nil {
		fmt.Printf("%v", err)
	}
	return
}

func (db *PgDb) NewInvoice(newInvoiceData *pb.Invoice) (data *pb.Invoice) {
	data = new(pb.Invoice)
	row := db.conn.QueryRow(
		context.Background(),
		`insert into invoices (
			issuer_id,
			reference,
			denom,
			amount,
			asking
		)
		values($1, $2, $3, $4, $5)
		returning
		    id,
			issuer_id,
			reference,
			denom,
			amount,
			asking`,
		newInvoiceData.IssuerId,
		newInvoiceData.Reference,
		newInvoiceData.Denom,
		newInvoiceData.Amount,
		newInvoiceData.Asking,
	)

	if err := row.Scan(&data.Id, &data.IssuerId, &data.Reference, &data.Denom, &data.Amount, &data.Asking); err != nil {
		fmt.Printf("%+v", err)
	}
	return
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
