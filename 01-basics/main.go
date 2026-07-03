package main

import "fmt"

func main() {
	fmt.Println("🚀 Go Basics Practice")

	// Variables
	y := 10
	z := 23
	n := 10

	// Function Examples
	fmt.Printf("Sum of numbers from 1 to %d: %d\n", n, sumToN(n))
	fmt.Printf("Addition of %d and %d: %d\n", y, z, add(y, z))
	fmt.Printf("Maximum of %d and %d: %d\n", y, z, max(y, z))
	fmt.Printf("Is %d even? %t\n\n", y, isEven(y))

	// FizzBuzz
	fmt.Println("FizzBuzz (1 - 20)")
	for i := 1; i <= 20; i++ {
		switch {
		case i%15 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}

	// slices example: append
	nums := []int{1, 2, 3}
	nums = append(nums, 4)
	fmt.Println("nums after append:", nums)

	// Maps example
	ages := map[string]int{"me": 20}
	ages["frnd"] = 21
	val, ok := ages["me"]
	fmt.Println("ages:", ages, "val:", val, "ok:", ok)
}

// sumToN returns the sum of numbers from 1 to n.
func sumToN(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

// add returns the sum of two integers.
func add(a, b int) int {
	return a + b
}

// max returns the larger of two integers.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// isEven returns true if the number is even.
func isEven(n int) bool {
	return n%2 == 0
}
