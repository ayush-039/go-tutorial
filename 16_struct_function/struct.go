package main

import (
	"fmt"
	"time"
)

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // has nanosecond precision
}

// receiver type
func (o *order) changeStatus(status string) {
	o.status = status
}

func (o *order) getAmount() float32 {
	return o.amount
}

func newOrder(id string, amount float32, status string) *order {
	myOrder := order{
		id:        id,
		amount:    amount,
		status:    status,
		createdAt: time.Now(),
	}
	return &myOrder
}

func main() {
	order1 := order{
		id:        "1",
		amount:    50.00,
		status:    "pending",
		createdAt: time.Now(),
	}
	fmt.Println(order1)
	order1.changeStatus("delivered")
	fmt.Println(order1)
	fmt.Println(order1.getAmount())

	myOrder := newOrder("1", 30.75, "pending")
	fmt.Println(myOrder)
}
