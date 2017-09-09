package cloneable

// Cloneable interface
type Cloneable interface {
	Clone(...interface{}) interface{}
}
