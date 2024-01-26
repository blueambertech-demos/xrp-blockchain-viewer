package mock

import "context"

type Store struct{}

func (m *Store) GetData(_ string) ([]byte, bool) {
	return nil, false
}
func (m *Store) SetData(_ context.Context, _ string, _ []byte, _ int64) {}
func (m *Store) DeleteData(_ string)                                    {}
