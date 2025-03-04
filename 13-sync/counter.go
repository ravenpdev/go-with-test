package sync

import "sync"

// Sync

// Mutex allows us to add locks to our data.
// WaitGroup is a means of waiting for goroutines to finish jobs.

// When to use locks over channels and goroutines?

// A common Go newbie mistake is to over-use channels and goroutines just because it's possible,
// and/or because it's fun. Don't be afraid to use sync.Mutex if that fits your problem best. Go
// is pragmatic in letting you use the tools that solve your problem best and not forcing you into
// one style of code.

// - Use channels when passing ownership of data
// - Use mutexes for managing state

// Don't use embedding because it's convenient
// - Think about the effect embedding has on your public API.
// - Do you really want to expose these methods and have people coupling their own code to them?
// - With respect to mutexes, this could potentially disastrous in very unpredictable and weird
// ways, imagine some nefarious code unlocking mutex when it shouldn't be; this would cause some
// very strange bugs that will be hard to track down.

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
