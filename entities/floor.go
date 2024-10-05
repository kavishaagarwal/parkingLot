package entities

type Floor struct {
	Spots   []*ParkingSpot
	FloorID int
}

func NewFloor(spotTypes map[string]int, floorID int) *Floor {
	spots := make([]*ParkingSpot, 0)
	for t, s := range spotTypes {
		isElectric := t == "Electric"
		for i := 0; i < s; i++ {
			spots = append(spots, NewParkingSpot(t, 1, isElectric))
		}
	}
	return &Floor{Spots: spots, FloorID: floorID}
}
