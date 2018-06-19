package app_test

import (
	"encoding/json"
	"fmt"
	"github.com/yuanbenio/universe-go-sdk/app"
	kts "github.com/yuanbenio/universe-go-sdk/types"
	"testing"
)

const node_url = "https://testnet.yuanbenlian.com"

func TestQueryMetadata(t *testing.T) {
	res := app.QueryMetadata(node_url, "", "80PJRBEAYI2IEXXZ4LZ8XE91JLCBO7OS9N1AOALABLFHJCBHG")
	if res.Code == "error" {
		fmt.Println("metadata query eror :", res.Msg)
	} else {
		js, _ := json.Marshal(res.Data)

		fmt.Println(string(js))
	}

}

func TestSaveMetadata(t *testing.T) {
	md := &kts.Metadata{
		Content: "原本链是一个分布式的底层数据网络；" +
			"原本链是一个高效的，安全的，易用的，易扩展的，全球性质的，企业级的可信联盟链；" +
			"原本链通过智能合约系统以及数字加密算法，实现了链上数据可持续性交互以及数据传输的安全；" +
			"原本链通过高度抽象的“DTCP协议”与世界上独一无二的“原本DNA”互锁，确保链上数据100%不可篡改；" +
			"原本链通过优化设计后的共识机制和独创的“闪电DNA”算法，已将区块写入速度提高至毫秒级别",
		BlockHash: "4D36473D2FF1FE0772A6C0C55D7911295D8E1E27",
		Type:      "article",
		Title:     "原本链go版本sdk测试",
		License: struct {
			Type       string            `json:"type,omitempty" binding:"required"`
			Parameters map[string]string `json:"parameters,omitempty"`
		}{Type: "cc", Parameters: map[string]string{
			"y": "4",
			"b": "2",
		}},
	}
	pri_key := "50ced2bc6bc71ddfa517121b9df107400c9ba866344567da6aef82fac7824ade"
	app.FullMetadata(pri_key, md)

	res := app.SaveMetadata(node_url, "", "true", md)
	if res.Code == "error" {
		fmt.Println("metadata post error : ", res.Msg)
	} else {
		fmt.Println("success~ dna: ", res.Data.Dna)
	}
}

func TestQueryLicense(t *testing.T) {
	res := app.QueryLicense(node_url, "", "cc")
	if res.Code == "error" {
		fmt.Println("metadata post error : ", res.Msg)
	} else {
		js, _ := json.Marshal(res)
		fmt.Println("success~ license: ", string(js))
	}
}

func TestQueryLastedBlockHash(t *testing.T) {
	res := app.QueryLatestBlockHash(node_url, "")
	if res.Code == "error" {
		fmt.Println("query error : ", res.Msg)
	} else {
		js, _ := json.Marshal(res)
		fmt.Println("success~ blockhHash: ", string(js))
	}
}

func TestCheckBlockHash(t *testing.T) {
	req := &kts.BlockHashCheckReq{
		Hash:   "4A7FCE024C64061D28BEB91A3FC935465BE54B3B",
		Height: 22102,
	}
	res := app.CheckBlockHash(node_url, "", req)
	if res.Code == "error" {
		fmt.Println("check error : ", res.Msg)
	} else {
		fmt.Println("success~ check result is ", res.Data)
	}
}