package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/cinar/indicator"
)

const (
	ticker         = "/ticker"
	trades         = "/trades"
	orderBook      = "/order_book"
	currency       = "/currency"
	candlesHistory = "/candles_history"
)

type CandlesHistory struct {
	Candles []Candle `json:"candles"`
}

type Candle struct {
	Timestamp int64   `json:"t"`
	Open      float64 `json:"o"`
	High      float64 `json:"h"`
	Low       float64 `json:"l"`
	Close     float64 `json:"c"`
	Volume    float64 `json:"v"`
}

type OrderBookPair struct {
	AskQuantity string     `json:"ask_quantity"`
	AskAmount   string     `json:"ask_amount"`
	AskTop      string     `json:"ask_top"`
	BidQuantity string     `json:"bid_quantity"`
	BidAmount   string     `json:"bid_amount"`
	BidTop      string     `json:"bid_top"`
	Ask         [][]string `json:"ask"`
	Bid         [][]string `json:"bid"`
}

type TickerValue struct {
	BuyPrice  string `json:"buy_price"`
	SellPrice string `json:"sell_price"`
	LastTrade string `json:"last_trade"`
	High      string `json:"high"`
	Low       string `json:"low"`
	Avg       string `json:"avg"`
	Vol       string `json:"vol"`
	VolCurr   string `json:"vol_curr"`
	Updated   int64  `json:"updated"`
}

type Pair struct {
	TradeID  int64  `json:"trade_id"`
	Date     int64  `json:"date"`
	Type     string `json:"type"`
	Quantity string `json:"quantity"`
	Price    string `json:"price"`
	Amount   string `json:"amount"`
}

type Indicator struct {
	exchange     Exchanger
	calculateSMA func(data []float64, period int) []float64
	calculateEMA func(data []float64, period int) []float64
}

type IndicatorOption func(*Indicator)
type Currencies []string
type OrderBook map[string]OrderBookPair
type Ticker map[string]TickerValue
type Trades map[string][]Pair
type Exchanger interface {
	GetTicker() (Ticker, error)
	GetTrades(pairs ...string) (Trades, error)
	GetOrderBook(limit int, pairs ...string) (OrderBook, error)
	GetCurrencies() (Currencies, error)
	GetCandlesHistory(pair string, limit int, start, end time.Time) (CandlesHistory, error)
	GetClosePrice(pair string, limit int, start, end time.Time) ([]float64, error)
}

type Indicatorer interface {
	SMA(pair string, resolution, period int, from, to time.Time) ([]float64, error)
	EMA(pair string, resolution, period int, from, to time.Time) ([]float64, error)
}

type Exmo struct {
	client *http.Client
	url    string
}

func NewExmo(opts ...func(exmo *Exmo)) *Exmo {
	exmo := &Exmo{client: &http.Client{}, url: "https://api.exmo.com/v1.1"}

	for _, option := range opts {
		option(exmo)
	}

	return exmo
}

func WithClient(client *http.Client) func(exmo *Exmo) {
	return func(exmo *Exmo) {
		exmo.client = client
	}
}

func WithURL(url string) func(exmo *Exmo) {
	return func(exmo *Exmo) {
		exmo.url = url
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

func NewIndicator(exchange Exchanger, opts ...IndicatorOption) Indicatorer {
	indicator := &Indicator{exchange: exchange}

	for _, option := range opts {
		option(indicator)
	}

	return indicator
}

func (i *Indicator) SMA(pair string, resolution, period int, from, to time.Time) ([]float64, error) {
	data, err := i.exchange.GetClosePrice(pair, resolution, from, to)

	if err != nil {
		return nil, err
	}

	return i.calculateSMA(data, period), nil
}

func (i *Indicator) EMA(pair string, resolution, period int, from, to time.Time) ([]float64, error) {
	data, err := i.exchange.GetClosePrice(pair, resolution, from, to)

	if err != nil {
		return nil, err
	}

	return i.calculateEMA(data, period), nil
}

func calculateSMA(closing []float64, period int) []float64 {

	return indicator.Sma(period, closing)
}

func calculateEMA(closing []float64, period int) []float64 {
	return indicator.Ema(period, closing)
}

func WithCalculateEMA(f func(closing []float64, period int) []float64) IndicatorOption {
	return func(i *Indicator) {
		i.calculateEMA = f
	}
}

func WithCalculateSMA(f func(closing []float64, period int) []float64) IndicatorOption {
	return func(i *Indicator) {
		i.calculateSMA = f
	}
}
