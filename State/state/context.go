package state

// Context is struct
type Context struct {
	state State
}

// NewContext func for initializing Context
func NewContext(stateObj State) *Context {
	return &Context{
		state: stateObj,
	}
}

// SetState func for change state
func (c *Context) SetState(obj State) {
	c.state = obj
}

// Handle func for handling state
func (c *Context) Handle() {
	c.state.handle(c)
}

// GetState func for fetching state
func (c *Context) GetState() string {
	return c.state.getConcreateState()
}
