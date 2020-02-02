package main

import "fmt"

func FindOne(data []int, value int) bool {
	size := len(data)
	for i := 0; i < size; i++ {
		if value == data[i] {
			return true
		}
	}
	return false
}
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Sequential search:", FindOne(arr, 7))
	fmt.Println("Sequential search:", FindOne(arr, 9))
}
