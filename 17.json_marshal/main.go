package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Person struct {
	Name         string        `json:"name"`
	Age          int           `json:"age"`
	Address      Address       `json:"address"`
	Reservations []Reservation `json:"reservations"`
}

type Address struct {
	StreetName string `json:"sn"`
	Number     int    `json:"n"`
	City       string `json:"c"`
}

type Reservation struct {
	HotelID  int       `json:"hotelId"`
	Room     int       `json:"roomNumber"`
	CheckIn  time.Time `json:"checkIn"`
	CheckOut time.Time `json:"checkOut"`
}

func main() {

	var person = Person{
		Name: "Mary",
		Age:  39,
		Address: Address{
			StreetName: "Rua Lauro Linhares",
			Number:     122,
			City:       "Florianopolis",
		},
		Reservations: []Reservation{
			{
				HotelID:  1,
				Room:     101,
				CheckIn:  time.Date(2022, 04, 29, 14, 0, 0, 0, time.UTC),
				CheckOut: time.Date(2022, 04, 30, 11, 0, 0, 0, time.UTC),
			},
		},
	}

	byteSlice, err := json.Marshal(&person)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Person: %+v\n", string(byteSlice))
}

// go run 17.json_marshal/main.go
