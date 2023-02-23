package notification

// a Stack that can be observed by others
type Observable_Stack struct {
	Items    []int
	Observer *Observer
}

// the Stack has changed; let the Observer know
func (stack Observable_Stack) Push(val int) {
	stack.Items = append(stack.Items, val)

	if stack.Observer != nil {
		stack.Observer.State_Changed(0, val)
	}
}

// the Stack has changed; let the Observer know
func (stack Observable_Stack) Pop() {
	len := len(stack.Items)

	removed_item := stack.Items[len-1]

	stack.Items = stack.Items[:len-1]

	if stack.Observer != nil {
		stack.Observer.State_Changed(1, removed_item)
	}
}
