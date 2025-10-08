package main

import "fmt"

type Node struct {
	prev *Node
	data int
	next *Node
}

// doubly linked list

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

// insert at head

func (d *DoublyLinkedList) InsertAtHead(value int) {
	NewNode := &Node{data: value}
	if d.head == nil {
		d.head = NewNode
		d.tail = NewNode
	} else {
		NewNode.next = d.head
		d.head.prev = NewNode
		d.head = NewNode
	}
}

// insert at tail
func (d *DoublyLinkedList) InsertAtTail(value int) {
	NewNode := &Node{data: value}
	if d.head == nil {
		d.head = NewNode
		d.tail = NewNode
	} else {
		NewNode.prev = d.tail
		d.tail.next = NewNode
		d.tail = NewNode
	}
}

// delete by value
func (d *DoublyLinkedList) DeleteByValue(value int) {
	current := d.head
	if current == nil {
		return
	}
	for current != nil {
		if current.data == value {
			if current.prev != nil {
				current.prev.next = current.next
			} else {
				d.head = current.next
			}
			if current.next != nil {
				current.next.prev = current.prev
			} else {
				d.tail = current.prev
			}
			return
		}
		current = current.next
	}
}

// delete by position
func (d *DoublyLinkedList) DeleteByPosition(pos int) {
	if pos < 0 {
		return
	}
	current := d.head
	for i := 0; current != nil && i < pos; i++ {
		current = current.next
	}
	if current == nil {
		return
	}
	if current.prev != nil {
		current.prev.next = current.next
	} else {
		d.head = current.next
	}
	if current.next != nil {
		current.next.prev = current.prev
	} else {
		d.tail = current.prev
	}
}

func (d *DoublyLinkedList) Reverse() {
	current := d.head
	var temp *Node
	for current != nil {
		temp = current.prev
		current.prev = current.next
		current.next = temp
		current = current.prev
	}
	if temp != nil {
		d.head = temp.prev
	}
}

func ArrayToDoublyLinkedList(arr []int) *DoublyLinkedList {
	list := &DoublyLinkedList{}
	for _, val := range arr {
		list.InsertAtTail(val)
	}
	return list
}

// display the list
func (d *DoublyLinkedList) Display() {
	current := d.head
	for current != nil {
		fmt.Printf("%d <->", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func (d *DoublyLinkedList) DisplayBackward() {
	if d.tail == nil {
		fmt.Println("List is empty")
		return
	}
	current := d.tail
	for current != nil {
		fmt.Printf("%d <->", current.data)
		current = current.prev
	}
	fmt.Println("nil")
}

func (d *DoublyLinkedList) Search(value int) bool {
	current := d.head
	for current != nil {
		if current.data == value {
			return true
		}
		current = current.next
	}
	return false
}

func main() {
	arr := []int{50, 60, 70}
	d := ArrayToDoublyLinkedList(arr)
	d.InsertAtHead(10)
	d.InsertAtHead(20)
	// d.InsertAtTail(60)
	// d.InsertAtTail(70)
	d.InsertAtHead(30)
	d.InsertAtHead(40)
	fmt.Println("Doubly linked list:")
	d.Display()
	d.Reverse()
	d.Display()
	d.DisplayBackward()
	d.DeleteByValue(20)
	fmt.Println("After deleting by value 20:")
	d.DeleteByPosition(2)
	fmt.Println("After deleting by position 3:")
	d.Display()
	// d.DisplayBackward()
	fmt.Println("search 30:", d.Search(30))

}
