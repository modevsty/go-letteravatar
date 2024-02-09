package inputpicker

import "fmt"

// ShapePicker allows picking a shape.
type ShapePicker struct{}

func (sp ShapePicker) AskInput() (string, error) {
	var shape, shcode string
	fmt.Println("Pick a shape from the list below:")
	fmt.Println("1. Circle")
	fmt.Println("2. Square")
	fmt.Print("Enter the number of the shape you want: ")
	fmt.Scanln(&shcode)
	fmt.Println()
	switch shcode {
	case "1":
		shape = "circle"
	case "2":
		shape = "square"
	default:
		return "", fmt.Errorf("Invalid shape")
	}
	return shape, nil
}
