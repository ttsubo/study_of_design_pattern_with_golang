package memento

// Memento is struct
type Memento struct {
	money  int
	fruits []string
}

// GetMoney func for fetching money in Memento
func (m *Memento) GetMoney() int {
	return m.money
}

func (m *Memento) addFruit(fruit string) {
	m.fruits = append(m.fruits, fruit)
}

// GetFruits func for fetching current fruits list in Memento
func (m *Memento) GetFruits() []string {
	return m.fruits
}
