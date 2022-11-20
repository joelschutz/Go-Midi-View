package internal

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten"

	"github.com/joelschutz/Go-Midi-View/internal/components"
)

type ScreenController struct {
	config     *ImageConfig
	lastDraw   time.Time
	lastUpdate time.Time
	components []components.Component
}

func NewScreenController() *ScreenController {
	o := ScreenController{}
	o.config = NewImageConfig()
	return &o
}

func (g *ScreenController) AddComponent(c components.Component) error {
	g.components = append(g.components, c)
	return nil
}

func (g *ScreenController) Update(_ *ebiten.Image) error {
	currentUpdate := time.Now()
	delta := currentUpdate.UnixMilli() - g.lastUpdate.UnixMilli()
	g.lastUpdate = currentUpdate
	for _, comp := range g.components {
		err := comp.Update(time.Duration(delta))
		if err != nil {
			if g.config.Debug {
				fmt.Println("Error on component update:", err)
			} else {
				return err
			}
		}
	}
	return nil
}

func (g *ScreenController) Draw(screen *ebiten.Image) {
	currentFrame := time.Now()
	delta := currentFrame.UnixMilli() - g.lastDraw.UnixMilli()
	g.lastDraw = currentFrame
	for _, comp := range g.components {
		err := comp.Draw(screen, time.Duration(delta))
		if err != nil {
			fmt.Println("Error on component drawing:", err)
		}
	}
}

func (g *ScreenController) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.config.Width, g.config.Height
}
