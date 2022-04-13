package main

import "fmt"

func main() {

	var numbers [5]int

	for i := 0; i < 5; i++ {
		numbers[i] = i
	}

	fmt.Println("Numbers:", numbers)
	fmt.Println("Total Numbers:", len(numbers))

	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

}

// go run 13.arrays/main.go
