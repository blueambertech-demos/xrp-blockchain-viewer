package ledger

type httpRequest struct {
	Method string        `json:"method"`
	Params []interface{} `json:"params"`
}

type ledgerInfoParam struct {
	LedgerHash   string `json:"ledger_hash,omitempty"`
	LedgerIndex  string `json:"ledger_index,omitempty"`
	Transactions bool   `json:"transactions,omitempty"`
	Expand       bool   `json:"expand,omitempty"`
	OwnerFunds   bool   `json:"owner_funds,omitempty"`
	Binary       bool   `json:"binary,omitempty"`
	Queue        bool   `json:"queue,omitempty"`
}

type LedgerInfoResponse struct {
	Result struct {
		Ledger struct {
			AccountHash         string `json:"account_hash"`
			CloseFlags          int    `json:"close_flags"`
			CloseTime           int    `json:"close_time"`
			CloseTimeHuman      string `json:"close_time_human"`
			CloseTimeResolution int    `json:"close_time_resolution"`
			Closed              bool   `json:"closed"`
			LedgerHash          string `json:"ledger_hash"`
			LedgerIndex         string `json:"ledger_index"`
			ParentCloseTime     int    `json:"parent_close_time"`
			ParentHash          string `json:"parent_hash"`
			TotalCoins          string `json:"total_coins"`
			TransactionHash     string `json:"transaction_hash"`
		} `json:"ledger"`
		LedgerHash  string `json:"ledger_hash"`
		LedgerIndex int    `json:"ledger_index"`
		Status      string `json:"status"`
		Validated   bool   `json:"validated"`
	} `json:"result"`
}
