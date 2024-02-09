package letteravatarapi

import (
	"fmt"
	"image"
	"strings"

	"modevsty/go/letteravatar/contpickerapi"

	"github.com/fogleman/gg"
)

type LetterAvatar struct {
	Name            string
	Shape           string
	Size            int
	Format          string
	BackgroundColor string
	ForegroundColor string
}

func NewLetterAvatar(name string) *LetterAvatar {
	return &LetterAvatar{
		Name:            name,
		Shape:           "circle",
		Size:            250,
		BackgroundColor: "#0289d1",
		ForegroundColor: "#fff",
		Format:          "png",
	}
}

func (la *LetterAvatar) WithShape(shape string) *LetterAvatar {
	if shape != "" {
		la.Shape = shape
	}
	return la
}

func (la *LetterAvatar) WithSize(size int) *LetterAvatar {
	if size > 0 {
		la.Size = size
	}
	return la
}

func (la *LetterAvatar) WithColor(backgroundColor string, foregroundColor string) *LetterAvatar {
	if backgroundColor != "" {
		la.BackgroundColor = backgroundColor
	}
	if foregroundColor != "" {
		la.ForegroundColor = foregroundColor
	}
	return la
}

func (la *LetterAvatar) WithFormat(format string) *LetterAvatar {
	if format != "" {
		la.Format = format
	}
	return la
}

func NewLetterAvatarFromContPicker(contPicker *contpickerapi.ContPicker) *LetterAvatar {
	la := NewLetterAvatar(contPicker.Name).
		WithColor(contPicker.Bgcolor, "#fff").
		WithShape(contPicker.Shape).
		WithSize(contPicker.Size).
		WithFormat(contPicker.Format)

	return la
}

func (la *LetterAvatar) Generate() image.Image {
	dc := gg.NewContext(la.Size, la.Size)

	bgColor := la.BackgroundColor
	fgColor := la.ForegroundColor

	if la.Shape == "circle" {
		dc.DrawCircle(float64(la.Size)/2, float64(la.Size)/2, float64(la.Size)/2)
		dc.SetHexColor(bgColor)
		dc.Fill()
	} else {
		dc.DrawRectangle(0, 0, float64(la.Size), float64(la.Size))
		dc.SetHexColor(bgColor)
		dc.Fill()
	}

	dc.SetHexColor(fgColor)
	if err := dc.LoadFontFace("./assets/fonts/arial-bold.ttf", float64(la.Size)/2); err != nil {
		panic(fmt.Sprintf("Could not load font: %v", err))
	}

	initials := la.getInitials(la.Name)
	dc.DrawStringAnchored(initials, float64(la.Size)/2, float64(la.Size)/2, 0.5, 0.5)

	return dc.Image()
}

func (la *LetterAvatar) SaveAs(path string) error {
	img := la.Generate()

	if strings.HasSuffix(strings.ToLower(path), ".jpg") || strings.HasSuffix(strings.ToLower(path), ".jpeg") {
		return gg.SaveJPG(path, img, 100)
	} else if strings.HasSuffix(strings.ToLower(path), ".png") {
		return gg.SavePNG(path, img)
	}
	return fmt.Errorf("Unsupported file format")
}

func (la *LetterAvatar) Save() {
	savePath := fmt.Sprintf("%s.%s", strings.ToLower(strings.Join(strings.Fields(la.Name), "-")), la.Format)
	if err := la.SaveAs(savePath); err != nil {
		fmt.Printf("Error generating avatar: %v\n", err)
	} else {
		fmt.Printf("Avatar generated successfully at %s!\n", savePath)
	}
}

func (la *LetterAvatar) getInitials(name string) string {
	words := strings.Fields(name)
	var initials string
	for _, word := range words {
		initials += strings.ToUpper(word[:1])
	}
	return initials
}
