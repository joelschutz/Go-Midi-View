package signalgenerators

import (
	"time"

	"github.com/joelschutz/Go-Midi-View/internal/components"
)

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

// Implements a simple Virtual Combiner Generator 1D-4D
type Combiner1D4D struct {
	value0, value1, value2, value3 components.SignalGenerator1D
}

func NewCombiner1DGenerator4D(v0, v1, v2, v3 components.SignalGenerator1D) *Combiner1D4D {
	return &Combiner1D4D{
		value0: v0,
		value1: v1,
		value2: v2,
		value3: v3,
	}
}

func (c *Combiner1D4D) GetValues(delta *time.Duration) (*float64, *float64, *float64, *float64) {
	return c.value0.GetValue(delta), c.value1.GetValue(delta), c.value2.GetValue(delta), c.value3.GetValue(delta)
}

// Implements a simple Virtual Combiner Generator 1D-3D
type Combiner1D3D struct {
	value0, value1, value2 components.SignalGenerator1D
}

func NewCombiner1DGenerator3D(v0, v1, v2 components.SignalGenerator1D) *Combiner1D3D {
	return &Combiner1D3D{
		value0: v0,
		value1: v1,
		value2: v2,
	}
}

func (c *Combiner1D3D) GetValues(delta *time.Duration) (*float64, *float64, *float64) {
	return c.value0.GetValue(delta), c.value1.GetValue(delta), c.value2.GetValue(delta)
}

// Implements a simple Virtual Combiner Generator 1D-2D
type Combiner1D2D struct {
	value0, value1 components.SignalGenerator1D
}

func NewCombiner1DGenerator2D(v0, v1 components.SignalGenerator1D) *Combiner1D2D {
	return &Combiner1D2D{
		value0: v0,
		value1: v1,
	}
}

func (c *Combiner1D2D) GetValues(delta *time.Duration) (*float64, *float64) {
	return c.value0.GetValue(delta), c.value1.GetValue(delta)
}
