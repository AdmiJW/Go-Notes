package main

import "fmt"

// Enumerated types (enums) are a special case of sum types.
// An enum is a type that can have a fixed number of possible values, each with a distinct name.
// Go doesn't have an enum type as a distinct language feature, but enums are simple to implement using existing language idioms.

// Our enum type `ServerState` has an underlying type of `int`.
type ServerState int

// The possible values for `ServerState` are defined as constants. The special keyword, `iota`, generates successive
// constant values automatically; in this case 0, 1, 2, and so on.
const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

// By implementing the `fmt.Stringer` interface, values of `ServerState` can be printed out or converted to strings.
var stateName = map[ServerState]string {
	StateIdle: "Idle",
	StateConnected: "Connected",
	StateError: "Error",
	StateRetrying: "Retrying",
}
func (ss ServerState) String() string {
	return stateName[ss]
}

// Implement the `EnumIndex` method to return the index of the enum value.
func (ss ServerState) EnumIndex() int {
	return int(ss)
}

// `transition` emulates a state transition for a server - It takes the existing state and returns a new state.
func transition(ss ServerState) ServerState {
	switch ss {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %d", ss))
	}
}

func Enums() {
	ns := transition(StateIdle)
	fmt.Println(ns)
	fmt.Println(ns.EnumIndex())

	ns2 := transition(ns)
	fmt.Println(ns2)
	fmt.Println(ns2.EnumIndex())
}
