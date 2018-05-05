package main

import (
	"fmt"
	"github.com/dannypsnl/actor"
)

type Echo struct {
	actor.Actor
}

func (e *Echo) Fun() {
	msg := <-e.Receive
	switch msg.(type) {
	case strWithSender:
		m := msg.(strWithSender)
		m.sender <- m.str
		e.Fun()
	case closing:
		close(e.Receive)
	default:
		panic("Just don't want to resolve U!")
	}
}

type closing struct{}

type strWithSender struct {
	sender chan string
	str    string
}

func main() {
	echoPid := actor.Spawn(&Echo{}, []interface{}{})
	self := make(chan string)
	echoPid <- strWithSender{self, "Hi"}
	result := <-self
	fmt.Println(result) // expected: Hi
	echoPid <- closing{}
}
