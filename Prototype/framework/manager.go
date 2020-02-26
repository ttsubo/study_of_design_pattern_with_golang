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

// Register func for registering proto
func (m *Manager) Register(name string, proto Prototype) {
	if m.showcase == nil {
		m.showcase = make(map[string]Prototype)
	}
	m.showcase[name] = proto
}

// Create func for creating proto
func (m *Manager) Create(protoname string) Prototype {
	p := m.showcase[protoname]
	return p.createClone()
}
