package machinery

type MachineryAction struct {
	name string
}

func Action(name string) *MachineryAction {
	return &MachineryAction{
		name: name,
	}
}
