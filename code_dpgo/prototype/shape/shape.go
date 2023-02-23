package shape

type shape struct {
	color string
	kind  string
}

type Circle struct {
	shape  // embedding
	radius float64
}

/*
Using a "mini" Factory pattern here
*/
func New_Circle(color string, r float64) Circle {
	return Circle{
		shape:  shape{color, "circle"},
		radius: r,
	}
}

/*
Prototype Pattern
*/
func (c Circle) Clone() Circle {
	return Circle{
		shape:  shape{c.color, c.kind},
		radius: c.radius,
	}
}
