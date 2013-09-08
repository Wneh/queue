package queue

import (
	"errors"
)

//Internal node for holding the value
//and the pointer to the next node
type node struct {
	Value    interface{}
	NextNode *node
}

//Queue is FIFO based queue with linked-list structure
type Queue struct {
	head *node
	tail *node
	size int
}

//Creates a new queue
func NewQueue() *Queue {
	return &Queue{}
}

//Push add to the end of the queue
func (q *Queue) Push(i interface{}) {

	//Create a new node a contaier for the value
	newNode := node{i, nil}

	//Check if the queue is empty or not
	if q.IsEmpty() {
		//Empty so set the head pointer to the new node
		q.head = &newNode
	} else {
		//Not empty direct the current last node
		//to the new last node
		q.tail.NextNode = &newNode
	}
	//And allways set the tail to the new last node
	q.tail = &newNode

	//Increase the size before return
	q.size++

	return
}

//Pop remove and return the oldest node in the queue 
func (q *Queue) Pop() (i interface{}, err error) {

	//Check if the queue is empty
	if q.IsEmpty() {
		return nil, errors.New("Queue is empty")	
	}

	//Take out the headNode
	var headNode = *q.head

	//Check if we only got one item left in the queue
	if q.head == q.tail {
		//Then set the tail to nil
		q.tail = nil
	}

	//Move the head pointer one step forward
	q.head = q.head.NextNode

	//Reduce the size before return
	q.size--

	return headNode.Value,nil
}

// The number of elements in the queue
func (q *Queue) Size() int {
	return q.size
}

// Returns true if there is elements in the queue, otherwise false
func (q *Queue) IsEmpty() bool {
	return q.head == nil
}
