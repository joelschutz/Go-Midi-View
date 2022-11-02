package main

import (
	"fmt"
	"time"

	"github.com/bspaans/bleep/audio"
	"github.com/bspaans/bleep/channels"
	"github.com/bspaans/bleep/controller"
	"github.com/bspaans/bleep/generators"
	"github.com/bspaans/bleep/synth"
	"gitlab.com/gomidi/midi/v2"
	_ "gitlab.com/gomidi/midi/v2/drivers/rtmididrv"
)

func main() {
	defer midi.CloseDriver()

	// Midi In Port
	in, err := midi.FindInPort("Arturia")
	if err != nil {
		fmt.Println("can't find Arturia")
		return
	}

	cfg := audio.NewAudioConfig()

	// Bleep uses generators as the fundamental building block;
	// generators have a configurable pitch, velocity, etc. and *generate*
	// a number of samples.
	// We want to play a sine wave so let's use a SineWaveOscillator
	instr := generators.NewSineWaveOscillator

	// By attaching it to a channel we can play multiple notes at once
	// Each note will have its own dedicated generator function.
	channel := channels.NewPolyphonicChannel()
	channel.SetInstrument(instr)

	// We can combine multiple channels into a Mixer, which allows us to set
	// volumes, etc. In this case it's a bit overkill, because we have
	// only one channel.
	mixer := synth.NewMixer()
	mixer.AddChannel(channel)

	// The Synth takes a mixer, handles input and output events, and is the
	// component that actually asks for samples to send to the output sinks.
	synth := synth.NewSynth(cfg)
	synth.Mixer = mixer

	// The Controller is a high level object that can be used to setup
	// and control the Synthesizer, Sequencer and the instrument banks.
	// In this case we'll just attach the Synth that we made previously.
	ctrl := controller.NewController(cfg)
	ctrl.Synth = synth

	stop, err := midi.ListenTo(in, func(msg midi.Message, timestampms int32) {
		var bt []byte
		var ch, key, vel uint8
		switch {
		case msg.GetSysEx(&bt):
			fmt.Printf("got sysex: % X\n", bt)
		case msg.GetNoteStart(&ch, &key, &vel):
			fmt.Printf("starting note %s on channel %v with velocity %v\n", midi.Note(key), ch, vel)
			mixer.NoteOn(int(ch), int(key), float64(vel))
		case msg.GetNoteEnd(&ch, &key):
			fmt.Printf("ending note %s on channel %v\n", midi.Note(key), ch)
			mixer.NoteOff(int(ch), int(key))
		default:
			fmt.Printf("message %s on channel %v\n", midi.Note(key), ch)
		}
	}, midi.UseSysEx())

	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return
	}

	go synth.Start()
	time.Sleep(time.Minute * 5)

	stop()
}
