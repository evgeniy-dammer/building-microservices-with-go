package grpcserver

import (
	"context"
	protos "github.com/evgeniy-dammer/building-microservices-with-go/currency/protos/currency"
	"github.com/hashicorp/go-hclog"
)

// Currency is a gRPC server it implements the methods defined by the CurrencyServer interface
type Currency struct {
	log hclog.Logger
	protos.UnimplementedCurrencyServer
}

// NewCurrency creates a new Currency server
func NewCurrency(l hclog.Logger) *Currency {
	return &Currency{log: l}
}

// GetRate implements the CurrencyServer GetRate method and returns the currency exchange rate for the two given currencies.
func (c *Currency) GetRate(ctx context.Context, request *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Info("Handle GetRate", "base", request.GetBase(), "destination", request.GetDestination())

	return &protos.RateResponse{Rate: 0.5}, nil
}
