package main

import (
	"mediator/component"
	"mediator/dialog"
)

func main() {

	dlg := dialog.Dialog{
		// create empty array to store components that
		// has implemented the Clickable interface
		Widgets: map[string]component.Clickable{},
	}

	// stores a Button
	dlg.Widgets["btn_next"] = &component.Button{
		Widget: component.Widget{
			Id:     "btn_next",
			X:      50,
			Y:      60,
			Notify: dlg.Notify,
			Text:   "Next",
		},
		Is_enabled: false,
	}

	// stores a Checkbox
	dlg.Widgets["cb_accept"] = &component.Checkbox{
		Widget: component.Widget{
			Id:     "cb_accept",
			X:      60,
			Y:      80,
			Notify: dlg.Notify,
			Text:   "Accept License Agreement",
		},
		Is_checked: false,
	}

	simulate(&dlg)
}

func simulate(dlg *dialog.Dialog) {
	// the user must accept the license agreement first;
	// by clicking on the Accept checkbox
	dlg.Widgets["cb_accept"].Click()

	// then the Next button would be enabled, then the user
	// can click on it to proceed to the next screen
	dlg.Widgets["btn_next"].Click()
}
