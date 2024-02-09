package main

import (
	"modevsty/go/letteravatar/contpickerapi"
	"modevsty/go/letteravatar/letteravatarapi"
)

func main() {
	contpicker := contpickerapi.NewContPicker().WithName().WithBgcolor().WithShape().WithFormat().WithSize()
	la := letteravatarapi.NewLetterAvatarFromContPicker(contpicker)
	la.Save()
}
