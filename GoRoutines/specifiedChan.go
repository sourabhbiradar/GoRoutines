package main

import (
	"fmt"
	"time"
)

func one(ch1 chan<- string) {
	ch1 <- "message from one"
}

func two(ch1 <-chan string) {
	fmt.Println("received at two :", <-ch1)
}

func three(ch3 <-chan string, ch4 chan<- string) {
	//	ch3 <- "message from three"
	// invalid operation: cannot send to receive-only channel ch1 (variable of type <-chan string)
	//fmt.Println("from four :", <-ch3)

	ch4 <- "message from three"
	//	fmt.Println("receiver at three :",<-ch4)
	// invalid operation: cannot receive from send-only channel ch2 (variable of type chan<- string)

	//
	//
	fmt.Println("received at three :", <-ch3) //line 19 doesnt work bt this workes ???
	//
	//

}

func four(ch3 chan<- string, ch4 <-chan string) {

	fmt.Println("received at four :", <-ch4)

	ch3 <- "message from four"

}

func main() {

	fmt.Println("channel specific type")

	ch1 := make(chan string)
	ch3 := make(chan string)
	ch4 := make(chan string)

	go one(ch1)
	go two(ch1)
	go three(ch3, ch4)
	go four(ch3, ch4)

	time.Sleep(4)

}
