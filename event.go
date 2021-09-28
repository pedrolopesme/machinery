package machinery

type Event interface {
	Event() string
	ToState() State
}

type BasicEvent struct {
	event   string
	toState State
}

func (e BasicEvent) Event() string {
	return e.event
}

func (e BasicEvent) ToState() State {
	return e.toState
}

// NewEvent add an Event that changes Machine to some state
func NewEvent(name string, to State) *BasicEvent {
	return &BasicEvent{
		event:   name,
		toState: to,
	}
}
