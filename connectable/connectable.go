package connectable

import "context"

// Message
type Message struct {
	Parameters []interface{}
	Ctx        context.Context
	Error      []error
}

type Connectable interface {
	In() chan *Message
	Out() chan *Message
	Connect(ctx context.Context, in chan<- *Message)
}

func MakeIn() chan *Message {
	return make(chan *Message)
}

func MakeOut() chan *Message {
	return make(chan *Message)
}
