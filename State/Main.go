package main

import (
	"flag"
	"fmt"
	"time"

	"./state"
)

func setConcreteState(operation string) state.State {
	var stateObj state.State

	if operation == "start" {
		stateObj = state.NewConcreteStateBooting("booting")
	} else if operation == "stop" {
		stateObj = state.NewConcreteStateShutDown("shutdown")
	} else if operation == "restart" {
		stateObj = state.NewConcreteStateRestart("restart")
	}
	return stateObj
}

func startMain(initialOperation, changeOperation string) {
	obj := state.NewContext(setConcreteState(initialOperation))
	fmt.Printf("### パソコンを、[%s]します\n", initialOperation)
	obj.Handle()
	fmt.Printf("### パソコンは、[%s]の動作状態になりました\n", obj.GetState())
	fmt.Println("")

	fmt.Println("... sleep 5 second")
	fmt.Println("")
	time.Sleep(time.Second * 5)
	obj.SetState(setConcreteState(changeOperation))
	fmt.Printf("### パソコンを、[%s]します\n", changeOperation)
	obj.Handle()
	fmt.Printf("### パソコンの動作状態は、[%s]になりました\n", obj.GetState())
}

func main() {
	flag.Parse()
	startMain(flag.Arg(0), flag.Arg(1))
}
