package instantiable

// Instantiable interface
type Instantiable interface {
	Instantiate() (interface{}, error)
}
