package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type SinglyLinkedList struct {
	head *Node
}

func (s *SinglyLinkedList) Insert(value int) {
	newNode := &Node{data: value}
	if s.head == nil {
		s.head = newNode
		return
	}
	current := s.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (s *SinglyLinkedList) Delete(value int) {
	if s.head == nil {
		return
	}
	if s.head.data == value {
		s.head = s.head.next
		return
	}
	current := s.head
	for current.next != nil && current.next.data != value {
		current = current.next
	}
	if current.next != nil {
		current.next = current.next.next
	}
}

func (s *SinglyLinkedList) Search(value int) bool {
	current := s.head
	for current != nil {
		if current.data == value {
			return true
		}
		current = current.next
	}
	return false
}

func (s *SinglyLinkedList) Display() {
	current := s.head
	for current != nil {
		fmt.Printf("%d ->", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	s := &SinglyLinkedList{}

	s.Insert(10)
	s.Insert(20)
	s.Insert(30)

	s.Display()
	s.Delete(20)
	fmt.Println("After deleting 20:")
	s.Display()
	fmt.Println("Search 20:", s.Search(20))

}
