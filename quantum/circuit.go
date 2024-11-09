package quantum

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"time"
)

// IDS identify qubit source and corresponding single photon detectors (SPD) numbered one to four.
type IDS int

const (
	Q1 IDS = iota
	Q2
	Q3
	Q4
	SPD1
	SPD2
	SPD3
	SPD4
)

func (i IDS) Strings() string {
	return [...]string{"Q1", "Q2", "Q3", "Q4", "SPD1", "SPD2", "SPD3", "SPD4"}[i]
}

// Circuit runs the assembled logic using the available qubit generators and single photon detectors.
type Circuit struct {
	ZeroCount  int
	OneCount   int
	Bits       []int
	Generators *[]Generator
	Detectors  *[]Detector
	Bias       *Bias
}

// NewCircuit composes new circuit using Generator and Detector slices and a Bias struct.
func NewCircuit(generatorQueue *[]Generator, singlePhotonDetectorQueue *[]Detector, bias *Bias) *Circuit {
	return &Circuit{
		Generators: generatorQueue,
		Detectors:  singlePhotonDetectorQueue,
		Bias:       bias,
	}
}

// Run starts the circuit.
func (c *Circuit) Run() {
	littleEndianByte := make([]byte, 4)
	for _, g := range *c.Generators {
		c.PulseGenerator(g)
		measurement := c.Measure(c.Detectors, g.GetID())
		if c.Bias.Postprocessor(measurement, g.GetID()) {
			binary.LittleEndian.PutUint32(littleEndianByte, uint32(measurement))
			go println(hex.EncodeToString(littleEndianByte))
			c.Bits = append(c.Bits, measurement)
		}
	}
	c.Statistics()
	c.Bits = make([]int, 0, 4)
}

// PulseGenerator pulses the qubit generator source by setting the digital pin LOW then HIGH at 200 microsecond
// intervals.
func (c *Circuit) PulseGenerator(generator Generator) {
	generator.SetLow()
	time.Sleep(200 * time.Microsecond)
	generator.SetHigh()
	time.Sleep(200 * time.Microsecond)
}

// Measure makes a measurement with the single photon detectors connected to the analogue output for the
// corresponding qubit generator source.
func (c *Circuit) Measure(detector *[]Detector, qubitID string) int {
	for _, d := range *detector {
		switch qubitID {
		case Q1.Strings():
			if d.GetID() == SPD1.Strings() {
				return c.Detect(d, qubitID)
			}
		case Q2.Strings():
			if d.GetID() == SPD2.Strings() {
				return c.Detect(d, qubitID)
			}
		case Q3.Strings():
			if d.GetID() == SPD3.Strings() {
				return c.Detect(d, qubitID)
			}
		case Q4.Strings():
			if d.GetID() == SPD4.Strings() {
				return c.Detect(d, qubitID)
			}
		}
	}
	return 0
}

// Detect takes a reading from the single photon detector and encodes the result based on the superpositions for which
// the system was prepared.
func (c *Circuit) Detect(detector Detector, qubitID string) int {
	verticalValue := detector.GetVertical()
	horizontalValue := detector.GetHorizontal()
	detectorID := detector.GetID()
	fmt.Printf("%s -> %s VERTICAL   |1>: %d \n", qubitID, detectorID, verticalValue)
	fmt.Printf("%s -> %s HORIZONTAL |0>: %d \n", qubitID, detectorID, horizontalValue)
	decodedMeasurement := c.Encode(horizontalValue, verticalValue)
	return decodedMeasurement
}

// Encode converts the detected values for the
func (c *Circuit) Encode(horizontalLight, verticalLight uint16) int {
	if horizontalLight > verticalLight {
		return 0
	} else if verticalLight >= horizontalLight {
		return 1
	} else {
		return c.Encode(horizontalLight, verticalLight)
	}
}

// Statistics capture simple statistics on the output of the circuit for the duration of the runtime.
func (c *Circuit) Statistics() {
	if len(c.Bits) > 0 {
		for i := range c.Bits {
			if c.Bits[i] == 0 {
				c.ZeroCount++
			}
			if c.Bits[i] == 1 {
				c.OneCount++
			}
		}
	}
	if c.ZeroCount >= 1 && c.OneCount >= 1 {
		total := c.ZeroCount + c.OneCount
		fmt.Printf("Summary - #Total: #%d#  0: #%d# 1: #%d#\n", total, c.ZeroCount, c.OneCount)
	}
}
