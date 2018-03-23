package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	kts "universe-go-sdk/types"
)

func Metadata_get(url string, version string, dna string) (res *kts.MetadataGetResp) {
	if version == "" {
		version = "v1"
	}
	if resp, err := http.Get(fmt.Sprintf("%s/%s/metadata/%s", url, version, dna)); err != nil {
		res = &kts.MetadataGetResp{
			Code: "error",
			Msg:  err.Error(),
		}
	} else {
		d, _ := ioutil.ReadAll(resp.Body)

		res = &kts.MetadataGetResp{}
		json.Unmarshal(d, res)
	}
	return
}

//metadata_post
// return:MetadataPostResp
func Metadata_post(url string, version string, md *kts.Metadata) (res *kts.MetadataPostResp) {
	if md == nil {
		return &kts.MetadataPostResp{
			Code: "error",
			Msg:  "metadata is nil",
		}
	}

	if md.Signature == "" {
		return &kts.MetadataPostResp{
			Code: "error",
			Msg:  "metadata signature is empty",
		}
	}
	if md.License.Type == "" || md.License.Parameters == nil {
		return &kts.MetadataPostResp{
			Code: "error",
			Msg:  "metadata license is empty",
		}
	}

	if version == "" {
		version = "v1"
	}
	_d, _ := json.Marshal(md)
	if resp, err := http.Post(fmt.Sprintf("%s/%s/metadata", url, version), "application/json", bytes.NewBuffer(_d)); err != nil {
		res = &kts.MetadataPostResp{
			Code: "error",
			Msg:  err.Error(),
		}
	} else {
		d, _ := ioutil.ReadAll(resp.Body)

		res = &kts.MetadataPostResp{}
		json.Unmarshal(d, res)
	}
	return

}

//license_get
// return:LicenseGetResp
func License_get(url string, version string, license_type string) (res *kts.LicenseGetResp) {
	if version == "" {
		version = "v1"
	}
	if resp, err := http.Get(fmt.Sprintf("%s/%s/license/%s", url, version, license_type)); err != nil {
		res = &kts.LicenseGetResp{
			Code: "error",
			Msg:  err.Error(),
		}
	} else {
		d, _ := ioutil.ReadAll(resp.Body)

		res = &kts.LicenseGetResp{}
		json.Unmarshal(d, res)
	}
	return
}