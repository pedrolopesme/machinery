package machinery

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMachinery(t *testing.T) {

	t.Run("WhenAnInitialStateWasProvided_ThenMachinery_ShouldKeepTheCurrentState", func(t *testing.T) {
		state1 := NewState("first")
		machine := NewMachinery(state1)
		assert.Equal(t, state1, machine.State())
	})

	t.Run("WhenAnNilInitialStateWasProvided_ThenMachinery_ShouldNotReturnANewMachine", func(t *testing.T) {
		machine := NewMachinery(nil)
		assert.Nil(t, machine)
	})
}

func TestEvents(t *testing.T) {

	t.Run("WhenMovingFromASimpleEventToAnother_ThenMachinery_ShouldChangeMachineState", func(t *testing.T) {
		state1 := State("begin")
		state2 := State("end")

		shoot := Action("gun_shoot")
		play := Event("attack").AddStateFrom(state1).StateTo(state2).AllowAction(shoot)

		machine := NewMachinery(state1)
		machine.AddEvent(play)
		changed, err := machine.Do(shoot)

		assert.Nil(t, err)
		assert.True(t, changed)
		assert.Equal(t, state2, machine.State())
	})

	t.Run("WhenMovingFromASimpleEventToAnother_ButEventDoenstExist_ThenMachinery_ShouldNotChangeMachineState", func(t *testing.T) {
		state1 := NewState("begin")
		machine := NewMachinery(state1)

		err := machine.Trigger("missing event")
		assert.NotNil(t, err)
		assert.Equal(t, state1, machine.State())
	})

}
