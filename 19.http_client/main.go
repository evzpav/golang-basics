package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	httpClient := http.Client{
		Timeout: time.Second * 60,
	}

	url := "https://book.omnibees.com/availability_v4/q/4937/16/2022-04-01/2022-05-31/2/0/0"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(string(bs))
}

// go run 19.http_client/main.go
