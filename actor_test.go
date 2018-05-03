package actor

import (
	"testing"

	"bytes"
	"fmt"
	"reflect"
)

type SayHello struct {
	Actor
}

func (s *SayHello) Recv() chan interface{} {
	s.receive = make(chan interface{})
	return s.receive
}

func (s *SayHello) Do(i int) {
	msg := <-s.receive
	fmt.Println(i + msg.(int))
	close(s.receive)
}

func TestSpawn(t *testing.T) {
	sayHi := &SayHello{}
	pid := Spawn(sayHi, []interface{}{3})
	pid <- 30
}

func TestReflect(t *testing.T) {
	sayHi := &SayHello{}
	method := reflect.ValueOf(sayHi).MethodByName("Do")

	fmt.Printf("%v\n", method)
}
