package signalgenerators

import (
	"time"

	"github.com/joelschutz/Go-Midi-View/internal/components"
)

// Implements a simple Virtual(Procedural) linear curve
type LinearCurve1D struct {
	scale, offset components.SignalGenerator1D
}

// Constructor of Virtual Version
func NewVLinearCurve1D(scale, offset components.SignalGenerator1D) *LinearCurve1D {
	return &LinearCurve1D{
		scale:  scale,
		offset: offset,
	}
}

// Constructor of Concrete Version
func NewCLinearCurve1D(scale, offset float64) *LinearCurve1D {
	return &LinearCurve1D{
		scale:  NewConstantGenerator1D(offset),
		offset: NewConstantGenerator1D(offset),
	}
}

func (lc *LinearCurve1D) GetValue(delta *time.Duration) *float64 {
	s := *(lc.scale.GetValue(delta))
	o := *(lc.offset.GetValue(delta))
	r := s*float64(delta.Milliseconds()) + o
	return &r
}
