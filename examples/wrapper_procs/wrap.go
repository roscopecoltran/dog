package main

import (
	sci "github.com/scipipe/scipipe"
)

func main() {
	wfl := sci.NewWorkflow("wrapperwf")

	// Fooer
	foo := NewFooer("fooer")
	wfl.Add(foo)

	// Foo2barer
	f2b := NewFoo2Barer("foo2barer")
	f2b.InFoo().Connect(foo.OutFoo())
	wfl.Add(f2b)

	// Sink
	snk := sci.NewSink("sink")
	snk.Connect(f2b.OutBar())
	wfl.SetDriver(snk)

	// Run
	wfl.Run()
}

// ------------------------------------------------
// Components
// ------------------------------------------------

// Fooer
// -----

type Fooer struct {
	*sci.SciProcess
	name string
}

func NewFooer(name string) *Fooer {
	innerFoo := sci.NewProc("fooer", "echo foo > {o:foo}")
	innerFoo.SetPathStatic("foo", "foo.txt")
	return &Fooer{innerFoo, name}
}

// Define static ports

func (p *Fooer) OutFoo() *sci.FilePort {
	return p.Out("foo")
}

// Foo2Barer
// ---------

type Foo2Barer struct {
	*sci.SciProcess
	name string
}

func NewFoo2Barer(name string) *Foo2Barer {
	innerFoo2Bar := sci.NewProc("foo2bar", "sed 's/foo/bar/g' {i:foo} > {o:bar}")
	innerFoo2Bar.SetPathExtend("foo", "bar", ".bar.txt")
	return &Foo2Barer{innerFoo2Bar, name}
}

// Define static ports

func (p *Foo2Barer) InFoo() *sci.FilePort {
	return p.In("foo")
}

func (p *Foo2Barer) OutBar() *sci.FilePort {
	return p.Out("bar")
}
