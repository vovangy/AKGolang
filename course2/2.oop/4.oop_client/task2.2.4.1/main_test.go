package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetTickerReturnsValidData(t *testing.T) {
	mockResponse := `{
		"BTC_USD": {
		  "buy_price": "589.06",
		  "sell_price": "592",
		  "last_trade": "591.221",
		  "high": "602.082",
		  "low": "584.51011695",
		  "avg": "591.14698808",
		  "vol": "167.59763535",
		  "vol_curr": "99095.17162071",
		  "updated": 1470250973
		}
	  }`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockResponse))
	}))
	defer server.Close()

	exmo := NewExmo(WithURL(server.URL))

	ticker, err := exmo.GetTicker()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	expected := Ticker{
		"BTC_USD": TickerValue{
			Price: "589.06",
		},
	}

	if !jsonEqual(ticker, expected) {
		t.Errorf("expected %v, got %v", expected, ticker)
	}
}

func TestGetTrades_SuccessfullyFetchTradesForSinglePair(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := map[string][]Trade{
			"BTC_USD": {
				{Timestamp: time.Now().Unix(), Price: "50000", Amount: "5000"},
			},
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	exmo := NewExmo(WithURL(server.URL))
	trades, err := exmo.GetTrades("BTC_USD")

	assert.NoError(t, err)
	assert.NotNil(t, trades)
	assert.Contains(t, trades, "BTC_USD")
	assert.Equal(t, 1, len(trades["BTC_USD"]))
}

func jsonEqual(a, b interface{}) bool {
	aJSON, _ := json.Marshal(a)
	bJSON, _ := json.Marshal(b)
	return string(aJSON) == string(bJSON)
}

func TestGetOrderBookSinglePair(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := map[string]OrderBookPair{
			"BTC_USD": {
				AskTop: "5000",
				BidTop: "5500",
			},
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	exmo := NewExmo(WithURL(server.URL))

	orderBook, err := exmo.GetOrderBook(1, "BTC_USD")

	assert.NoError(t, err)
	assert.NotNil(t, orderBook)
	assert.Equal(t, orderBook["BTC_USD"].AskTop, "5000")
	assert.Equal(t, orderBook["BTC_USD"].BidTop, "5500")
}

func TestGetCurrencies_Success(t *testing.T) {
	mockResponse := `["USD"]`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockResponse))
	}))
	defer server.Close()

	exmo := NewExmo(WithURL(server.URL))
	currencies, err := exmo.GetCurrencies()

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expected := Currencies{"USD"}
	if len(currencies) != len(expected) {
		t.Fatalf("Expected %v, got %v", expected, currencies)
	}

	for i, currency := range currencies {
		if currency != expected[i] {
			t.Errorf("Expected %v, got %v", expected[i], currency)
		}
	}
}

func TestGetClosePriceReturnsCorrectPrices(t *testing.T) {
	mockResponse := CandlesHistory{
		Candles: []Candle{
			{Close: 10.0},
			{Close: 20.0},
			{Close: 30.0},
		},
	}
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(mockResponse)
	}))
	defer mockServer.Close()

	exmo := NewExmo(WithURL(mockServer.URL))
	pair := "BTC_USD"
	limit := 1
	start := time.Now().Add(-24 * time.Hour)
	end := time.Now()

	prices, err := exmo.GetClosePrice(pair, limit, start, end)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	expectedPrices := []float64{10.0, 20.0, 30.0}

	for i, price := range prices {
		if price != expectedPrices[i] {
			t.Errorf("expected price %v, got %v", expectedPrices[i], price)
		}
	}
}
