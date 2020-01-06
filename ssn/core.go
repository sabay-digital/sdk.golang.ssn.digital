package ssn

// Core describes the JSON structure related to the output of the Stellar Core info endpoint
type Core struct {
	Info coreInfo `json:"info"`
}

type coreInfo struct {
	Build                string     `json:"build"`
	History_failure_rate string     `json:"history_failure_rate"`
	Ledger               coreLedger `json:"ledger"`
	Network              string     `json:"network"`
	Peers                corePeers  `json:"peers"`
	Protocol_version     int        `json:"protocol_version"`
	Quorum               coreQuorum `json:"quorum"`
	StartedOn            string     `json:"startedOn"`
	State                string     `json:"state"`
}

type coreLedger struct {
	Age          int    `json:"age"`
	BaseFee      int    `json:"baseFee"`
	BaseReserve  int    `json:"baseReserve"`
	CloseTime    int    `json:"closeTime"`
	Hash         string `json:"hash"`
	MaxTxSetSize int    `json:"maxTxSetSize"`
	Num          int    `json:"num"`
	Version      int    `json:"version"`
}

type corePeers struct {
	Authenticated_count int `json:"authenticated_count"`
	Pending_count       int `json:"pending_count"`
}

type coreQuorum struct {
	Node       string         `json:"node"`
	Qset       coreQset       `json:"qset"`
	Transitive coreTransitive `json:"transitive"`
}

type coreQset struct {
	Agree     int    `json:"agree"`
	Delayed   int    `json:"delayed"`
	Disagree  int    `json:"disagree"`
	Fail_at   int    `json:"fail_at"`
	Hash      string `json:"hash"`
	Ledger    int    `json:"ledger"`
	Missing   int    `json:"missing"`
	Phase     string `json:"phase"`
	Validated bool   `json:"validated"`
}

type coreTransitive struct {
	Critical          bool `json:"critical"`
	Intersection      bool `json:"intersection"`
	Last_check_ledger int  `json:"last_check_ledger"`
	Node_count        int  `json:"node_count"`
}
