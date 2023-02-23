package dialog

import (
	"fmt"
	"mediator/component"
)

type Dialog struct {
	Widgets map[string]component.Clickable
}

func (dlg *Dialog) Notify(id string) {
	if id == "btn_next" {
		fmt.Println("Proceed to next screen.")
	} else if id == "cb_accept" {
		// cast widget to checkbox
		cb := dlg.Widgets["cb_accept"].(*component.Checkbox)

		// cast widget to button
		btn := dlg.Widgets["btn_next"].(*component.Button)

		// only if checkbox is checked, will the button be enabled
		btn.Is_enabled = cb.Is_checked
	}
}
