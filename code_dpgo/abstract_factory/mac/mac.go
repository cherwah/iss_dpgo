package mac

import (
	"abstract_factory/widget"
	"fmt"
)

/*
Implementations for Mac OSX.
*/

type Mac_Factory struct {
}

func (m Mac_Factory) Make_Button() widget.Button {
	fmt.Println("Creating Mac OSX button.")

	return widget.Button{
		Widget: widget.Widget{
			Width:  10,
			Height: 10,
			Color:  "gray",
		},
		Outline: true,
	}
}

func (m Mac_Factory) Make_Checkbox() widget.Checkbox {
	fmt.Println("Creating Mac OSX checkbox.")

	return widget.Checkbox{
		Button: widget.Button{
			Widget: widget.Widget{
				Width:  10,
				Height: 10,
				Color:  "gray",
			},
		},
		Is_Tick: false,
	}
}
