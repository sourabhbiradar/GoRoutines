package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("select timeout-1")

	ch := make(chan string)

	go func() {
		ch <- "message"
	}()

	select {
	case v := <-ch:
		fmt.Println(v)

	case <-time.After(1): //try 0 sec || 1 sec
		//select waits for 1 sec if there is no channel ready it will create GoR in background and send value on channel after sfecified time (here,1 sec)
		fmt.Println("timeout")

	}

}
