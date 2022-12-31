package data

import (
	"encoding/xml"
	"fmt"
	"github.com/hashicorp/go-hclog"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// Cubes contents a slice of Cube
type Cubes struct {
	CubeData []Cube `xml:"Cube>Cube>Cube"`
}

// Cube contents a single item of Cubes
type Cube struct {
	Currency string `xml:"currency,attr"`
	Rate     string `xml:"rate,attr"`
}

// ExchangeRate contents exchange rates
type ExchangeRate struct {
	log   hclog.Logger
	rates map[string]float64
}

// NewExchangeRate creates a new ExchangeRate object
func NewExchangeRate(l hclog.Logger) (*ExchangeRate, error) {
	er := &ExchangeRate{log: l, rates: map[string]float64{}}

	// get rates from European Bank API
	err := er.GetRates()

	return er, err
}

// GetRate checks exchange rates in ExchangeRate.rates list and return division
func (e *ExchangeRate) GetRate(base, dest string) (float64, error) {
	// base rate
	br, ok := e.rates[base]
	if !ok {
		return 0, fmt.Errorf("rate not found for currency %s", base)
	}

	// destination rate
	dr, ok := e.rates[base]
	if !ok {
		return 0, fmt.Errorf("rate not found for currency %s", dest)
	}

	return dr / br, nil
}

// GetRates fetches exchange rates from European Bank API
func (e *ExchangeRate) GetRates() error {
	// request to European Bank API for fetching exchange rates
	resp, err := http.DefaultClient.Get("https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml")

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("expected code 200, but got %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	md := &Cubes{}

	// decode xml
	err = xml.NewDecoder(resp.Body).Decode(&md)

	if err != nil {
		return fmt.Errorf("XML decoding failed %s", resp.Body)
	}

	for _, c := range md.CubeData {
		r, err := strconv.ParseFloat(c.Rate, 64)

		if err != nil {
			return err
		}

		e.rates[c.Currency] = r
	}

	e.rates["EUR"] = 1

	return nil
}

// MonitorRates checks the rates in the ECB API every interval and sends a message to the
// returned channel when there are changes
// Note: the ECB API only returns data once a day, this function only simulates the changes
// in rates for demonstration purposes
func (e *ExchangeRate) MonitorRates(interval time.Duration) chan struct{} {
	ret := make(chan struct{})

	go func() {
		ticker := time.NewTicker(interval)
		for {
			select {
			case <-ticker.C:
				// just add a random difference to the rate and return it
				// this simulates the fluctuations in currency rates
				for k, v := range e.rates {
					// change can be 10% of original value
					change := rand.Float64() / 10
					// is this a positive or negative change
					direction := rand.Intn(1)

					if direction == 0 {
						// new value with be min 90% of old
						change = 1 - change
					} else {
						// new value will be 110% of old
						change = 1 + change
					}

					// modify the rate
					e.rates[k] = v * change
				}

				// notify updates, this will block unless there is a listener on the other end
				ret <- struct{}{}
			}
		}
	}()

	return ret
}
