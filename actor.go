package actor

import (
	"bytes"
	"fmt"
	"reflect"
)

type Actor struct {
	receive chan interface{}
}

type Actorable interface {
	Recv() chan interface{}
}

func Spawn(actor Actorable, startArgs []interface{}) chan interface{} {
	act := reflect.ValueOf(actor).MethodByName("Fun")
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
		panic("expected no return!!!")
	}

	inputs := make([]reflect.Value, len(startArgs))
	for k, in := range startArgs {
		inputs[k] = reflect.ValueOf(in)
	}
	go act.Call(inputs)
	return actor.Recv()
}
