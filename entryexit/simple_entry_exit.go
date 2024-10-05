package entryexit

import (
	"awesomeProject/entities"
	"fmt"
	"time"
)

type SimpleEntryExit struct {
	parkingLot *entities.ParkingLot
}

func NewSimpleEntryExit(parkingLot *entities.ParkingLot) *SimpleEntryExit {
	return &SimpleEntryExit{parkingLot: parkingLot}
}

func (ee *SimpleEntryExit) Enter(vehicle entities.Vehicle) *entities.Ticket {
	if ee.parkingLot.GetAvailableSpots() == 0 {
		fmt.Println("Parking lot is full!")
		return nil
	}
	for i := range ee.parkingLot.Floors {
		for j := range ee.parkingLot.Floors[i].Spots {
			if !ee.parkingLot.Floors[i].Spots[j].IsOccupied && ee.parkingLot.Floors[i].Spots[j].Size >= vehicle.GetSize() {
				ee.parkingLot.Floors[i].Spots[j].IsOccupied = true
				ticket := &entities.Ticket{Vehicle: vehicle, EntryTime: time.Now()}
				fmt.Printf("Vehicle entered at floor %d, spot %d\n", i, j)
				return ticket
			}
		}
	}
	fmt.Println("No available parking spot")
	return nil
}

func (ee *SimpleEntryExit) Exit(ticket *entities.Ticket, paymentMethod entities.Payment) bool {
	fee := ticket.CalculateFee()
	paymentContext := &entities.PaymentContext{}
	paymentContext.SetPaymentMethod(paymentMethod)

	if paymentContext.ExecutePayment(fee) {
		for i := range ee.parkingLot.Floors {
			for j := range ee.parkingLot.Floors[i].Spots {
				if ee.parkingLot.Floors[i].Spots[j].IsOccupied {
					ee.parkingLot.Floors[i].Spots[j].IsOccupied = false
					fmt.Printf("Vehicle exited from floor %d, spot %d\n", i, j)
					return true
				}
			}
		}
	}
	return false
}
