package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("select order")

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2)
		ch1 <- "one"
	}()

	// go func(ch2 <-chan string){
	//    time.Sleep(3)
    // }("two")

	go func() {
		time.Sleep(3)
		//when time.Sleep() > time.Sleep() of GoR1 ,GoR1 executes first
		//when time.Sleep() < time.Sleep() of GoR1 ,GoR2 executes first
		ch2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case m1, ok := <-ch1:
			fmt.Println(m1, ok)
		case m2 := <-ch2:
			fmt.Println(m2)
		}
	}
}
