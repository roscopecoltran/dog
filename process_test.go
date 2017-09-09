package scipipe

import (
	"fmt"
	"testing"
)

func TestNewProc(t *testing.T) {
	p1 := NewProc("echo", "echo {p:text}")
	if p1.ParamPort("text") == nil {
		t.Error(`p.ParamPorts("text") = nil. want: not nil`)
	}

	p2 := NewProc("cat", "cat {i:infile} > {o:outfile}")
	if p2.In("infile") == nil {
		t.Error(`p.In("infile") = nil. want: not nil`)
	}
	if p2.Out("outfile") == nil {
		t.Error(`p.Out("outfile") = nil. want: not nil`)
	}
}

func ExampleExpandParams() {
	fmt.Println(ExpandParams("echo {p:msg}", map[string]string{"msg": "Hello"}))
	// Output:
	// echo Hello
}

func TestShellExpand_OnlyParams(t *testing.T) {
	p1 := ShellExpand("echo", "echo {p:foo}", nil, nil, map[string]string{"foo": "bar"})
	if p1.CommandPattern != "echo bar" {
		t.Error(`p.CommandPattern != "echo bar", want: echo bar`)
	}
}

func TestShellExpand_InputOutput(t *testing.T) {
	p := ShellExpand("cat", "cat {i:foo} > {o:bar}", map[string]string{"foo": "foo.txt"}, map[string]string{"bar": "bar.txt"}, nil)
	if p.CommandPattern != "cat foo.txt > bar.txt" {
		t.Error(`if p.CommandPattern != "cat foo.txt > bar.txt"`)
	}
}

func TestSetPathStatic(t *testing.T) {
	p := NewProc("echo_foo", "echo foo > {o:bar}")
	p.SetPathStatic("bar", "bar.txt")

	mockTask := NewSciTask("echo_foo_task", "", nil, nil, nil, nil, "", p.ExecMode)

	if p.PathFormatters["bar"](mockTask) != "bar.txt" {
		t.Error(`p.PathFormatters["bar"]() != "bar.txt"`)
	}
}

func TestSetPathExtend(t *testing.T) {
	p := NewProc("cat_foo", "cat {i:foo} > {o:bar}")
	p.SetPathExtend("foo", "bar", ".bar.txt")

	mockTask := NewSciTask("echo_foo_task", "", map[string]*InformationPacket{"foo": NewInformationPacket("foo.txt")}, nil, nil, nil, "", p.ExecMode)

	if p.PathFormatters["bar"](mockTask) != "foo.txt.bar.txt" {
		t.Error(`p.PathFormatters["bar"]() != "foo.txt.bar.txt"`)
	}
}

func TestSetPathReplace(t *testing.T) {
	p := NewProc("cat_foo", "cat {i:foo} > {o:bar}")
	p.SetPathReplace("foo", "bar", ".txt", ".bar.txt")

	mockTask := NewSciTask("echo_foo_task", "", map[string]*InformationPacket{"foo": NewInformationPacket("foo.txt")}, nil, nil, nil, "", p.ExecMode)

	if p.PathFormatters["bar"](mockTask) != "foo.bar.txt" {
		t.Error(`p.PathFormatters["bar"]() != "foo.bar.txt"`)
	}
}
