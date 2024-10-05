package entities

type ParkingSpot struct {
	SpotType   string
	Size       int
	IsOccupied bool
	IsElectric bool
}

func NewParkingSpot(spotType string, size int, isElectric bool) *ParkingSpot {
	return &ParkingSpot{SpotType: spotType, Size: size, IsOccupied: false, IsElectric: isElectric}
}
