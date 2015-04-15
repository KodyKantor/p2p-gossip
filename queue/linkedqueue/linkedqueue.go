package linkedqueue

import (
	"github.com/Sirupsen/logrus"
	"github.com/kodykantor/p2p/queue"
)

type Queue struct {
	front *Node
	rear  *Node
}

type Node struct {
	Data interface{}
	Next *Node
}

func init() {
	log := logrus.New()
	log.Println("Initialized linkedqueue")
}

func (q *Queue) RemoveAll() {
	q.front = nil
	q.rear = nil
}

func (q *Queue) IsEmpty() bool {
	return q.front == nil
}

func (q *Queue) Peek() interface{} {
	return q.front.Data
}

func (q *Queue) DeQueue() interface{} {
	result := q.front.Data // hold on to the data
	q.front = q.front.Next //move the pointer 'up'

	return result
}

func (q *Queue) EnQueue(data interface{}) {
	if q.front == nil {
		q.front = &Node{data, nil}
	}
	q.rear = q.front
}

func (q *Queue) EnQueueQueue(que queue.Queue) {
	nuqu := &Queue(que)
	q.rear.Next = nuqu.front
}
