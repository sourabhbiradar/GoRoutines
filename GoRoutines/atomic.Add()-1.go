package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("atomic.Add()")

	runtime.GOMAXPROCS(4)
	//max cores of cpu to be used

	var counter uint64
	var wg sync.WaitGroup

	wg.Add(50)

	for i := 0; i < 50; i++ {

		go func() {
			defer wg.Done()
			for c := 0; c < 1000; c++ {

				// counter++ // not accurate and inconsistent output (should be 50000)

				atomic.AddUint64(&counter, 1)
				//address of var counter uint64
				//data type allowed by atomic.Add
				//(here , var counter uint64 so atomic.AddUint64)
			}
		}()
	}

	wg.Wait()
	fmt.Println(counter)

}
