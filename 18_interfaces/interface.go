package main

import "fmt"

type paymenter interface {
	pay(amount float32)
	// refund(amount float32, amount string)
}

type payment struct {
	gateway paymenter
}

func (p payment) makePayment(amount float32) {
	p.gateway.pay(amount)
}

type paypal struct{}

func (p paypal) pay(amount float32) {
	fmt.Println("making payment using paypal ", amount)
}

type razorpay struct{}

func (p razorpay) pay(amount float32) {
	fmt.Println("making payment using razorpay ", amount)
}

func main() {

	paypalgw := paypal{}
	newPayment := payment{
		gateway: paypalgw,
	}
	newPayment.makePayment(5000)

	razorgw := razorpay{}
	newPayment1 := payment{
		gateway: razorgw,
	}
	newPayment1.makePayment(15000)
}
