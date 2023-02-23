package component

type Clickable interface {
	Click()
}

type Widget struct {
	Id     string
	X      int
	Y      int
	Text   string
	Notify func(s string)
}

/**************************
 * Button
 *************************/
type Button struct {
	Widget
	Is_enabled bool
}

func (btn *Button) Click() {
	// inform mediator
	if btn.Is_enabled {
		btn.Notify(btn.Id)
	}
}

/**************************
 * Checkbox
 *************************/
type Checkbox struct {
	Widget
	Is_checked bool
}

func (cb *Checkbox) Click() {
	// toggle checkbox
	cb.Is_checked = !cb.Is_checked

	// inform mediator
	cb.Notify(cb.Id)
}
