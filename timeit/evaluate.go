package timeit

import (
	"time"
)

type Clock struct {
	start time.Time
	end   time.Time
}

type Timer interface {
	Evaluate() int
	Start()
	Stop()
}

func (clock *Clock) Start() {
	clock.start = time.Now()
}

func (clock *Clock) Stop() {
	clock.end = time.Now()
}

func (clock *Clock) Evaluate() int {
	return clock.end.Second() - clock.start.Second()
}
