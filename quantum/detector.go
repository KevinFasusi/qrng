package quantum

import "machine"

// Detector is a single photon detector.
type Detector interface {
	GetVertical() uint16
	GetHorizontal() uint16
	GetID() string
}

// SinglePhotonDetector provides the analogue pins corresponding to the HorizontalPin and VerticalPin.
type SinglePhotonDetector struct {
	ID            string
	HorizontalPin machine.ADC
	VerticalPin   machine.ADC
}

// NewSinglePhotonDetector returns a new single photon detectors (SPD).
func NewSinglePhotonDetector(verticalPin machine.Pin, horizontalPin machine.Pin, id string) *SinglePhotonDetector {
	v := machine.ADC{Pin: verticalPin}
	h := machine.ADC{Pin: horizontalPin}
	v.Configure(machine.ADCConfig{})
	h.Configure(machine.ADCConfig{})

	return &SinglePhotonDetector{
		ID:            id,
		HorizontalPin: h,
		VerticalPin:   v,
	}
}

// GetHorizontal returns the value for the single photon detector prepared in the horizontal orientation.
func (p *SinglePhotonDetector) GetHorizontal() uint16 {
	return p.HorizontalPin.Get()
}

// GetVertical returns the value for the single photon detector prepared in the vertical orientation.
func (p *SinglePhotonDetector) GetVertical() uint16 {
	return p.VerticalPin.Get()
}

// GetID returns the id for the single photon detector.
func (p *SinglePhotonDetector) GetID() string {
	return p.ID
}
