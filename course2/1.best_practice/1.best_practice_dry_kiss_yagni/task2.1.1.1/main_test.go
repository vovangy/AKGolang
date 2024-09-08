package main

import (
	"reflect"
	"testing"
)

func TestStatisticProfit_GetAverageProfit(t *testing.T) {
	product := &Product{
		Sells:         []float64{10, 20, 30},
		Buys:          []float64{5, 15, 25},
		CurrentPrice:  35,
		ProfitPercent: 10,
	}
	statProfit := NewStatisticProfit(WithAverageProfit).(*StatisticProfit)
	statProfit.SetProduct(product)

	expected := 5.0
	if avgProfit := statProfit.GetAverageProfit(); avgProfit != expected {
		t.Errorf("Expected average profit to be %v, but got %v", expected, avgProfit)
	}
}

func TestStatisticProfit_GetAverageProfitPercent(t *testing.T) {
	product := &Product{
		Sells:         []float64{10, 20, 30},
		Buys:          []float64{5, 15, 25},
		CurrentPrice:  35,
		ProfitPercent: 10,
	}
	statProfit := NewStatisticProfit(WithAverageProfit, WithAverageProfitPercent).(*StatisticProfit)
	statProfit.SetProduct(product)

	expected := 100 - 66.66666666666667
	if avgProfitPercent := statProfit.GetAverageProfitPercent(); avgProfitPercent != expected {
		t.Errorf("Expected average profit percent to be %v, but got %v", expected, avgProfitPercent)
	}
}

func TestStatisticProfit_GetCurrentProfit(t *testing.T) {
	product := &Product{
		Sells:         []float64{10, 20, 30},
		Buys:          []float64{5, 15, 25},
		CurrentPrice:  35,
		ProfitPercent: 10,
	}
	statProfit := NewStatisticProfit(WithCurrentProfit).(*StatisticProfit)
	statProfit.SetProduct(product)

	expected := 3.5 // 35 - 35 * (100-10)/100
	if currentProfit := statProfit.GetCurrentProfit(); currentProfit != expected {
		t.Errorf("Expected current profit to be %v, but got %v", expected, currentProfit)
	}
}

func TestStatisticProfit_GetDifferenceProfit(t *testing.T) {
	product := &Product{
		Sells:         []float64{10, 20, 30},
		Buys:          []float64{5, 15, 25},
		CurrentPrice:  35,
		ProfitPercent: 10,
	}
	statProfit := NewStatisticProfit(WithDifferenceProfit).(*StatisticProfit)
	statProfit.SetProduct(product)

	expected := 15.0 // 35 - (10+20+30)/3
	if diffProfit := statProfit.GetDifferenceProfit(); diffProfit != expected {
		t.Errorf("Expected difference profit to be %v, but got %v", expected, diffProfit)
	}
}

func TestStatisticProfit_GetAllData(t *testing.T) {
	product := &Product{
		Sells:         []float64{10, 20, 30},
		Buys:          []float64{5, 15, 25},
		CurrentPrice:  35,
		ProfitPercent: 10,
	}
	statProfit := NewStatisticProfit(
		WithAverageProfit,
		WithAverageProfitPercent,
		WithCurrentProfit,
		WithDifferenceProfit,
		WithAllData,
	).(*StatisticProfit)
	statProfit.SetProduct(product)

	expected := []float64{33.33333333333333, 3.5, 15.0}
	if allData := statProfit.GetAllData(); !reflect.DeepEqual(allData, expected) {
		t.Errorf("Expected all data to be %v, but got %v", expected, allData)
	}
}

func TestStatisticProfit_Average(t *testing.T) {
	statProfit := &StatisticProfit{}
	prices := []float64{10, 20, 30}
	expected := 20.0

	if avg := statProfit.Average(prices); avg != expected {
		t.Errorf("Expected average to be %v, but got %v", expected, avg)
	}
}

func TestStatisticProfit_Sum(t *testing.T) {
	statProfit := &StatisticProfit{}
	prices := []float64{10, 20, 30}
	expected := 60.0

	if sum := statProfit.Sum(prices); sum != expected {
		t.Errorf("Expected sum to be %v, but got %v", expected, sum)
	}
}
