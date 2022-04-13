package main

import "fmt"

func main() {
	s := make([]int, 3)

	fmt.Println("slice with 3 zeros:", s)

	var numbers []int

	for i := 0; i < 5; i++ {
		numbers = append(numbers, i)
	}

	fmt.Println("Numbers:", numbers)

	sequence := []int{1, 2, 3, 4, 5, 6}

	fmt.Println("first 3 items:", sequence[:3])
	fmt.Println("last 3 items:", sequence[3:])
	fmt.Println("last item:", sequence[len(sequence)-1])

}

// go run 14.slices/main.go
