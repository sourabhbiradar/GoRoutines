package main

import (
	"fmt"
)

func main() {

	fmt.Println("unbuffered channel loop")

	ch := make(chan int)

	go func() {

		for i := 0; i < 5; i++ {
			fmt.Println("sending :", i)
			ch <- i
		}
		close(ch)
		//once channel is closed 'range ch' will close
		//if not closed fatal error:all GoRs are asleep - deadlock!
	}()

	for v := range ch {
		fmt.Println("received :", v)
	}

}
