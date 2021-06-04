package ssn

import "encoding/json"

// Ledger describes the JSON structure related to an individual ledger
type Ledger struct {
	Id                           string                     `json:"id,omitempty"`
	Paging_token                 string                     `json:"paging_token,omitempty"`
	Hash                         string                     `json:"hash,omitempty"`
	Prev_hash                    string                     `json:"prev_hash,omitempty"`
	Sequence                     int                        `json:"sequence,omitempty"`
	Successful_transaction_count int                        `json:"successful_transaction_count"`
	Failed_transaction_count     int                        `json:"failed_transaction_count"`
	Operation_count              int                        `json:"operation_count"`
	Tx_set_operation_count       int                        `json:"tx_set_operation_count"`
	Closed_at                    string                     `json:"closed_at,omitempty"`
	Total_coins                  string                     `json:"total_coins,omitempty"`
	Fee_pool                     string                     `json:"fee_pool,omitempty"`
	Base_fee_in_stroops          int                        `json:"base_fee_in_stroops,omitempty"`
	Base_reserve_in_stroops      int                        `json:"base_reserve_in_stroops,omitempty"`
	Max_tx_set_size              int                        `json:"max_tx_set_size,omitempty"`
	Protocol_version             int                        `json:"protocol_version,omitempty"`
	Header_xdr                   string                     `json:"header_xdr,omitempty"`
	Type                         string                     `json:"type,omitempty"`
	Title                        string                     `json:"title,omitempty"`
	Status                       int                        `json:"status,omitempty"`
	Detail                       string                     `json:"detail,omitempty"`
	Extras                       map[string]json.RawMessage `json:"extras,omitempty"`
}
