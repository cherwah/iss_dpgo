package main

import (
  "fmt"
	"abstract_factory/factory"
	"runtime"
)

func main() {
	os := runtime.GOOS

  var factory factory.Factory = factory.Get_factory(os)
  if factory != nil {
    factory.Make_Button()
    factory.Make_Checkbox()
    factory.Make_Dropdown()
  } else {
    fmt.Println("Unsupported OS.")
  }
}
