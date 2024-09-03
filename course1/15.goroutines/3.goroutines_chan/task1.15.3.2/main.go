package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	url := "https://httpbin.org/get"
	parallelRequest := 5
	requestCount := 50
	result := benchRequest(url, parallelRequest, requestCount)
	for i := 0; i < requestCount; i++ {
		statusCode := <-result
		if statusCode != 200 {
			fmt.Printf(fmt.Sprintf("Ошибка при отправке запроса: %d", statusCode))
		}
	}
	fmt.Println("Все горутины завершили работу")
}

func benchRequest(url string, parallelRequest, requestCount int) <-chan int {
	results := make(chan int, requestCount)
	wg := sync.WaitGroup{}
	sem := make(chan struct{}, parallelRequest)

	for i := 0; i < requestCount; i++ {
		wg.Add(1)
		sem <- struct{}{}
		go func() {
			defer wg.Done()
			statusCode, err := httpRequest(url)
			if err != nil {
				statusCode = 0
			}
			results <- statusCode
			<-sem
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	return results
}

func httpRequest(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	return resp.StatusCode, nil
}
