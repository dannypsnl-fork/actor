package actor

import (
	"testing"

	"fmt"
)

type SayHello struct {
	Actor
}

func (s *SayHello) Recv() chan interface{} {
	s.receive = make(chan interface{})
	return s.receive
}

func (s *SayHello) Fun(i int) {
	msg := <-s.receive
	fmt.Println(i + msg.(int))
	close(s.receive)
}

func TestSpawn(t *testing.T) {
	sayHi := &SayHello{}
	pid := Spawn(sayHi, []interface{}{3})
	pid <- 30
}

func TestActorUsingSelf(t *testing.T) {
	sayHi := &SayHello{}
	go sayHi.Fun(10)
	pid := sayHi.Recv()
	pid <- 30
}
