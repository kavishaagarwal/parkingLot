package admin

import (
	"awesomeProject/entities"
)

type ParkingAdmin struct {
	parkingLot *entities.ParkingLot
}

func NewParkingAdmin(parkingLot *entities.ParkingLot) *ParkingAdmin {
	return &ParkingAdmin{parkingLot: parkingLot}
}

func (pa *ParkingAdmin) AddFloor(spotTypes map[string]int) {
	pa.parkingLot.AddFloor(spotTypes)
}

func (pa *ParkingAdmin) RemoveFloor(index int) {
	pa.parkingLot.RemoveFloor(index)
}

func (pa *ParkingAdmin) AddParkingSpot(floorIndex int, spotType string, isElectric bool) {
	pa.parkingLot.AddSpot(floorIndex, spotType, isElectric)

}
