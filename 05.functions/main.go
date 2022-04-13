package main

import "fmt"

func addInt(x, y int) int {
	return x + y
}

func addFloat64(x, y float64) float64 {
	return x + y
}

func resolvePayment(total, payment, feePercentage float64) (float64, float64) {
	fee := payment * feePercentage

	remaining := total - payment - fee

	return remaining, fee
}

func main() {
	var sumInt = addInt(1, 2)
	sumFloat := addFloat64(1.1, 2.2)

	fmt.Printf("sumInt: %d\n", sumInt)
	fmt.Printf("sumFloat: %f\n", sumFloat)

	remaining, fee := resolvePayment(10000.0, 500.0, 0.1)
	fmt.Printf("remaining:%.f; fee: %f\n", remaining, fee)
}

// go run 05.functions/main.go
