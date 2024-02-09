package inputpicker

import "fmt"

func AskContinuousInput(p Pickable) string {
	var input string
	var err error
	for {
		input, err = p.AskInput()
		if err != nil {
			fmt.Println("\nError:", err)
			continue
		}
		if input != "" {
			break
		}
	}
	return input
}
