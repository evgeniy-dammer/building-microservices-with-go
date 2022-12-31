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
	rates         *data.ExchangeRate
	log           hclog.Logger
	subscriptions map[protos.Currency_SubscribeRatesServer][]*protos.RateRequest
	protos.UnimplementedCurrencyServer
}

// NewCurrency creates a new Currency server
func NewCurrency(r *data.ExchangeRate, l hclog.Logger) *Currency {
	c := &Currency{rates: r, log: l, subscriptions: make(map[protos.Currency_SubscribeRatesServer][]*protos.RateRequest)}
	go c.handleUpdate()

	return c
}

// handleUpdate get exchange rates and send them to subscribers
func (c *Currency) handleUpdate() {
	ru := c.rates.MonitorRates(5 * time.Second)

	for range ru {
		c.log.Info("Got updated rates")

		// loop over subscribe clients
		for k, v := range c.subscriptions {
			// loop over rates
			for _, rr := range v {
				rate, err := c.rates.GetRate(rr.GetBase().String(), rr.GetDestination().String())
				if err != nil {
					c.log.Error("Unable to get updated rate", "base", rr.GetBase().String(), "destination", rr.GetDestination().String())
				}

				err = k.Send(&protos.RateResponse{Base: rr.Base, Destination: rr.Destination, Rate: rate})
				if err != nil {
					c.log.Error("Unable to send updated rate", "base", rr.GetBase().String(), "destination", rr.GetDestination().String())
				}
			}
		}
	}
}

// GetRate implements the CurrencyServer GetRate method and returns the currency exchange rate for the two given currencies.
func (c *Currency) GetRate(ctx context.Context, request *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Info("Handle GetRate", "base", request.GetBase(), "destination", request.GetDestination())

	rate, err := c.rates.GetRate(request.GetBase().String(), request.GetDestination().String())

	if err != nil {
		return nil, err
	}
	return &protos.RateResponse{Base: request.Base, Destination: request.Destination, Rate: rate}, nil
}

// SubscribeRates implements the gRPC bi-direction streaming method for the server
func (c *Currency) SubscribeRates(src protos.Currency_SubscribeRatesServer) error {
	// handle client messages
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
			return err
		}
		c.log.Info("Handle client request", "request", rr.String())
		rrs, ok := c.subscriptions[src]

		if !ok {
			rrs = []*protos.RateRequest{}
		}

		rrs = append(rrs, rr)
		c.subscriptions[src] = rrs
	}
	return nil
}
