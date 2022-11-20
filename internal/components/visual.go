package components

import (
	"time"

	"github.com/hajimehoshi/ebiten"
)

type VisualComponent struct {
	layers []VisualGenerator
}

func NewVisualComponent(layers ...VisualGenerator) (comp *VisualComponent, err error) {
	obj := &VisualComponent{
		layers: layers,
	}
	return obj, nil
}

func (ma *VisualComponent) Update(delta time.Duration) error { return nil }
func (ma *VisualComponent) Draw(screen *ebiten.Image, delta time.Duration) error {
	for _, layer := range ma.layers {
		layer.Draw(screen, delta)
	}
	return nil
}
