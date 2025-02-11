package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("return the total of two numbers", func(t *testing.T) {
		sum := Adder(5, 3)
		expected := 8

		if sum != expected {
			t.Errorf("sum is %d, expected is %d", sum, expected)
		}
	})
}

func ExampleAdder() {
	sum := Adder(5, 4)
	fmt.Println(sum)
	// Output: 9
}
