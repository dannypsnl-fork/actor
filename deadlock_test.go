package actor

import (
	"testing"
)

type EchoServer struct {
	Actor
}

func (e *EchoServer) Fun() {
	msg := <-e.Receive
	if v, ok := msg.(msgWith); ok {
		v.sender <- v.content
		go e.Fun()
	}
	if _, ok := msg.(closing); ok {
		close(e.Receive)
	}
}

type msgWith struct {
	sender  chan string
	content string
}

type closing struct{}

func TestActorShouldNotAtReAskPidPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error(r)
		}
	}()
	echo := &EchoServer{}
	echo.Init()
	echoServer := echo.Pid()

	go echo.Fun()

	self := make(chan string)
	defer close(self)

	echoServer <- msgWith{self, "Hello"}
	<-self
	echoServer = echo.Pid()
	echoServer <- msgWith{self, "Hello"}
	<-self
	echoServer <- msgWith{self, "Hello"}
	<-self
	echoServer <- closing{}
}
