package types

type Metadata struct {
	PubKey      string `json:"pubkey,omitempty"`
	BlockHash   string `json:"block_hash,omitempty"`
	BlockHeight int64  `json:"block_height,omitempty"`
	Signature   string `json:"signature,omitempty"`
	ID          string `json:"id,omitempty"`

	// 用逗号隔开
	Category    string `json:"category,omitempty"`
	ContentHash string `json:"content_hash,omitempty"`
	Type        string `json:"type,omitempty" binding:"required"`
	Title       string `json:"title,omitempty" binding:"required"`
	Content     string `json:"content,omitempty" binding:"required"`

	// 时间戳
	Created   string      `json:"created,omitempty"`
	Abstract  string      `json:"abstract,omitempty"`
	DNA       string      `json:"dna,omitempty"`
	ParentDna string      `json:"parent_dna,omitempty"`
	Language  string      `json:"language,omitempty"`
	Source    string      `json:"source,omitempty"`
	Extra     interface{} `json:"extra,omitempty"`
	Data      interface{} `json:"data,omitempty"`
	License   struct {
		Type       string            `json:"type,omitempty" binding:"required"`
		Parameters map[string]string `json:"parameters,omitempty" binding:"required"`
	} `json:"license,omitempty" binding:"required"`
}

func (a *Metadata) Dumps() []byte {
	d, _ := json.Marshal(a)
	return d
}

func (a *Metadata) DumpsLicense() []byte {
	d, _ := json.Marshal(a.License)
	return d
}

func (a *Metadata) DumpsRmSignSort() []byte {
	//remove signature attribute
	sign := a.Signature
	dna := a.DNA
	a.Signature = ""
	a.DNA = ""
	a.Content = ""
	//struct -- > json
	js, _ := json.Marshal(a)
	//json -- > map
	var re map[string]interface{}
	json.Unmarshal(js, &re)
	//map --> json
	js, _ = json.Marshal(re)

	a.DNA = dna
	a.Signature = sign
	return js
}
