package types

import (
	"errors"
	"fmt"
	"time"
)

type License struct {
	// 元数据
	PubKey      string `json:"pubkey,omitempty"`
	BlockHash   string `json:"block_hash,omitempty"`
	BlockHeight int64  `json:"block_height,omitempty"`
	Signature   string `json:"signature,omitempty" binding:"required"`
	ID          string `json:"id,omitempty" binding:"required"`

	// 数据类型
	Type    string            `json:"type,omitempty" binding:"required"`
	Params  map[string]string `json:"params,omitempty" binding:"required"`
	Desc    string            `json:"desc,omitempty"`
	Created string            `json:"created,omitempty"`
	Extra   string            `json:"extra,omitempty"`
}

func (a *License) Dumps() []byte {
	//struct -- > json
	js, _ := json.Marshal(a)
	//json -- > map
	var re map[string]interface{}
	json.Unmarshal(js, &re)
	//map --> json
	js, _ = json.Marshal(re)
	return js
}
func (li *License) validate_params() error {
	if li == nil {
		return errors.New("license is nil")
	}

	if li.Signature == "" {
		return errors.New("license`s signature is empty")
	}
	if li.ID == "" {
		return errors.New("license`s id is empty")
	}
	if li.Type == "" {
		return errors.New("license`s type is empty")
	}
	if li.Created == "" {
		li.Created = fmt.Sprintf("%d", time.Now().Unix())
	}
	return nil
}
