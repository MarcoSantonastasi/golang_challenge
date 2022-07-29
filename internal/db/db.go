package db

import (
	"context"
	"fmt"
	"os"

	pg "github.com/jackc/pgx/v4"
	pb "github.com/marcosantonastasi/arex_challenge/api/arex/v1"
)

var DockerPG = Db{}

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

func init() {
	//
	// POSTGRES_USER=postgres
	// POSTGRES_PASSWORD=postgres
	// POSTGRES_DB=postgres
	// POSTGRES_HOSTNAME=localhost
	//

	conn, err := pg.Connect(context.Background(), os.Getenv("POSTGRES_HOSTNAME"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	DockerPG.Conn = conn
}
