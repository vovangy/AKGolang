package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/cinar/indicator"
)

type Exchanger interface {
	GetCandlesHistory(pair string, limit int, start, end time.Time) (CandlesHistory, error)
}

type CandlesHistory struct {
	Candles []Candle `json:"candles"`
}

type Candle struct {
	T int64   `json:"t"`
	O float64 `json:"o"`
	C float64 `json:"c"`
	H float64 `json:"h"`
	L float64 `json:"l"`
	V float64 `json:"v"`
}

type Indicator interface {
	SMA(period int) []float64
	EMA(period int) []float64
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

type Indicatorer interface {
	GetData(period int, history CandlesHistory) ([]float64, error)
}

type IndicatorSMA struct {
	candleHistory CandlesHistory
	exmo          Exchanger
}

type IndicatorEMA struct {
	candleHistory CandlesHistory
	exmo          Exchanger
}

func (i *IndicatorSMA) GetData(period int, history CandlesHistory) ([]float64, error) {
	i.candleHistory = history
	return i.SMA(period), nil
}

func (i *IndicatorSMA) SMA(period int) []float64 {
	closingPrices := []float64{}
	for _, candle := range i.candleHistory.Candles {
		closingPrices = append(closingPrices, candle.C)
	}
	return indicator.Sma(period, closingPrices)
}

func (i *IndicatorEMA) GetData(period int, history CandlesHistory) ([]float64, error) {
	i.candleHistory = history
	return i.EMA(period), nil
}

func (i *IndicatorEMA) EMA(period int) []float64 {
	closingPrices := []float64{}
	for _, candle := range i.candleHistory.Candles {
		closingPrices = append(closingPrices, candle.C)
	}
	return indicator.Ema(period, closingPrices)
}

func getResponseBody(e *Exmo, request string) ([]byte, error) {
	resp, err := e.client.Get(request)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data from URL: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	return body, nil
}

func unmarshalCandles(data []byte) (CandlesHistory, error) {
	var candlesHistory CandlesHistory
	if err := json.Unmarshal(data, &candlesHistory); err != nil {
		return candlesHistory, fmt.Errorf("failed to parse candle data: %v", err)
	}
	return candlesHistory, nil
}

func (e *Exmo) GetCandlesHistory(pair string, limit int, start, end time.Time) (CandlesHistory, error) {
	request := fmt.Sprintf("%s/candles_history?symbol=%s&resolution=%d&from=%d&to=%d", e.url, pair, limit, start.Unix(), end.Unix())
	body, err := getResponseBody(e, request)
	if err != nil {
		return CandlesHistory{}, fmt.Errorf("error fetching candle history: %v", err)
	}

	history, err := unmarshalCandles(body)
	if err != nil {
		return history, fmt.Errorf("error unmarshalling candle history: %v", err)
	}

	return history, nil
}

func NewIndicatorSMA(exchange Exchanger) *IndicatorSMA {
	return &IndicatorSMA{exmo: exchange}
}

func NewIndicatorEMA(exchange Exchanger) *IndicatorEMA {
	return &IndicatorEMA{exmo: exchange}
}

type Dashboarder interface {
	GetDashboard(pair string, opts ...IndicatorOpt) (DashboardData, error)
}

type DashboardData struct {
	Name           string
	CandlesHistory CandlesHistory
	Indicators     map[string][]IndicatorData
	From           time.Time
	To             time.Time
}

type IndicatorData struct {
	Name     string
	Period   int
	Indicate []float64
}

type IndicatorOpt struct {
	Name      string
	Periods   []int
	Indicator Indicatorer
}

type Dashboard struct {
	exchange           Exchanger
	withCandlesHistory bool
	IndicatorOpts      []IndicatorOpt
	period             int
	from               time.Time
	to                 time.Time
}

func NewDashboard(exchange Exchanger) *Dashboard {
	return &Dashboard{exchange: exchange}
}

func (d *Dashboard) WithCandlesHistory(period int, from, to time.Time) {
	d.period = period
	d.from = from
	d.to = to
	d.withCandlesHistory = true
}

func (d *Dashboard) GetDashboard(pair string, opts ...IndicatorOpt) (DashboardData, error) {
	data := DashboardData{
		Name:       pair,
		From:       d.from,
		To:         d.to,
		Indicators: make(map[string][]IndicatorData),
	}

	candlesHistory, err := d.exchange.GetCandlesHistory(pair, d.period, d.from, d.to)
	if err != nil {
		return data, fmt.Errorf("error fetching dashboard candles history: %v", err)
	}

	data.CandlesHistory = candlesHistory

	for _, opt := range opts {
		var indicators []IndicatorData
		for _, period := range opt.Periods {
			currentData := IndicatorData{Name: opt.Name, Period: period}
			indicateData, err := opt.Indicator.GetData(period, data.CandlesHistory)
			if err != nil {
				fmt.Printf("error calculating indicator %s for period %d: %v\n", opt.Name, period, err)
				break
			}
			currentData.Indicate = indicateData
			indicators = append(indicators, currentData)
		}
		data.Indicators[opt.Name] = indicators
	}

	return data, nil
}

func main() {
	exchange := NewExmo()
	dashboard := NewDashboard(exchange)
	dashboard.WithCandlesHistory(30, time.Now().Add(-time.Hour*24*5), time.Now())
	opts := []IndicatorOpt{
		{
			Name:      "SMA",
			Periods:   []int{5, 10, 20},
			Indicator: NewIndicatorSMA(exchange),
		},
		{
			Name:      "EMA",
			Periods:   []int{5, 10, 20},
			Indicator: NewIndicatorEMA(exchange),
		},
	}
	data, err := dashboard.GetDashboard("BTC_USD", opts...)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}
