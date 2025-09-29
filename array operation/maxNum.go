package main

import "fmt"

// find highest number in array
func FindMax(arr []int) int {
	max := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
	}
	return max
}

func main() {
	numbers := []int{1, 2, 3, 11, 4, 5, 6}
	fmt.Println("highest number:", FindMax(numbers))
}
