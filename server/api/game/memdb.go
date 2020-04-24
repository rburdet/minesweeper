package game

type MemDbImpl struct {
	data map[string]Game
}

func NewMemDbImpl() *MemDbImpl {
	return &MemDbImpl{make(map[string]Game)}
}

func (m *MemDbImpl) SaveGame(name string, b Game) {
	m.data[name] = b
}

func (m *MemDbImpl) GetGame(name string) Game {
	return m.data[name]
}

func (m *MemDbImpl) DeleteGame(name string) {
	delete(m.data, name)
}

func (m *MemDbImpl) Exists(name string) bool {
	_, ok := m.data[name]
	return ok
}
