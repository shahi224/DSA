package main

import "fmt"

// two sum given target
func TwoSum(arr []int, target int) []int {
	n := len(arr)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if arr[i]+arr[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func main() {
	nums := []int{2, 3, 5, 6, 4}
	target := 10
	result := TwoSum(nums, target)
	fmt.Println(result)
}
