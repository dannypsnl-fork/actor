package actor

import (
	"bytes"
	"fmt"
	"reflect"
)

// Actor contains a receive channel help user don't need to understand how actor work at first.
// Just embedded Actor into their custom actor to use it
type Actor struct {
	Receive chan interface{}
}

func (a *Actor) Init() {
	a.Receive = make(chan interface{})
}

// Pid return actor's pid
func (a *Actor) Pid() chan interface{} {
	return a.Receive
}

// Actorable is interface prepare for Spawn.
// With it, you can have a way to return PID(concept, is a channel in fact) in Spawn!
type Actorable interface {
	Init()
	Pid() chan interface{}
}

// Spawn help user start an actor like Erlang way.
// The Actor can execute in Spawn require method Fun & match interface Actorable
func Spawn(actor Actorable, startArgs []interface{}) chan interface{} {
	act := reflect.ValueOf(actor).MethodByName("Fun")
	if act.Kind() == reflect.Invalid {
		panic("Spawn expected actor have method Fun")
	}
	var buf bytes.Buffer
	for _, v := range startArgs {
		t := reflect.TypeOf(v)
		buf.WriteString(t.Name())
	}
	var buf2 bytes.Buffer
	for i := 0; i < act.Type().NumIn(); i++ {
		t := act.Type().In(i)
		buf2.WriteString(t.Name())
	}
	expected := buf2.String()
	input := buf.String()
	if expected != input {
		panic(fmt.Sprintf("expected: %s, but receive: %s", expected, input))
	}
	if act.Type().NumOut() > 0 {
		panic("expected actor Fun no return!!!")
	}

	inputs := make([]reflect.Value, len(startArgs))
	for k, in := range startArgs {
		inputs[k] = reflect.ValueOf(in)
	}
	actor.Init()
	go act.Call(inputs)
	return actor.Pid()
}
