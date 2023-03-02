package main

import (
	"abstract_factory/factory"
	"runtime"
)

func main() {
	os := which_os()

	if os == "win" || os == "mac" {
		var factory factory.Factory = factory.Get_factory(os)
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
