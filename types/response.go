package types

type BaseResponse struct {
	Code string `json:"code,omitempty"`
	Err  error  `json:"err,omitempty"`
	Msg  string `json:"msg,omitempty"`
}

type MetadataPostResp struct {
	Code string `json:"code,omitempty"`
	Msg  string `json:"msg,omitempty"`
	Data struct {
		Dna string `json:"dna,omitempty" binding:"required"`
	} `json:"data,omitempty" binding:"required"`
}

type MetadataGetResp struct {
	Code string      `json:"code,omitempty"`
	Msg  string      `json:"msg,omitempty"`
	Data Metadata    `json:"data,omitempty"`
	Tx   Transaction `json:"tx,omitempty"`
}
type LicenseGetResp struct {
	Code string  `json:"code,omitempty"`
	Msg  string  `json:"msg,omitempty"`
	Data map[string]interface{} `json:"data,omitempty"`
}

type LicensePostResp struct {
	Code string `json:"code,omitempty"`
	Msg  string `json:"msg,omitempty"`
}
