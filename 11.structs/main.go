package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func newPerson(name string) *Person {
	return &Person{Name: name}
}

func main() {

	seanPerson := Person{
		Name: "Sean",
		Age:  50,
	}
	fmt.Printf("Name: %s; Age: %d\n", seanPerson.Name, seanPerson.Age)

	seanPerson.Age = 51
	fmt.Printf("Name: %s; Age: %d\n", seanPerson.Name, seanPerson.Age)

	johnPerson := newPerson("John")

	johnPerson.Age = 30

	fmt.Printf("Name: %s; Age: %d\n", johnPerson.Name, johnPerson.Age)

	jp := johnPerson
	jp.Age = 99
	fmt.Printf("Name: %s; Age: %d\n", jp.Name, jp.Age)

}

// go run 11.structs/main.go
