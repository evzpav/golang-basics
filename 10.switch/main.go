package main

import "fmt"

func main() {
	// switch don't need break, and it only goes to default if no case match

	i := 2
	fmt.Print("Write ", i, " as ")

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("zero")
	}

	userType := ""

	switch userType {
	case "company":
		fmt.Println("company")
	case "customer":
		fmt.Println("customer")
	default:
		fmt.Println("user type not found")
	}

}

// go run 10.switch/main.go
