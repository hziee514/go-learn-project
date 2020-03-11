package queue

type Element struct {
	next  *Element
	Value interface{}
}

type Queue struct {
	head *Element
	tail *Element
	len  int
}

func New() *Queue {
	return new(Queue).Init()
}

func (q *Queue) Init() *Queue {
	q.head = nil
	q.tail = nil
	q.len = 0
	return q
}

func (q *Queue) IsEmpty() bool {
	return q.len == 0
}

func (q *Queue) Len() int {
	return q.len
}

func (q *Queue) Enqueue(v interface{}) {
	e := &Element{Value: v}
	if q.len == 0 {
		q.head = e
		q.tail = e
	} else {
		q.tail.next = e
		q.tail = e
	}
	q.len++
}

func (q *Queue) Dequeue() interface{} {
	if q.len == 0 {
		return nil
	}
	e := q.head
	if q.len == 1 {
		q.head = nil
		q.tail = nil
	} else {
		q.head = e.next
		e.next = nil
	}
	q.len--
	return e.Value
}

func (q *Queue) Each(visit func(interface{}) bool) {
	p := q.head
	for p != nil {
		if !visit(p.Value) {
			break
		}
		p = p.next
	}
}

func (q *Queue) Head() interface{} {
	if q.len == 0 {
		return nil
	}
	return q.head.Value
}

func (q *Queue) Tail() interface{} {
	if q.len == 0 {
		return nil
	}
	return q.tail.Value
}
