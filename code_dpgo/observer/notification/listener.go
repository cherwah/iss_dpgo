package notification

import "fmt"

// a Listener type
type Listener struct {
	Name string
}

// a Add event has occured
func (o Listener) on_add(val int) {
	fmt.Print("Listener '", o.Name,
		"' received ADD notification.\n")
}

// a Remove event has occured
func (o Listener) on_remove(val int) {
	fmt.Print("Listener '", o.Name,
		"' received REMOVE notification.\n")
}
