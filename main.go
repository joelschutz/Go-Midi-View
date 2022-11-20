// This file merge together the gitlab.com/gomidi/midi example and github.com/bspaans/bleep example
package main

import (
	"fmt"
	"log"

	"github.com/bspaans/bleep/audio"
	"github.com/bspaans/bleep/controller"
	"github.com/hajimehoshi/ebiten"
	"gitlab.com/gomidi/midi/v2"
	_ "gitlab.com/gomidi/midi/v2/drivers/rtmididrv"

	"github.com/joelschutz/Go-Midi-View/internal"
	"github.com/joelschutz/Go-Midi-View/internal/components"
	signalgenerators "github.com/joelschutz/Go-Midi-View/internal/components/SignalGenerators"
	visualgenerators "github.com/joelschutz/Go-Midi-View/internal/components/VisualGenerators"
)

func main() {
	defer midi.CloseDriver()

	// Midi In Port
	in, err := midi.FindInPort("Arturia")
	if err != nil {
		fmt.Println("can't find Arturia")
		panic(err)
	}

	cfg := audio.NewAudioConfig()

	ctrl := controller.NewController(cfg)

	err = ctrl.LoadInstrumentBank("./audio/bank.yml")
	if err != nil {
		panic(err)
	}
	err = ctrl.LoadPercussionBank("./audio/percussion_bank.yml")
	if err != nil {
		panic(err)
	}
	ctrl.Synth.Mixer.ChangeInstrument(ctrl.Config, 0, 0)
	ctrl.Synth.Mixer.ChangeInstrument(ctrl.Config, 9, 52)

	stop, midiCtrl, err := components.NewMidiController(in, nil)
	if err != nil {
		panic(err)
	}

	audioComp, err := components.NewMidiAudioComponent(ctrl, midiCtrl)
	if err != nil {
		panic(err)
	}
	colorSig := signalgenerators.NewConstantGenerator4D(255, 0, 0, 255)
	lengthSig := signalgenerators.NewMidiVelocityGenerator(0)
	visualComp, err := components.NewVisualComponent(visualgenerators.NewSquareGenerator(lengthSig, colorSig))
	if err != nil {
		panic(err)
	}
	sc := internal.NewScreenController()
	sc.AddComponent(audioComp)
	sc.AddComponent(visualComp)
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(sc); err != nil {
		log.Fatal(err)
		ctrl.Quit()
		stop()
	}
}
