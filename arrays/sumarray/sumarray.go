package main

import "fmt"

func SumArray(data []int) int {
	size := len(data)
	total := 0
	for i := 0; i < size; i++ {
		total = total + data[i]
	}
	return total
}

func main() {
	data := []int{1, 2, 3, 4, 5}
	fmt.Println("Sum is:", SumArray(data))
}
