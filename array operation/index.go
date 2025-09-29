package main

import "fmt"

// inser a value at given index
func InsertAt(arr []int, index, value int) []int {
	arr = append(arr[:index], append([]int{value}, arr[index:]...)...)
	return arr
}

// delete a value at given index
func DeleteAt(arr []int, index int) []int {
	return append(arr[:index], arr[index+1:]...)
}

func main() {
	arr := []int{1, 3, 4, 5, 6}
	fmt.Println("Origina array:", arr)

	arr = InsertAt(arr, 1, 2)
	fmt.Println("after inserting 2:", arr)

	arr = DeleteAt(arr, 5)
	fmt.Println("After deleting:", arr)

}
