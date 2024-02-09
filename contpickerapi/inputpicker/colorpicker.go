package inputpicker

import (
	"fmt"
)

// ColorPicker allows picking a color.
type ColorPicker struct{}

func (cp ColorPicker) AskInput() (string, error) {
	var bgcolor, bgcode string
	fmt.Println("Pick a background color from the list below:")
	fmt.Println("1. Magenta")
	fmt.Println("2. Green")
	fmt.Println("3. Blue")
	fmt.Print("Enter the number of the color you want: ")
	fmt.Scanln(&bgcode)
	fmt.Println()

	switch bgcode {
	case "1":
		bgcolor = "#c2185c"
	case "2":
		bgcolor = "#609234"
	case "3":
		bgcolor = "#0289d1"
	default:
		return "", fmt.Errorf("Invalid color")
	}
	return bgcolor, nil
}
