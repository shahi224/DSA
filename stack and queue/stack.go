package main

import "fmt"

type Stack struct {
	items []int
}

// push operation
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

//pop operation

func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		fmt.Println("Stack is empty")
		return -1
	}

	lastIndex := len(s.items) - 1
	top := s.items[lastIndex]
	s.items = s.items[:lastIndex] // remove the last items
	return top
}

// peek operation (get top element without removing)
func (s *Stack) Peek() int {
	if len(s.items) == 0 {
		fmt.Println("Stack is empty")
		return -1
	}
	return s.items[len(s.items)-1]
}

// is empty operation
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func main() {
	s := Stack{}

	s.Push(10)
	s.Push(20)
	s.Push(30)

	fmt.Println("Stack:", s.items)
	fmt.Println("Top element:", s.Peek())
	fmt.Println("Popped:", s.Pop())
	fmt.Println("stack after popping:", s.items)
	fmt.Println("is stacj is empty:", s.IsEmpty())

}
