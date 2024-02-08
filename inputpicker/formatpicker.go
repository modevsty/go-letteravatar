package inputpicker

import "fmt"

// FormatPicker allows picking a save format.
type FormatPicker struct{}

func (fp FormatPicker) AskInput() (string, error) {
	var format, fcode string
	fmt.Println("Pick a save format from the list below:")
	fmt.Println("1. PNG")
	fmt.Println("2. JPG")
	fmt.Print("Enter the number of the format you want: ")
	fmt.Scan(&fcode)
	switch fcode {
	case "1":
		format = "png"
	case "2":
		format = "jpg"
	default:
		return "", fmt.Errorf("Invalid format")
	}
	return format, nil
}
