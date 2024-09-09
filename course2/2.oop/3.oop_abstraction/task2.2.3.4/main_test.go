package main

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func TestGetTicker(t *testing.T) {
	tests := []struct {
		name           string
		response       string
		expectedResult Ticker
		expectError    bool
	}{
		{
			name: "Valid Ticker",
			response: `{
				"BTC_USD": {
					"buy_price": "589.06",
					"sell_price": "592",
					"last_trade": "591.221",
					"high": "602.082",
					"low": "584.51011695",
					"avg": "591.14698808",
					"vol": "167.59763535",
					"vol_curr": "99095.17162071",
					"updated": 1617879240
				}
			}`,
			expectedResult: Ticker{
				"BTC_USD": TickerValue{
					BuyPrice:  "589.06",
					SellPrice: "592",
					LastTrade: "591.221",
					High:      "602.082",
					Low:       "584.51011695",
					Avg:       "591.14698808",
					Vol:       "167.59763535",
					VolCurr:   "99095.17162071",
					Updated:   1617879240,
				},
			},
			expectError: false,
		},
		{
			name:           "Invalid JSON",
			response:       `{"BTC_USD": {`,
			expectedResult: Ticker{},
			expectError:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(tt.response))
			}))
			defer ts.Close()

			exmo := NewExmo(WithURL(ts.URL))
			got, err := exmo.GetTicker()

			if (err != nil) != tt.expectError {
				t.Errorf("GetTicker() error = %v, expectError %v", err, tt.expectError)
				return
			}

			if !reflect.DeepEqual(got, tt.expectedResult) {
				t.Errorf("GetTicker() = %v, expected %v", got, tt.expectedResult)
			}
		})
	}
}

func TestGetTrades(t *testing.T) {
	tests := []struct {
		name           string
		response       string
		pairs          []string
		expectedResult Trades
		expectError    bool
	}{
		{
			name: "Valid Trades",
			response: `{
				"BTC_USD": [
					{
						"trade_id": 12345,
						"type": "buy",
						"price": "589.06",
						"quantity": "1.0",
						"amount": "589.06",
						"date": 1617879240
					}
				]
			}`,
			pairs: []string{"BTC_USD"},
			expectedResult: Trades{
				"BTC_USD": []TradePair{
					{
						TradeID:  12345,
						Type:     "buy",
						Price:    "589.06",
						Quantity: "1.0",
						Amount:   "589.06",
						Date:     1617879240,
					},
				},
			},
			expectError: false,
		},
		{
			name:           "Invalid JSON",
			response:       `{"BTC_USD": [}`,
			pairs:          []string{"BTC_USD"},
			expectedResult: Trades{},
			expectError:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(tt.response))
			}))
			defer ts.Close()

			exmo := NewExmo(WithURL(ts.URL))
			got, err := exmo.GetTrades(tt.pairs...)

			if (err != nil) != tt.expectError {
				t.Errorf("GetTrades() error = %v, expectError %v", err, tt.expectError)
				return
			}

			if !reflect.DeepEqual(got, tt.expectedResult) {
				t.Errorf("GetTrades() = %v, expected %v", got, tt.expectedResult)
			}
		})
	}
}

func TestGetOrderBook(t *testing.T) {
	tests := []struct {
		name           string
		response       string
		pairs          []string
		limit          int
		expectedResult OrderBook
		expectError    bool
	}{
		{
			name: "Valid OrderBook",
			response: `{
				"BTC_USD": {
					"ask_quantity": "10",
					"ask_amount": "5890.60",
					"ask_top": "591.22",
					"bid_quantity": "5",
					"bid_amount": "2945.30",
					"bid_top": "589.06",
					"ask": [[591, 10]],
					"bid": [[589, 5]]
				}
			}`,
			pairs: []string{"BTC_USD"},
			limit: 1,
			expectedResult: OrderBook{
				"BTC_USD": OrderBookPair{
					AskQuantity: "10",
					AskAmount:   "5890.60",
					AskTop:      "591.22",
					BidQuantity: "5",
					BidAmount:   "2945.30",
					BidTop:      "589.06",
					Ask:         [][]int{{591, 10}},
					Bid:         [][]int{{589, 5}},
				},
			},
			expectError: false,
		},
		{
			name:           "Invalid JSON",
			response:       `{"BTC_USD": {`,
			pairs:          []string{"BTC_USD"},
			limit:          1,
			expectedResult: OrderBook{},
			expectError:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(tt.response))
			}))
			defer ts.Close()

			exmo := NewExmo(WithURL(ts.URL))
			got, err := exmo.GetOrderBook(tt.limit, tt.pairs...)

			if (err != nil) != tt.expectError {
				t.Errorf("GetOrderBook() error = %v, expectError %v", err, tt.expectError)
				return
			}

			if !reflect.DeepEqual(got, tt.expectedResult) {
				t.Errorf("GetOrderBook() = %v, expected %v", got, tt.expectedResult)
			}
		})
	}
}

func TestGetCurrencies(t *testing.T) {
	tests := []struct {
		name           string
		response       string
		expectedResult Currencies
		expectError    bool
	}{
		{
			name:           "Valid Currencies",
			response:       `["USD", "BTC", "ETH"]`,
			expectedResult: Currencies{"USD", "BTC", "ETH"},
			expectError:    false,
		},
		{
			name:           "Invalid JSON",
			response:       `["USD", "BTC", "ETH"`,
			expectedResult: Currencies{},
			expectError:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(tt.response))
			}))
			defer ts.Close()

			exmo := NewExmo(WithURL(ts.URL))
			got, err := exmo.GetCurrencies()

			if (err != nil) != tt.expectError {
				t.Errorf("GetCurrencies() error = %v, expectError %v", err, tt.expectError)
				return
			}

			if !reflect.DeepEqual(got, tt.expectedResult) {
				t.Errorf("GetCurrencies() = %v, expected %v", got, tt.expectedResult)
			}
		})
	}
}

func TestGetCandlesHistory(t *testing.T) {
	tests := []struct {
		name           string
		response       string
		pair           string
		limit          int
		start, end     time.Time
		expectedResult CandlesHistory
		expectError    bool
	}{
		{
			name: "Valid CandlesHistory",
			response: `{
				"candles": [
					{"t": 1617879240, "o": 589.06, "c": 591.22, "h": 593.00, "l": 588.00, "v": 1000}
				]
			}`,
			pair:           "BTC_USD",
			limit:          1,
			start:          time.Unix(1617879240, 0),
			end:            time.Unix(1617882840, 0),
			expectedResult: CandlesHistory{{Time: 1617879240, Open: 589.06, Close: 591.22, High: 593.00, Low: 588.00, Volume: 1000}},
			expectError:    false,
		},
		{
			name:           "Invalid JSON",
			response:       `{"candles": [}`,
			pair:           "BTC_USD",
			limit:          1,
			start:          time.Unix(1617879240, 0),
			end:            time.Unix(1617882840, 0),
			expectedResult: CandlesHistory{},
			expectError:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(tt.response))
			}))
			defer ts.Close()

			exmo := NewExmo(WithURL(ts.URL))
			got, err := exmo.GetCandlesHistory(tt.pair, tt.start, tt.end, tt.limit)

			if (err != nil) != tt.expectError {
				t.Errorf("GetCandlesHistory() error = %v, expectError %v", err, tt.expectError)
				return
			}

			if !reflect.DeepEqual(got, tt.expectedResult) {
				t.Errorf("GetCandlesHistory() = %v, expected %v", got, tt.expectedResult)
			}
		})
	}
}
