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

func (db *Db) GetAllInvestors() []*pb.Investor {
	var greeting string
	_ = db.Conn.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	fmt.Println(greeting)
	return nil
}

func (db *Db) GetAllIssuers() []*pb.Issuer {
	var greeting string
	_ = db.Conn.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	fmt.Println(greeting)
	return nil
}
func (db *Db) GetAllInvoices() []*pb.Invoice {
	var greeting string
	_ = db.Conn.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	fmt.Println(greeting)
	return nil
}

var DockerPG = Db{}

func init() {
	//
	// POSTGRES_USER=postgres
	// POSTGRES_PASSWORD=postgres
	// POSTGRES_DB=postgres
	// POSTGRES_HOSTNAME=localhost
	//

	conn, err := pg.Connect(context.Background(), fmt.Sprintf("postgres://%s:%s@%s:5432/%s", os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_HOSTNAME"), os.Getenv("POSTGRES_DB")))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	DockerPG.Conn = conn
}
