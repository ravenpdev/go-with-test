package arrayslices

// func Sum(numbers [5]int) (sum int) {
// 	for i := 0; i < len(numbers); i++ {
// 		sum += numbers[i]
// 	}

// 	return
// }

// declaring array
// numbers = numbers[5]int{1, 2, 3, 4, 5}
// numbers = numbers[...]int{1, 2, 3, 4, 5}

// An interesting property of arrays is that the size is encoded in its type.
//If you try to pass an [4]int into a function that expects [5]int, it won't compile.
// They are different types so it's just the same as trying to pass a string into a function that wants an int.

func Sum(numbers []int) int {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}

	return sum
}

// recursive version
// func Sum(numbers []int) int {
// 	if len(numbers) == 0 {
// 		return 0
// 	}

// 	return numbers[0] + Sum(numbers[1:])
// }

// variadic functions that can take a variable number of arguments
func SumAll(numbersToSum ...[]int) []int {
	// lengOfNumbers := len(numbersToSum)
	// sums := make([]int, lengOfNumbers)
	// for i, numbers := range numbersToSum {
	// 	sums[i] = Sum(numbers)
	// }

	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTail(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(numbers[1:]))
		}
	}

	return sums
}
