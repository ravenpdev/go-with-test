package iterations

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 6)
	expected := "aaaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func TestContain(t *testing.T) {
	got := strings.Contains("golang", "go")
	expected := true

	if expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}
}

func TestClone(t *testing.T) {
	s := "abc"
	got := strings.Clone(s)
	expected := "abc"

	if expected != got {
		t.Errorf("expedted %q, got %q", expected, got)
	}
}

func ExampleRepeat() {
	result := Repeat("r", 8)
	fmt.Println(result)
	// Output: rrrrrrrr
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 8)
	}
}

// go test -bench=.
// 136 ns/op means is our function takes an average 136 nanoseconds to run.

// go test -bench=. -benchmem
// B/op: the number of bytes allocated per iteration
// allocs/op: the number of memeory allocation per iteration
