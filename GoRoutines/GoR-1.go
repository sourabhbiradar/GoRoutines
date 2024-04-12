package main

import (
	"fmt"
	"time"
)

func fun(s string) {
	fmt.Println(s)
}

func main() {
	fmt.Println("GoR udemy")

	fun("Direct call")

	go fun("GoR-1") //GoR function call

	go func() { // GoR anonymous func
		fun("GoR-2")
	}()

	fv := fun
	go fv("GoR-3") //GoR with func value call

	time.Sleep(1)
	// asking mainR to wait for go fun() to execute
	// if time.Sleep is nt used mainR executes without go fun()

	fmt.Println("Done...")
}
