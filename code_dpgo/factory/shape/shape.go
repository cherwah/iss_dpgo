package shape

import (
	"math"
)

/*
Base Struct.
*/
type shape struct {
	color string
	kind  string
}

func (s shape) Get_Kind() string {
	return s.kind
}

/*
Interface.
*/
type Geometry interface {
	Area() float64
	Perimeter() float64
	Get_Kind() string
}

/*
Circle.
*/
type circle struct {
	shape
	radius float64
}

func (c circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

/*
Equilateral Triangle (or sides are the same length).
*/
type equ_triangle struct {
	shape
	length float64
}

func (et equ_triangle) Area() float64 {
	one_radian := 180 / math.Pi
	height := math.Sin(60/one_radian) * et.length
	return 0.5 * et.length * height
}

func (et equ_triangle) Perimeter() float64 {
	return 3 * et.length
}

/*
Rectangle.
*/
type rectangle struct {
	shape
	width  float64
	height float64
}

func (r rectangle) Area() float64 {
	return r.width * r.height
}

func (r rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

/*
Factory for creating Shapes.
*/
func New_Shape(shape_name string, color string,
	attrs map[string]float64) Geometry {
	if shape_name == "circle" {
		// returns a circle
		return circle{
			shape:  shape{color, "circle"},
			radius: attrs["radius"],
		}
	} else if shape_name == "equ_triangle" {
		// returns a equilateral triangle
		return equ_triangle{
			shape:  shape{color, "equ_triangle"},
			length: attrs["length"],
		}
	} else if shape_name == "rectangle" {
		// returns a rectangle
		return rectangle{
			shape:  shape{color, "rectangle"},
			width:  attrs["width"],
			height: attrs["height"],
		}
	}

	return nil
}
