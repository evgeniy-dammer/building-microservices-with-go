package grpcserver

import (
	"context"
	"github.com/evgeniy-dammer/building-microservices-with-go/currency/data"
	protos "github.com/evgeniy-dammer/building-microservices-with-go/currency/protos/currency"
	"github.com/hashicorp/go-hclog"
	"io"
	"time"
)

// Currency is a gRPC server it implements the methods defined by the CurrencyServer interface
type Currency struct {
	rates *data.ExchangeRate
	log   hclog.Logger
	protos.UnimplementedCurrencyServer
}

// NewCurrency creates a new Currency server
func NewCurrency(r *data.ExchangeRate, l hclog.Logger) *Currency {
	return &Currency{rates: r, log: l}
}

// GetRate implements the CurrencyServer GetRate method and returns the currency exchange rate for the two given currencies.
func (c *Currency) GetRate(ctx context.Context, request *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Info("Handle GetRate", "base", request.GetBase(), "destination", request.GetDestination())

	rate, err := c.rates.GetRate(request.GetBase().String(), request.GetDestination().String())

	if err != nil {
		return nil, err
	}
	return &protos.RateResponse{Rate: rate}, nil
}

// SubscribeRates implements the gRPC bi-direction streaming method for the server
func (c *Currency) SubscribeRates(src protos.Currency_SubscribeRatesServer) error {
	// handle client messages
	go func() {
		for {
			// Recv is a blocking method which returns on client data
			rr, err := src.Recv()
			// io.EOF signals that the client has closed the connection
			if err == io.EOF {
				c.log.Info("Client has closed connection")
				break
			}
			// any other error means the transport between the server and client is unavailable
			if err != nil {
				c.log.Error("Unable to read from client")
				break
			}
			c.log.Info("Handle client request", "request", rr.String())
		}
	}()

	// handle server responses, we block here to keep the connection open
	for {
		err := src.Send(&protos.RateResponse{Rate: 12.1})
		if err != nil {
			return err
		}
		// send a message back to the client
		time.Sleep(5 * time.Second)
	}
}
