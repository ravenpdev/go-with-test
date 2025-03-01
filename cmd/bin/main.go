package main

import (
	"fmt"
	"io"
	"log"
	"maps"
	"math"
	"net/http"
	"slices"
	"time"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))

	// goBasicTypes()

	// defaultValues()

	// goVariables()

	// goConstants()

	// goFor()

	// forRangeLoop()

	// goIfElse()

	// goSwitch()

	// goArrays()

	// goSlices()

	// goMaps()

	// goVariadicFunctions(10, 20, 30)
	// nums := []int{1, 2, 3, 4, 5}
	// goVariadicFunctions(nums...)

	// goPointers()

	// goStructs()

	// l := LinkedList{}
	// l.Append(1)
	// l.Append(2)
	// l.Append(3)
	// l.Prepend(0)
	// l.Delete(2)
	// l.PrintList()

	// binary search
	// nums := make([]int, 100)
	// for i := range nums {
	// 	nums[i] = i + 1
	// }
	// found, index := binarySearch(nums, 43)
	// fmt.Printf("is found: %t, index: %d\n", found, index)
	// found, index = binarySearch(nums, 77)
	// fmt.Printf("is found: %t, index: %d\n", found, index)

	// // find the smallest value
	// randomNumber := []int{99, 50, 75, 44, 1}
	// idx := findSmallest(randomNumber)
	// fmt.Printf("index: %d\n", idx)
	// fmt.Printf("randomNumber: %v\n", randomNumber)

	// sortedNumbers := selectionSort(randomNumber)
	// fmt.Printf("sortedNumbers: %v\n", sortedNumbers)

}

func goBasicTypes() {
	// bool

	// string

	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64

	// byte // alias for uint8

	// rune // alias for int32 // represents a Unicode code point

	// float32 float64

	// complex64 complex128
}

func goVariables() {
	// Variables declared without initialization are zero-valued.
	// For example, the zero value for an int is 0
	var e int
	fmt.Printf("e: %d\n", e)

	var a, b = 5, 10
	fmt.Printf("a: %d, b: %d\n", a, b)

	f, g := 1, 2
	fmt.Printf("f: %d, g: %d\n", f, g)
}

const s string = "constant"

func goConstants() {
	fmt.Println(s)

	// const declares a constant value.
	const n = 500_000_000
	fmt.Printf("type of n is: %T\n", n)

	// A constant statement can appear anywhere a var can.

	// Constant expressions perform arithmetic with arbitrary precision
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Printf("type of d is %T\n", d)

	// A numeric constant has no type until it's given one, such as by an explicit conversion
	fmt.Println(float64(d))

	// A number can be given a type by using it in a context that requires one,
	// such as a variable assignment or function call. For example, here math.Sin expects a float64.
	fmt.Println(math.Sin(n))
}

func goFor() {
	// for is Go's only looping construct.

	// the most basic type, with a single condition
	i := 1
	for i < 3 {
		fmt.Println(i)
		i++
	}

	fmt.Println()

	// A classic initial / condition / after for loop.
	for i := 1; i <= 3; i++ {
		fmt.Println(i)
	}

	fmt.Println()

	// Another way of accomplishing the basic "do this N times"
	// iteration is range over an integer.
	for i := range 3 {
		fmt.Println("Range", i)
	}

	fmt.Println()

	// For without a condition will loop repeatedly until you break out of the loop
	// or return from the enclosing function.
	for {
		fmt.Println("loop")
		break
	}

	fmt.Println()

	// You can also continue to the next iteration of the loop.
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}

func goIfElse() {
	// Basic example
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// You can even have an if without an else.
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// Logical operators like && and || are often useful in conditions.
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("Either 8 or 7 are even")
	}

	// A statement can precede conditionals; any variables declared in this statement are available in the current
	// and all subsequent branches.
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	// There is no ternary if in Go.
}

func goSwitch() {
	// Here's a basic switch
	i := 2
	fmt.Println("Write ", i, " as ")

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// You can use commas to separate multiple expressions in the same
	// case statement. We use the optional default case in this example as well.
	fmt.Println(time.Now())
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// Switch without an expression is an alternative way to
	// express if/else logic. Here we also show how the case expressions
	// can be non-constants.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// A type switch compares type instead of values. You can use this to
	// discover the type of an interface value. In this example, the variable t
	// will have the type corresponding to its caluse.
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("string")
}

func goArrays() {
	// In Go, an array is a numbered sequence of element of a specific length.

	// Here we create an array that will hold exactly 5 ints. The type of elements
	// and length are both part of the arrays type. By default an array is zero-value,
	// which for int means 0.
	var a [5]int
	fmt.Println(a)

	// We can set a value at an index using the array[index] = value syntax, and get a value
	// with array[index].
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// The builtin len returns the length of an array.
	fmt.Println("len:", len(a))

	// Use this syntax to declare and initialize an array in one line.
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// You can also have the compiler count the number of elements for you with ...
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// If you specify the index with :, the elements in between will be zeroed.
	b = [...]int{10, 3: 400, 500}
	fmt.Println("idx:", b)

	// Array types are one-dimensional, but you can compose types to build
	// multi-dimensional data structures.
	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// You can create and initialize multi-dimensional arrays at once too.
	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)
}

