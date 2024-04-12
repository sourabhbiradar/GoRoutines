package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("select caution !!!")

	//nil channel and
	//empty select will block forever !!!

	ch := make(chan string) // nil channel (if below GoR commented)

	// go func() {
	// 	ch <- "message"
	// }()

	time.Sleep(1)

	//select {} //empty select

	select { // comment out GoR
	case v := <-ch:
		fmt.Println(v)
	}

}
