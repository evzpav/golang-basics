package main

import "fmt"

func main() {
	var counter int
	var price float64
	var name string
	var isReserved bool

	fmt.Printf("int: %d\n", counter)
	fmt.Printf("float64: %.f\n", price)
	fmt.Printf("string: %s\n", name)
	fmt.Printf("bool: %t\n", isReserved)

	// Reference: https://pkg.go.dev/fmt#pkg-types
}

// go run 02.values/main.go
