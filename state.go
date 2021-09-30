package machinery

type MachineryState interface {
	State() string
}

func State(state string) *BasicState {
	return &BasicState{
		state: state,
	}
}

type BasicState struct {
	state string
}

func (s BasicState) State() string {
	return s.state
}
