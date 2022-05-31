package main

import (
	"fmt"

	"golang.tutorial/array"
	"golang.tutorial/timeit"
)

func main() {
	var counter int = 5
	var step int64 = 100000000
	var sum int64 = 0
	results := make(chan int64)
	var i int
	var maintimer timeit.Clock
	maintimer.Start()
	for i = 0; i < counter; i++ {
		var v array.Array
		v.SetStart(int64(i) * step)
		v.SetEnd(int64(i+1)*step - 1)
		go v.Compute(results)
	}
	for i = 0; i < counter; i++ {
		sum += <-results
	}
	maintimer.Stop()
	fmt.Printf("Total sum : %v computed in %v seconds\n", sum, maintimer.Evaluate())
}
