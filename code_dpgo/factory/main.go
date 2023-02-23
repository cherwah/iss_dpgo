package main

import (
	"factory/shape"
	"fmt"
)

func main() {
	/*
	 creates a circle with our factory
	*/
	attrs1 := map[string]float64{
		"radius": 3,
	}
	shape1 := shape.New_Shape("circle", "blue", attrs1)

	/*
	 creates a equilateral triangle with our factory
	*/
	attrs2 := map[string]float64{
		"length": 3,
	}
	shape2 := shape.New_Shape("equ_triangle", "red", attrs2)

	/*
	 creates a rectangle with the factory
	*/
	attrs3 := map[string]float64{
		"width":  2,
		"height": 3,
	}
	shape3 := shape.New_Shape("rectangle", "green", attrs3)

	/*
		display shapes' info
	*/
	display_shape_info(shape1)
	display_shape_info(shape2)
	display_shape_info(shape3)
}

func display_shape_info(shape shape.Geometry) {
	fmt.Println("Area of", shape.Get_Kind(), "=", shape.Area())
	fmt.Println("Perimeter of", shape.Get_Kind(), "=", shape.Perimeter())
}
