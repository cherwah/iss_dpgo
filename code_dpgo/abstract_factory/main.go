package main

import (
  "fmt"
	"abstract_factory/factory"
	"runtime"
)

func main() {
	os := runtime.GOOS

  var factory factory.Factory = factory.Get_factory(os)
  if factory == nil {
    fmt.Println("Unsupported OS.")
    return
  }

  // create OS-specific widgets
  factory.Make_Button()
  factory.Make_Checkbox()
  factory.Make_Dropdown()
}
