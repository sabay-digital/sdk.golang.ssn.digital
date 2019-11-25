package ssn

import "encoding/json"

// Transaction describes the JSON structure related to an individual transaction
type Transaction struct {
	Id                      string                     `json:"id,omitempty"`
	Paging_token            string                     `json:"paging_token,omitempty"`
	Successful              bool                       `json:"successful,omitempty"`
	Hash                    string                     `json:"hash,omitempty"`
	Ledger                  int                        `json:"ledger,omitempty"`
	Created_at              string                     `json:"created_at,omitempty"`
	Source_account          string                     `json:"source_account,omitempty"`
	Source_account_sequence string                     `json:"source_account_sequence,omitempty"`
	Max_fee                 int                        `json:"max_fee,omitempty"`
	Fee_charged             int                        `json:"fee_charged,omitempty"`
	Operation_count         int                        `json:"operation_count,omitempty"`
	Envelope_xdr            string                     `json:"envelope_xdr,omitempty"`
	Result_xdr              string                     `json:"result_xdr,omitempty"`
	Result_meta_xdr         string                     `json:"result_meta_xdr,omitempty"`
	Fee_meta_xdr            string                     `json:"fee_meta_xdr,omitempty"`
	Memo_type               string                     `json:"memo_type,omitempty"`
	Signatures              []string                   `json:"signatures,omitempty"`
	Type                    string                     `json:"type,omitempty"`
	Title                   string                     `json:"title,omitempty"`
	Status                  int                        `json:"status,omitempty"`
	Detail                  string                     `json:"detail,omitempty"`
	Extras                  map[string]json.RawMessage `json:"extras,omitempty"`
}
