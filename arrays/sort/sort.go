package main

// Problem: Given an array of length N. It contains unique elements from 1 to N.
// Sort the elements of the array

// Solution: For each index value pick the element and then put it into its
// proper place and the element which is in its proper position then pick it
// and repeat the process again

// Time complexity: O(n), it looks like quadratic time complexity but the inner
// loop traverse elements only once. Once the inner loop elements are processed
// then the elements at that position will be either that index value or it will
// contain -1

import "fmt"

func SortA(arr []int, size int) {
	curr, value, next := 0, 0, 0
	for i := 0; i < size; i++ {
		curr = i
		value = -1
		for curr >= 0 && curr < size && arr[curr] != curr+1 {
			next = arr[curr]
			arr[curr] = value
			value = next
			curr = next - 1
		}
	}
}

func SortB(arr []int) {
	temp := 0
	size := len(arr)
	for i := 0; i < size; i++ {
		for arr[i] != i+1 && arr[i] > 1 {
			temp = arr[i]
			arr[i] = arr[temp-1]
			arr[temp-1] = temp
		}
	}
}

func main() {
	arr := []int{8, 5, 6, 1, 9, 3, 2, 7, 4, 10}
	arr2 := []int{15, 12, 5, 8, 1, 14, 9, 7, 3, 2, 4, 6, 10, 11, 13}
	size := len(arr)

	SortA(arr, size)
	SortB(arr2)
	fmt.Println(arr)
	fmt.Println(arr2)

}
