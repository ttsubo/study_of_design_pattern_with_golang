package state

import "fmt"

// State is interface
type State interface {
	handle(context *Context)
	getConcreateState() string
}

type concreteState struct {
	state string
}

func (c *concreteState) getConcreateState() string {
	return c.state
}

// ConcreteStateBooting is struct
type ConcreteStateBooting struct {
	*concreteState
}

// NewConcreteStateBooting func for initializing ConcreteStateBooting
func NewConcreteStateBooting(state string) *ConcreteStateBooting {
	return &ConcreteStateBooting{
		concreteState: &concreteState{
			state: state,
		},
	}
}

func (c *ConcreteStateBooting) handle(context *Context) {
	fmt.Println("*** パソコンは、起動中です")
	context.SetState(NewConcreteStateRun("running"))
}

// ConcreteStateRun is struct
type ConcreteStateRun struct {
	*concreteState
}

// NewConcreteStateRun func for initializing ConcreteStateRun
func NewConcreteStateRun(state string) *ConcreteStateRun {
	return &ConcreteStateRun{
		concreteState: &concreteState{
			state: state,
		},
	}
}

func (c *ConcreteStateRun) handle(context *Context) {
	fmt.Println("*** パソコンは、動作中です")
}

// ConcreteStateShutDown is struct
type ConcreteStateShutDown struct {
	*concreteState
}

// NewConcreteStateShutDown func for initializing ConcreteStateShutDown
func NewConcreteStateShutDown(state string) *ConcreteStateShutDown {
	return &ConcreteStateShutDown{
		concreteState: &concreteState{
			state: state,
		},
	}
}

func (c *ConcreteStateShutDown) handle(context *Context) {
	fmt.Println("*** パソコンは、停止しています")
}

// ConcreteStateRestart is struct
type ConcreteStateRestart struct {
	*concreteState
}

// NewConcreteStateRestart func for initializing ConcreteStateRestart
func NewConcreteStateRestart(state string) *ConcreteStateRestart {
	return &ConcreteStateRestart{
		concreteState: &concreteState{
			state: state,
		},
	}
}

func (c *ConcreteStateRestart) handle(context *Context) {
	fmt.Println("*** パソコンは、再起動をはじめます")
	context.SetState(NewConcreteStateBooting("booting"))
	context.Handle()
}
