package main

import "fmt"

// Find the Greatest Common Divisor

// Here we use Euclid's algorithm to find the gcd.

func GCD(m int, n int) int {
	if m < n {
		return GCD(n, m)
	}
	if m%n == 0 {
		return n
	}
	return GCD(n, m%n)
}

func main() {
	fmt.Println("GCD is :", GCD(24, 18))
}
