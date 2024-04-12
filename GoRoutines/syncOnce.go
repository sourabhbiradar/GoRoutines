package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("sync.Once")

	var wg sync.WaitGroup

	var once sync.Once

	load := func() {
		fmt.Println("print only once intilization function")
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			once.Do(load)
		}()

	}
	wg.Wait()

}
