package actor

import (
	"testing"
)

// SayHello is actor implementation for testing
type SayHello struct {
	Actor
}

// Recv to match Spawn required
func (s *SayHello) Recv() chan interface{} {
	s.receive = make(chan interface{})
	return s.receive
}

// Fun is another standard
// Start an actor by Spawn need to implement this method
func (s *SayHello) Fun(i int) {
	msg := <-s.receive
	switch msg.(type) {
	case intWithRecv:
		msg.(intWithRecv).recv <- msg.(intWithRecv).value + i
	default:
		panic("This actor do not handle this kind of message")
	}
	close(s.receive)
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
	sayHi := &SayHello{}
	go sayHi.Fun(10)
	recv := make(chan int)
	defer close(recv)
	pid := sayHi.Recv()
	pid <- intWithRecv{recv, 10}
	result := <-recv
	if result != 20 {
		t.Errorf("result: %d", result)
	}
}
