package actor

import (
	"testing"
)

// SayHello is actor implementation for testing
type SayHello struct {
	Actor
}

// Fun is another standard
// Start an actor by Spawn need to implement this method
func (s *SayHello) Fun(i int) {
	defer close(s.Receive)
	msg := <-s.Receive
	switch msg.(type) {
	case intWithRecv:
		resp := msg.(intWithRecv).value + i
		msg.(intWithRecv).recv <- resp
	default:
		panic("This actor do not handle this kind of message")
	}
}

type intWithRecv struct {
	recv  chan int
	value int
}

func TestSpawn(t *testing.T) {
	sayHi := &SayHello{}
	pid := Spawn(sayHi, []interface{}{3})
	recv := make(chan int)
	defer close(recv)
	pid <- intWithRecv{recv, 10}
	result := <-recv
	if result != 13 {
		t.Errorf("result: %d", result)
	}
}

func TestActorUsingSelf(t *testing.T) {
	sayHi := &SayHello{
		Actor: Actor{make(chan interface{})},
	}
	go sayHi.Fun(10)
	recv := make(chan int)
	defer close(recv)
	pid := sayHi.Pid()
	pid <- intWithRecv{recv, 10}
	result := <-recv
	if result != 20 {
		t.Errorf("result: %d", result)
	}
}
