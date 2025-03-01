package main

import (
	"os"
	"time"

	"github.com/ravenpdev/go-with-test/09-mocking/mocking"
)

func main() {
	sleeper := &mocking.ConfigurableSleeper{
		Duration: 1 * time.Second,
		SleepBy:  time.Sleep,
	}
	mocking.Countdown(os.Stdout, sleeper)
}
