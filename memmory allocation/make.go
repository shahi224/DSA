package main

import "fmt"

func main() {
	slice := make([]int, 5)
	fmt.Println(slice)

	m := make(map[string]int)
	m["age"] = 25

	fmt.Println(m)
}
