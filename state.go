package machinery

type MachineryState interface {
	String() string
}

func State(state string) *BasicState {
	return &BasicState{
		state: state,
	}
}

type BasicState struct {
	state string
}

func (s BasicState) String() string {
	return s.state
}
