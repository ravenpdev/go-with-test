package mocking

import (
	"bytes"
	"slices"
	"testing"
	"time"
)

// In main we will send to os.Stdout so our users see the countdown printed to the terminal.
// In test we will send to bytes.Buffer so our tests can capture what data is being generated.

// Mocking
// The test still pass and the software works as intended but we have some problems:
// - Our tests take 3 seconds to run.
//    - Every forward-thinking post about software development emphasises the importance of quick feeback loops.
//    - Slow tests ruin developer productivity.
//    - Imagine if the requirements get more sophisticated warranting more tests. Are we happy with 3s added to the test run for every new test of Countdown?
// - We have not tested an important propery of our functions.

// We have a dependency on Sleep ing which we need to extract so we can then control it in our tests.

// If we can mock time.Sleep we can use dependency injection to use it instead of a "real!" time.Sleep and then
// we can spy on the calls to make assertions on them.

// Spies shich are kind of mock. Mocks area a type of "test double"
// Test Double is a generic term for any case where you replace a production object for testing purposes.

const (
	write = "write"
	sleep = "sleep"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type SpycountdownOperations struct {
	Calls []string
}

func (s *SpycountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpycountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func TestCoundDown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer, &SpycountdownOperations{})

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &SpycountdownOperations{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !slices.Equal(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})

	// 	buffer := &bytes.Buffer{}
	// 	spySleeper := &SpySleeper{}

	// 	Countdown(buffer, spySleeper)

	// 	got := buffer.String()
	// 	want := `3
	// 2
	// 1
	// Go!`

	// 	if got != want {
	// 		t.Errorf("got %q want %q", got, want)
	// 	}

	//	if spySleeper.Calls != 3 {
	//		t.Errorf("not enough calls to sleeper, want 3 got %d", spySleeper.Calls)
	//	}
}

type SpyTime struct {
	durationSlept time.Duration
}

func (st *SpyTime) Sleep(duration time.Duration) {
	st.durationSlept = duration
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
