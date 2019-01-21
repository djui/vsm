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

func TestDefaultEvents(t *testing.T) {
	cases := []struct {
		initState     State
		givenEvent    Event
		givenRole     Role
		expectedState State
	}{
		{StateReady, EventStartRide, RoleEndUser, StateRiding},
		{StateRiding, EventEndRide, RoleEndUser, StateReady},

		{StateReady, EventNighttime, RoleAutomatic, StateBounty},
		{StateReady, EventExpire, RoleAutomatic, StateUnknown},
		{StateRiding, EventSafeBattery, RoleAutomatic, StateBatteryLow},
		{StateBatteryLow, EventNotifyHunter, RoleAutomatic, StateBounty},

		{StateReady, EventStartRide, RoleHunter, StateRiding},
		{StateRiding, EventEndRide, RoleHunter, StateReady},
		{StateBounty, EventCollect, RoleHunter, StateCollected},
		{StateCollected, EventReturn, RoleHunter, StateDropped},
		{StateDropped, EventDistribute, RoleHunter, StateReady},
	}

	m := New(DefaultTransitions)

	for _, tc := range cases {
		t.Run(tc.givenEvent.String(), func(t *testing.T) {
			// Arrange

			m.state = tc.initState

			// Act

			err := m.Transition(tc.givenEvent.State, tc.givenRole)

			// Assert

			require.NoError(t, err)
			assert.Equal(t, tc.expectedState, m.state)
		})
	}

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
