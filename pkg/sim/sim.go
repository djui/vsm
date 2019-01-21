package sim

import (
	"time"

	"github.com/djui/vsm/pkg/vsm"
)

// Simulator defines a stepper simulator for VSM transitions.
type Simulator interface {
	Now() time.Time
	Step(d time.Duration)
	State() vsm.State
	Transition(state vsm.State, role vsm.Role) error
}

var _ Simulator = (*ContinuousStepper)(nil)

// ContinuousStepper is a continuously stepping simulator, aka reality.
type ContinuousStepper struct {
	Machine *vsm.VSM
}

// Now implements the Simulator interface.
func (c *ContinuousStepper) Now() time.Time {
	return time.Now()
}

// Step implements the Simulator interface.
func (c *ContinuousStepper) Step(time.Duration) {
	// No-op
}

// State implements the Simulator interface.
func (c *ContinuousStepper) State() vsm.State {
	return c.Machine.State()
}

// Transition implements the Simulator interface.
func (c *ContinuousStepper) Transition(s vsm.State, r vsm.Role) error {
	return nil
}

var _ Simulator = (*DiscreteStepper)(nil)

// DiscreteStepper is a discrete stepping simulator.
type DiscreteStepper struct {
	Clock   *FixedClock
	Machine *vsm.VSM
}

// Now implements the Simulator interface.
func (d *DiscreteStepper) Now() time.Time {
	return d.Clock.Now()
}

// Step implements the Simulator interface.
func (d *DiscreteStepper) Step(duration time.Duration) {
	d.Clock.Time = d.Clock.Time.Add(duration)
}

// State implements the Simulator interface.
func (d *DiscreteStepper) State() vsm.State {
	return d.Machine.State()
}

// Transition implements the Simulator interface.
func (d *DiscreteStepper) Transition(s vsm.State, r vsm.Role) error {
	return d.Machine.Transition(s, r)
}
