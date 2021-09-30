package machinery

type MachineryAction interface {
	String() string
}

type BasicAction struct {
	action string
}

func Action(action string) *BasicAction {
	return &BasicAction{
		action: action,
	}
}

func (s BasicAction) String() string {
	return s.action
}
