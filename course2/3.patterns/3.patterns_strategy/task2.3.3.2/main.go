package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/cinar/indicator"
)

type Exchanger interface {
	GetCandlesHistory(pair string, limit int, start, end time.Time) (CandlesHistory, error)
}

func UnmarshalCandles(data []byte) (CandlesHistory, error) {
	var r CandlesHistory
	err := json.Unmarshal(data, &r)
	return r, err
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
	StochPrice() ([]float64, []float64)
	RSI(period int) ([]float64, []float64)
	StochRSI(rsiPeriod int) ([]float64, []float64)
	SMA(period int) []float64
	MACD() ([]float64, []float64)
	EMA() []float64
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

type GeneralIndicatorer interface {
	GetData(pair string, period int, from, to time.Time, indicator Indicatorer) ([]float64,
		error)
}

type Indicatorer interface {
	GetData(pair string, limit, period int, from, to time.Time) ([]float64, error)
}

type IndicatorSMA struct {
	candleHistory CandlesHistory
	exmo          Exchanger
}

type IndicatorEMA struct {
	candleHistory CandlesHistory
	exmo          Exchanger
}

func (indicator *IndicatorSMA) GetData(pair string, limit, period int, from, to time.Time) ([]float64, error) {
	history, err := indicator.exmo.GetCandlesHistory(pair, limit, from, to)

	if err != nil {
		return nil, err
	}

	indicator.candleHistory = history
	return indicator.SMA(), nil
}

func (t *IndicatorSMA) SMA() []float64 {
	closing := []float64{}

	for _, i := range t.candleHistory.Candles {
		closing = append(closing, i.C)
	}

	return indicator.Sma(5, closing)
}

func (indicator *IndicatorEMA) GetData(pair string, limit, period int, from, to time.Time) ([]float64, error) {
	history, err := indicator.exmo.GetCandlesHistory(pair, limit, from, to)

	if err != nil {
		return nil, err
	}

	indicator.candleHistory = history
	return indicator.EMA(), nil
}

func (t *IndicatorEMA) EMA() []float64 {
	closing := []float64{}

	for _, i := range t.candleHistory.Candles {
		closing = append(closing, i.C)
	}

	return indicator.Ema(5, closing)
}

func GetRespBody(e *Exmo, request string) ([]byte, error) {
	resp, err := e.client.Get(request)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func (e *Exmo) GetCandlesHistory(pair string, limit int, start, end time.Time) (CandlesHistory, error) {
	request := e.url + "/candles_history?symbol=" + pair + "&resolution=" + strconv.Itoa(limit) + "&from=" +
		strconv.FormatInt(start.Unix(), 10) + "&to=" + strconv.FormatInt(end.Unix(), 10)
	body, err := GetRespBody(e, request)
	var history CandlesHistory

	if err != nil {
		return history, err
	}

	history, err = UnmarshalCandles(body)

	if err != nil {
		return history, err
	}

	return history, nil
}

func NewIndicatorSMA(exchange Exchanger) *IndicatorSMA {
	return &IndicatorSMA{exmo: exchange}
}

func NewIndicatorEMA(exchange Exchanger) *IndicatorEMA {
	return &IndicatorEMA{exmo: exchange}
}

type GeneralIndicator struct {
}

func (generalIndicator *GeneralIndicator) GetData(pair string, period int, from, to time.Time,
	indicator Indicatorer) ([]float64, error) {
	return indicator.GetData(pair, 15, period, from, to)
}
