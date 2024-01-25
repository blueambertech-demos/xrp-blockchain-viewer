package ledger

type MemoryStore interface {
	SetData(key string, data []byte)
	GetData(key string) ([]byte, bool)
}
