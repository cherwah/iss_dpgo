package win

import (
	"fmt"
)

/*
Implementations for Windows.
*/
type Win_Factory struct {
}

func (w Win_Factory) Make_Button() {
	fmt.Println("Creating Windows button.")
}

func (w Win_Factory) Make_Checkbox() {
	fmt.Println("Creating Windows checkbox.")
}

func (w Win_Factory) Make_Dropdown() {
  fmt.Println("Creating Windows dropdown.")
}
