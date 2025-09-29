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

func (d *DoublyLinkedList) InsertAtTaile(value int) {
	NewNode := &Node{data: value}
	if d.head == nil {
		d.head = NewNode
		d.tail = NewNode
	} else {
		NewNode.prev = d.tail
		d.head.next = NewNode
		d.tail = NewNode
	}
}

// delete by value
func (d DoublyLinkedList) DeleteByValue(value int) {
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
	d := &DoublyLinkedList{}
	d.InsertAtHead(10)
	d.InsertAtHead(20)
	d.InsertAtHead(30)
	d.InsertAtHead(40)
	fmt.Println("Doubly linked list:")
	d.Display()
	d.DisplayBackward()
	d.DeleteByValue(20)
	fmt.Println("After deleting by value 20:")
	d.DisplayBackward()
	fmt.Println("search 30:", d.Search(30))
}
