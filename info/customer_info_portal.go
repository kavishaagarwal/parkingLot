package info

//
//import (
//	"awesomeProject/entities"
//	"fmt"
//)
//
//type CustomerInfoPortal struct {
//	parkingLot *entities.ParkingLot
//}
//
//func NewCustomerInfoPortal(parkingLot *entities.ParkingLot) *CustomerInfoPortal {
//	return &CustomerInfoPortal{parkingLot: parkingLot}
//}
//
//func (cip *CustomerInfoPortal) DisplayAvailableSpots() {
//	for i, floor := range cip.parkingLot.Floors {
//		fmt.Printf("Floor %d:\n", i)
//		for _, spot := range floor.Spots {
//			if !spot.IsOccupied {
//				fmt.Printf("Available %s spot\n", spot.SpotType)
//			}
//		}
//	}
//}
