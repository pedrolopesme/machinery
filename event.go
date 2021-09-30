package machinery

type MachineryEvent interface {
	String() string
	AddStateFrom(...MachineryState)
	StateTo(MachineryState)
	AllowAction(...MachineryAction)
}

type BasicEvent struct {
	event      string
	statesFrom []MachineryState
	stateTo    MachineryState
	actions    []MachineryAction
}

func (e BasicEvent) Event() string {
	return e.event
}

// AddStateFrom allows to set one (or more) states which this event can move the machine FROM
func (e *BasicEvent) AddStateFrom(states ...MachineryState) *BasicEvent {
	e.statesFrom = append(e.statesFrom, states...)
	return e
}

// StateTo allows to set one state which this event can move the machine TO
func (e *BasicEvent) StateTo(state MachineryState) *BasicEvent {
	e.stateTo = state
	return e
}

// AllowAction allows to set one (or more) actions that trigger this event to happen
func (e *BasicEvent) AllowAction(actions ...MachineryAction) *BasicEvent {
	e.actions = append(e.actions, actions...)
	return e
}

// NewEvent add an Event that changes Machine to some state
func Event(name string) *BasicEvent {
	return &BasicEvent{
		event: name,
	}
}
