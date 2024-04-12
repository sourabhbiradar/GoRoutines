package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("select order-2")

	ch1 := make(chan string)

	go func() {
		time.Sleep(2)
		//2 sec
		ch1 <- "one"
	}()

	select {
	case m1, ok := <-ch1:
		fmt.Println(m1, ok)
	case <-time.After(3):
		//when time is > 2 sec case m1 executes,
		//when time is < 2 sec this case executes
		fmt.Println("timeout")
	}

}
