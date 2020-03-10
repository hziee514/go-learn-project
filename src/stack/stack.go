package stack

type Element struct {
	prev  *Element
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

func (stack *Stack) PushValue(v interface{}) {
	stack.Push(&Element{Value: v})
}

func (stack *Stack) Push(e *Element) {
	if stack.len == 0 {
		e.prev = nil
	} else {
		e.prev = stack.top
	}
	stack.top = e
	stack.len++
}

func (stack *Stack) PopValue() interface{} {
	e := stack.Pop()
	if e == nil {
		return nil
	}
	return e.Value
}

func (stack *Stack) Pop() *Element {
	if stack.len == 0 {
		return nil
	}
	p := stack.top
	stack.top = p.prev
	stack.len--
	return p
}

func (stack *Stack) TopValue() interface{} {
	if stack.len == 0 {
		return nil
	}
	return stack.Top().Value
}

func (stack *Stack) Top() *Element {
	return stack.top
}
