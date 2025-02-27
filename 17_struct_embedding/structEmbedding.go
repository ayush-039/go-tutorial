package main

import (
	"fmt"
	"time"
)

type Customer struct {
	name    string
	phoneNo string
}

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
	Customer
}

func newCustomer(namee string, phoneNom string) *Customer {
	c := Customer{
		name:    namee,
		phoneNo: phoneNom,
	}
	return &c
}

func newOrder(id string, amount float32, status string, customer Customer) *order {
	myOrder := order{
		id:        id,
		amount:    amount,
		status:    status,
		createdAt: time.Now(),
		Customer:  customer,
	}
	return &myOrder
}

func main() {

	cust1 := newCustomer("Ayush Barot", "+91 9909764364")

	myOrder := newOrder("1", 30.75, "pending", *cust1)
	fmt.Println(myOrder)
}
