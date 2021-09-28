package machinery

import "errors"

type Machine interface {
	AddEvent(event Event)
	State() State
	Trigger(eventName string) error
}

type Machinery struct {
	currentState State
	events       map[string]Event
}

func NewMachinery(initialState State) *Machinery {
	if initialState == nil {
		return nil
	}

	return &Machinery{
		currentState: initialState,
		events:       make(map[string]Event),
	}
}

// AddEvent registers an event into machine
func (m *Machinery) AddEvent(event Event) {
	m.events[event.Event()] = event
}

// Trigger fires an event that can change machine current state
func (m *Machinery) Trigger(eventName string) error {
	if event, found := m.events[eventName]; found {
		m.currentState = event.ToState()
		return nil
	}
	return errors.New("Event not found")
}

// State returns the current state
func (m Machinery) State() State {
	return m.currentState
}
