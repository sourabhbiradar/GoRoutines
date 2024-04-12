package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("select unblocking-2")

	ch := make(chan string)

	go func() {

		for i := 0; i < 2; i++ {
			time.Sleep(3)
			ch <- "message"
		}

	}()

	time.Sleep(2500 * time.Millisecond)

	for i := 0; i < 2; i++ {
		select {

		case m := <-ch:
			fmt.Println(m)
		default:
			fmt.Println("message not yet received")
		}

		fmt.Println("processing...")

	}

}
