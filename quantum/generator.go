package quantum

import (
	"machine"
)

// Generator implements controls for the qubit source, setting the configured designated digital pins high and low.
type Generator interface {
	SetLow()
	SetHigh()
	GetID() string
}

// QubitGenerator is a single source of photons controlled by a digital pin.
type QubitGenerator struct {
	ID         string
	VoltagePin machine.Pin
}

// NewQubitGenerator generates a new QubitGenerator with a voltage pin configured as an output pin.
func NewQubitGenerator(pin machine.Pin, id string) *QubitGenerator {
	pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	return &QubitGenerator{
		ID:         id,
		VoltagePin: pin,
	}
}

// SetLow sets the voltage low for the corresponding QubitGenerator
func (g *QubitGenerator) SetLow() {
	g.VoltagePin.Low()
}

// SetHigh sets the voltage high for the corresponding QubitGenerator
func (g *QubitGenerator) SetHigh() {
	g.VoltagePin.High()
}

// GetID returns the unique ID for the QubitGenerator
func (g *QubitGenerator) GetID() string {
	return g.ID
}
