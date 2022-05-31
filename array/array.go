package array

import (
	"fmt"

	"golang.tutorial/timeit"
)

type Array struct {
	start int64
	end   int64
	store []int64
}

type Calculator interface {
	SetStart(x int64)
	SetEnd(x int64)
	FillStore() error
	Compute(result chan<- int64) (int64, error)
}

func (array *Array) SetStart(x int64) {
	array.start = x
}

func (array *Array) SetEnd(x int64) {
	array.end = x
}

func (array *Array) FillStore() error {
	var i int64
	for i = array.start; i <= array.end; i++ {
		if i < 0 {
			return fmt.Errorf("Value %x can't be negative", i)
		}
		array.store = append(array.store, i)
	}
	return nil
}

func (array *Array) Compute(result chan<- int64) {
	var timer timeit.Clock
	timer.Start()
	var sum int64 = 0
	err := array.FillStore()
	if err != nil {
		result <- 0
		fmt.Printf("Error %v", err)
		return
	}
	for _, k := range array.store {
		sum += k
	}
	result <- sum
	array = nil
	timer.Stop()
	fmt.Printf("Time used for computation : %v\n", timer.Evaluate())
}
