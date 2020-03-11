package queue

import "errors"

type CircularQueue struct {
	ring []interface{}
	head uint
	tail uint
	cap  uint
	mask uint
}

func NewCQ(cap uint) (*CircularQueue, error) {
	if cap&(cap-1) != 0 {
		return nil, errors.New("capacity require 2^n")
	}
	return new(CircularQueue).Init(cap), nil
}

func (q *CircularQueue) Init(cap uint) *CircularQueue {
	q.cap = cap
	q.mask = cap - 1
	q.head = 0
	q.tail = 0
	q.ring = make([]interface{}, cap)
	return q
}

func (q *CircularQueue) Len() uint {
	return q.tail - q.head
}

func (q *CircularQueue) IsFull() bool {
	return q.Len() == q.cap
}

func (q *CircularQueue) IsEmpty() bool {
	return q.head == q.tail
}

func (q *CircularQueue) Enqueue(v interface{}) (int, error) {
	if q.IsFull() {
		return 0, errors.New("full")
	}
	position := q.tail & q.mask
	q.tail++
	q.ring[position] = v
	return 1, nil
}

func (q *CircularQueue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	position := q.head & q.mask
	q.head++
	return q.ring[position]
}

func (q *CircularQueue) Each(visit func(interface{}) bool) {
	for idx := q.head; idx < q.tail; idx++ {
		if !visit(q.ring[idx&q.mask]) {
			break
		}
	}
}

func (q *CircularQueue) Head() interface{} {
	if q.IsEmpty() {
		return nil
	}
	position := q.head & q.mask
	return q.ring[position]
}

func (q *CircularQueue) Tail() interface{} {
	if q.IsEmpty() {
		return nil
	}
	position := (q.tail - 1) & q.mask
	return q.ring[position]
}
