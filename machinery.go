package machinery

type State interface {
	State() string
}

type Machine interface {
	AddTransition(from State, to []State)
	Change(from State, to State) error
}

type Machinery struct {
	currentState State
	transitions  map[State][]State
}

func NewMachinery(initialState State) *Machinery {
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
