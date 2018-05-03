package actor

import (
	"testing"
)

type noFun struct {
	Actor
}

func TestSpawnReceiveActorNoFun(t *testing.T) {
	defer func() {
		expected := "Spawn expected actor have method Fun"
		if r := recover(); r != expected {
			t.Errorf("expected: `%s`, actual: `%s`", expected, r)
		}
	}()
	a := &noFun{}
	Spawn(a, []interface{}{})
}

func TestSpawnInvaildParameters(t *testing.T) {
	defer func() {
		expected := "expected: int, but receive: string"
		if r := recover(); r != expected {
			t.Errorf("expected: `%s`, actual: `%s`", expected, r)
		}
	}()
	sayHi := &SayHello{}
	// expected use int start sayHi.Fun
	Spawn(sayHi, []interface{}{"Hi"})
}
