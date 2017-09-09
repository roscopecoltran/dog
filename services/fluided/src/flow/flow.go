package flow

import (
	"fluided/cloneable"
	"fluided/component"
	"fluided/controllable"
)

func init() {

}

// Flow interface
type Flow interface {
	controllable.Controllable
	cloneable.Cloneable

	Name() string
	AddComponent(component.Component) error
	Connect(producer component.Component, consumer component.Component) error
}
