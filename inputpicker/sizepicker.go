package inputpicker

import (
	"fmt"
	"strconv"
)

// SizePicker allows picking a size between 16 and 1024.
type SizePicker struct{}

func (sp SizePicker) AskInput() (string, error) {
	var sizeInput string
	var size int
	var err error
	fmt.Println("Enter a size for the avatar between 16 and 1024:")
	fmt.Scan(&sizeInput)
	size, err = strconv.Atoi(sizeInput)
	if err != nil {
		return "", fmt.Errorf("Invalid input: please enter a number")
	}
	if size < 16 || size > 1024 {
		return "", fmt.Errorf("Invalid size: please enter a number between 16 and 1024")
	}
	return sizeInput, nil
}
