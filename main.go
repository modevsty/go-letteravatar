package main

import (
	"bufio"
	"fmt"
	"modevsty/go/letteravatar/letteravatarapi"
	"os"
	"strings"
)

func pickColor() (string, error) {
	var bgcolor, bgcode string
	// pick foreground color from list of colors
	fmt.Println("Pick a background color from the list below:")
	fmt.Println("1. Magenta")
	fmt.Println("2. Green")
	fmt.Println("3. Blue")
	fmt.Print("Enter the number of the color you want: ")
	fmt.Scan(&bgcode)
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

func pickShape() (string, error) {
	var shape, shcode string
	fmt.Println("Pick a shape from the list below:")
	fmt.Println("1. Circle")
	fmt.Println("2. Square")
	fmt.Print("Enter the number of the shape you want: ")
	fmt.Scan(&shcode)
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

func main() {
	var name string
	var bgcolor, shape string
	var bgerr, sherr error

	for len(name) == 0 {
		fmt.Print("Enter your name: ")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			name = scanner.Text()
			if len(strings.Split(name, " ")) > 2 {
				fmt.Println("Please enter your first and last name only")
				fmt.Println()
				name = ""
			}
		}
	}

	for len(bgcolor) == 0 {
		bgcolor, bgerr = pickColor()

		if bgerr != nil {
			fmt.Println(bgerr)
			fmt.Println()
		}
	}

	for len(shape) == 0 {
		shape, sherr = pickShape()

		if sherr != nil {
			fmt.Println(sherr)
			fmt.Println()
		}
	}

	la := letteravatarapi.NewLetterAvatar(name).WithShape(shape)
	la.SetColor(bgcolor, "#fff")
	la.SaveAs(fmt.Sprintf("%s.png", strings.ToLower(strings.Replace(name, " ", "-", -1))))
	fmt.Println("Avatar generated successfully!")
}
