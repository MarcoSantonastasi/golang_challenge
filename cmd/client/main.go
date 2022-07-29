package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/marcosantonastasi/arex_challenge/api/arex/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not not connect to the gRPC server: %v", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	investorServiceClient := pb.NewInvestorServiceClient(conn)
	invoiceServiceClient := pb.NewInvoiceServiceClient(conn)
	issuerServiceClient := pb.NewIssuerServiceClient(conn)

	res1, err1 := investorServiceClient.GetAllInvestors(ctx, &pb.Empty{})
	if err1 != nil {
		log.Fatalf("could not get Investors: %v", err1)
	}
	log.Printf("Greeting: %v", res1.GetData())

	res2, err2 := invoiceServiceClient.GetAllInvoices(ctx, &pb.Empty{})
	if err2 != nil {
		log.Fatalf("could not get Invoices: %v", err2)
	}
	log.Printf("Greeting: %v", res2.GetData())

	res3, err3 := issuerServiceClient.GetAllIssuers(ctx, &pb.Empty{})
	if err3 != nil {
		log.Fatalf("could not get Issuers: %v", err3)
	}
	log.Printf("Greeting: %v", res3.GetData())

}
