package stack

import "testing"

func TestStack_Push(t *testing.T) {
	s := New()
	v := s.Pop()
	t.Logf("IsEmpty[%v], Len[%d], Top[%v], Pop[%v]", s.IsEmpty(), s.Len(), s.Top(), v)
	s.Push(1)
	t.Logf("IsEmpty[%v], Len[%d], Top[%v]", s.IsEmpty(), s.Len(), s.Top())
	s.Push(2)
	t.Logf("IsEmpty[%v], Len[%d], Top[%v]", s.IsEmpty(), s.Len(), s.Top())
	s.Push(3)
	t.Logf("IsEmpty[%v], Len[%d], Top[%v]", s.IsEmpty(), s.Len(), s.Top())
	v = s.Pop()
	t.Logf("IsEmpty[%v], Len[%d], Top[%v], Pop[%v]", s.IsEmpty(), s.Len(), s.Top(), v)
	v = s.Pop()
	t.Logf("IsEmpty[%v], Len[%d], Top[%v], Pop[%v]", s.IsEmpty(), s.Len(), s.Top(), v)
	v = s.Pop()
	t.Logf("IsEmpty[%v], Len[%d], Top[%v], Pop[%v]", s.IsEmpty(), s.Len(), s.Top(), v)
	v = s.Pop()
	t.Logf("IsEmpty[%v], Len[%d], Top[%v], Pop[%v]", s.IsEmpty(), s.Len(), s.Top(), v)
}
