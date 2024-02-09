package inputpicker

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// NamePicker allows picking a name.
type NamePicker struct{}

func (np NamePicker) AskInput() (string, error) {
	fmt.Print("Enter your name: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		fmt.Println()
		name := scanner.Text()
		if wordCount := len(strings.Fields(name)); wordCount > 2 {
			return "", fmt.Errorf("Please enter your first and last name only")
		}
		for _, r := range name {
			if !unicode.IsLetter(r) && !unicode.IsSpace(r) {
				return "", fmt.Errorf("Invalid name: please enter only letters and space")
			}
		}
		return name, nil
	}
	return "", fmt.Errorf("Failed to read name")
}
