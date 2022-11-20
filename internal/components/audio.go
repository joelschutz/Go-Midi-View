package components

import (
	"time"

	"github.com/bspaans/bleep/controller"
	"github.com/hajimehoshi/ebiten"
)

type MidiAudioComponent struct {
	audioController *controller.Controller
}

func NewMidiAudioComponent(audio *controller.Controller, midi *MidiController) (comp *MidiAudioComponent, err error) {
	obj := &MidiAudioComponent{
		audioController: audio,
	}
	obj.audioController.EnableSDLSink()
	go obj.audioController.Synth.Start()
	midi.AddHandler(obj)
	return obj, nil
}

func (ma *MidiAudioComponent) OnNoteOn(channel, key, velocity *uint8) {
	ma.audioController.Synth.Mixer.NoteOn(int(*channel), int(*key), float64(*velocity))
}
func (ma *MidiAudioComponent) OnNoteOff(channel, key, _ *uint8) {
	ma.audioController.Synth.Mixer.NoteOff(int(*channel), int(*key))
}
func (ma *MidiAudioComponent) OnPolyAfterTouch(_, _, _ *uint8)           {}
func (ma *MidiAudioComponent) OnAfterTouch(_, _ *uint8)                  {}
func (ma *MidiAudioComponent) OnProgramChange(_, _ *uint8)               {}
func (ma *MidiAudioComponent) OnPitchBend(_ *uint8, _ *int16, _ *uint16) {}
func (ma *MidiAudioComponent) OnControlChange(_, _, _ *uint8)            {}
func (ma *MidiAudioComponent) OnSysEx(_ *[]byte)                         {}

func (ma *MidiAudioComponent) Update(delta time.Duration) error                     { return nil }
func (ma *MidiAudioComponent) Draw(screen *ebiten.Image, delta time.Duration) error { return nil }
