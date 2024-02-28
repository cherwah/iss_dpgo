package mac

import (
	"fmt"
)

/*
Implementations for Mac OSX.
*/

type Mac_Factory struct {
}

func (m Mac_Factory) Make_Button() {
	fmt.Println("Creating Mac button.")
}

func (m Mac_Factory) Make_Checkbox() {
	fmt.Println("Creating Mac checkbox.")
}

func (m Mac_Factory) Make_Dropdown() {
  fmt.Println("Creating Mac dropdown.")
}
