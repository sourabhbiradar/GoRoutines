package main

import (
	"fmt"
	"sync"
)

func main() {

	fmt.Println("clouser-1")

	var wg sync.WaitGroup

	fn := func(wg *sync.WaitGroup) {

		var i int
		wg.Add(1)

		go func() {
			defer wg.Done()
			i++
			fmt.Println(i)
		}()

		fmt.Println("return from func")
		return
	}

	fn(&wg)

	wg.Wait()

	fmt.Println("Done...")

}
