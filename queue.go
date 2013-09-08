//Package queue provides a FIFO queue with mutex lock for safe read and write
//
//Queue uses linked list internal structure
package queue

import (
	"errors"
	"sync"
)

//Internal node for holding the value
//and the pointer to the next node
type node struct {
	Value    interface{} //The data
	NextNode *node       //Next node
}

//Queue is FIFO based queue with linked-list structure
type Queue struct {
	head *node        //Pointer to the first node
	tail *node        //Pointer to the last node
	size int          //Number of nodes in the queue
	mu   sync.RWMutex //Mutex lock for atomic operations
}

//Creates a new queue
func NewQueue() *Queue {
	return &Queue{}
}

//Push add to the end of the queue
func (q *Queue) Push(i interface{}) {

	//Lock it
	q.mu.Lock()

	//Create a new node a contaier for the value
	newNode := node{i, nil}

	//Check if the queue is empty or not
	if q.head == nil {
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

	//Unlock it
	q.mu.Unlock()

	return
}

//Pop remove and return the oldest node in the queue. 
//If the queue is empty error will contain "Queue is empty"
func (q *Queue) Pop() (i interface{}, err error) {

	//Check if the queue is empty
	if q.IsEmpty() {
		return nil, errors.New("Queue is empty")
	}

	//Lock it
	q.mu.Lock()

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

	//Unlock it
	q.mu.Unlock()

	return headNode.Value, nil
}

// The number of elements in the queue
func (q *Queue) Size() int {
	//Lock it
	q.mu.RLock()
	//Get the data
	size := q.size
	//Unlock it
	q.mu.RUnlock()

	return size
}

// Returns true if there is elements in the queue, otherwise false
func (q *Queue) IsEmpty() bool {
	//Lock it
	q.mu.RLock()
	//Get the data
	empty := q.head == nil
	//Unlock it
	q.mu.RUnlock()
	return empty
}
