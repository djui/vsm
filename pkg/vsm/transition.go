package vsm

// Transition defines a valid state transition.
type Transition struct {
	From State
	To   State

	// Inputs
	Roles []Role
}

// DefaultTransitions defines the default abstract vehicle life-cycle.
var DefaultTransitions = []Transition{}
