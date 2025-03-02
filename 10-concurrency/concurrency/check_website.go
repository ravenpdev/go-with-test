package concurrency

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	// results := make(map[string]bool)

	// for _, url := range urls {
	// 	results[url] = wc(url)
	// }

	// return results

	// To tell Go to start a new goroutine we turn a function call into a go statement by putting
	// the keyword go in front of it: go doSomething()

	// Anonymous functions have a number of features which make them useful, two of which we're using above.
	// Firstly, they can be executed at the same time that they're declared - this is what the () at the end
	// of the anonymous function doing.
	// Secondly they maintain access in the lexical scope in which they are defined - All variables that are available
	// at the point when you declare the anonymous function are also available in the body of the function.

	// race condition might occurs when go routines try to write at the same time.
	// for _, url := range urls {
	// 	go func() {
	// 		result[url] = wc(url)
	// 	}()
	// }

	// we can solve this data race by coordinating our goroutines using channels.
	// Channels are go data structure that can both receive and send values. These operations, along
	// with their details, allow communication between different processes.
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func() {
			resultChannel <- result{url, wc(url)}
		}()
	}

	for range urls {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
