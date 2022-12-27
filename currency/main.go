package main

import (
	protos "github.com/evgeniy-dammer/building-microservices-with-go/currency/protos/currency"
	"github.com/evgeniy-dammer/building-microservices-with-go/currency/server"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
)

func main() {
	log := hclog.Default()

	// create a new gRPC server, use WithInsecure to allow http connections
	gs := grpc.NewServer()

	// create an instance of the Currency server
	cs := grpcserver.NewCurrency(log)

	// register the currency server
	protos.RegisterCurrencyServer(gs, cs)

	// register the reflection service which allows clients to determine the methods for this gRPC service
	reflection.Register(gs)

	// create a TCP socket for inbound server connections
	l, err := net.Listen("tcp", ":9092")

	if err != nil {
		log.Error("Unable to listen", "error", err)
		os.Exit(1)
	}

	log.Info("Currency service is started", "port", "9092")

	// listen for requests
	err = gs.Serve(l)

	if err != nil {
		log.Error("Unable to serve", "error", err)
		os.Exit(1)
	}
}
