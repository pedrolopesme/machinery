package machinery

type State interface {
	State() string
}

type BasicState struct {
	state string
}

func NewState(state string) *BasicState {
	return &BasicState{
		state: state,
	}
}

func (s BasicState) State() string {
	return s.state
}

type Machine interface {
	AddTransition(from State, to []State)
	Change(from State, to State) error
	State() State
}

type Machinery struct {
	currentState State
	transitions  map[State][]State
}

func NewMachinery(initialState State) *Machinery {
	if initialState == nil {
		return nil
	}

	return &Machinery{
		currentState: initialState,
		transitions:  make(map[State][]State),
	}
}

// TODO implementation
// TODO add tests
func (m *Machinery) AddTransition(from State, to []State) {
}

// TODO implementation
// TODO add tests
func (m *Machinery) Change(from State, to State) error {
	return nil
}

// State returns the current state
func (m Machinery) State() State {
	return m.currentState
}
