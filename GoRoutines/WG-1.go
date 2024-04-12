package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	fmt.Println("WG-1")

	var i int
	var wg sync.WaitGroup

	wg.Add(2) //numbert of GoR

	go func() {
		i++

		fmt.Printf("GoR-1 : value of i : %v\n", i)

		defer wg.Done()
	}()

	fmt.Printf("print1 : value of i : %v\n", i)

	go func() {
		i++

		fmt.Printf("GoR-2 : value of i : %v\n", i)

		defer wg.Done()
	}()

	if i == 0 {
		fmt.Printf("if stat : value of i : %v\n", i)
	}

	fmt.Printf("print2 : value of i : %v\n", i)

	time.Sleep(1)

	wg.Wait()
}
