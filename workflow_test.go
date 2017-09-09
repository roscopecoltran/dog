package scipipe

import (
	"github.com/stretchr/testify/assert"
	"sync"
	t "testing"
)

func TestAdd(t *t.T) {
	InitLogError()
	wf := NewWorkflow("TestAddProcsWf")

	proc1 := NewBogusProcess("bogusproc1")
	wf.Add(proc1)
	proc2 := NewBogusProcess("bogusproc2")
	wf.Add(proc2)

	assert.EqualValues(t, len(wf.procs), 2)

	assert.IsType(t, &BogusProcess{}, wf.procs["bogusproc1"], "Process 1 was not of the right type!")
	assert.IsType(t, &BogusProcess{}, wf.procs["bogusproc2"], "Process 2 was not of the right type!")
}

// --------------------------------
// Helper stuff
// --------------------------------

// A process with does just satisfy the Process interface, without doing any
// actual work.
type BogusProcess struct {
	Process
	name       string
	WasRun     bool
	WasRunLock sync.Mutex
}

var bogusWg sync.WaitGroup

func NewBogusProcess(name string) *BogusProcess {
	return &BogusProcess{WasRun: false, name: name}
}

func (p *BogusProcess) Run() {
	p.WasRunLock.Lock()
	p.WasRun = true
	p.WasRunLock.Unlock()
}

func (p *BogusProcess) Name() string {
	return p.name
}

func (p *BogusProcess) IsConnected() bool {
	return true
}
