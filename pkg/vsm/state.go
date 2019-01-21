package vsm

// State defines a vehicle state.
type State int

// String returns a human readable state text.
//
// String implements the Stringer interface.
func (s State) String() string {
	return stateNames[s]
}

// Enumeration of vehicle states.
const (
	// Operational statutes

	StateReady      State = iota // The vehicle is operational and can be claimed by an end ­user
	StateBatteryLow              // The vehicle is low on battery but otherwise operational. The vehicle cannot be claimed by an end­user but can be claimed by a hunter.
	StateBounty                  // Only available for “Hunters” to be picked up for charging.
	StateRiding                  // An end user is currently using this vehicle; it can not be claimed by another user or hunter.
	StateCollected               // A hunter has picked up a vehicle for charging.
	StateDropped                 // A hunter has returned a vehicle after being charged.

	// Not commissioned for service, not claimable by either end ­users nor hunters.

	StateServiceMode // The vehicle is getting maintained
	StateTerminated  // The vehicle is terminated
	StateUnknown     // The vehicle is in an unknown state
)

var stateNames = map[State]string{
	StateReady:      "Ready",
	StateRiding:     "Riding",
	StateBatteryLow: "Battery-low",
	StateBounty:     "Bounty",
	StateCollected:  "Collected",
	StateDropped:    "Dropped",

	StateServiceMode: "Service_mode",
	StateTerminated:  "Terminated",
	StateUnknown:     "Unknown",
}
