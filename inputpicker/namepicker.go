package inputpicker

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// NamePicker allows picking a name.
type NamePicker struct{}

func (np NamePicker) AskInput() (string, error) {
	fmt.Print("Enter your name: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		name := scanner.Text()
		if len(strings.Split(name, " ")) > 2 {
			return "", fmt.Errorf("Please enter your first and last name only")
		}
		return name, nil
	}
	return "", fmt.Errorf("Failed to read name")
}
