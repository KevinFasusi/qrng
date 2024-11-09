package quantum

// System implements a quantum systems for controlling several circuits
//
// Currently the circuit uses four qubit generators to compose a single circuit. The system interface allows control of
// multiple circuits running different logic and rules.
type System interface {
	Run()
	Detect(detector Detector, qubitID string) int
	Encode(zero, one uint16) int
	Statistics()
}
