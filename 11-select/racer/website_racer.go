package racer

import (
	"fmt"
	"net/http"
	"time"
)

// Wrapping up

// select
// - Helps you wait on multiple channels.
// - Sometimes you'll want to include time.After in one of your cases to prevent your system blocking forever.

// httptest
// - A conveninent way of creating test servers so you cna have reliable and controllable tests.
// - Uses the same interfaces as the "real" net/http servers which is consisten and less for you to learn.

var tenSecondTimeout = 10 * time.Second

// You'll recall from the currency chapter that you can wait for values to be send to a
// channel with myVar := <-ch. This is a blocking call, as you're waiting for a value.

// select allows you to wait on multiople channels. The first one to send a value "wins"
// and the code underneath the case is executed.

// We use ping in our select to set up two channels, one for each of our URL'S. Whichever one
// writes to its channel first will have its code executed in the select, which results in its
// URL being returned (and being the winner).
func Racer(a, b string) (string, error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)

	// aDuration := measureResponseTime(a)
	// bDuration := measureResponseTime(a)

	// if aDuration < bDuration {
	// 	return a
	// }

	// return b
}

func ConfigurableRacer(a, b string, timeout time.Duration) (string, error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

// We have defined a function ping which creates a chan struct{} and returns it.

// In our case, we don't care what type is end to the channel, we just want to signal we are
// done and closing the channel works perfectly!

// Why struct{} and not another type like a bool? Well, a chan struct is the smallest data type
// available from a memory perspective so we get no allocation versus a bool. Since we are closing and
// not sending anything on the chan, why allocate anything?

// Notice how we have to use make when creating a channel;rather than say var ch chan struct{}.
// When you use var the variable will be initialised with the "zero" value of the type. Slo for
// string it is "", int it is 0, etc.

// For channels the zero value is nil and if you try and send to it with <- it will block
// forever because you cannot send to nil channels.
func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()

	return ch
}
