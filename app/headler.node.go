package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	kts "universe-go-sdk/types"
)

//QueryMetadata : query metadata by metadata`s dna on yuanben chain
//return:MetadataQueryResp.Code == "error" representative query failure
func QueryMetadata(url string, version string, dna string) (res *kts.MetadataQueryResp) {
	if version == "" {
		version = "v1"
	}
	if resp, err := http.Get(fmt.Sprintf("%s/%s/metadata/%s", url, version, dna)); err != nil {
		res = &kts.MetadataQueryResp{
			Code: "error",
			Msg:  err.Error(),
		}
	} else {
		d, _ := ioutil.ReadAll(resp.Body)
		res = &kts.MetadataQueryResp{}
		json.Unmarshal(d, res)
	}
	return
}

//SaveMetadata : register metadata on yuanben chain node
// return:MetadataSaveResp.Code == "error" representative save failure
func SaveMetadata(url string, version string, async string, md *kts.Metadata) (res *kts.MetadataSaveResp) {
	if md == nil {
		return &kts.MetadataSaveResp{
			Code: "error",
			Msg:  "metadata is nil",
		}
	}

	if md.Signature == "" {
		return &kts.MetadataSaveResp{
			Code: "error",
			Msg:  "metadata signature is empty",
		}
	}
	if md.License.Type == "" || md.License.Parameters == nil {
		return &kts.MetadataSaveResp{
			Code: "error",
			Msg:  "metadata license is empty",
		}
	}

	if version == "" {
		version = "v1"
	}
	_d, _ := json.Marshal(md)
	if resp, err := http.Post(fmt.Sprintf("%s/%s/metadata?async=%s", url, version, string(async)), "application/json", bytes.NewBuffer(_d)); err != nil {
		res = &kts.MetadataSaveResp{
			Code: "error",
			Msg:  err.Error(),
		}
	} else {
		d, _ := ioutil.ReadAll(resp.Body)

		res = &kts.MetadataSaveResp{}
		json.Unmarshal(d, res)
	}
	return

}

//QueryLicense : query the license by license.Type on yuanben chain
//return:LicenseQueryResp.Code == "error" representative query failure
func QueryLicense(url string, version string, license_type string) (res *kts.LicenseQueryResp) {
	if version == "" {
		version = "v1"
	}
	if resp, err := http.Get(fmt.Sprintf("%s/%s/license/%s", url, version, license_type)); err != nil {
		res = &kts.LicenseQueryResp{
			Code: "error",
			Msg:  err.Error(),
		}
	} else {
		d, _ := ioutil.ReadAll(resp.Body)

		res = &kts.LicenseQueryResp{}
		json.Unmarshal(d, res)
	}
	return
}

//QueryLatestBlockHash : query the lasted block hash on the yuaben chain
//return:BlockHashQueryResp.Code == "error" representative query failure
func QueryLatestBlockHash(url string, version string) (res *kts.BlockHashQueryResp) {
	if version == "" {
		version = "v1"
	}
	if resp, err := http.Get(fmt.Sprintf("%s/%s/block_hash", url, version)); err != nil {
		res = &kts.BlockHashQueryResp{
			Code: "error",
			Msg:  err.Error(),
		}
	} else {
		d, _ := ioutil.ReadAll(resp.Body)

		res = &kts.BlockHashQueryResp{}
		json.Unmarshal(d, res)
	}
	return
}

//CheckBlockHash : check whether blcokHash is on the yuanben chain
//return:BlockHashCheckResp.Code == "error" representative check failure
//check result:BlockHashCheckResp.Data
func CheckBlockHash(url string, version string, req *kts.BlockHashCheckReq) (res *kts.BlockHashCheckResp) {
	if version == "" {
		version = "v1"
	}
	if req == nil || req.Hash == "" {
		res = &kts.BlockHashCheckResp{
			Code: "error",
			Msg:  "request body is empty",
		}
		return
	}
	_d, _ := json.Marshal(req)
	if resp, err := http.Post(fmt.Sprintf("%s/%s/check_block_hash", url, version), "application/json", bytes.NewBuffer(_d)); err != nil {
		res = &kts.BlockHashCheckResp{
			Code: "error",
			Msg:  err.Error(),
		}
	} else {
		d, _ := ioutil.ReadAll(resp.Body)

		res = &kts.BlockHashCheckResp{}
		json.Unmarshal(d, res)
	}
	return
}
