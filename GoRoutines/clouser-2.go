package main

import (
	"fmt"
	"sync"
)

func main() {

	fmt.Println("clouser-2")

	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)

		// go func() {
		// 	fmt.Println("value of I :", i)
		// 	defer wg.Done()
		// 	//Here every value of i is 3 as the i value already been incremented to max. before GoR executes
		// 	//even tho i < 3 it prints '3'   ????
		// }()

		go func(i int) {
			fmt.Println("value of i :", i)
			defer wg.Done()
			// here i is passed as parameter to GoR so loop happens
			//here i has all values
		}(i)

	}
	wg.Wait()

	fmt.Println("Done...")

}
