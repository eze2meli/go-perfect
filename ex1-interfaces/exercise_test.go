package ex1_interfaces

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Effected interface {
	effect()
}

var (
	// Type checking
	_ Effected = (*Entity)(nil)
)

type SideEffectPerformer struct {}

func (SideEffectPerformer) sideEffect(e Effected) {
	e.effect()
}

type Entity struct {
	attribute string
}

func (e Entity) effect() {
	e.attribute += "sas"
}

// TODO 1: Make this test pass
// HINT: An interface IS a pointer
func TestExercise(t *testing.T) {
	effected := Entity{}
	SideEffectPerformer{}.sideEffect(effected)

	// Dont change assert statement
	assert.Equal(t, "sas", effected.attribute)
}