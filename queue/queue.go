package queue

type Queue interface {
	//RemoveAll removes every element from the queue
	RemoveAll()

	//IsEmpty tells whether or not the queue is empty
	IsEmpty() bool

	//Peek returns the element at the beginning of the queue without removing it
	Peek() interface{}

	//DeQueue removes the first element from the queue and returns it
	DeQueue() interface{}

	//EnQueue adds an element to the end of the queue
	EnQueue(elem interface{})

	//EnQueueQueue appends the beginning of parameter queue
	//to the end of the target queue
	EnQueueQueue(queue Queue)
}
