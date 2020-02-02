package main

// Time complexity is O(n)

import "fmt"

func Factorial(n int) int {
	// Termination condition
	if n <= 1 {
		return 1
	}
	// Recursive Expansion
	return n * Factorial(n-1)
}

func main() {
	fmt.Println("factorail 5 is :: ", Factorial(5))
}
