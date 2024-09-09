package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestLinesProxy_StochPrice(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockIndicator := NewMockIndicator(ctrl)
	lp := &LinesProxy{
		lines: mockIndicator,
		cache: make(map[string][]float64),
	}

	expectedK := []float64{10, 20, 30}
	expectedD := []float64{40, 50, 60}

	mockIndicator.EXPECT().StochPrice().Return(expectedK, expectedD).Times(1)

	k, d := lp.StochPrice()

	assert.Equal(t, expectedK, k)
	assert.Equal(t, expectedD, d)
	assert.Equal(t, expectedK, lp.cache["k"])
	assert.Equal(t, expectedD, lp.cache["d"])
}

func TestLinesProxy_EMA(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockIndicator := NewMockIndicator(ctrl)

	expectedEMA := []float64{1.1, 2.2, 3.3}

	mockIndicator.EXPECT().EMA().Return(expectedEMA).Times(1)

	linesProxy := &LinesProxy{lines: mockIndicator, cache: make(map[string][]float64)}

	result := linesProxy.EMA()
	assert.Equal(t, expectedEMA, result)

	resultCached := linesProxy.EMA()
	assert.Equal(t, expectedEMA, resultCached)
}

func TestLinesProxy_RSI(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockIndicator := NewMockIndicator(ctrl)

	period := 14
	expectedRS := []float64{0.1, 0.2, 0.3}
	expectedRSI := []float64{30.5, 40.6, 50.7}

	mockIndicator.EXPECT().RSI(period).Return(expectedRS, expectedRSI).Times(1)

	linesProxy := &LinesProxy{lines: mockIndicator, cache: make(map[string][]float64)}

	rs, rsi := linesProxy.RSI(period)
	assert.Equal(t, expectedRS, rs)
	assert.Equal(t, expectedRSI, rsi)

	rsCached, rsiCached := linesProxy.RSI(period)
	assert.Equal(t, expectedRS, rsCached)
	assert.Equal(t, expectedRSI, rsiCached)
}

func TestLoadKlines(t *testing.T) {
	data := []byte(`{"pair":"BTCUSD","candles":[{"t":1703056979,"o":50000,"c":55000,"h":60000,"l":49000,"v":1000}]}`)

	lines := LoadKlines(data)

	assert.Equal(t, []float64{60000}, lines.high)
	assert.Equal(t, []float64{49000}, lines.low)
	assert.Equal(t, []float64{55000}, lines.closing)
}

func TestLinesProxy_SMA(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockIndicator := NewMockIndicator(ctrl)

	period := 10
	expectedSMA := []float64{10, 20, 30}

	mockIndicator.EXPECT().SMA(period).Return(expectedSMA).Times(1)

	linesProxy := &LinesProxy{lines: mockIndicator, cache: make(map[string][]float64)}

	result := linesProxy.SMA(period)
	assert.Equal(t, expectedSMA, result)

	resultCached := linesProxy.SMA(period)
	assert.Equal(t, expectedSMA, resultCached)
}

func TestLinesProxy_StochRSI(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockIndicator := NewMockIndicator(ctrl)

	rsiPeriod := 14
	expectedK := []float64{20.5, 30.6, 40.7}
	expectedD := []float64{50.5, 60.6, 70.7}

	mockIndicator.EXPECT().StochRSI(rsiPeriod).Return(expectedK, expectedD).Times(1)

	linesProxy := &LinesProxy{lines: mockIndicator, cache: make(map[string][]float64)}

	k, d := linesProxy.StochRSI(rsiPeriod)
	assert.Equal(t, expectedK, k)
	assert.Equal(t, expectedD, d)

	kCached, dCached := linesProxy.StochRSI(rsiPeriod)
	assert.Equal(t, expectedK, kCached)
	assert.Equal(t, expectedD, dCached)
}

func TestLinesProxy_MACD(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockIndicator := NewMockIndicator(ctrl)

	expectedMACD := []float64{1.1, 2.2, 3.3}
	expectedSignal := []float64{0.5, 1.5, 2.5}

	mockIndicator.EXPECT().MACD().Return(expectedMACD, expectedSignal).Times(1)

	linesProxy := &LinesProxy{lines: mockIndicator, cache: make(map[string][]float64)}

	macd, signal := linesProxy.MACD()
	assert.Equal(t, expectedMACD, macd)
	assert.Equal(t, expectedSignal, signal)

	macdCached, signalCached := linesProxy.MACD()
	assert.Equal(t, expectedMACD, macdCached)
	assert.Equal(t, expectedSignal, signalCached)
}

func TestLines_StochPrice(t *testing.T) {
	lines := Lines{
		high:    []float64{10, 20, 30},
		low:     []float64{5, 15, 25},
		closing: []float64{12, 18, 28},
	}

	k, d := lines.StochPrice()

	expectedK := []float64{140, 86.66666666666667, 92}
	expectedD := []float64{140, 113.33333333333334, 106.22222222222223}

	assert.Equal(t, expectedK, k)
	assert.Equal(t, expectedD, d)
}

func TestLines_RSI(t *testing.T) {
	lines := Lines{
		closing: []float64{1, 8, 12},
	}

	_, rsi := lines.RSI(30)

	expectedRSI := []float64{math.NaN(), 100, 100}

	for i, val := range expectedRSI {
		if val != rsi[i] && (!math.IsNaN(val) && !math.IsNaN(rsi[i])) {
			assert.Equal(t, 1, 0)
		}
	}
}
