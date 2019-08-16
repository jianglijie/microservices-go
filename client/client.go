package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	pb "services/proto"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewFinanceClient(conn)
	c1 := pb.NewGreeterClient(conn)

	r, err := c.GetCoinFinance(context.Background(), &pb.CoinFinanceRequest{Coin: "BTC", StartTime: 1562284800, EndTime: 1562295600})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("data: %s", r.GetFinanceList())
	r1, err1 := c1.SayHello(context.Background(), &pb.HelloRequest{Name: "aaaa"})
	if err1 != nil {
		log.Fatalf("could not greet: %v", err1)
	}
	log.Printf("data: %s", r1.Message)
}
