package main

import "fmt"

func main() {
	var counter int = 1
	var price float64 = 1000.12345678
	var name string = "Dora"
	var isReserved bool = true

	fmt.Printf("int: %d\n", counter)
	fmt.Printf("float64: %f\n", price)
	fmt.Printf("float64 - 2 decimals: %.2f\n", price)
	fmt.Printf("string: %s\n", name)
	fmt.Printf("bool: %t\n", isReserved)
}

// go run 03.variables/main.go
