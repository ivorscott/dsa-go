package main

// Problem: Given an array of integers are increasing order,
// find if some value is present in the array suing recursion

// Note: the speed of a recursive program is slower because of
// stack overheads. If the same task can be done using an iterative
// solution (using loops), then we should perfer an iterative solution
// in place of recursion to avoid stack overhead.

import "fmt"

func BinarySearchRecursive(data []int, value int) bool {
	size := len(data)
	return BinarySearchRecursiveUtil(data, 0, size-1, value)
}

func BinarySearchRecursiveUtil(data []int, low int, high int, value int) bool {
	if low > high {
		return false
	}
	mid := (low + high) / 2
	if data[mid] == value {
		return true
	} else if data[mid] < value {
		return BinarySearchRecursiveUtil(data, mid+1, high, value)
	} else {
		return BinarySearchRecursiveUtil(data, low, mid-1, value)
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 9}
	fmt.Println("BinarySearchRecursive:", BinarySearchRecursive(arr, 7))
	fmt.Println("BinarySearchRecursive:", BinarySearchRecursive(arr, 8))
}
