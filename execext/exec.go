package execext

import (
	"context"
	"errors"
	"io"
	"strings"

	"mvdan.cc/sh/interp"
	"mvdan.cc/sh/syntax"
)

// RunCommandOptions is the options for the RunCommand func
type RunCommandOptions struct {
	Context context.Context
	Command string
	Dir     string
	Env     []string
	Stdin   io.Reader
	Stdout  io.Writer
	Stderr  io.Writer
}

var (
	// ErrNilOptions is returned when a nil options is given
	ErrNilOptions = errors.New("execext: nil options given")
)

// RunCommand runs a shell command
func RunCommand(opts *RunCommandOptions) error {
	if opts == nil {
		return ErrNilOptions
	}

	p, err := syntax.NewParser().Parse(strings.NewReader(opts.Command), "")
	if err != nil {
		return err
	}

	r := interp.Runner{
		Context: opts.Context,
		Dir:     opts.Dir,
		Env:     opts.Env,
		Stdin:   opts.Stdin,
		Stdout:  opts.Stdout,
		Stderr:  opts.Stderr,
	}
	if err = r.Reset(); err != nil {
		return err
	}
	return r.Run(p)
}
