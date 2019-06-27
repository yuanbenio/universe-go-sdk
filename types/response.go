package types

import "encoding/json"

const (
	APISuccessCode = "ok"
)

type BaseResp struct {
	Code string `json:"code,omitempty"`
	Msg  string `json:"msg,omitempty"`
}

func (resp *BaseResp) Success() bool {
	return resp.Code == APISuccessCode
}

type MetadataSaveResp struct {
	BaseResp
	Data struct {
		Dna string `json:"dna,omitempty" binding:"required"`
	} `json:"data,omitempty" binding:"required"`
}

type MetadataQueryResp struct {
	BaseResp
	Data Metadata    `json:"data,omitempty"`
	Tx   Transaction `json:"tx,omitempty"`
}

type LicenseQueryResp struct {
	BaseResp
	Data map[string]interface{} `json:"data,omitempty"`
	Tx   Transaction            `json:"tx,omitempty"`
}

type Transaction struct {
	BlockHash   string `json:"block_hash,omitempty"`
	BlockHeight int64  `json:"block_height,omitempty"`
	DataHeight  int64  `json:"data_height,omitempty"`
	Sender      string `json:"sender,omitempty"`
	Time        int64  `json:"time,omitempty"`
}

func (t *Transaction) Dumps() []byte {
	d, _ := json.Marshal(t)
	return d
}

type BlockHashQueryResp struct {
	BaseResp
	Data BlockHashResp `json:"data,omitempty"`
}

type BlockHashResp struct {
	LatestBlockHash   string `json:"latest_block_hash,omitempty"`
	LatestBlockHeight int64  `json:"latest_block_height,omitempty"`
	LatestBlockTime   string `json:"latest_block_time,omitempty"`
}

type BlockHashCheckResp struct {
	BaseResp
	Data bool `json:"data,omitempty"`
}

type BlockHashCheckReq struct {
	Hash   string `json:"hash,omitempty"`
	Height int64  `json:"height,omitempty"`
}

type RegisterAccountReq struct {
	Signature string   `json:"signature"`
	Pubkey    string   `json:"pubkey"`
	Subkeys   []string `json:"subkeys"`
}

type RegisterAccountResp struct {
	BaseResp
}
