package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Order struct {
	ID       int
	Complete bool
}

var (
	orders              []*Order
	completeOrders      map[int]bool
	wg                  sync.WaitGroup
	processTimes        chan time.Duration
	sinceProgramStarted time.Duration
	count, limitCount   int
)

func main() {
	count = 30
	limitCount = 5
	processTimes = make(chan time.Duration, count)
	orders = GenerateOrders(count)
	completeOrders = GenerateCompleteOrders(count)
	programStart := time.Now()
	LimitSpawnOrderProcessing(limitCount)
	wg.Wait()
	sinceProgramStarted = time.Since(programStart)
	go func() {
		time.Sleep(1 * time.Second)
		close(processTimes)
	}()
	checkTimeDifference(limitCount)
}

func checkTimeDifference(limitCount int) {
	// do not edit
	var averageTime time.Duration
	var orderProcessTotalTime time.Duration
	var orderProcessedCount int
	for v := range processTimes {
		orderProcessedCount++
		orderProcessTotalTime += v
	}
	if orderProcessedCount != count {
		panic("orderProcessedCount != count")
	}
	averageTime = orderProcessTotalTime / time.Duration(orderProcessedCount)
	fmt.Println("orderProcessTotalTime", orderProcessTotalTime/time.Second)
	fmt.Println("averageTime", averageTime/time.Second)
	fmt.Println("sinceProgramStarted", sinceProgramStarted/time.Second)
	fmt.Println("sinceProgramStarted average", sinceProgramStarted/(time.Duration(orderProcessedCount)*time.Second))
	fmt.Println("orderProcessTotalTime - sinceProgramStarted", (orderProcessTotalTime-sinceProgramStarted)/time.Second)
	if (orderProcessTotalTime/time.Duration(limitCount)-sinceProgramStarted)/time.Second > 0 {
		panic("(orderProcessTotalTime-sinceProgramStarted)/time.Second > 0")
	}
}

func LimitSpawnOrderProcessing(limitCount int) {
	limit := make(chan struct{}, limitCount)
	for _, order := range orders {
		wg.Add(1)
		limit <- struct{}{}
		go func(order *Order) {
			defer wg.Done()
			t := time.Now()
			OrderProcessing(order)
			processTimes <- time.Since(t)
			<-limit
		}(order)
	}
}

func OrderProcessing(order *Order) {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	if complete, ok := completeOrders[order.ID]; ok && complete {
		order.Complete = true
	}
}

func GenerateOrders(count int) []*Order {
	orders := make([]*Order, count)
	for i := 0; i < count; i++ {
		orders[i] = &Order{ID: i, Complete: false}
	}
	return orders
}

func GenerateCompleteOrders(count int) map[int]bool {
	completeOrders := make(map[int]bool, count)
	for i := 0; i < count; i++ {
		completeOrders[i] = rand.Intn(2) == 0
	}
	return completeOrders
}
