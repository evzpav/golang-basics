package main

import "fmt"

func main() {

	names := []string{"John", "Mary", "Katie", "Harry"}

	newPerson := "Julia"
	names = append(names, newPerson)

	fmt.Printf("items: %d\n", len(names))
	fmt.Printf("firstItem: %s\n", names[0])
	fmt.Printf("lastItem: %s\n", names[len(names)-1])
	fmt.Printf("panic: %s\n", names[5])
}

// go run 06.slices/main.go
