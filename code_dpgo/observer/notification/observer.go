package notification

import "fmt"

// interface for Listeners to implement
type Notification interface {
	on_add(val int)
	on_remove(val int)
}

// an Observer type
type Observer struct {
	listeners []Listener
}

// able to Subscribe to be notified
func (o *Observer) Subscribe(listener Listener) {
	fmt.Print("Listener '", listener.Name, "' subscribed.\n")
	o.listeners = append(o.listeners, listener)
}

// able to Unsubscribe to no longer be notified
func (o *Observer) Unsubscribe(listener Listener) {
	len := len(o.listeners)

	for i := 0; i < len; i++ {
		if o.listeners[i] == listener {
			if i == 0 {
				// remove first element
				o.listeners = o.listeners[1:]
			} else if i == len-1 {
				// remove last element
				o.listeners = o.listeners[:len-1]
			} else {
				o.listeners = append(o.listeners[:i-1],
					o.listeners[i+1:]...)
			}

			fmt.Print("Listener '", listener.Name, "' unsubscribed.\n")
			break
		}
	}
}

// object telling Observer that it has changed
func (o *Observer) State_Changed(action int, val int) {
	if action == 0 {
		for _, listener := range o.listeners {
			listener.on_add(val)
		}
	} else if action == 1 {
		for _, listener := range o.listeners {
			listener.on_remove(val)
		}
	}
}
