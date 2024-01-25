package ledger

import (
	"encoding/json"

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
)

var storage MemoryStore

func SetMemoryStore(store MemoryStore) {
	storage = store
}

func Info(serverAddr string, idxType LedgerIndexType) (*LedgerInfoResponse, error) {
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
		storage.SetData(key, b)
	}
	return buildResponse(b)
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

func buildResponse(b []byte) (*LedgerInfoResponse, error) {
	var response LedgerInfoResponse
	err := json.Unmarshal(b, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
