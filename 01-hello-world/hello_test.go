package main

import "testing"

// func TestHello(t *testing.T) {
// 	got := Hello()
// 	want := "Hello, world"

// 	if got != want {
// 		t.Errorf("got %q want %q", got, want)
// 	}
// }

func TestHello(t *testing.T) {
	got := Hello("raven")
	want := "Hello, raven"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
