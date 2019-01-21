package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/djui/vsm/pkg/sim"
	"github.com/djui/vsm/pkg/vsm"
)

var stockholm *time.Location

func init() {
	var err error
	if stockholm, err = time.LoadLocation("Europe/Stockholm"); err != nil {
		panic(err.Error())
	}
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("")
	flag.Parse()

	s := &sim.DiscreteStepper{
		Clock:   &sim.FixedClock{Time: time.Now().In(stockholm)},
		Machine: vsm.New(vsm.DefaultTransitions),
	}

	repl(s)
}

func repl(s sim.Simulator) {
	reader := bufio.NewReader(os.Stdin)

	// REPL
	for {
		// Print

		fmt.Printf("%s %-13s [D,S|E,R]: ", s.Now().Format(time.RFC3339), s.State())

		// Read

		text, _ := reader.ReadString('\n')

		parts := strings.SplitN(text, ",", 3)
		if len(parts) != 3 {
			log.Printf("Error: invalid input: require `<duration>,<state|event>,<role>`")
			continue
		}

		d, err := time.ParseDuration(strings.TrimSpace(parts[0]))
		if err != nil {
			log.Printf("Error: invalid duration: %v", err)
			continue
		}

		state, err := parseStateOrEvent(strings.TrimSpace(parts[1]))
		if err != nil {
			log.Printf("Error: %v", err)
			continue
		}

		role, err := parseRole(strings.TrimSpace(parts[2]))
		if err != nil {
			log.Printf("Error: %v", err)
			continue
		}

		// Evaluate

		past := s.Now()
		uncertaintyPeriod := past.Add(48 * time.Hour)
		evening := time.Date(past.Year(), past.Month(), past.Day(), 21, 30, 0, 0, past.Location())

		s.Step(d)

		// Automatic pre-transitions

		if s.State() == vsm.StateReady && s.Now().After(evening) {
			if err := s.Transition(vsm.StateBounty, vsm.RoleAutomatic); err != nil {
				log.Printf("Error: automatic transition failed: %v", err)
			}
			continue
		}
		if s.State() == vsm.StateReady && s.Now().After(uncertaintyPeriod) {
			if err := s.Transition(vsm.StateUnknown, vsm.RoleAutomatic); err != nil {
				log.Printf("Error: automatic transition failed: %v", err)
			}
			continue
		}

		// Manual transition

		if err := s.Transition(state, role); err != nil {
			log.Printf("Error: transition failed: %v", err)
			continue
		}

		// Automatic post-transitions

		if s.State() == vsm.StateBatteryLow {
			if err := s.Transition(vsm.StateBounty, vsm.RoleAutomatic); err != nil {
				log.Printf("Error: automatic transition failed: %v", err)
				continue
			}
		}
	}
}

var allStates = []vsm.State{
	vsm.StateReady,
	vsm.StateBatteryLow,
	vsm.StateBounty,
	vsm.StateRiding,
	vsm.StateCollected,
	vsm.StateDropped,
	vsm.StateServiceMode,
	vsm.StateTerminated,
	vsm.StateUnknown,
}

var allEvents = []vsm.Event{
	vsm.EventNighttime,
	vsm.EventExpire,
	vsm.EventSafeBattery,
	vsm.EventNotifyHunter,
	vsm.EventStartRide,
	vsm.EventEndRide,
	vsm.EventCollect,
	vsm.EventReturn,
	vsm.EventDistribute,
	vsm.EventTerminate,
	vsm.EventServiceMode,
}

func parseStateOrEvent(stateOrEvent string) (vsm.State, error) {
	for _, s := range allStates {
		if canonicalize(stateOrEvent) == canonicalize(s.String()) {
			return s, nil
		}
	}

	for _, e := range allEvents {
		if canonicalize(stateOrEvent) == canonicalize(e.String()) {
			return e.State, nil
		}
	}

	var allStateAndEventNames []string
	for _, s := range allStates {
		allStateAndEventNames = append(allStateAndEventNames, s.String())
	}
	for _, e := range allEvents {
		allStateAndEventNames = append(allStateAndEventNames, e.String())
	}

	return -1, fmt.Errorf("invalid state/event: %s not in %v", stateOrEvent, strings.Join(allStateAndEventNames, ","))
}

var allRoles = []vsm.Role{
	vsm.RoleAutomatic,
	vsm.RoleAdmin,
	vsm.RoleEndUser,
	vsm.RoleHunter,
}

func parseRole(role string) (vsm.Role, error) {
	for _, r := range allRoles {
		if canonicalize(role) == canonicalize(r.String()) {
			return r, nil
		}
	}

	var allRoleNames []string
	for _, r := range allRoles {
		allRoleNames = append(allRoleNames, r.String())
	}

	return -1, fmt.Errorf("invalid role: %s not in %v", role, strings.Join(allRoleNames, ","))
}

func canonicalize(s string) string {
	s = strings.TrimSpace(s)
	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, "_", "", -1)
	s = strings.Replace(s, "-", "", -1)
	s = strings.ToLower(s)
	return s
}
