package signalgenerators

import (
	"fmt"
	"time"

	"github.com/joelschutz/Go-Midi-View/internal/components"
)

// Implements a Virtual Addition Generator 1D
type Addition1D struct {
	values []components.SignalGenerator1D
}

func NewAdditionGenerator1D(v ...components.SignalGenerator1D) *Addition1D {
	return &Addition1D{values: v}
}

func (c *Addition1D) GetValue(delta *time.Duration) *float64 {
	var r float64
	for _, v := range c.values {
		r += *v.GetValue(delta)
	}
	return &r
}

// Implements a Virtual Multiplication Generator 1D
type Multiplication1D struct {
	values []components.SignalGenerator1D
}

func NewMultiplicationGenerator1D(v ...components.SignalGenerator1D) *Multiplication1D {
	return &Multiplication1D{values: v}
}

func (c *Multiplication1D) GetValue(delta *time.Duration) *float64 {
	r := 1.
	for _, v := range c.values {
		r *= *v.GetValue(delta)
	}
	fmt.Println("Result Value", r)
	return &r
}
