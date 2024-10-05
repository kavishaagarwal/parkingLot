package main

import (
	"awesomeProject/admin"
	"awesomeProject/display"
	"awesomeProject/entities"
	"awesomeProject/entryexit"
	"encoding/json"
	"fmt"
)

func main() {
	parkingLot := entities.NewParkingLot()

	admin := admin.NewParkingAdmin(parkingLot)

	// Adding floors with different parking spot configurations
	admin.AddFloor(map[string]int{
		"Compact":     5,
		"Regular":     10,
		"Handicapped": 2,
		"Electric":    3,
	})
	admin.AddFloor(map[string]int{
		"Large":      10,
		"Motorcycle": 5,
		"Electric":   3,
	})

	ss, _ := json.Marshal(parkingLot)
	fmt.Println("ss:", string(ss))

	entryExit := entryexit.NewSimpleEntryExit(parkingLot)
	// Show initial available spots
	displayBoard := display.NewDisplayBoard(parkingLot)
	displayBoard.ShowAvailableSpots()
	displayBoard.DisplayAvailableSpots()

	car := &entities.Car{}
	ticket := entryExit.Enter(car)
	if ticket != nil {
		// Pay with cash
		//entryExit.Exit(ticket, &entities.Cash{})

		// Or pay with a credit card
		entryExit.Exit(ticket, &entities.CreditCard{})
	}
}
