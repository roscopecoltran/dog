package main

import (
	"bytes"
	"fmt"

	r "reflect"

	"strings"

	sci "github.com/scipipe/scipipe"
)

func main() {
	p := NewFooToBarReplacer()
	fmt.Println("Process: ", p)
}

// -------------------------------------------
//  Example of defining a new wrapper task
// -------------------------------------------

type FooToBarReplacer struct {
	sci.Process
	InnerProcess *sci.Process
	Run          func(p *FooToBarReplacer)
	InFoo        chan *sci.InformationPacket
	OutBar       chan *sci.InformationPacket
}

func NewFooToBarReplacer() interface{} {
	execFunc := func(task *sci.SciTask) {
		indata := task.InTargets["foo"].Read()
		indataReplaced := bytes.Replace(indata, []byte("foo"), []byte("bar"), -1)
		task.OutTargets["bar"].WriteTempFile(indataReplaced)
	}
	pathFuncs := map[string]func(*sci.SciTask) string{
		"bar": func(t *sci.SciTask) string { return t.InTargets["foo"].GetPath() + ".bar.txt" },
	}
	return NewProcessFromStruct(&FooToBarReplacer{}, execFunc, pathFuncs)
}

// -------------------------------------------
//  New helper methods
// -------------------------------------------

func NewProcessFromStruct(procStruct interface{}, execFunc func(*sci.SciTask), pathFuncs map[string]func(*sci.SciTask) string) interface{} {
	// Get in-ports of struct
	inPorts := map[string]chan *sci.InformationPacket{}
	outPorts := map[string]chan *sci.InformationPacket{}

	procStructVal := r.ValueOf(procStruct).Elem()
	for i := 0; i < procStructVal.NumField(); i++ {
		structFieldName := procStructVal.Type().Field(i).Name
		structFieldType := procStructVal.Type().Field(i).Type
		exampleChan := make(chan *sci.InformationPacket)
		if strings.HasPrefix(structFieldName, "In") && structFieldType == r.TypeOf(exampleChan) {
			fmt.Println("In-port:", structFieldName)
			inPorts[strings.ToLower(structFieldName)] = exampleChan // TODO: Change this!
		} else if strings.HasPrefix(structFieldName, "Out") && structFieldType == r.TypeOf(exampleChan) {
			fmt.Println("Out-port:", structFieldName)
			outPorts[strings.ToLower(structFieldName)] = exampleChan // TODO: Change this!
		}
	}
	return procStruct
}
