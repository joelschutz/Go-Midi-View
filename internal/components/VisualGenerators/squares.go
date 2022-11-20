package visualgenerators

import (
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"

	"github.com/joelschutz/Go-Midi-View/internal/components"
)

type Square struct {
	side  components.SignalGenerator1D
	color components.SignalGenerator4D
}

type SquareGenerator struct {
	shapes [256]*Square
	active bool
}

func NewSquareGenerator(sideLength components.SignalGenerator1D, color components.SignalGenerator4D) *SquareGenerator {
	sg := SquareGenerator{}
	for i := 0; i < len(sg.shapes); i++ {
		sg.shapes[i] = &Square{side: sideLength, color: color}
	}
	return &sg
}

func (sg *SquareGenerator) Draw(screen *ebiten.Image, delta time.Duration) {
	for key, shape := range sg.shapes {
		if shape != nil {
			sWidth, sHeight := screen.Size()
			r, g, b, a := shape.color.GetValues(&delta)
			ebitenutil.DrawRect(screen, float64((sWidth/24)*(key%24)), float64(sHeight/2), float64(sWidth/24), *shape.side.GetValue(&delta), color.RGBA{uint8(*r), uint8(*g), uint8(*b), uint8(*a)})
		}
	}
}