func goSlices() {
	// Slices are an important data type in Go, giving a more powerful interface
	// to sequences than arrays.

	// Unline arrays, slices are typed only by the elements they contain (not the number of elements).
	// An uninitialized slice equals to nil and has length 0.
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	// To create an empty slice with non-zero length, use the builtin make. Here we make
	// a slice of strings of length 3 (initially zero-valued). By default a new slice's capacity
	// is equal to its length; if we know the slice is going to grow ahead of time, it's possible
	// to pass a capacity explicitly as an additional parameter to make.
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// We can set and get just like with arrays.
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// len returns the length of the slice as expected
	fmt.Println("len:", len(s))

	// In addition to these basic operations, slices support several more that make them richer
	// than arrays. One is the builtin append, whih returns a slice containing one or more new values.
	// Note that we need to accept a return value from append as we may get a new slice value.
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// Slices can also be copy'd. Here we create an empty slice c
	// of the same length as s and copy into c from s.
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// Slices support a "slice" operator with the syntax
	// slice[low:high]. For example, this gets a slice of the elements s[2],s[3], and s[4].
	l := s[2:5]
	fmt.Println("sl1:", l, "len:", len(l), "cap:", cap(l))

	// This slices up to (but excluding) s[5].
	l = s[:5]
	fmt.Println("sl1:", l, "len:", len(l), "cap:", cap(l))

	// We can declare and initialiize a variable for slice in a single line as well.
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// The slices package contains a number of useful utility functions for slices.
	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	// Slices can composed into multi-dimensional data structures. The length of the inner
	// slices can vary, unline with multi-dimensional arrays.
	twoD := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)

	// len(b)=0, cap(b)=5
	b := make([]int, 0, 5)
	// fmt.Println("b:", b, "len:", len(b), "cap:", cap(b))
	fmt.Printf("b: %v, len: %d, cap: %d\n", b, len(b), cap(b))
	b = b[:cap(b)] // len(b)=5, cap(b)=5
	fmt.Printf("b: %v, len: %d, cap: %d\n", b, len(b), cap(b))
	b = b[1:]
	fmt.Printf("b: %v, len: %d, cap: %d\n", b, len(b), cap(b))

	// Appending to a slice
	// append works on ni slices.
	var myInts []int
	fmt.Printf("myInts: %v, len: %d, cap: %d\n", myInts, len(myInts), cap(myInts))
	myInts = append(myInts, 0)
	fmt.Printf("myInts: %v, len: %d, cap: %d\n", myInts, len(myInts), cap(myInts))

	// This create a slice given an array
	x := [3]string{"apple", "banana", "grape"}
	newSlice := x[:]
	fmt.Printf("newSlice: %v, len: %d, cap: %d\n", newSlice, len(newSlice), cap(newSlice))

	// Slice cannot be grown beyond its capacity.
	s3 := []byte{1, 2, 3, 4, 5}
	fmt.Printf("s3: %v, len: %d, cap: %d\n", s3, len(s3), cap(s3))
	s3 = s3[:4]
	fmt.Printf("s3: %v, len: %d, cap: %d\n", s3, len(s3), cap(s3))
	s3 = s3[:cap(s3)]
	fmt.Printf("s3: %v, len: %d, cap: %d\n", s3, len(s3), cap(s3))

	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}
	for idx, v := range pow {
		fmt.Printf("2**%d = %d\n", idx, v)
	}
}

func goMaps() {
	// An interesting propery of maps is that you can modify them without passing as an address to it
	// (e.g &myMap)
	// This may make them feel like a "reference type"
	// A map value is a pointer to a runtime.hmap structure

	// A gotcha with maps is that hey can be an il value. A nil map behaves like an
	// empty map when reading, but attempts to write to a nil map will cause a runtime panic.

	// Therefor, you should never initialize a nil map variable:
	// var m map[string]string

	// Instead, you can initialize an empty map or use the make keyword to create a map for you:
	// var dictionary = map[string]string{}
	// or
	// var dictionary = make(map[string]string)
	// Both approaches create an emtpy hash map and point dictionary at it. Which ensures that you
	// will never get a runtime panic.
	var m = map[string]int{}

	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)

	// Get a value for a key with name[key] .
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// If the key doesn't exist, the zero value of the value type is returned.
	v3 := m["k2"]
	fmt.Println("v3:", v3)
	v4 := m["k3"]
	fmt.Println("v4:", v4)

	// The built in len returns the number of key/value pairs when called on a map
	fmt.Println("len:", len(m))

	// The builtin delete removes key/value pairs from a map
	delete(m, "k2")
	fmt.Println("map:", m)

	// To remove all the key/value pairs from a map, use the clear builtin
	clear(m)
	fmt.Println("map:", m)

	m["k4"] = 20
	// The optional second return value when getting a value from a map indicates if the key
	// was present in the map. This can be used to disambiguate between missing keys and keys with zero value 0 or "".
	// Here we didn't need the value itself, so we ignored it with the blank indentifiers.
	_, prs := m["k4"]
	fmt.Println("prs:", prs)

	// You can also declare and initialize a new map in the same line with this syntax.
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	// The maps package contains a number of useful utility functions.
	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}

