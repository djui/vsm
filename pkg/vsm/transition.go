package vsm

// Transition defines a valid state transition.
type Transition struct {
	From State
	To   State

	// Inputs
	Roles []Role
}

// DefaultTransitions defines the default abstract vehicle life-cycle.
var DefaultTransitions = []Transition{
	{From: StateReady, To: StateRiding, Roles: RoleAllUsers},
	{From: StateReady, To: StateBounty, Roles: []Role{RoleAutomatic}},
	{From: StateReady, To: StateUnknown, Roles: []Role{RoleAutomatic}},
	{From: StateRiding, To: StateReady, Roles: RoleAllUsers},
	{From: StateRiding, To: StateBatteryLow, Roles: []Role{RoleAutomatic}},
	{From: StateBatteryLow, To: StateBounty, Roles: []Role{RoleAutomatic}},
	{From: StateBounty, To: StateCollected, Roles: []Role{RoleHunter}},
	{From: StateCollected, To: StateDropped, Roles: []Role{RoleHunter}},
	{From: StateDropped, To: StateReady, Roles: []Role{RoleHunter}},
}
