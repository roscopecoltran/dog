package component

import (
	"context"
	"fluided/cloneable"
	"fluided/connectable"
	"fluided/controllable"
	"fluided/instantiable"
)

func init() {

}

// Component interface
type Component interface {
	controllable.Controllable
	instantiable.Instantiable
	cloneable.Cloneable
	Connect(context.Context, chan<- *connectable.Message) error
	Process(*connectable.Message) (*connectable.Message, error)
	Name() string
	In() chan *connectable.Message
	Out() chan *connectable.Message
}
