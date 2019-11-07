package app

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/yuanbenio/universe-go-sdk/utils/base36"
	"io"
	"io/ioutil"
	"net/http"

	kts "github.com/yuanbenio/universe-go-sdk/types"
)

var (
	ErrMDNil      = errors.New("metadata is nil")
	ErrSignNil    = errors.New("signature is nil")
	ErrLicenseNil = errors.New("license is nil")
	ErrParameters = errors.New("illegal parameters ")
	ErrDNAInvalid = errors.New("invalid DNA")
)

const (
	DefaultChainVersion = "v1"
	ContentType         = "application/json; charset=utf-8"
)

const (
	URIMetadata         = "metadata"
	URIQueryLicense     = "queryLicense"
	URIQueryLatestBlock = "block_hash"
	URICheckBlock       = "check_block_hash"
	URIRegisterAccount  = "accounts"
)

type NodeProcessor struct {
	url          string
	chainVersion string
}

//init
//url : Yuanben Chain node URL
//chainVersion : Yuanben Chain API version, default:v1
func InitNodeProcessor(url, chainVersion string) *NodeProcessor {
	if chainVersion == "" {
		chainVersion = DefaultChainVersion
	}
	return &NodeProcessor{
		url:          url,
		chainVersion: chainVersion,
	}
}

//QueryMetadata : query metadata by metadata`s dna from Yuanben Chain
func (processor *NodeProcessor) QueryMetadata(dna string) (res *kts.MetadataQueryResp, err error) {
	if dna == "" || len(base36.DecodeToBytes(dna)) != 32 {
		return nil, ErrDNAInvalid
	}
	var response *http.Response
	response, err = Request(http.MethodGet, processor.getURIWithDNA(URIMetadata, dna), nil)
	if err != nil {
		return nil, err
	}
	res = &kts.MetadataQueryResp{}
	err = ReadResponse(response, res)
	return
}

//SaveMetadata : save metadata to Yuanben Chain
func (processor *NodeProcessor) SaveMetadata(md *kts.Metadata) (res *kts.MetadataSaveResp, err error) {
	if md == nil {
		return nil, ErrMDNil
	}

	if md.Signature == "" {
		return nil, ErrSignNil
	}
	if md.License.Type == "" || (md.License.Type != kts.NoneLicense.Type && md.License.Params == nil) {
		return nil, ErrLicenseNil
	}

	_d, err := json.Marshal(md)
	if err != nil {
		return nil, err
	}

	var response *http.Response
	response, err = Request(http.MethodPost, processor.getURI(URIMetadata), _d)
	if err != nil {
		return nil, err
	}
	res = &kts.MetadataSaveResp{}
	err = ReadResponse(response, res)
	return
}

//QueryLicense : query the license from Yuanben Chain
func (processor *NodeProcessor) QueryLicense(licenseType, licenseVersion string) (res *kts.LicenseQueryResp, err error) {
	if licenseType == "" {
		return nil, ErrLicenseNil
	}
	params := map[string]string{
		"type":    licenseType,
		"version": licenseVersion,
	}
	_d, e := json.Marshal(params)
	if e != nil {
		return nil, err
	}

	var response *http.Response
	response, err = Request(http.MethodPost, processor.getURI(URIQueryLicense), _d)
	if err != nil {
		return nil, err
	}
	res = &kts.LicenseQueryResp{}
	err = ReadResponse(response, res)
	return
}

//QueryLatestBlockHash : query the lasted block information from the Yuaben Chain node
func (processor *NodeProcessor) QueryLatestBlockHash() (res *kts.BlockHashQueryResp, err error) {
	var response *http.Response
	response, err = Request(http.MethodGet, processor.getURI(URIQueryLatestBlock), nil)
	if err != nil {
		return nil, err
	}
	res = &kts.BlockHashQueryResp{}
	err = ReadResponse(response, res)
	return
}

//CheckBlockHash : check whether block is on the Yuanben Chain
//check result:BlockHashCheckResp.Data
func (processor *NodeProcessor) CheckBlockHash(req *kts.BlockHashCheckReq) (res *kts.BlockHashCheckResp, err error) {
	if req == nil || req.Hash == "" {
		return nil, ErrParameters
	}
	_d, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	var response *http.Response
	response, err = Request(http.MethodPost, processor.getURI(URICheckBlock), _d)
	if err != nil {
		return nil, err
	}
	res = &kts.BlockHashCheckResp{}
	err = ReadResponse(response, res)
	return
}

//RegisterAccount: register sub accounts to special YuanBen chain node
//For nodes that have authentication enabled, if you need to store data, you must registered
//param req : see dtcp.go ---> GenRegisterAccountReq
func (processor *NodeProcessor) RegisterAccount(req *kts.RegisterAccountReq) (res *kts.RegisterAccountResp, err error) {
	if req == nil || req.Subkeys == nil {
		return nil, ErrParameters
	}
	_d, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	var response *http.Response
	response, err = Request(http.MethodPost, processor.getURI(URIRegisterAccount), _d)
	if err != nil {
		return nil, err
	}
	res = &kts.RegisterAccountResp{}
	err = ReadResponse(response, res)
	return
}

func Request(method, url string, params []byte) (*http.Response, error) {
	var body io.Reader
	if params != nil {
		body = bytes.NewBuffer(params)
	}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", ContentType)
	return http.DefaultClient.Do(req)
}

//read response
func ReadResponse(response *http.Response, result interface{}) error {
	_d, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	return json.Unmarshal(_d, result)
}

func (processor *NodeProcessor) getURI(goroutine string) string {
	return fmt.Sprintf("%s/%s/%s", processor.url, processor.chainVersion, goroutine)
}

func (processor *NodeProcessor) getURIWithDNA(goroutine string, dna string) string {
	return fmt.Sprintf("%s/%s/%s/%s", processor.url, processor.chainVersion, goroutine, dna)
}
