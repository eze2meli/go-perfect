package solved

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

func (e *Entity) effect() {
	e.attribute += "sas"
}

func TestExercise(t *testing.T) {
	effected := &Entity{}
	SideEffectPerformer{}.sideEffect(effected)

	// Dont change assert statement
	assert.Equal(t, "sas", effected.attribute)
}

/** Explanation (Esp)

Los receptores con punteros hacen obligatorio la llamada de esos métodos
a partir de estructuras instanciadas como punteros

Al tener el método con puntero, cualquier efecto de lado impacta en el objeto apuntado

*/
