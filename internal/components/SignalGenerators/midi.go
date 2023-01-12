package signalgenerators

import (
	"time"

	"github.com/joelschutz/Go-Midi-View/internal/components"
)

// Implements a simple Monoponic Concrete Midi Generator
type MidiKey1D struct {
	key     uint8
	channel uint8
}

func NewMidiKeyGenerator(channel uint8) *MidiKey1D {
	return &MidiKey1D{channel: channel}
}

func (mk *MidiKey1D) GetValue(_ *time.Duration) *float64 {
	r := float64(mk.key)
	return &r
}

func (mk *MidiKey1D) OnNoteOn(channel, key, _ *uint8) {
	if mk.channel == *channel {
		mk.key = *key
	}
}
func (mk *MidiKey1D) OnNoteOff(channel, key, _ *uint8) {
	if mk.channel == *channel && mk.key == *key {
		mk.key = 0
	}
}
func (mk *MidiKey1D) OnPolyAfterTouch(_, _, _ *uint8)           {}
func (mk *MidiKey1D) OnAfterTouch(_, _ *uint8)                  {}
func (mk *MidiKey1D) OnProgramChange(_, _ *uint8)               {}
func (mk *MidiKey1D) OnPitchBend(_ *uint8, _ *int16, _ *uint16) {}
func (mk *MidiKey1D) OnControlChange(_, _, _ *uint8)            {}
func (mk *MidiKey1D) OnSysEx(_ *[]byte)                         {}

// Implements a simple Monoponic Concrete Midi Generator
type MidiVelocity1D struct {
	vel     uint8
	channel uint8
}

func NewMidiVelocityGenerator(channel uint8, midi *components.MidiController) *MidiVelocity1D {
	obj := &MidiVelocity1D{channel: channel}
	midi.AddHandler(obj)
	return obj
}

func (mk *MidiVelocity1D) GetValue(_ *time.Duration) *float64 {
	r := float64(mk.vel)
	return &r
}

func (mk *MidiVelocity1D) OnNoteOn(channel, _, vel *uint8) {
	if mk.channel == *channel {
		mk.vel = *vel
	}
}
func (mk *MidiVelocity1D) OnNoteOff(channel, _, vel *uint8) {
	if mk.channel == *channel {
		mk.vel = *vel
	}
}
func (mk *MidiVelocity1D) OnPolyAfterTouch(_, _, _ *uint8)           {}
func (mk *MidiVelocity1D) OnAfterTouch(_, _ *uint8)                  {}
func (mk *MidiVelocity1D) OnProgramChange(_, _ *uint8)               {}
func (mk *MidiVelocity1D) OnPitchBend(_ *uint8, _ *int16, _ *uint16) {}
func (mk *MidiVelocity1D) OnControlChange(_, _, _ *uint8)            {}
func (mk *MidiVelocity1D) OnSysEx(_ *[]byte)                         {}

// Implements a simple Monoponic Concrete Midi Generator
type MidiControl1D struct {
	value, channel, control uint8
}

func NewMidiControlGenerator(channel, control uint8, midi *components.MidiController) *MidiControl1D {
	obj := &MidiControl1D{channel: channel, control: control}
	midi.AddHandler(obj)
	return obj
}

func (mk *MidiControl1D) GetValue(_ *time.Duration) *float64 {
	r := float64(mk.value)
	return &r
}

func (mk *MidiControl1D) OnNoteOn(_, _, _ *uint8)                   {}
func (mk *MidiControl1D) OnNoteOff(_, _, _ *uint8)                  {}
func (mk *MidiControl1D) OnPolyAfterTouch(_, _, _ *uint8)           {}
func (mk *MidiControl1D) OnAfterTouch(_, _ *uint8)                  {}
func (mk *MidiControl1D) OnProgramChange(_, _ *uint8)               {}
func (mk *MidiControl1D) OnPitchBend(_ *uint8, _ *int16, _ *uint16) {}
func (mk *MidiControl1D) OnControlChange(channel, control, value *uint8) {
	if mk.channel == *channel && mk.control == *control {
		mk.value = *value
	}
}
func (mk *MidiControl1D) OnSysEx(_ *[]byte) {}
