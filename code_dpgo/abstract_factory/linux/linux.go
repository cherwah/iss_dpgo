package linux 

import (
	"fmt"
)

/*
Implementations for Linux.
*/
type Linux_Factory struct {
}

func (l Linux_Factory) Make_Button() {
	fmt.Println("Creating Linux button.")
}

func (l Linux_Factory) Make_Checkbox() {
	fmt.Println("Creating Linux checkbox.")
}

func (l Linux_Factory) Make_Dropdown() {
  fmt.Println("Create Linux dropdown.")
}
