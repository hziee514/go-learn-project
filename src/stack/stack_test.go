package stack

import "testing"

func TestStack_Push(t *testing.T) {
	s := New()
	v := s.PopValue()
	t.Logf("IsEmpty[%v], Len[%d], Top[%v], Pop[%v]", s.IsEmpty(), s.Len(), s.TopValue(), v)
	s.PushValue(1)
	t.Logf("IsEmpty[%v], Len[%d], Top[%v]", s.IsEmpty(), s.Len(), s.TopValue())
	s.PushValue(2)
	t.Logf("IsEmpty[%v], Len[%d], Top[%v]", s.IsEmpty(), s.Len(), s.TopValue())
	s.PushValue(3)
	t.Logf("IsEmpty[%v], Len[%d], Top[%v]", s.IsEmpty(), s.Len(), s.TopValue())
	v = s.PopValue()
	t.Logf("IsEmpty[%v], Len[%d], Top[%v], Pop[%v]", s.IsEmpty(), s.Len(), s.TopValue(), v)
	v = s.PopValue()
	t.Logf("IsEmpty[%v], Len[%d], Top[%v], Pop[%v]", s.IsEmpty(), s.Len(), s.TopValue(), v)
	v = s.PopValue()
	t.Logf("IsEmpty[%v], Len[%d], Top[%v], Pop[%v]", s.IsEmpty(), s.Len(), s.TopValue(), v)
	v = s.PopValue()
	t.Logf("IsEmpty[%v], Len[%d], Top[%v], Pop[%v]", s.IsEmpty(), s.Len(), s.TopValue(), v)
}
