package vsm

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInitState(t *testing.T) {
	m := New(DefaultTransitions)
	assert.Equal(t, m.state, StateReady)
}

func TestInvalidTransition(t *testing.T) {
	// Arrange

	m := New(DefaultTransitions)

	// Act

	actualErr := m.Transition(StateCollected, RoleEndUser)

	// Assert

	assert.Error(t, actualErr, ErrInvalidTransition{StateReady, StateCollected}.Error())
}

func TestDeniedTransition(t *testing.T) {
	// Arrange

	m := New(DefaultTransitions)
	m.state = StateBounty

	// Act

	actualErr := m.Transition(StateCollected, RoleEndUser)

	// Assert

	assert.Error(t, actualErr, ErrInvalidPermission{role: RoleEndUser, permitted: []Role{RoleHunter}}.Error())
}

func TestAutomaticTransition(t *testing.T) {
	// Arrange

	m := New(DefaultTransitions)

	// Act

	actualErr := m.Transition(StateUnknown, RoleAutomatic)

	// Assert

	require.NoError(t, actualErr)
	assert.Equal(t, StateUnknown, m.state)
}

func TestDecommissionedStateTransition(t *testing.T) {
	// Arrange

	m := New(DefaultTransitions)
	m.state = StateUnknown

	// Act

	actualErr := m.Transition(StateRiding, RoleEndUser)

	// Assert

	assert.Error(t, actualErr, ErrInvalidTransition{StateUnknown, StateRiding}.Error())
}

func TestAdminImplicitTransition(t *testing.T) {
	// Arrange

	m := New(DefaultTransitions)

	// Act

	actualErr := m.Transition(StateServiceMode, RoleAdmin)

	// Assert

	require.NoError(t, actualErr)
	assert.Equal(t, StateServiceMode, m.state)
}
