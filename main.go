package main

import (
	"fmt"
	"modevsty/go/letteravatar/inputpicker"
	"modevsty/go/letteravatar/letteravatarapi"
	"strings"
)

func main() {
	var (
		name, bgcolor, shape, format string
	)

	pickers := []inputpicker.Pickable{
		inputpicker.NamePicker{},
		inputpicker.ColorPicker{},
		inputpicker.ShapePicker{},
		inputpicker.FormatPicker{},
	}

	for _, picker := range pickers {
		var input string
		var err error

		for len(input) == 0 {
			input, err = picker.AskInput()

			if err != nil {
				fmt.Println(err)
				fmt.Println()
			}
		}

		switch picker.(type) {
		case inputpicker.NamePicker:
			name = input
		case inputpicker.ColorPicker:
			bgcolor = input
		case inputpicker.ShapePicker:
			shape = input
		case inputpicker.FormatPicker:
			format = input
		}
	}

	la := letteravatarapi.NewLetterAvatar(name).WithShape(shape).WithColor(bgcolor, "#fff")
	saveAs := fmt.Sprintf("%s.%s", strings.ToLower(strings.Replace(name, " ", "-", -1)), format)
	la.SaveAs(saveAs)
	fmt.Printf("Avatar generated successfully at %s!\n", saveAs)
}
