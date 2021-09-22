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
