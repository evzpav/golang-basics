package main

import "fmt"

func main() {
	counter := 1
	price := 1000.12345678
	name := "Dora"
	isReserved := true // initializing var
	isReserved = false // setting value to var

	fmt.Printf("int: %d\n", counter)
	fmt.Printf("float64: %f\n", price)
	fmt.Printf("float64 - 2 decimals: %.2f\n", price)
	fmt.Printf("string: %s\n", name)
	fmt.Printf("bool: %t\n", isReserved)

}

// go run 04.short_variables/main.go
