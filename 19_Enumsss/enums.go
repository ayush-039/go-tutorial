package main

import "fmt"

// enumerated types

type OrderStatus int

const (
	Received OrderStatus = iota
	Confirmed
	Prepared
	Delivered
)


func chageOrderStatus(status string){
	fmt.Println("changing order status to", status)
}

func main()  {
	
	chageOrderStatus("confirmed")
}