package queue

import (
	"reflect"
	"testing"
)

func TestQueue(t *testing.T) {
	q := New()
	t.Logf("IsEmpty[%v], Len[%v], Head[%v], Tail[%v]",
		q.IsEmpty(), q.Len(), q.Head(), q.Tail())
	q.Enqueue(1)
	t.Logf("IsEmpty[%v], Len[%v], Head[%v], Tail[%v]",
		q.IsEmpty(), q.Len(), q.Head(), q.Tail())
	q.Enqueue(2)
	t.Logf("IsEmpty[%v], Len[%v], Head[%v], Tail[%v]",
		q.IsEmpty(), q.Len(), q.Head(), q.Tail())
	q.Enqueue(3)
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
