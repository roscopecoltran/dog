package component

import (
	"context"
	"fluided/connectable"
	"fluided/controllable"
	"fmt"
	"sync"
	"time"
)

/*
	Simple component implementation
*/
type simpleComponent struct {
	name      string
	status    controllable.Status
	ctx       context.Context
	in        <-chan *connectable.Message
	out       chan<- *connectable.Message
	initiated *sync.Once
}

// Cloneable interface implementation
func (sc *simpleComponent) Clone(params ...interface{}) interface{} {
	return &simpleComponent{name: sc.name + fmt.Sprint(time.Now()), initiated: &sync.Once{}, status: controllable.NotInitialized, ctx: params[0].(context.Context), in: make(<-chan *connectable.Message), out: make(chan<- *connectable.Message)}
}

// Component interface implementation1
func (sc *simpleComponent) Name() string {
	return sc.name
}

func (sc *simpleComponent) In() <-chan *connectable.Message {
	return sc.in
}

func (sc *simpleComponent) Out() chan<- *connectable.Message {
	return sc.out
}

func (sc *simpleComponent) Connect(ctx context.Context, in <-chan *connectable.Message) error {
	sc.ctx = ctx
	sc.in = in
	return nil
}

func (sc *simpleComponent) Process(message *connectable.Message) (*connectable.Message, error) {
	// default behavior
	if sc.status == controllable.Started {
		return message, nil
	}
	return nil, fmt.Errorf("Component %s not started (%s", sc.name, fmt.Sprint(sc.status))
}

// Controllable interface implementation
func (sc *simpleComponent) Initialize() (controllable.Status, error) {

	f := func() {
		go func() {
			defer sc.Shutdown()

			for {
				select {
				/* Check for broadcast shutdown */
				case <-sc.ctx.Done():
					return
				/* Launching the consuming loop */
				case message := <-sc.in:
					{
						var err error
						if message, err = sc.Process(message); err != nil {
							message.Error = append(message.Error, err)
						}
						sc.out <- message
					}
				}
			}
		}()
	}

	sc.initiated.Do(f)

	sc.status = controllable.Initialized
	return sc.status, nil
}

func (sc *simpleComponent) Start() (controllable.Status, error) {
	sc.status = controllable.Started
	return sc.status, nil
}

func (sc *simpleComponent) Stop() (controllable.Status, error) {
	sc.status = controllable.Stopped
	return sc.status, nil
}

func (sc *simpleComponent) Shutdown() (controllable.Status, error) {
	sc.status = controllable.Shuwndown
	return sc.status, nil
}
