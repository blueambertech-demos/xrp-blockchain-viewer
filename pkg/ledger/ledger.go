package ledger

import (
	"context"
	"time"

	"github.com/blueambertech-demos/xrp-blockchain-viewer/pkg/httputil"
)

type LedgerIndexType string

const (
	httpMethod = "ledger"
)

const (
	Validated LedgerIndexType = "validated"
	Closed    LedgerIndexType = "closed"
	Current   LedgerIndexType = "current"

	cacheExpirySecs = 5
)

var (
	storage         MemoryStore
	cancellationCtx context.Context
)

// SetMemoryStore sets the memory store to use to cache data retrieved from the ledger along with a context that the store
// can use to monitor cancellation requests
func SetMemoryStore(ctx context.Context, store MemoryStore) {
	storage = store
	cancellationCtx = ctx
}

// Info retrieves ledger information from a specified ledger type (either Validated, Closed or Current)
// at the specified address. This func will attempt to use cached data in the store, if cached data does not
// exist it will be retrieved from the XRP API and the cache will be updated.
func Info(serverAddr string, idxType LedgerIndexType) ([]byte, error) {
	key := "ledger.info." + string(idxType)
	b, ok := storage.GetData(key)
	if !ok {
		var err error
		b, err = info(serverAddr, []interface{}{
			ledgerInfoParam{
				LedgerIndex: string(idxType),
			},
		})
		if err != nil {
			return nil, err
		}
		storage.SetData(cancellationCtx, key, b, time.Now().Unix()+cacheExpirySecs)
	}
	return b, nil
}

func info(serverAddr string, params []interface{}) ([]byte, error) {
	req := httpRequest{
		Method: httpMethod,
		Params: params,
	}

	j, err := httputil.PostJson(serverAddr, req)
	if err != nil {
		return nil, err
	}
	return j, nil
}