func goVariadicFunctions(nums ...int) {
	total := 0
	for _, val := range nums {
		total += val
	}

	fmt.Println(total)
}

func goPointers() {
	// Go has pointers. A pointer holds the memory address of a value.

	// The type *T is a pointer to a T value. Its zero value is nil.
	var p *int
	fmt.Printf("type: %T, value: %v\n", p, p)

	// The & operator generates a pointer to its operarnt
	i := 42
	p = &i // p points to i memory
	fmt.Printf("i: %d, p address: %p, p value: %d\n", i, p, *p)
	i = 35
	fmt.Printf("i: %d, p address: %p, p value: %d\n", i, p, *p)

	// The * operator denotes the pointer's underlying value.
	// This is known a s"dereferencing"
	fmt.Println(*p) // 42
	// set i through the pointer
	*p = 21
	fmt.Printf("i: %d, p address: %p, p value: %d\n", i, p, *p)
}

type Vertex struct {
	X int
	Y int
}

func goStructs() {
	// A struct is a collection of fields.
	fmt.Println(Vertex{1, 2})

	// Struct fields are accessed using a dot.
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	// Struct fields can be accessed through a struct pointer.

	// To access the field X of a struct when we have the struct pointer p we could write
	// (*p).X. However, that notation is cumbersome, so the language permits us instead
	// to write just p.X, without the explicit dereference.
	v1 := Vertex{1, 2}
	p1 := &v1
	p1.X = 1e9
	fmt.Println(v1)

	// A struct literal denotes a newly allocated struct value by listing the values
	// of its fields.
	// You can list just a subset of fields by using the Name: syntax. (And the order of named
	// fields is irrelevant.)
	// The special prefix & returns a pointer to the struct value.
	v2 := Vertex{1, 2}  // has type Vertex
	v3 := Vertex{X: 1}  // Y:0 is implicit
	v4 := Vertex{}      // X:0 and Y:0
	p2 := &Vertex{1, 2} // has type *Vertex

	fmt.Println(v2, v3, v4, p2)
}

func defaultValues() {
	var myInteger int
	var myFloat float64
	var myBool bool
	var myString string
	var myArray [5]int
	var mySlice []int
	var myMap map[int]int

	// interface zero value is nil. That means it holds no value and type
	var myInterface interface{}

	fmt.Printf("int default value: %d\n", myInteger)
	fmt.Printf("float64 default value: %.2f\n", myFloat)
	fmt.Printf("bool default value: %t\n", myBool)
	fmt.Printf("string default value: %q, type: %T\n", myString, myString)
	fmt.Printf("array default value: %v, type: %T\n", myArray, myArray)
	fmt.Printf("slice default value: %v, type: %T, nil: %t\n", mySlice, mySlice, mySlice == nil)
	fmt.Printf("map default value: %v, type: %T, nil: %t\n", myMap, myMap, myMap == nil)
	fmt.Printf("interface default value: %v, type: %T, nil: %t\n", myInterface, myInterface, myInterface == nil)
}

func forRangeLoop() {
	// for range loop with array
	numbers := [5]int{1, 2, 3, 4, 5}

	for index, value := range numbers {
		fmt.Printf("index: %d, value: %d\n", index, value)
	}

	// for range loop with slices
	names := []string{"raven", "kristine", "zianna", "zaniya"}

	for _, name := range names {
		fmt.Printf("name: %s\n", name)
	}

	for index := range 5 {
		fmt.Printf("index: %d\n", index)
	}
}

func binarySearch(values []int, search int) (bool, int) {
	isFound := false
	index := -1

	high := len(values) - 1
	low := 0

	for low < high {
		mid := int(math.Round(float64(low)+float64(high)) / 2)
		guess := values[mid]

		if guess == search {
			return true, mid
		}

		if guess < search {
			low = mid
		}

		if guess > search {
			high = mid - 1
		}
	}

	return isFound, index
}

func findSmallest(values []int) int {

	smallestIndex := 0

	for index, value := range values {
		if values[smallestIndex] > value {
			smallestIndex = index
		}
	}

	return smallestIndex
}

func selectionSort(values []int) []int {
	size := len(values)
	sorted := make([]int, size)

	for idx := range size {
		index := findSmallest(values)

		sorted[idx] = values[index]
		values = slices.Delete(values, index, index+1)
	}

	return sorted
}
