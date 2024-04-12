package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("mutex-1")

	runtime.GOMAXPROCS(4)
	//max processors // max os threads to be used

	var balance int
	var wg sync.WaitGroup
	var mu sync.Mutex

	deposite := func(amt int) {
		mu.Lock()
		balance = balance + amt
		mu.Unlock()
		defer wg.Done()
	}

	withdrawl := func(amt int) {
		mu.Lock()
		defer mu.Unlock()
		balance -= amt
	}

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go deposite(1)
	}

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			withdrawl(1)
		}()
	}

	wg.Wait()
	fmt.Println("Balance :", balance)

}
