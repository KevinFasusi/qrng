package main

import (
	"machine"
	"qrng-firmware/quantum"
)

func main() {

	qb1 := quantum.NewQubitGenerator(machine.D2, quantum.Q1.Strings())
	qb2 := quantum.NewQubitGenerator(machine.D3, quantum.Q2.Strings())
	qb3 := quantum.NewQubitGenerator(machine.D4, quantum.Q3.Strings())
	qb4 := quantum.NewQubitGenerator(machine.D5, quantum.Q4.Strings())

	var generatorQueue []quantum.Generator

	generatorQueue = append(generatorQueue, qb1)
	generatorQueue = append(generatorQueue, qb2)
	generatorQueue = append(generatorQueue, qb3)
	generatorQueue = append(generatorQueue, qb4)

	// pairs
	spd1 := quantum.NewSinglePhotonDetector(machine.A0, machine.A1, quantum.SPD1.Strings())
	spd2 := quantum.NewSinglePhotonDetector(machine.A2, machine.A3, quantum.SPD2.Strings())
	spd3 := quantum.NewSinglePhotonDetector(machine.A4, machine.A5, quantum.SPD3.Strings())
	spd4 := quantum.NewSinglePhotonDetector(machine.A6, machine.A7, quantum.SPD4.Strings())

	var singlePhotonDetectorQueue []quantum.Detector

	singlePhotonDetectorQueue = append(singlePhotonDetectorQueue, spd1)
	singlePhotonDetectorQueue = append(singlePhotonDetectorQueue, spd2)
	singlePhotonDetectorQueue = append(singlePhotonDetectorQueue, spd3)
	singlePhotonDetectorQueue = append(singlePhotonDetectorQueue, spd4)

	bias := quantum.NewBias([]string{
		quantum.Q1.Strings(),
		quantum.Q2.Strings(),
		quantum.Q3.Strings(),
		quantum.Q4.Strings()})

	qc := quantum.NewCircuit(&generatorQueue, &singlePhotonDetectorQueue, bias)

	for {
		qc.Run()
	}
}
