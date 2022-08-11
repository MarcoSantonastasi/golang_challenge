package db

import (
	"context"
	"fmt"
	"os"

	pg "github.com/jackc/pgx/v4"
	pb "github.com/marcosantonastasi/arex_challenge/api/arex/v1"
)

type IDb interface {
	GetAllInvestors() []*pb.Investor
	GetAllIssuers() []*pb.Issuer
	GetAllInvoices() []*pb.Invoice
}

type Db struct {
	Conn *pg.Conn
}

func (db *Db) GetAllInvestors() (data []*pb.Investor) {
	rows, err := db.Conn.Query(context.Background(), "select id, name, balance from accounts where type = 'INVESTOR'")
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
	fmt.Printf("Query results: %v\n", data)
	return
}

func (db *Db) GetAllIssuers() (data []*pb.Issuer) {
	rows, err := db.Conn.Query(context.Background(), "select id, name, balance from accounts where type = 'ISSUER'")
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
	fmt.Printf("Query results: %v\n", data)
	return
}
func (db *Db) GetAllInvoices() (data []*pb.Invoice) {
	rows, err := db.Conn.Query(context.Background(), "select id, denom, amount, asking from invoices")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
	}
	defer rows.Close()
	for rows.Next() {
		row := new(pb.Invoice)
		if err := rows.Scan(&row.Id, &row.Denom, &row.Amount, &row.Asking); err != nil {
			fmt.Printf("%v", err)
		}
		data = append(data, row)
	}
	if err := rows.Err(); err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf("Query results: %v\n", data)
	return
}

func NewPGDB(
	pg_user string,
	pg_pwd string,
	pg_hostname string,
	pg_dbname string,
) *Db {
	newPgDb := Db{}

	conn, err := pg.Connect(context.Background(), fmt.Sprintf("postgres://%s:%s@%s:5432/%s", pg_user, pg_pwd, pg_hostname, pg_dbname))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	newPgDb.Conn = conn

	return &newPgDb
}

//
// POSTGRES_USER=postgres
// POSTGRES_PASSWORD=postgres
// POSTGRES_DB=postgres
// POSTGRES_HOSTNAME=localhost
//

var DockerPG = NewPGDB(os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_HOSTNAME"), os.Getenv("POSTGRES_DB"))
