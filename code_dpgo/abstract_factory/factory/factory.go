package factory

import (
	"abstract_factory/mac"
	"abstract_factory/widget"
	"abstract_factory/win"
)

/*
Interface.
*/

type Factory interface {
	Make_Button() widget.Button
	Make_Checkbox() widget.Checkbox
}

/*
Helper function to get the correct factory to use
*/
func Get_factory(os string) Factory {
	if os == "mac" {
		return mac.Mac_Factory{}
	} else if os == "win" {
		return win.Win_Factory{}
	}

	return nil
}
