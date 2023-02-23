package win

import (
	"abstract_factory/widget"
	"fmt"
)

/*
Implementations for Windows.
*/
type Win_Factory struct {
	widget.Factory
}

func (w Win_Factory) Make_Button() widget.Button {
	fmt.Println("Creating Windows button.")

	return widget.Button{
		Widget: widget.Widget{
			Width:  10,
			Height: 10,
			Color:  "gray",
		},
		Outline: true,
	}
}

func (w Win_Factory) Make_Checkbox() widget.Checkbox {
	fmt.Println("Creating Windows checkbox.")

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
