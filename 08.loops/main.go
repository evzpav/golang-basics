package main

import "fmt"

func main() {
	fmt.Println("--- Loop 1")

	for i := 0; i < 10; i++ {
		fmt.Printf("i:%d\n", i)
	}

	fmt.Println("--- Loop 2")

	i := 0
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	fmt.Println("--- Loop range slice")

	var names = []string{"John", "Mary", "Katie", "Harry"}
	for i, n := range names {
		fmt.Printf("i:%d - name: %s\n", i, n)
	}

	fruitWeights := map[string]int{
		"Apple":  45,
		"Mango":  24,
		"Orange": 34,
	}

	fmt.Println("--- Loop range map")
	for fruitName, weight := range fruitWeights {
		fmt.Printf("fruit:%s - weight: %d\n", fruitName, weight)
	}

}

// go run 08.loops/main.go
