package signalgenerators

import (
	"fmt"
	"time"

	"github.com/joelschutz/Go-Midi-View/internal/components"
)

// Implements a simple Virtual(Procedural) linear curve
type LinearCurve1D struct {
	scale, offset components.SignalGenerator1D
	delta         time.Duration
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
		scale:  NewConstantGenerator1D(scale),
		offset: NewConstantGenerator1D(offset),
	}
}

func (lc *LinearCurve1D) GetValue(delta *time.Duration) *float64 {
	lc.delta += *delta
	s := *(lc.scale.GetValue(delta))
	o := *(lc.offset.GetValue(delta))
	r := s + o
	// r := s*float64(lc.delta.Seconds()) + o
	fmt.Println("Curve Value", r)
	return &r
}

// Implements a simple Virtual(Procedural) attack curve
type AttackCurve1D struct {
	scale, offset components.SignalGenerator1D
	delta         time.Duration
	oldOffset     float64
}

// Constructor of Virtual Version
func NewVAttackCurve1D(scale, offset components.SignalGenerator1D) *AttackCurve1D {
	return &AttackCurve1D{
		scale:  scale,
		offset: offset,
	}
}

func (lc *AttackCurve1D) GetValue(delta *time.Duration) *float64 {
	lc.delta += *delta
	s := *(lc.scale.GetValue(delta))
	o := *(lc.offset.GetValue(delta))
	if o != lc.oldOffset {
		fmt.Println("Diferent")
		lc.oldOffset = o
		lc.delta = 0
	}
	r := s*float64(lc.delta.Nanoseconds()) + o
	fmt.Println("Curve Value", r)
	return &r
}

// Implements a simple Virtual(Procedural) decay curve
type DecayCurve1D struct {
	scale, offset components.SignalGenerator1D
	delta         time.Duration
	oldOffset     float64
}

// Constructor of Virtual Version
func NewVDecayCurve1D(scale, offset components.SignalGenerator1D) *DecayCurve1D {
	return &DecayCurve1D{
		scale:  scale,
		offset: offset,
	}
}

func (lc *DecayCurve1D) GetValue(delta *time.Duration) *float64 {
	lc.delta += *delta
	s := *(lc.scale.GetValue(delta))
	o := *(lc.offset.GetValue(delta))
	if o != lc.oldOffset {
		fmt.Println("Diferent")
		lc.oldOffset = o
		lc.delta = 0
	}
	r := -s*float64(lc.delta.Nanoseconds()) + o
	fmt.Println("Curve Value", r)
	return &r
}
