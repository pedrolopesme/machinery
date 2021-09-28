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
