package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("GoR-2")

	var i int

	go func() {
		i++
		fmt.Printf("GoR : value of i : %v\n", i)
	}()

	fmt.Printf("print1 : value of i : %v\n", i)

	if i == 0 {
		fmt.Printf("if stat : value of i : %v\n", i)
	}

	fmt.Printf("print2 : value of i : %v\n", i)

	time.Sleep(1)
}
