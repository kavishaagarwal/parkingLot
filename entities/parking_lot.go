package entities

type ParkingLot struct {
	Floors []Floor
}

func NewParkingLot() *ParkingLot {
	return &ParkingLot{}
}

func (pl *ParkingLot) AddFloor(spotTypes map[string]int) {
	floor := NewFloor(spotTypes, len(pl.Floors))
	pl.Floors = append(pl.Floors, *floor)
}

func (pl *ParkingLot) RemoveFloor(index int) {
	if index >= 0 && index < len(pl.Floors) {
		pl.Floors = append(pl.Floors[:index], pl.Floors[index+1:]...)
	}
}

func (pl *ParkingLot) AddSpot(floorIndex int, spotType string, isElectric bool) {
	if floorIndex >= 0 && floorIndex < len(pl.Floors) {
		newSpot := NewParkingSpot(spotType, 1, isElectric)
		pl.Floors[floorIndex].Spots = append(pl.Floors[floorIndex].Spots, newSpot)
	}
}

func (pl *ParkingLot) GetAvailableSpots() (available int) {
	for _, floor := range pl.Floors {
		for _, spot := range floor.Spots {
			if !spot.IsOccupied {
				available++
			}
		}
	}
	return
}
