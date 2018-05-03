package actor

import (
	"testing"
)

func TestSpawnVaildParameters(t *testing.T) {
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
