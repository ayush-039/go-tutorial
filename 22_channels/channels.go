package main

import (
	"fmt"
	"time"
)

func processNum(numchan chan int) {

	for num := range numchan {
		fmt.Println("processing number ", num)
		time.Sleep(time.Second)
	}
}

func sum(result chan int, num1 int, num2 int) {
	numresult := num1 + num2
	result <- numresult
}

// gorouting synchronizer
func task(done chan bool) {
	defer func() { done <- true }()
	fmt.Println("precessing....")
}

func emailSender(emailChan chan string, done chan bool) {
	defer func() { done <- true }()

	for email := range emailChan {
		fmt.Println("sending email to", email)
		time.Sleep(time.Second)
	}
}

func main() {



	

	// chan1 := make(chan int)
	// chan2 := make(chan string)

	// go func() {
	// 	chan1 <- 10
	// }()
	// go func() {
	// 	chan2 <- "pong"
	// }()

	// for i := 0; i < 2; i++ {
	// 	select {
	// 	case chan1val := <-chan1:
	// 		fmt.Println("Received data from chan1", chan1val)
	// 	case chan2val := <-chan2:
	// 		fmt.Println("Received data from chan2", chan2val)

	// 	}
	// }

	// emailChan := make(chan string, 100)
	// done := make(chan bool)

	// go emailSender(emailChan, done)
	// for i := 0; i < 10; i++ {
	// 	emailChan <- fmt.Sprintf("%d@gmail.com", i+1)
	// }
	// // emailSender(emailChan, done)
	// fmt.Println("done sending.")

	// close(emailChan) // important find why??
	// <-done

	// done := make(chan bool)
	// go task(done)
	// <-done // blocking call

	// numchan := make(chan int)
	// go processNum(numchan)

	// for {
	// 	numchan <- rand.Intn(100)
	// }
	// numchan <- 6
	// numchan <- 7
	// numchan <- 8
	// numchan <- 9
	// numchan <- 10
	// numchan <- 11
	// numchan <- 12

	// result := make(chan int)
	// go sum(result, 4, 5)
	// res := <-result
	// fmt.Println(res)

	// time.Sleep(time.Second * 2)
}
