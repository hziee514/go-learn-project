package stack

type Element struct {
	next  *Element
	Value interface{}
}

type Stack struct {
	top *Element
	len int
}

func New() *Stack {
	return new(Stack).Init()
}

func (stack *Stack) Init() *Stack {
	stack.top = nil
	stack.len = 0
	return stack
}

func (stack *Stack) IsEmpty() bool {
	return stack.len == 0
}

func (stack *Stack) Len() int {
	return stack.len
}

func (stack *Stack) Push(v interface{}) {
	e := &Element{Value: v}
	if stack.len > 0 {
		e.next = stack.top
	}
	stack.top = e
	stack.len++
}

func (stack *Stack) Pop() interface{} {
	if stack.len == 0 {
		return nil
	}
	p := stack.top
	stack.top = p.next
	stack.len--
	p.next = nil
	return p.Value
}

func (stack *Stack) Top() interface{} {
	if stack.len == 0 {
		return nil
	}
	return stack.top.Value
}

func (stack *Stack) Each(visit func(interface{}) bool) {
	p := stack.top
	for p != nil {
		if !visit(p.Value) {
			break
		}
		p = p.next
	}
}
