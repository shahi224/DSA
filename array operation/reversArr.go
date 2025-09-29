package main

import "fmt"

// revrse an array
func ReverseArr(arr []int) []int {
	left, right := 0, len(arr)-1

	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
	return arr
}

func main() {
	arr := []int{1, 23, 3, 4, 5}
	fmt.Println("original arr:", arr)

	ReverseArr(arr)
	fmt.Println("Reversed arr :", arr)
}
