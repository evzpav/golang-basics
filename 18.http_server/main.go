package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "hey Asksuite\n")
}

func main() {

	http.HandleFunc("/hello", hello)

	http.ListenAndServe(":8989", nil)
}

// go run 18.http_server/main.go
// curl localhost:8989/hello
