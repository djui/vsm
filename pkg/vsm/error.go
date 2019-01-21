package vsm

import "fmt"

// ErrInvalidTransition used for invalid state transitions.
type ErrInvalidTransition struct {
	from State
	to   State
}

// Error implements the error interface.
func (e ErrInvalidTransition) Error() string {
	switch e.from {
	case StateServiceMode,
		StateTerminated,
		StateUnknown:
		return fmt.Sprintf("invalid transition: from %s to %s: vehicle decomissioned", e.from, e.to)
	default:
		return fmt.Sprintf("invalid transition: from %s to %s", e.from, e.to)
	}
}

// ErrInvalidPermission is used for invalid role permissions.
type ErrInvalidPermission struct {
	role      Role
	permitted []Role
}

// Error implements the error interface.
func (e ErrInvalidPermission) Error() string {
	return fmt.Sprintf("invalid permission: %s not member of %s", e.role, e.permitted)
}
