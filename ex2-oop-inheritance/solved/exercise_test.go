package solved

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Abstract struct {
	impl interface{}
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

func (a Abstract) OverrideableMethod() string {
	impl, castOk := a.impl.(interface{ OverrideableMethod() string })
	if castOk {
		return impl.OverrideableMethod()
	}

	// default behavior
	return "I am abstract"
}

type Particular struct {

}

func (Particular) OverrideableMethod() string {
	return "I am particular"
}

type Third struct {
	component Abstract
}

func TestExercise(t *testing.T) {

	third := Third{}
	third.component = Abstract{ impl: Particular{} }

	assert.Equal(t, "I am particular", third.component.OverrideableMethod())
}