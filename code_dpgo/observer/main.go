package main

import "observer/notification"

func main() {

	arr := []int{1, 2, 3}

	observer := notification.Observer{}

	stack := notification.Observable_Stack{
		Items:    arr,
		Observer: &observer,
	}

	// create listeners to receive events
	listener1 := notification.Listener{
		Name: "listener1",
	}
	listener2 := notification.Listener{
		Name: "listener2",
	}

	observer.Subscribe(listener1)
	observer.Subscribe(listener2)

	stack.Push(1)
	stack.Push(2)

	observer.Unsubscribe(listener1)

	stack.Pop()
	stack.Pop()

	observer.Unsubscribe(listener2)
}
