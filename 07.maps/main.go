package main

import "fmt"

func main() {

	fruitWeights := map[string]int{
		"Apple":  45,
		"Mango":  24,
		"Orange": 34,
	}

	fruitWeights["Banana"] = 50

	mangoWeight, ok := fruitWeights["Mango"]
	if ok {
		fmt.Printf("Mango weight: %d\n", mangoWeight)
	}

	_, ok = fruitWeights["Pineapple"]
	if !ok {
		fmt.Println("Pineapple not found in map")
	}

	fmt.Printf("items in map: %d\n", len(fruitWeights))
	fmt.Printf("%#v\n", fruitWeights)
}

// go run 07.maps/main.go
