package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// Synchronising processes

// To do this, we're going to introduce a new construct called select which helps us
// synchronise processes really easily and clearly.

func TestRacer(t *testing.T) {

	// slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	time.Sleep(20 * time.Millisecond)
	// 	w.WriteHeader(http.StatusOK)
	// }))

	// fastServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	w.WriteHeader(http.StatusOK)
	// }))

	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
		slowServer := makeDelayServer(20 * time.Millisecond)
		fastServer := makeDelayServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		want := fastUrl
		got, err := Racer(slowUrl, fastUrl)

		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns an error if a server doesn't response within the specified time", func(t *testing.T) {
		server := makeDelayServer(25 * time.Millisecond)
		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}

		// serverA := makeDelayServer(11 * time.Second)
		// serverB := makeDelayServer(12 * time.Second)

		// defer serverA.Close()
		// defer serverB.Close()

		// _, err := Racer(serverA.URL, serverB.URL)

		// if err == nil {
		// 	t.Errorf("expected an error but didn't get one")
		// }
	})
}

func makeDelayServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
