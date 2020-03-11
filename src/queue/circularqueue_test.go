package queue

import (
	"reflect"
	"testing"
)

func TestCircularQueue_Cap(t *testing.T) {
	q, _ := NewCQ(2)
	r, e := q.Enqueue(1)
	if r == 0 {
		t.Error(e)
	}
	r, e = q.Enqueue(2)
	if r == 0 {
		t.Error(e)
	}
	r, e = q.Enqueue(3)
	if r == 1 {
		t.Fail()
	}
}

func TestCircularQueue(t *testing.T) {
	q, _ := NewCQ(4)
	t.Logf("IsEmpty[%v], Len[%v], Head[%v], Tail[%v]",
		q.IsEmpty(), q.Len(), q.Head(), q.Tail())
	_, _ = q.Enqueue(1)
	t.Logf("IsEmpty[%v], Len[%v], Head[%v], Tail[%v]",
		q.IsEmpty(), q.Len(), q.Head(), q.Tail())
	_, _ = q.Enqueue(2)
	t.Logf("IsEmpty[%v], Len[%v], Head[%v], Tail[%v]",
		q.IsEmpty(), q.Len(), q.Head(), q.Tail())
	_, _ = q.Enqueue(3)
	t.Logf("IsEmpty[%v], Len[%v], Head[%v], Tail[%v]",
		q.IsEmpty(), q.Len(), q.Head(), q.Tail())

	q.Each(func(v interface{}) bool {
		t.Logf("Visit: Value[%v], Type[%v]", v, reflect.TypeOf(v))
		return true
	})

	v := q.Dequeue()
	t.Logf("IsEmpty[%v], Len[%v], Head[%v], Tail[%v], Dequeue[%v]",
		q.IsEmpty(), q.Len(), q.Head(), q.Tail(), v)
	v = q.Dequeue()
	t.Logf("IsEmpty[%v], Len[%v], Head[%v], Tail[%v], Dequeue[%v]",
		q.IsEmpty(), q.Len(), q.Head(), q.Tail(), v)
	v = q.Dequeue()
	t.Logf("IsEmpty[%v], Len[%v], Head[%v], Tail[%v], Dequeue[%v]",
		q.IsEmpty(), q.Len(), q.Head(), q.Tail(), v)
	v = q.Dequeue()
	t.Logf("IsEmpty[%v], Len[%v], Head[%v], Tail[%v], Dequeue[%v]",
		q.IsEmpty(), q.Len(), q.Head(), q.Tail(), v)
}
