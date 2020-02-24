package queue

import "fmt"

var (
	ErrQueueFull    = fmt.Errorf("Queue is full")
	ErrQueueEmpty   = fmt.Errorf("Queue is empty")
	ErrNodeNotFound = fmt.Errorf("A node with given value not found")
)

type Queue struct {
	rear   *node
	front  *node
	length int
	size   int
}

// NewQueue creates a new queue of size s
func NewQueue(s int) *Queue {
	return &Queue{
		rear:   nil,
		front:  nil,
		length: 0,
		size:   s,
	}
}

type node struct {
	data string
	next *node
}

func newNode(d string) *node {
	return &node{
		data: d,
		next: nil,
	}
}

// Enqueue adds a new node to the end of
// the queue
func (q *Queue) Enqueue(value string) error {
	if q.length == q.size {
		return ErrQueueFull
	}
	newNode := NewQueue(value)
	if q.front == q.rear && q.front == nil {
		q.front = newNode
		r.rear = newNode
	}
	q.rear.next = newNode
	q.rear = newNode
	// increment the length of the queue
	q.length++
	return nil
}

// Dequeue removes a node from the start
// of the queue
func (q *Queue) Dequeue() error {
	if q.length == 0 {
		return ErrQueueEmpty
	}
	q.front = q.front.next
	q.length--
	return nil
}

func (q *Queue) MoveNodeToEnd(value string) error {
	if q.length == 0 {
		return ErrQueueEmpty
	}
	temp := q.front
	tempPrev := q.front
	for i := 0; temp != nil; i++ {
		if temp.data == value {
			break
		}
		if i > 0 {
			tempPrev = tempPrev.next
		}
		temp = temp.next
	}
	if temp == nil {
		return ErrNodeNotFound
	}
	if tempPrev == temp {

	}

}
