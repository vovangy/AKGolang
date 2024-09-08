package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestIndicatorSMA_GetData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExchanger := NewMockExchanger(ctrl)

	pair := "BTC_USD"
	limit := 15
	period := 5
	from := time.Now().Add(-time.Hour)
	to := time.Now()

	history := CandlesHistory{
		Candles: []Candle{
			{T: 1, C: 100.0},
			{T: 2, C: 105.0},
			{T: 3, C: 110.0},
		},
	}

	mockExchanger.EXPECT().
		GetCandlesHistory(pair, limit, from, to).
		Return(history, nil).
		Times(1)

	indicatorSMA := NewIndicatorSMA(mockExchanger)

	data, err := indicatorSMA.GetData(pair, limit, period, from, to)

	assert.NoError(t, err)
	assert.Equal(t, []float64{100.0, 102.5, 105}, data)
}

func TestIndicatorEMA_GetData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockExchanger := NewMockExchanger(ctrl)

	pair := "BTC_USD"
	limit := 15
	period := 5
	from := time.Now().Add(-time.Hour)
	to := time.Now()

	history := CandlesHistory{
		Candles: []Candle{
			{T: 1, C: 100.0},
			{T: 2, C: 105.0},
			{T: 3, C: 110.0},
		},
	}

	mockExchanger.EXPECT().
		GetCandlesHistory(pair, limit, from, to).
		Return(history, nil).
		Times(1)

	indicatorEMA := NewIndicatorEMA(mockExchanger)

	data, err := indicatorEMA.GetData(pair, limit, period, from, to)

	assert.NoError(t, err)
	assert.Equal(t, []float64{100.0, 101.66666666666667, 104.44444444444446}, data)
}

func TestGeneralIndicator_GetData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockIndicator := NewMockIndicatorer(ctrl)

	pair := "BTC_USD"
	period := 5
	from := time.Now().Add(-time.Hour)
	to := time.Now()

	mockIndicator.EXPECT().
		GetData(pair, 15, period, from, to).
		Return([]float64{100.0, 105.0, 110.0}, nil).
		Times(1)

	generalIndicator := &GeneralIndicator{}
	data, err := generalIndicator.GetData(pair, period, from, to, mockIndicator)

	assert.NoError(t, err)
	assert.Equal(t, []float64{100.0, 105.0, 110.0}, data)
}

func TestExmo_GetCandlesHistory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/candles_history", r.URL.Path)
		assert.Contains(t, r.URL.RawQuery, "symbol=BTC_USD")
		assert.Contains(t, r.URL.RawQuery, "resolution=15")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"candles":[{"t":1609459200, "c":100.0}, {"t":1609459260, "c":105.0}]}`))
	}))
	defer server.Close()

	exmo := NewExmo(func(e *Exmo) {
		e.url = server.URL
	})

	start := time.Unix(1609459200, 0)
	end := time.Unix(1609459260, 0)

	history, err := exmo.GetCandlesHistory("BTC_USD", 15, start, end)

	assert.NoError(t, err)
	assert.Len(t, history.Candles, 2)
	assert.Equal(t, 100.0, history.Candles[0].C)
	assert.Equal(t, 105.0, history.Candles[1].C)
}

func TestGetRespBody(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"candles":[{"t":1609459200, "c":100.0}]}`))
	}))
	defer server.Close()

	exmo := NewExmo(func(e *Exmo) {
		e.url = server.URL
	})

	body, err := GetRespBody(exmo, server.URL)

	assert.NoError(t, err)
	assert.NotNil(t, body)
	assert.Contains(t, string(body), `"candles"`)
}
