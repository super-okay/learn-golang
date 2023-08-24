package dataStructures

import (
	"fmt"
)

type NewCar struct {
	seats  int
	color  string
	weight int
	make   string
	model  string
}

type SUV struct {
	NewCar
	towingCapacity int
	payload        int
}

type Truck struct {
	NewCar
	towingCapacity int
	payload        int
	bedLength      int
}

func CreateNewCar() {
	myTacoma := Truck{
		towingCapacity: 6000,
		payload:        1500,
		bedLength:      5,
	}
	myTacoma.seats = 5
	myTacoma.color = "cement"
	myTacoma.weight = 4000
	myTacoma.make = "Toyota"
	myTacoma.model = "Tacoma"
	fmt.Println(myTacoma)
}
