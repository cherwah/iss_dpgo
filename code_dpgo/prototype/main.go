package main

import (
	"fmt"
	"prototype/shape"
)

/*
making a copy of Circle without having to expose
the internals of Circle
*/
func main() {
	fmt.Println("Running 'Prototype' pattern demo")

	circle1 := shape.New_Circle("red", 3)
	circle2 := circle1.Clone()

	fmt.Println("circle1==circle2? =>",
		circle1 == circle2)
}
