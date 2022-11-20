package components

import (
	"time"

	"github.com/hajimehoshi/ebiten"
)

// Component Interfaces
type Component interface {
	Update(delta time.Duration) error
	Draw(screen *ebiten.Image, delta time.Duration) error
}

// Midi Interfaces
type MidiHandler interface {
	OnNoteOn(channel, key, velocity *uint8)
	OnNoteOff(channel, key, velocity *uint8)
	OnPolyAfterTouch(channel, key, pressure *uint8)
	OnAfterTouch(channel, pressure *uint8)
	OnProgramChange(channel, program *uint8)
	OnPitchBend(channel *uint8, relative *int16, absolute *uint16)
	OnControlChange(channel, controller, value *uint8)
	OnSysEx(bt *[]byte)
}

// Visual interfaces
type VisualGenerator interface {
	Draw(screen *ebiten.Image, delta time.Duration)
}

// Data interfaces
type SignalGenerator1D interface {
	GetValue(delta *time.Duration) *float64
}

type SignalGenerator2D interface {
	GetValues(delta *time.Duration) (*float64, *float64)
}

type SignalGenerator3D interface {
	GetValues(delta *time.Duration) (*float64, *float64, *float64)
}

type SignalGenerator4D interface {
	GetValues(delta *time.Duration) (*float64, *float64, *float64, *float64)
}
