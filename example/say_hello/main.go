package main

import (
	"fmt"

	"github.com/dannypsnl/actor"
)

type Hello struct {
	actor.Actor
}

func (hi *Hello) Fun() {
	defer close(hi.Receive)
	msg := <-hi.Receive
	switch msg.(type) {
	case nameWith:
		resp := fmt.Sprint("Hello,", msg.(nameWith).value)
		msg.(nameWith).recv <- resp
	default:
		panic("Hello do not handle this kind of message")
	}
}

type nameWith struct {
	recv  chan string
	value string
}

func main() {
	hello := &Hello{}
	helloService := actor.Spawn(hello, []interface{}{})

	recv := make(chan string)
	defer close(recv)
	helloService <- nameWith{recv, "Danny"}
	fmt.Printf("Hello says: %s\n", <-recv)
}
