package notification

import "fmt"

// a Listener type
type Listener struct {
	Name string
}

// a Add event has occured
func (o Listener) on_add(val int) {
	fmt.Printf("Listener '%s' notified: ADDED %d\n",
		o.Name, val)
}

// a Remove event has occured
func (o Listener) on_remove(val int) {
	fmt.Printf("Listener '%s' notified: REMOVED: %d\n",
		o.Name, val)
}
