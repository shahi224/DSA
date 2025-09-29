package main

import "fmt"

type Queue struct {
	items []int
}

// Enqueue operation
func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

// Dequeue operation

func (q *Queue) Dequeue() int {
	if len(q.items) == 0 {
		fmt.Println("Queue is empty")
		return -1
	}

	front := q.items[0]
	q.items = q.items[1:] // remove the first element
	return front
}

// front operation
func (q *Queue) Front() int {
	if len(q.items) == 0 {
		fmt.Println("Queue is empty")
		return -1
	}
	return q.items[0]
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func main() {
	q := Queue{}
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)

	fmt.Println("Queue:", q.items)

	fmt.Println("Front element:", q.Front())
	fmt.Println("Dequeued:", q.Dequeue())
	fmt.Println("Queue after dequeueing:", q.items)

	fmt.Println("is queue is empty:", q.IsEmpty())

}
