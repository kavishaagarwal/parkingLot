package entities

import "time"

type Ticket struct {
	Vehicle   Vehicle
	EntryTime time.Time
}

func (t *Ticket) CalculateFee() float64 {
	duration := time.Since(t.EntryTime)
	hours := int(duration.Hours())
	var fee float64

	if hours == 0 {
		return 4.0 // For the first hour
	} else if hours == 1 {
		fee = 4.0
	} else if hours <= 3 {
		fee = 4.0 + float64(hours-1)*3.5
	} else {
		fee = 4.0 + 3.5*2 + float64(hours-3)*2.5
	}
	return fee
}
