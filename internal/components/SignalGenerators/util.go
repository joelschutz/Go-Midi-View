package signalgenerators

import "time"

// Implements a simple Constant Concrete Generator 1D
type Constant1D struct {
	value float64
}

func NewConstantGenerator1D(v float64) *Constant1D {
	return &Constant1D{value: v}
}

func (c *Constant1D) GetValue(_ *time.Duration) *float64 {
	return &c.value
}

// Implements a simple Constant Concrete Generator 4D
type Constant4D struct {
	value0, value1, value2, value3 float64
}

func NewConstantGenerator4D(v0, v1, v2, v3 float64) *Constant4D {
	return &Constant4D{
		value0: v0,
		value1: v1,
		value2: v2,
		value3: v3,
	}
}

func (c *Constant4D) GetValues(_ *time.Duration) (*float64, *float64, *float64, *float64) {
	return &c.value0, &c.value1, &c.value2, &c.value3
}

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

func NewMidiVelocityGenerator(channel uint8) *MidiVelocity1D {
	return &MidiVelocity1D{channel: channel}
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
	if mk.channel == *channel && mk.vel == *vel {
		mk.vel = 0
	}
}
func (mk *MidiVelocity1D) OnPolyAfterTouch(_, _, _ *uint8)           {}
func (mk *MidiVelocity1D) OnAfterTouch(_, _ *uint8)                  {}
func (mk *MidiVelocity1D) OnProgramChange(_, _ *uint8)               {}
func (mk *MidiVelocity1D) OnPitchBend(_ *uint8, _ *int16, _ *uint16) {}
func (mk *MidiVelocity1D) OnControlChange(_, _, _ *uint8)            {}
func (mk *MidiVelocity1D) OnSysEx(_ *[]byte)                         {}
