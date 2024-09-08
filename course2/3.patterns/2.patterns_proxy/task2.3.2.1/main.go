package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/cinar/indicator"
)

type Indicator interface {
	StochPrice() ([]float64, []float64)
	RSI(period int) ([]float64, []float64)
	StochRSI(rsiPeriod int) ([]float64, []float64)
	SMA(period int) []float64
	MACD() ([]float64, []float64)
	EMA() []float64
}

func UnmarshalKLines(data []byte) (KLines, error) {
	var r KLines
	err := json.Unmarshal(data, &r)
	return r, err
}
func (k *KLines) Marshal() ([]byte, error) {
	return json.Marshal(k)
}

type KLines struct {
	Pair    string   `json:"pair"`
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

type Lines struct {
	high    []float64
	low     []float64
	closing []float64
}

func (l *Lines) StochPrice() ([]float64, []float64) {
	k, d := indicator.StochasticOscillator(l.high, l.low, l.closing)
	return k, d
}

func (l *Lines) RSI(period int) ([]float64, []float64) {
	rs, rsi := indicator.RsiPeriod(period, l.closing)
	return rs, rsi
}

func (l *Lines) StochRSI(rsiPeriod int) ([]float64, []float64) {
	_, rsi := l.RSI(rsiPeriod)
	k, d := indicator.StochasticOscillator(rsi, rsi, rsi)
	return k, d
}

func (l *Lines) MACD() ([]float64, []float64) {
	return indicator.Macd(l.closing)
}

func (l *Lines) EMA() []float64 {
	return indicator.Ema(5, l.closing)
}

func (l *Lines) SMA(period int) []float64 {
	return indicator.Sma(period, l.closing)
}

type LinesProxy struct {
	lines Indicator
	cache map[string][]float64
}

func LoadKlinesProxy(data []byte) *LinesProxy {
	lines := LoadKlines(data)
	cache := make(map[string][]float64)
	proxy := LinesProxy{lines: lines, cache: cache}
	return &proxy
}

func (lp *LinesProxy) StochPrice() ([]float64, []float64) {
	kStochPrice := "k"
	dStochPrice := "d"

	if value, ok := lp.cache[kStochPrice]; ok {
		return value, lp.cache[dStochPrice]
	}

	k, d := lp.lines.StochPrice()

	lp.cache[kStochPrice] = k
	lp.cache[dStochPrice] = d

	return k, d
}

func (lp *LinesProxy) RSI(period int) ([]float64, []float64) {
	rsKey := fmt.Sprintf("rs_%v", period)
	rsiKey := fmt.Sprintf("rsi_%v", period)

	if value, ok := lp.cache[rsKey]; ok {
		return value, lp.cache[rsiKey]
	}

	rs, rsi := lp.lines.RSI(period)

	lp.cache[rsKey] = rs
	lp.cache[rsiKey] = rsi

	return rs, rsi
}

func (lp *LinesProxy) StochRSI(rsiPeriod int) ([]float64, []float64) {
	kKey := fmt.Sprintf("k_stochrsi_%v", rsiPeriod)
	dKey := fmt.Sprintf("d_stochrsi_%v", rsiPeriod)

	if value, ok := lp.cache[kKey]; ok {
		return value, lp.cache[dKey]
	}

	k, d := lp.lines.StochRSI(rsiPeriod)

	lp.cache[kKey] = k
	lp.cache[dKey] = d

	return k, d
}

func (lp *LinesProxy) MACD() ([]float64, []float64) {
	macdKey := "macd"
	signalKey := "signal"

	if value, ok := lp.cache[macdKey]; ok {
		return value, lp.cache[signalKey]
	}

	macd, signal := lp.lines.MACD()
	lp.cache[macdKey] = macd
	lp.cache[signalKey] = signal

	return macd, signal
}

func (lp *LinesProxy) EMA() []float64 {
	emaKey := "ema"

	if value, ok := lp.cache[emaKey]; ok {
		return value
	}

	ema := lp.lines.EMA()
	lp.cache[emaKey] = ema

	return ema
}

func (lp *LinesProxy) SMA(period int) []float64 {
	smaKey := fmt.Sprintf("sma_%v", period)

	if value, ok := lp.cache[smaKey]; ok {
		return value
	}

	sma := lp.lines.SMA(period)
	lp.cache[smaKey] = sma

	return sma
}

func LoadKlines(data []byte) *Lines {
	klines, err := UnmarshalKLines(data)
	if err != nil {
		log.Fatal(err)
	}

	t := &Lines{}
	for _, v := range klines.Candles {
		t.closing = append(t.closing, v.C)
		t.low = append(t.low, v.L)
		t.high = append(t.high, v.H)
	}

	return t
}

func LoadCandles(pair string) []byte {
	req, err := http.NewRequest("GET", fmt.Sprintf(
		"https://api.exmo.com/v1.1/candles_history?symbol=%s&resolution=30&from=1703056979&to=1705476839", pair), nil)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return data
}
