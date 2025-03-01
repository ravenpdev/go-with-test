package mocking

import (
	"fmt"
	"io"
	"iter"
	"time"
)

const (
	finalWord      = "Go!"
	countdownStart = 3
)

type ConfigurableSleeper struct {
	Duration time.Duration
	SleepBy  func(time.Duration)
}

func (cs *ConfigurableSleeper) Sleep() {
	cs.SleepBy(cs.Duration)
}

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// iterators
func countDownFrom(from int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := from; i > 0; i-- {
			if !yield(i) {
				return
			}
		}
	}
}

func Countdown(out io.Writer, sleeper Sleeper) {
	// for i := countdownStart; i > 0; i-- {
	// 	sleeper.Sleep()
	// }

	// for i := countdownStart; i > 0; i-- {
	// 	fmt.Fprintln(out, i)
	// 	sleeper.Sleep()
	// }

	for i := range countDownFrom(3) {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}

	fmt.Fprint(out, finalWord)
}
