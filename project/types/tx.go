package types

type Transaction struct {
	PubKey    string      `json:"pubkey,omitempty"`
	Signature string      `json:"sign,omitempty"`
	Key       string      `json:"key,omitempty"`
	Value     interface{} `json:"value,omitempty"`
	Path      string      `json:"path,omitempty"`
	Timestamp int64       `json:"time,omitempty"`
}

func (t *Transaction) Dumps() []byte {
	d, _ := json.Marshal(t)
	return d
}
