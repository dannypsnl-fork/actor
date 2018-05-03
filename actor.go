package actor

type Actor struct {
	receive chan interface{}
}

type Actorable interface {
	Recv() chan interface{}
}
