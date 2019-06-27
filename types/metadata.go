package types

import "encoding/json"

type Metadata struct {
	ContentHash string `json:"content_hash,omitempty" binding:"required"`
	Created     string `json:"created,omitempty" binding:"required"`
	BlockHash   string `json:"block_hash,omitempty" binding:"required"`
	BlockHeight string `json:"block_height,omitempty"  binding:"required"`
	Language    string `json:"language,omitempty"  binding:"required"`
	Signature   string `json:"signature,omitempty" binding:"required"`
	PubKey      string `json:"pubkey,omitempty" binding:"required"`
	Type        string `json:"type,omitempty" binding:"required"`
	License     struct {
		Params map[string]string `json:"parameters,omitempty"`
		Type   string            `json:"type,omitempty" binding:"required"`
	} `json:"license,omitempty" binding:"required"`

	ID        string      `json:"id,omitempty"`
	Abstract  string      `json:"abstract,omitempty"`
	Category  string      `json:"category,omitempty"`
	Content   string      `json:"content,omitempty"`
	Data      interface{} `json:"data,omitempty"`
	DNA       string      `json:"dna,omitempty"`
	ParentDna string      `json:"parent_dna,omitempty"`
	Extra     interface{} `json:"extra,omitempty"`
	Source    string      `json:"source,omitempty"`
	Title     string      `json:"title,omitempty"`
}

var (
	NoneLicense = License{Type: "none"}
)

type License struct {
	Params map[string]string `json:"parameters,omitempty"`
	Type   string            `json:"type,omitempty" binding:"required"`
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
	// remove signature attribute
	sign := a.Signature
	dna := a.DNA
	content := a.Content

	a.Signature = ""
	a.DNA = ""
	a.Content = ""
	// struct -- > json
	js, _ := json.Marshal(a)
	// json -- > map
	var re map[string]interface{}
	json.Unmarshal(js, &re)
	// map --> json
	js, _ = json.Marshal(re)

	a.Content = content
	a.DNA = dna
	a.Signature = sign
	return js
}

type MetadataType string

func (c MetadataType) Value() string {
	return string(c)
}

const (
	ARTICLE MetadataType = "article"
	IMAGE   MetadataType = "image"
	AUDIO   MetadataType = "audio"
	VIDEO   MetadataType = "video"
	PRIVATE MetadataType = "private"
	CUSTOM  MetadataType = "custom"
)
