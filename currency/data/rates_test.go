package data

import (
	"github.com/hashicorp/go-hclog"
	"testing"
)

// TestNewExchangeRate tests NewExchangeRate function
func TestNewExchangeRate(t *testing.T) {
	_, err := NewExchangeRate(hclog.Default())

	if err != nil {
		t.Fatal(err)
	}
}
