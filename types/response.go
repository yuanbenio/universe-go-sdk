package types

import "encoding/json"

type MetadataSaveResp struct {
	Code string `json:"code,omitempty"`
	Msg  string `json:"msg,omitempty"`
	Data struct {
		Dna string `json:"dna,omitempty" binding:"required"`
	} `json:"data,omitempty" binding:"required"`
}

type MetadataQueryResp struct {
	Code string      `json:"code,omitempty"`
	Msg  string      `json:"msg,omitempty"`
	Data Metadata    `json:"data,omitempty"`
	Tx   Transaction `json:"tx,omitempty"`
}
type LicenseQueryResp struct {
	Code string                 `json:"code,omitempty"`
	Msg  string                 `json:"msg,omitempty"`
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
	Code string        `json:"code,omitempty"`
	Msg  string        `json:"msg,omitempty"`
	Data BlockHashResp `json:"data,omitempty"`
}
type BlockHashResp struct {
	LatestBlockHash   string `json:"latest_block_hash,omitempty"`
	LatestBlockHeight int64  `json:"latest_block_height,omitempty"`
	LatestBlockTime   string `json:"latest_block_time,omitempty"`
}

type BlockHashCheckResp struct {
	Code string `json:"code,omitempty"`
	Msg  string `json:"msg,omitempty"`
	Data bool   `json:"data,omitempty"`
}

type BlockHashCheckReq struct {
	Hash   string `json:"hash,omitempty"`
	Height int64  `json:"height,omitempty"`
}
