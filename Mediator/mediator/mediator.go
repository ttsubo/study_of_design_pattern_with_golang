package mediator

import "fmt"

// Mediator is interface
type Mediator interface {
	onChange(component *colleague)
	SetColleagues(inputIDObj, inputPwObj *ConcreteColleagueTextArea, buttonObj *ConcreteColleagueButton)
	getAuthentication() bool
}

// ConcreteMediator is struct
type ConcreteMediator struct {
	authentication         bool
	inputIDObj, inputPwObj *ConcreteColleagueTextArea
	buttonObj              *ConcreteColleagueButton
}

// NewConcreteMediator func for initializing ConcreteMediator
func NewConcreteMediator() Mediator {
	return &ConcreteMediator{
		authentication: false,
	}
}

// SetColleagues func for setting Objects
func (c *ConcreteMediator) SetColleagues(inputIDObj, inputPwObj *ConcreteColleagueTextArea, buttonObj *ConcreteColleagueButton) {
	c.inputIDObj = inputIDObj
	c.inputPwObj = inputPwObj
	c.buttonObj = buttonObj
}

func (c *ConcreteMediator) onChange(component *colleague) {
	if component.name == "ID" || component.name == "PW" {
		c.refreshButton()
	} else if component.name == "Login" {
		c.tryAuthentication()
	}
}

func (c *ConcreteMediator) refreshButton() {
	if c.inputIDObj.text != "" && c.inputPwObj.text != "" {
		fmt.Println("(Active login button)")
		c.buttonObj.active = true
	}
}

func (c *ConcreteMediator) tryAuthentication() {
	if c.inputIDObj.text == "hoge" && c.inputPwObj.text == "fuga" {
		fmt.Println("(ID/PW is confirmed)")
		c.authentication = true
	} else {
		fmt.Println("(ID/PW is incorrect)")
	}
}

func (c *ConcreteMediator) getAuthentication() bool {
	return c.authentication
}
