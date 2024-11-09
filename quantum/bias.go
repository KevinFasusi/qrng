package quantum

// Bias maps the state for generators which have been measured.
type Bias struct {
	states map[string]int
}

// NewBias returns a Bias struct with states all set to 2, representing an impossible state and placeholder
func NewBias(generatorIDSs []string) *Bias {
	states := make(map[string]int)
	for _, id := range generatorIDSs {
		states[id] = 2
	}
	return &Bias{states: states}
}

// Postprocessor implements Jon von Neumann's skew correction on for each qubit source.
//
// The Postprocessor compares the last current state for a qubit source, based on their id, returning true for states that
// are not equal.
func (b *Bias) Postprocessor(state int, id string) bool {
	if b.states[id] == state {
		b.states[id] = 2
		return false
	}
	if b.states[id] != state {
		b.states[id] = state
	}
	return true
}
