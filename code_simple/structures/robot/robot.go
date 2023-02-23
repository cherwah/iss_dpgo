package robot

import "fmt"

type Robot struct {
	Name        string
	Battery     int
	Speed       int
	isActivated bool
}

func (r *Robot) Activate() {
	r.isActivated = true
	fmt.Printf("%s has been activated.\n", r.Name)
}

func (r *Robot) Deactivate() {
	r.isActivated = false
	fmt.Printf("%s has been deactivated.\n", r.Name)
}

func (r *Robot) Move() {
	if r.isActivated {
		fmt.Printf("%s is moving at speed %d.\n", r.Name, r.Speed)
		r.Battery--
	} else {
		fmt.Printf("%s is not activated.\n", r.Name)
		return
	}

	if r.Battery == 0 {
		fmt.Printf("%s has no battery left.\n", r.Name)
		r.Deactivate()
	}
}
