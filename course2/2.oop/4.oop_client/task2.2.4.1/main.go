package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	ticker         = "/ticker"
	trades         = "/trades"
	orderBook      = "/order_book"
	currency       = "/currency"
	candlesHistory = "/candles_history"
)

type Candle struct {
	Timestamp int64   `json:"t"`
	Open      float64 `json:"o"`
	High      float64 `json:"h"`
	Low       float64 `json:"l"`
	Close     float64 `json:"c"`
	Volume    float64 `json:"v"`
}

type CandlesHistory struct {
	Candles []Candle `json:"candles"`
}

type Currencies []string

type OrderBook map[string]OrderBookPair

type OrderBookPair struct {
	BidTop string `json:"bid_top"`
	AskTop string `json:"ask_top"`
}

type Ticker map[string]TickerValue

type TickerValue struct {
	Price string `json:"last_trade"`
}

type Trades map[string][]Trade

type Trade struct {
	Price     string `json:"price"`
	Amount    string `json:"amount"`
	Timestamp int64  `json:"date"`
}

type Exchanger interface {
	GetTicker() (Ticker, error)
	GetTrades(pairs ...string) (Trades, error)
	GetOrderBook(limit int, pairs ...string) (OrderBook, error)
	GetCurrencies() (Currencies, error)
	GetCandlesHistory(pair string, limit int, start, end time.Time) (CandlesHistory, error)
	GetClosePrice(pair string, limit int, start, end time.Time) ([]float64, error)
}

type Exmo struct {
	client *http.Client
	url    string
}

func NewExmo(opts ...func(*Exmo)) *Exmo {
	exmo := &Exmo{
		client: &http.Client{},
		url:    "https://api.exmo.com/v1",
	}
	for _, opt := range opts {
		opt(exmo)
	}
	return exmo
}

func WithClient(client *http.Client) func(*Exmo) {
	return func(e *Exmo) {
		e.client = client
	}
}

func WithURL(url string) func(*Exmo) {
	return func(e *Exmo) {
		e.url = url
	}
}

func (e *Exmo) GetTicker() (Ticker, error) {
	resp, err := e.client.Get(e.url + ticker)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var tickerData Ticker
	if err := json.NewDecoder(resp.Body).Decode(&tickerData); err != nil {
		return nil, err
	}

	return tickerData, nil
}

func (e *Exmo) GetTrades(pairs ...string) (Trades, error) {

	resp, err := e.client.Get(e.url + trades + "?" + "pair=" + strings.Join(pairs, ","))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var tradesData Trades
	if err := json.NewDecoder(resp.Body).Decode(&tradesData); err != nil {
		return nil, err
	}

	return tradesData, nil
}

func (e *Exmo) GetOrderBook(limit int, pairs ...string) (OrderBook, error) {
	resp, err := e.client.Get(e.url + orderBook + "?pair=" + strings.Join(pairs, ",") + "&limit=" + strconv.Itoa(limit))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var orderBookData OrderBook
	if err := json.NewDecoder(resp.Body).Decode(&orderBookData); err != nil {
		return nil, err
	}

	return orderBookData, nil
}

func (e *Exmo) GetCurrencies() (Currencies, error) {
	resp, err := e.client.Get(e.url + currency)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var currenciesData Currencies
	if err := json.NewDecoder(resp.Body).Decode(&currenciesData); err != nil {
		return nil, err
	}

	return currenciesData, nil
}

func (e *Exmo) GetCandlesHistory(pair string, limit int, start, end time.Time) (CandlesHistory, error) {
	resp, err := e.client.Get(e.url + candlesHistory + "?symbol=" + pair + "&resolution=" + strconv.Itoa(limit) + "&from=" +
		strconv.FormatInt(start.Unix(), 10) + "&to=" + strconv.FormatInt(end.Unix(), 10))

	if err != nil {

		return CandlesHistory{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {

		return CandlesHistory{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var candlesHistoryData CandlesHistory
	if err := json.NewDecoder(resp.Body).Decode(&candlesHistoryData); err != nil {
		return CandlesHistory{}, err
	}

	return candlesHistoryData, nil
}

func (e *Exmo) GetClosePrice(pair string, limit int, start, end time.Time) ([]float64, error) {
	candles, err := e.GetCandlesHistory(pair, limit, start, end)
	if err != nil {
		return nil, err
	}

	var closePrices []float64
	for _, candle := range candles.Candles {
		closePrices = append(closePrices, candle.Close)
	}

	return closePrices, nil
}
