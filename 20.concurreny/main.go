package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var websites = []string{
	"https://github.com/",
	"https://gitlab.com/",
	"https://stackoverflow.com/",
	"https://www.linkedin.com/",
	"https://medium.com/",
	"https://www.udemy.com/",
}

func main() {
	var wg sync.WaitGroup

	result := make(chan string)

	for _, url := range websites {
		wg.Add(1)
		go sendRequest(&wg, url, result)
	}

	go func() {
		for {
			res := <-result
			fmt.Println(res)
		}
	}()

	wg.Wait()
}

func sendRequest(wg *sync.WaitGroup, url string, result chan<- string) {
	defer wg.Done()
	httpClient := &http.Client{
		Timeout: time.Second * time.Duration(5),
	}

	req, _ := http.NewRequest(http.MethodGet, url, nil)
	res, err := httpClient.Do(req)
	if err != nil {
		result <- fmt.Sprintf("[nok] %s - error [%v]\n", url, err)
		return
	}

	if res.StatusCode >= 300 {
		result <- fmt.Sprintf("[nok] %s - status [%d]\n", url, res.StatusCode)
		return
	}

	result <- fmt.Sprintf("[ok] %s - status [%d]\n", url, res.StatusCode)
}

// // sync example - without concurrency
// func main() {
// 	for _, url := range websites {
// 		httpClient := &http.Client{
// 			Timeout: time.Second * time.Duration(5),
// 		}

// 		req, _ := http.NewRequest(http.MethodGet, url, nil)
// 		res, err := httpClient.Do(req)
// 		if err != nil {
// 			fmt.Printf("[nok] %s - error [%v]\n", url, err)
// 			return
// 		}

// 		if res.StatusCode >= 300 {
// 			fmt.Printf("[nok] %s - status [%d]\n", url, res.StatusCode)
// 			return
// 		}

// 		fmt.Printf("[ok] %s - status [%d]\n", url, res.StatusCode)
// 	}

// }

// go run 20.concurreny/main.go
