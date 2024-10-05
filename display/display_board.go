package display

import (
	"awesomeProject/entities"
	"fmt"
)

type DisplayBoard struct {
	parkingLot *entities.ParkingLot
}

func NewDisplayBoard(parkingLot *entities.ParkingLot) *DisplayBoard {
	return &DisplayBoard{parkingLot: parkingLot}
}

func (db *DisplayBoard) ShowAvailableSpots() {
	totalAvailable := 0
	for _, floor := range db.parkingLot.Floors {
		for _, spot := range floor.Spots {
			if !spot.IsOccupied {
				totalAvailable++
			}
		}
	}
	fmt.Printf("Total available spots in the parking lot: %d\n", totalAvailable)
}
