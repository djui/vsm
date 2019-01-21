package sim

import "time"

// Clock is a time source.
type Clock interface {
	Now() time.Time
}

var _ Clock = (*WallClock)(nil)

// WallClock defines a system's wall clock.
type WallClock struct{}

// Now implements the Clock interface.
func (c *WallClock) Now() time.Time {
	return time.Now()
}

var _ Clock = (*FixedClock)(nil)

// FixedClock defines a clock with a fixed time.
type FixedClock struct {
	Time time.Time
}

// Now implements the Clock interface.
func (c *FixedClock) Now() time.Time {
	return c.Time
}
