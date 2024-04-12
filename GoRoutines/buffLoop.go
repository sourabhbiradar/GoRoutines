package main

import (
	"fmt"
)

func main() {

	fmt.Println("buffered channel loop")

	//ch := make(chan int)//unbuffered (check result)
	//here , we sent one value and received it immediatly
	//undetermined output

	ch := make(chan int, 5) //buffered (check result)
	//here, we received all values after sending all
	// sequential output

	go func() {

		for i := 0; i < 5; i++ {
			fmt.Println("sending :", i)
			ch <- i
		}
		defer close(ch)
		//once channel is closed 'range ch' will close
		//if not closed fatal error:all GoRs are asleep - deadlock!
	}()

	for v := range ch {
		fmt.Println("received :", v)
	}

}
