package flow

import (
	"context"
	"fluided/component"
	"fluided/controllable"
)

func init() {

}

type simpleFlow struct {
	name       string
	components map[string]component.Component
	status     controllable.Status
	ctx        context.Context
}

// Flow interface implementation
func (flow *simpleFlow) AddComponent(newcomponent component.Component) error {
	flow.components[newcomponent.Name()] = newcomponent
	return nil
}

func (flow *simpleFlow) Connect(producer component.Component, consumer component.Component) error {
	consumer.Connect(nil, producer.Out())
	return nil
}

func (flow *simpleFlow) Name() string {
	return flow.name
}

// Cloneable interface implementation
func (flow *simpleFlow) Clone(params ...interface{}) interface{} {
	components := make(map[string]component.Component)
	for _, _component := range flow.components {
		components[_component.Name()] = _component.Clone(flow.ctx).(component.Component)
	}
	return &simpleFlow{components: components, status: controllable.NotInitialized, ctx: params[0].(context.Context)}
}

/*
TODO: To implement controllable.Controllable interface
*/
// Controllable interface implementation
func (flow *simpleFlow) Initialize() (controllable.Status, error) {
	flow.status = controllable.Initialized

	// TODO: Initialize all of its components

	return flow.status, nil
}

func (flow *simpleFlow) Start() (controllable.Status, error) {

	// TODO: Start all of its components

	flow.status = controllable.Started
	return flow.status, nil
}

func (flow *simpleFlow) Stop() (controllable.Status, error) {

	// TODO: Stop all of its components

	flow.status = controllable.Stopped
	return flow.status, nil
}

func (flow *simpleFlow) Shutdown() (controllable.Status, error) {

	// TODO: Shutdown all of its components

	flow.status = controllable.Shuwndown
	return flow.status, nil
}
