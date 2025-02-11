package integers

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	sum := Sum(1, 3)
	expected := 4

	if sum != expected {
		t.Errorf("sum is %d, expected is %d", sum, expected)
	}
}

func ExampleSum() {
	sum := Sum(1, 5)
	fmt.Println(sum)
	// Output: 6
}
