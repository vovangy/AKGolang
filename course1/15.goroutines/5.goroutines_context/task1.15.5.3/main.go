package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func CallService() string {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	data := make(chan string, 2)

	serviceLocator := NewServiceLocator()

	go func() {
		defer func() { serviceLocator.slow = true }()
		result, _ := serviceLocator.SlowService(ctx)
		data <- result
		fmt.Println("slow service done")
	}()

	go func() {
		defer func() { serviceLocator.fast = true }()
		result, _ := serviceLocator.FastService(ctx)
		data <- result
		fmt.Println("fast service done")
	}()

	var result string
	select {
	case res := <-data:
		result = res
	case <-ctx.Done():
		if len(data) > 1 {
			panic("error: more than one result")
		}
		if !serviceLocator.slow {
			panic("error: slow service called")
		}
		if !serviceLocator.fast {
			panic("error: fast service not called")
		}
	}

	checkService(serviceLocator)

	return result
}

type ServiceLocator struct {
	client *http.Client
	fast   bool
	slow   bool
}

func NewServiceLocator() *ServiceLocator {
	return &ServiceLocator{client: &http.Client{Timeout: 5 * time.Second}}
}

func (s *ServiceLocator) FastService(ctx context.Context) (string, error) {
	return s.doRequest(ctx, "https://api.exmo.com/v1/ticker")
}

func (s *ServiceLocator) SlowService(ctx context.Context) (string, error) {
	defer func() { s.slow = true }()
	time.Sleep(2 * time.Second)
	return s.doRequest(ctx, "https://api.exmo.com/v1/ticker")
}

func (s *ServiceLocator) doRequest(ctx context.Context, url string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func checkService(s *ServiceLocator) {
	if !s.slow {
		panic("error: slow service called")
	}
	if !s.fast {
		panic("error: fast service not called")
	}
}

func main() {
	res := CallService()
	fmt.Println(res)
}
