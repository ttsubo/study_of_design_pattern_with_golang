package command

import (
	"fmt"
	"os"
)

type command interface {
	execute()
	display()
}

// FileOperator is struct
type FileOperator struct {
}

// NewFileOperator func for initializing FileOperator
func NewFileOperator() *FileOperator {
	return &FileOperator{}
}

func (f *FileOperator) createFile(filename string) {
	os.Create(filename)
}

func (f *FileOperator) changeFileMode(filename string, permission uint64) {
	os.Chmod(filename, os.FileMode(permission))
}

// CompositeCommand is struct
type CompositeCommand struct {
	cmds []command
}

//NewCompositeCommand func for initializing CompositeCommand
func NewCompositeCommand() *CompositeCommand {
	return &CompositeCommand{}
}

// AppendCmd func for appending command
func (c *CompositeCommand) AppendCmd(cmd command) {
	c.cmds = append(c.cmds, cmd)
}

// Execute func for executing command
func (c *CompositeCommand) Execute() {
	for _, cmd := range c.cmds {
		cmd.execute()
	}
}

// Display func for executing command
func (c *CompositeCommand) Display() {
	for _, cmd := range c.cmds {
		cmd.display()
	}
}

// FileTouchCommand is struct
type FileTouchCommand struct {
	filename string
	receiver *FileOperator
}

// NewFileTouchCommand func for initializing FileTouchCommand
func NewFileTouchCommand(filename string, receiverObj *FileOperator) *FileTouchCommand {
	return &FileTouchCommand{
		filename: filename,
		receiver: receiverObj,
	}
}

func (f *FileTouchCommand) execute() {
	f.receiver.createFile(f.filename)
}

func (f *FileTouchCommand) display() {
	fmt.Printf("%% touch %s\n", f.filename)
}

// ChmodCommand is struct
type ChmodCommand struct {
	filename   string
	permission uint64
	receiver   *FileOperator
}

// NewChmodCommand func for initializing ChmodCommand
func NewChmodCommand(filename string, permission uint64, receiverObj *FileOperator) *ChmodCommand {
	return &ChmodCommand{
		filename:   filename,
		permission: permission,
		receiver:   receiverObj,
	}
}

func (c *ChmodCommand) execute() {
	c.receiver.changeFileMode(c.filename, c.permission)
}

func (c *ChmodCommand) display() {
	fmt.Printf("%% chmod %o %s\n", c.permission, c.filename)
}
