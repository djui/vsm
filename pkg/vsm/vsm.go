package vsm

// VSM is a finite state machine for an abstract vehicle.
type VSM struct {
	graph map[State]map[State][]Role

	state   State
}

// New instantiates a new VSM with initial state from the first transition.
//
// It builds a graph with state vertices given a set of transition edges.
func New(transitions []Transition) *VSM {
	g := map[State]map[State][]Role{}

	for _, t := range transitions {
		if _, found := g[t.From]; !found {
			g[t.From] = map[State][]Role{}
		}
		g[t.From][t.To] = t.Roles
	}

	var initState State
	if len(transitions) > 0 {
		initState = transitions[0].From
	}

	return &VSM{
		state: initState,
		graph: g,
	}
}

// Transition executes a state transition after successful validation.
func (m *VSM) Transition(state State, role Role) error {
	// Allow admin override
	if role != RoleAdmin {
		// Check if transition is valid
		roles, valid := m.graph[m.state][state]
		if !valid {
			return &ErrInvalidTransition{from: m.state, to: state}
		}

		// Check if permissions are valid
		if !role.IsMember(roles) {
			return &ErrInvalidPermission{role: role, permitted: roles}
		}
	}

	m.state = state
	return nil
}

// State returns the current state.
func (m *VSM) State() State {
	return m.state
}
