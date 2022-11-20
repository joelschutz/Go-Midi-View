package components

import (
	"fmt"

	"gitlab.com/gomidi/midi/v2"
	"gitlab.com/gomidi/midi/v2/drivers"
)

type MidiController struct {
	inPort drivers.In
	// outPort  drivers.Out
	handlers []MidiHandler
}

func NewMidiController(in drivers.In, out drivers.Out) (stop func(), ctrl *MidiController, err error) {
	ctrl = &MidiController{
		inPort: in,
		// outPort: out,
	}
	stop, err = ctrl.init()
	return stop, ctrl, err
}

func (ma *MidiController) AddHandler(handler MidiHandler) {
	ma.handlers = append(ma.handlers, handler)
}

func (ma *MidiController) init() (func(), error) {
	stop, err := midi.ListenTo(ma.inPort, ma.handleMessage, midi.UseSysEx())

	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return func() {}, err
	}
	return stop, nil

}

func (ma *MidiController) handleMessage(msg midi.Message, timestampms int32) {
	var ch, key, vel, pres, prg, ctl, val uint8
	var abs uint16
	var rel int16
	var bt []byte
	switch {
	case msg.GetNoteOn(&ch, &key, &vel):
		for _, handler := range ma.handlers {
			handler.OnNoteOn(&ch, &key, &vel)
		}
	case msg.GetNoteOff(&ch, &key, &vel):
		for _, handler := range ma.handlers {
			handler.OnNoteOff(&ch, &key, &vel)
		}
	case msg.GetPolyAfterTouch(&ch, &key, &pres):
		for _, handler := range ma.handlers {
			handler.OnPolyAfterTouch(&ch, &key, &pres)
		}
	case msg.GetAfterTouch(&ch, &pres):
		for _, handler := range ma.handlers {
			handler.OnAfterTouch(&ch, &pres)
		}
	case msg.GetProgramChange(&ch, &prg):
		for _, handler := range ma.handlers {
			handler.OnProgramChange(&ch, &prg)
		}
	case msg.GetPitchBend(&ch, &rel, &abs):
		for _, handler := range ma.handlers {
			handler.OnPitchBend(&ch, &rel, &abs)
		}
	case msg.GetControlChange(&ch, &ctl, &val):
		for _, handler := range ma.handlers {
			handler.OnControlChange(&ch, &ctl, &val)
		}
	case msg.GetSysEx(&bt):
		for _, handler := range ma.handlers {
			handler.OnSysEx(&bt)
		}
	default:
		fmt.Printf("message %s\n", msg.String())
	}
}
