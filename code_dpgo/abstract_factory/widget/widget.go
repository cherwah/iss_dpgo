package widget

type Widget struct {
	Width  int
	Height int
	Color  string
}

type Button struct {
	Widget
	Outline bool
}

type Checkbox struct {
	Button
	Is_Tick bool
}
