package prototype

// ManagerInterface is interface
type ManagerInterface interface {
	Register(name string, producter Prototype)
	Create(protoname string) Prototype
}

// Manager is struct
type Manager struct {
	showcase map[string]Prototype
}

// NewManager func for initializing Manager
func NewManager() *Manager {
	return &Manager{
		showcase: make(map[string]Prototype),
	}
}

// Register func for registering proto
func (m *Manager) Register(name string, proto Prototype) {
	m.showcase[name] = proto
}

// Create func for creating proto
func (m *Manager) Create(protoname string) Prototype {
	p := m.showcase[protoname]
	return p.createClone()
}
