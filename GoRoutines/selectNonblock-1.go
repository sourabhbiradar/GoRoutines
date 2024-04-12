package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("select unblocking-1")

	// as select blocks if there are no channels ready , avoid blocking by using 'default:'

	ch := make(chan string)

	go func() { // try commenting this GoR
		ch <- "message"
	}()

	time.Sleep(1)

	select {

	case v := <-ch:
		fmt.Println(v)

	default:
		fmt.Println("no message received")
		//when GoR commented there is no GoR sending value for 'v' to be received but still select wont block instead 'default:' is executed

	}

}
