package vsm

// Event defines a transition event.
type Event struct {
	State State
	Role  Role
}

// String returns a human readable event text.
//
// String implements the Stringer interface.
func (e Event) String() string {
	return eventNames[e]
}

// Enumeration of transition events.
var (
	EventNighttime    = Event{State: StateBounty}                   // Automatically after 9.30 pm local time
	EventExpire       = Event{State: StateUnknown}                  // Automatically after 48 hours without a state change
	EventSafeBattery  = Event{State: StateBatteryLow}               // Automatically if battery level drops below 20%
	EventNotifyHunter = Event{State: StateBounty, Role: RoleHunter} // Automatically notify a hunter when battery is low

	EventStartRide   = Event{State: StateRiding}                    // Manually when an end user starts riding
	EventEndRide     = Event{State: StateReady}                     // Manually when an end user ends a ride
	EventCollect     = Event{State: StateCollected, Role: RoleHunter}  // Manually when a hunter collects the vehicle
	EventReturn      = Event{State: StateDropped, Role: RoleHunter} // Manually when a hunter dropped the vehicle
	EventDistribute  = Event{State: StateReady, Role: RoleHunter}   // Manually when the vehicle gets distributed
	EventTerminate   = Event{State: StateTerminated, Role: RoleAdmin}  // Manually when admin terminates vehicle
	EventServiceMode = Event{State: StateServiceMode, Role: RoleAdmin} // Manually when admin maintains vehicle
)

var eventNames = map[Event]string{
	EventNighttime:    "night_time",
	EventExpire:       "expire",
	EventSafeBattery:  "safe_battery",
	EventNotifyHunter: "notify_hunter",

	EventStartRide:   "start_ride",
	EventEndRide:     "end_ride",
	EventCollect:     "collect",
	EventReturn:      "return",
	EventDistribute:  "distribute",
	EventTerminate:   "terminate",
	EventServiceMode: "maintenance",
}
