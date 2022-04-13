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

	input := `{	
		"name": "Mary",	
		"age": 39,
		"address": {
			"sn": "Rua Lauro Linhares",
			"n": 122,
			"c": "Florianopolis"
		},
		"reservations":[
			{
				"hotelId":1,
				"roomNumber":101,
				"checkIn":"2022-04-29T14:00:00Z",
				"checkOut":"2022-04-30T11:00:00Z"
			}
		]	
	}`

	var person Person
	err := json.Unmarshal([]byte(input), &person)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Person: %+v\n", person)
}

// go run 16.json_unmarshal/main.go
