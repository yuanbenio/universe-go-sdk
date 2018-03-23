package main

import (
	"fmt"
	"universe-go-sdk/app"
	kts "universe-go-sdk/types"
	//"encoding/json"
	"encoding/json"
)

var (
	node_url = "http://119.23.22.129:9000"
)

//test result
//{"pubkey":"02e70aebdffd7b8f7495f35b504f0f3053024e44a87b8c419f7d886659e1475e19","block_hash":"4D36473D2FF1FE0772A6C0C55D7911295D8E1E27","signature":"b6600824bfd45338acd54fdfc9740a2fb2e4ad6489482339e8cff079b2fcf5004e1b5fb155066ca432a1948ba77047c7191b5f53c195bb66e3ac57d517ff7bc201","id":"ecad9b2f29774cd98a12dc9c82cfcb4d","category":"原本,DNA,链是,加密算法,数据网络","content_hash":"54ce1d0eb4759bae08f31d00095368b239af91d0dbb51f233092b65788f2a526","type":"article","title":"原本链go版本sdk测试","created":"1521771876","abstract":"原本链是一个分布式的底层数据网络；原本链是一个高效的，安全的，易用的，易扩展的，全球性质的，企业级的可信联盟链；原本链通过智能合约系统以及数字加密算法，实现了链上数据可持续性交互以及数据传输的安全；原本链通过高度抽象的“DTCP协议”与世界上独一无二的“原本DNA”互锁，确保链上数据100%不可篡改；原本链通过优化设计后的共识机制和独创的“闪电DNA”算法，已将区块写入速度提高至毫秒级别","dna":"3Y8DCROXXKLDSN3TPSBS20FELZA5PE26A001YKR5EAHHUJIXEX","language":"zh-cn","license":{"type":"cc","parameters":{"b":"2","y":"4"}}}
func Metadata_get_test() {
	res := app.Metadata_get(node_url, "", "3Y8DCROXXKLDSN3TPSBS20FELZA5PE26A001YKR5EAHHUJIXEX")
	if res.Code == "error" {
		fmt.Println("metadata query eror :", res.Msg)
	} else {
		js, _ := json.Marshal(res.Data)
		fmt.Println(string(js))
	}

}

//test result:
//success~ dna:  3Y8DCROXXKLDSN3TPSBS20FELZA5PE26A001YKR5EAHHUJIXEX
func Metadata_post_test() {

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
			Parameters map[string]string `json:"parameters,omitempty" binding:"required"`
		}{Type: "cc", Parameters: map[string]string{
			"y": "4",
			"b": "2",
		}},
	}
	pri_key := "50ced2bc6bc71ddfa517121b9df107400c9ba866344567da6aef82fac7824ade"
	app.GenerateMetadataFromContent(pri_key, md)

	res := app.Metadata_post(node_url, "", md)
	if res.Code == "error" {
		fmt.Println("metadata post error : ", res.Msg)
	} else {
		fmt.Println("success~ dna: ", res.Data.Dna)
	}
}

//test result:
//success~ dna:  {"block_hash":"794d30a2984e14e62cfd082a83877351ca3f0c1b","block_height":375,"data":{"description":"Creative Commons","id":"123456","parameters":[{"description":"是否允许演绎","name":"adaptation","type":"select","values":"y,n,sa"},{"description":"是否允许商用","name":"commercial","type":"select","values":"y,n"},{"description":"有效期","name":"expire","type":"timestamp","values":"0"},{"description":"授权价格","name":"price","type":"decimal","values":"0"}],"type":"cc","version":"4.0"},"data_height":12,"sender":"01218d7e82a0f2c4b31d7089a4dee33deba34899cc3924e99c1cd32d71ba25eb3a","time":1521698372}
func license_get_test () {
	res := app.License_get(node_url,"","cc")
	if res.Code == "error" {
		fmt.Println("metadata post error : ", res.Msg)
	} else {
		js,_ := json.Marshal(res.Data)
		fmt.Println("success~ dna: ", string(js))
	}
}
func main() {
	license_get_test()
}
