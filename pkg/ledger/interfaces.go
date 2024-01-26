package ledger

import "context"

type MemoryStore interface {
	SetData(ctx context.Context, key string, data []byte, expiry int64)
	GetData(key string) ([]byte, bool)
	DeleteData(key string)
}
