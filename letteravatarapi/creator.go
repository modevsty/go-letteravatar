package letteravatarapi

import (
	"fmt"
	"image"
	"strings"

	"github.com/fogleman/gg"
)

type LetterAvatar struct {
	Name            string
	Shape           string
	Size            int
	BackgroundColor string
	ForegroundColor string
}

func NewLetterAvatar(name string) *LetterAvatar {
	return &LetterAvatar{
		Name:            name,
		Shape:           "circle", // Default shape
		Size:            250,      // Default size
		BackgroundColor: "#0289d1",
		ForegroundColor: "#fff",
	}
}

func (la *LetterAvatar) WithShape(shape string) *LetterAvatar {
	la.Shape = shape
	return la
}

func (la *LetterAvatar) WithSize(size int) *LetterAvatar {
	la.Size = size
	return la
}

func (la *LetterAvatar) WithColor(backgroundColor string, foregroundColor string) *LetterAvatar {
	la.BackgroundColor = backgroundColor
	la.ForegroundColor = foregroundColor
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

	// Determine format from path extension, default to PNG
	if strings.HasSuffix(strings.ToLower(path), ".jpg") || strings.HasSuffix(strings.ToLower(path), ".jpeg") {
		return gg.SaveJPG(path, img, 100)
	} else if strings.HasSuffix(strings.ToLower(path), ".png") {
		return gg.SavePNG(path, img)
	}
	return fmt.Errorf("Unsupported file format")
}

func (la *LetterAvatar) getInitials(name string) string {
	words := strings.Fields(name)
	var initials string
	for _, word := range words {
		initials += strings.ToUpper(word[:1])
	}
	return initials
}
