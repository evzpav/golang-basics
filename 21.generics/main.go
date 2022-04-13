package main

import "fmt"
// Needs Go 1.18

// func Print(s []string) {
// 	for _, v := range s {
// 		fmt.Print(v)
// 	}
// }

// func Print(s []int) {
// 	for _, v := range s {
// 		fmt.Print(v)
// 	}
// }

func Print[T any](s []T) {
	for _, v := range s {
		fmt.Print(v)
	}
}


func main() {
	Print([]string{"Hello, ", "playground\n"})
	Print([]int{1,2,3})
}

// go run 21.generics/main.go 