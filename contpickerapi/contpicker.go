package contpickerapi

import (
	"fmt"
	"modevsty/go/letteravatar/contpickerapi/inputpicker"
	"strconv"
)

// ShapePicker allows picking a shape.
type ContPicker struct {
	Name      string
	Bgcolor   string
	Shape     string
	Format    string
	Size      int
	Customize string
}

func NewContPicker() *ContPicker {
	return &ContPicker{}
}

func (cp *ContPicker) WithName() *ContPicker {
	namePicker := inputpicker.NamePicker{}
	name := inputpicker.AskContinuousInput(namePicker)
	cp.Name = name
	return cp
}
func (cp *ContPicker) WithBgcolor() *ContPicker {
	colorPicker := inputpicker.ColorPicker{}
	bgcolor, err := colorPicker.AskInput()
	if err != nil {
		fmt.Println("No color found. Default color will be used")
	}
	cp.Bgcolor = bgcolor
	return cp
}

func (cp *ContPicker) WithShape() *ContPicker {
	shapePicker := inputpicker.ShapePicker{}
	shape, err := shapePicker.AskInput()
	if err != nil {
		fmt.Println("No shape found. Default shape will be used")
	}
	cp.Shape = shape
	return cp
}

func (cp *ContPicker) WithFormat() *ContPicker {
	formatPicker := inputpicker.FormatPicker{}
	format, err := formatPicker.AskInput()
	if err != nil {
		fmt.Println("No format found. Default format will be used")
	}
	cp.Format = format
	return cp
}

func (cp *ContPicker) WithSize() *ContPicker {
	sizePicker := inputpicker.SizePicker{}
	size, err := sizePicker.AskInput()
	if err != nil {
		fmt.Println("No size found. Default size will be used")
	}
	cp.Size, _ = strconv.Atoi(size)
	return cp
}
