package main

import (
	"fmt"
	"io"
	"os"
)

// Reflection

// golang challenge: write a function walk(x interface{}, fn func(string)) which takes a struct
// x and calls fn for all string fields found inside. difficulty level: recursively.

// To do this we will need to use reflection:

// Reflection in computing is the ability of a program examine its own structure, particularly through types;
// it's form of metaprogramming. It's also a great source of confusion.

// The Representation of an interface

// A variable of interface type stores a pair: the concrete value assigned to a variable, and that value's type descriptor.
// To be more precise, the value is the underlying conrete data item that implements the interface and the type describes
// the full type of that item.

func main() {
	var r io.Reader
	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
	if err != nil {
		return
	}
	r = tty

	fmt.Printf("r type is %T", r)
}
