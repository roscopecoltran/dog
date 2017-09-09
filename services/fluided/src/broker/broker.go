package broker

import(
	"fluided/component"
	"fluded/flow"
	"fluided/connectable"
	"fluided/controllable"
)



// Broker interface
type Broker interface {
	Controllable

	AddFlow(flow flow.Flow)
	StartFlow(name string, message *connectable.Message) error
	ChangeFlowStatus(name string, controllable.Status) error
	GetFlowStatus(name string) (controllable.Status,error)
	GetActiveFlows() []string)
}


/*
	Simple broker implementation
*/
type simpleBroker struct{
	satus Status
}

func (sb *simpleBroker) Initialize() (Status, error){
	sb.satus = Initialized
	return sb.satus, nil
}

func (sb *simpleBroker) Start() (Status, error) (Status, error){
	sb.status = Started
	return sb.satus, nil
}

func (sb *simpleBroker) Stop() (Status, error) (Status, error){
	sb.status = Stopped
	return sb.satus, nil
}

func (sb *simpleBroker) Shutdown() (Status, error) (Status, error){
	sb.status = Shuwndown
	return sb.satus, nil
}