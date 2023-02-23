package main

import (
	"abstract_factory/mac"
	"abstract_factory/widget"
	"abstract_factory/win"
	"runtime"
)

func main() {
	os := which_os()

	if os == "win" || os == "mac" {
		var factory widget.Factory = get_factory(os)
		if factory != nil {
			factory.Make_Button()
			factory.Make_Checkbox()
		}
	}
}

/*
Helper function to determine the OS app is currently running on
*/
func which_os() string {
	os := runtime.GOOS

	switch os {
	case "windows":
		return "win"
	case "darwin":
		return "mac"
	}

	return os
}

/*
Helper function to get the correct factory to use
*/
func get_factory(os string) widget.Factory {
	if os == "mac" {
		return mac.Mac_Factory{}
	} else if os == "win" {
		return win.Win_Factory{}
	}

	return nil
}
