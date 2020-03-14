package mediator

type colleague struct {
	mediator Mediator
	name     string
}

func (c *colleague) onChange() {
	if c.mediator != nil {
		c.mediator.onChange(c)
	}
}

// ConcreteColleagueButton is struct
type ConcreteColleagueButton struct {
	*colleague
	active bool
}

// NewConcreteColleagueButton func for initializing ConcreteColleagueButton
func NewConcreteColleagueButton(mediatorObj Mediator, name string) *ConcreteColleagueButton {
	return &ConcreteColleagueButton{
		colleague: &colleague{
			mediator: mediatorObj,
			name:     name},
		active: false,
	}
}

// ClickButton func for detecting whether button is active or not
func (c *ConcreteColleagueButton) ClickButton() bool {
	if c.active {
		c.onChange()
	}
	return c.mediator.getAuthentication()
}

// CheckButtonStatus func for detecting whether button is active or not
func (c *ConcreteColleagueButton) CheckButtonStatus() bool {
	return c.active
}

// ConcreteColleagueTextArea is struct
type ConcreteColleagueTextArea struct {
	*colleague
	text string
}

// NewConcreteColleagueTextArea func for initializing ConcreteColleagueTextArea
func NewConcreteColleagueTextArea(mediatorObj Mediator, name string) *ConcreteColleagueTextArea {
	return &ConcreteColleagueTextArea{
		colleague: &colleague{
			mediator: mediatorObj,
			name:     name},
		text: "",
	}
}

// InputText func for putting text
func (c *ConcreteColleagueTextArea) InputText(text string) {
	c.text = text
	c.onChange()
}
