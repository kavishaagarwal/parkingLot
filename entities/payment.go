package entities

import "fmt"

// Payment is the strategy interface
type Payment interface {
	ProcessPayment(amount float64) bool
}

// Cash is a concrete strategy for cash payments
type Cash struct{}

func (c *Cash) ProcessPayment(amount float64) bool {
	fmt.Printf("Processed cash payment of $%.2f\n", amount)
	return true
}

// CreditCard is a concrete strategy for credit card payments
type CreditCard struct{}

func (cc *CreditCard) ProcessPayment(amount float64) bool {
	fmt.Printf("Processed credit card payment of $%.2f\n", amount)
	return true
}

// PaymentContext holds a reference to a Payment strategy
type PaymentContext struct {
	PaymentMethod Payment
}

// SetPaymentMethod allows changing the payment strategy at runtime
func (pc *PaymentContext) SetPaymentMethod(paymentMethod Payment) {
	pc.PaymentMethod = paymentMethod
}

// ExecutePayment uses the current payment strategy to process the payment
func (pc *PaymentContext) ExecutePayment(amount float64) bool {
	if pc.PaymentMethod == nil {
		fmt.Println("No payment method set")
		return false
	}
	return pc.PaymentMethod.ProcessPayment(amount)
}
