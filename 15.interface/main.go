package main

import "fmt"

type Notifier interface {
	SendMessage(string) error
}

type Service1 struct{}

func (s1 *Service1) SendMessage(msg string) error {
	fmt.Println("Sending msg from s1:", msg)
	return nil
}

type Service2 struct{}

func (s2 *Service2) SendMessage(msg string) error {
	fmt.Println("Sending msg from s2:", msg)
	return nil
}

func SendGreetings(notifier Notifier) {
	notifier.SendMessage("Hello World")
}

func main() {

	s1 := &Service1{}
	SendGreetings(s1)

	s2 := &Service2{}
	SendGreetings(s2)
}

//go run 15.interface/main.go 