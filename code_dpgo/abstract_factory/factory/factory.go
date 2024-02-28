package factory

import (
	"abstract_factory/mac"
	"abstract_factory/win"
  "abstract_factory/linux"
)

/*
Interface.
*/

type Factory interface {
	Make_Button() 
	Make_Checkbox()
  Make_Dropdown()
}

/*
Helper function to get the correct factory to use
*/
func Get_factory(os string) Factory {
	if os == "darwin" {
		return mac.Mac_Factory{}
	} else if os == "windows" {
		return win.Win_Factory{}
	} else if os == "linux" {
    return linux.Linux_Factory{}
  } else {
    // OS not supported
    return nil
  }
}
