package queue

import (
	"testing"
)

//Test size on a new queue
func Test_Queue_NewEmpty(t *testing.T) {
	q := NewQueue()

	if q.Size() != 0 {
		t.Error("Size != 0")
	} else {
		t.Log("NewEmpty: test passed")
	}

	//The queue should be empty now
	if q.IsEmpty() {
		t.Log("NewEmpty: test passed")
	} else {
		t.Error("The queue was not empty")		
	}
}

// Push and Pop one str
func Test_Queue_One(t *testing.T) {
	q := NewQueue()

	q.Push("Foo")

	//Size should now be 1
	if q.Size() != 1 {
		t.Error("Size != 1")
	} else {
		t.Log("One: size test passed")
	}

	pop, err := q.Pop()

	if err != nil {
	 	t.Error(err)
	}

	//We should now get Foo back when we Pop
	if pop != "Foo" {
		t.Error("One: Pop() did not pass")
	} else {
		t.Log("One: Passed")
	}

	//Size should now be 0
	if q.Size() != 0 {
		t.Error("Size != 0")
	} else {
		t.Log("One: size test passed")
	}
}

// Push and Pop two string
func Test_Queue_Two(t *testing.T) {
	q := NewQueue()

	q.Push("Foo")
	q.Push("Bar")

	//Size should now be 2
	if q.Size() != 2 {
		t.Error("Size != 2")
	} else {
		t.Log("Two: Size == 2 test passed")
	}

	pop, err := q.Pop()
	if err != nil {
	 	t.Error(err)
	}

	//We should now get Foo back when we Pop
	if pop != "Foo" {
		t.Error("Two: Pop() did not pass")
	} else {
		t.Log("Two: Foo Passed")
	}

	//Size should now be 1
	if q.Size() != 1 {
		t.Error("Size != 1")
	} else {
		t.Log("Two: Size == 1 test passed")
	}

	//Pop another from the queue
	pop, err = q.Pop()
	if err != nil {
	 	t.Error(err)
	}

	//We should now get bar back when we Pop
	if pop != "Bar" {
		t.Error("Two: Pop() did not pass")
	} else {
		t.Log("Two: Bar Passed")
	}

	//Size should now be 1
	if q.Size() != 0 {
		t.Error("Size != 0")
	} else {
		t.Log("Two: Size == 0 test passed")
	}
}

//Test calling Pop on empty queue
func Test_Queue_NewEmptyError(t *testing.T) {
	q := NewQueue()

	//Test the size
	if q.Size() != 0 {
		t.Error("Size != 0")
	} else {
		t.Log("NewEmptyError: test passed, size was 0")
	}

	//The queue should be empty now
	if q.IsEmpty() {
		t.Log("NewEmptyError: test passed")
	} else {
		t.Error("NewEmptyError: The queue was not empty")		
	}

	//Do a Pop on a empty queueu
	pop,err := q.Pop()

	//Pop should be nil
	if pop != nil {
		t.Error("NewEmptyError: Pop() did not pass - was not nil")
	} else {
		t.Log("NewEmptyError: Pop Passed")
	}

	//And we should have an error instead
	if err.Error() != "Queue is empty" {
		t.Error("NewEmptyError: Pop() did not pass - err did not contain Queue is empty")	
		t.Log("NewEmptyError: Bar Passed")
	}
}

//Testing 10 integers in the queueu
func Test_Queue_Counting(t *testing.T) {

	q := NewQueue()

	//Push 1,2,3,4,5,6,7,8,9,10 to the queue
	for i := 1; i <= 10; i++ {
		q.Push(i)
	} 

	result := 0

	//Pop and sum them up
	for !q.IsEmpty() {
		tempValue,_ := q.Pop()
		value,_ := tempValue.(int)
		result += value
	}

	//The sum should be 55
	if result != 55 {
		t.Error("Counting: The total sum was not 55")
	} else {
		t.Log("Counting: Passed")
	}
}
