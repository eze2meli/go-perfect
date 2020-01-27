package ex2_oop_inheritance

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Abstract struct {

}

func (Abstract) RepeatedCode1() {
	// Do something
}

func (Abstract) RepeatedCode2() {
	// Do something
}

func (Abstract) RepeatedCode3() {
	// Do something
}

func (Abstract) OverrideableMethod() string {
	return "I am abstract"
}

type Particular struct {
	Abstract
}

func (Particular) OverrideableMethod() string {
	return "I am particular"
}

// Our software requires that component must be typed as Abstract
type Third struct {
	component Abstract
}

// TODO 2: Make this test pass
func TestExercise(t *testing.T) {
	// HINT implement an "inheritance" relation

	third := Third{}
	third.component = Abstract{} // component MUST be Abstract, but must be Particular too

	// Dont change anything here
	assert.Equal(t, "I am particular", third.component.OverrideableMethod())
	anotherAbstract := Abstract{}
	assert.Equal(t, "I am abstract", anotherAbstract.OverrideableMethod())
}