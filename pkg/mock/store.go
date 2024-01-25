package mock

type Store struct{}

func (m *Store) GetData(key string) ([]byte, bool) {
	return nil, false
}
func (m *Store) SetData(key string, data []byte) {}
