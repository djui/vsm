# vsm

Vehicle state machine.

A library for state transitions on an abstract vehicle with role permissions.

The state graph is constructed by providing adjacent edges (`Transitions`) to
the `VSM` struct. A transition contains a start and end end state as well as
set of permitted roles. From then on, states can be transitioned
(`VSM.Transition`) by providing the destination state and a role. The transition
is validated to that only known edges with permitted roles can traverse the
graph.

The state machine is desired to be stateless, but currently holds the current
state.

A subset of the graph, namely all transitions additionally allowed by the
administrator role, is implemented in the validation logic instead of
constructing a fully connected bi-directly graph.

All automated transitions must be initiated externally and with the explicit
`RoleAutomatic` role.

The repository contains a command-line simulator which can simulate all
automated and manual transitions; see `cmd/vsm`.


## Usage

For help:

    make

To compile:

    make build

To manually try:

    ./vsm


## Testing

    make test
