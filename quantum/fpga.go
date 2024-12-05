package quantum

import "machine"

type FpgaWire struct {
	ID   string
	Wire machine.Pin
}

// NewFpgaWire generates a new FpgaWire with a voltage pin configured as an output pin.
func NewFpgaWire(pin machine.Pin, id string) *FpgaWire {
	pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	return &FpgaWire{
		ID:   id,
		Wire: pin,
	}
}

// SetLow sets the voltage low for the corresponding FpgaWire wire
func (f *FpgaWire) SetLow() {
	f.Wire.Low()
}

// SetHigh sets the voltage high for the corresponding FpgaWire wire
func (f *FpgaWire) SetHigh() {
	f.Wire.High()
}

// GetID returns the unique ID for the FpgaWire
func (f *FpgaWire) GetID() string {
	return f.ID
}
