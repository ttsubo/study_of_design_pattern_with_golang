package main

import (
	"flag"
	"strconv"

	"./command"
)

func startMain(filename string, permission uint64) {
	recv := command.NewFileOperator()
	cc := command.NewCompositeCommand()
	cc.AppendCmd(command.NewFileTouchCommand(filename, recv))
	cc.AppendCmd(command.NewChmodCommand(filename, permission, recv))
	cc.Execute()
	cc.Display()
}

func main() {
	flag.Parse()
	perm32, _ := strconv.ParseUint(flag.Arg(1), 8, 32)
	startMain(flag.Arg(0), perm32)
}
