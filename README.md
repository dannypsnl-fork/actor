# actor

[![Build Status](https://travis-ci.org/dannypsnl/actor.svg?branch=master)](https://travis-ci.org/dannypsnl/actor)
[![Go Report Card](https://goreportcard.com/badge/github.com/dannypsnl/actor)](https://goreportcard.com/report/github.com/dannypsnl/actor)
[![GoDoc](https://godoc.org/github.com/dannypsnl/actor?status.svg)](https://godoc.org/github.com/dannypsnl/actor)
[![GitHub license](https://img.shields.io/github/license/dannypsnl/actor.svg)](https://github.com/dannypsnl/actor/blob/master/LICENSE)

Implementation Actor model in Go.

## Design

Actor = {receive chan interface{}}

## Usage

### Install

```bash
go get github.com/dannypsnl/actor
```

### Example

```go
import (
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
    println(result) // expected: Hi
    echoPid <- closing{}
}
```

## Note

Erlang won't receive one message twice.
