package main

import (
	"fmt"
)

func main() {

	fmt.Println("channel-1")

	//var ch chan int
	ch := make(chan int)

	go func(a, b int) {
		c := a + b
		fmt.Println(c)
		ch <- c
		// sending value into channel 'ch'
	}(3, 4)

	ans := <-ch
	//receiving value from channel and coping into local variable 'ans'
	fmt.Printf("ans to c : %v", ans)

	//time.Sleep(1) not reqd. as GoR and mainR are communicating ,mainR will wait for value to be received frm GoR channel
}
