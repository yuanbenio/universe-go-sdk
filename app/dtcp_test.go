package app_test

import (
	"encoding/json"
	"fmt"
	"github.com/yuanbenio/universe-go-sdk/app"
	kts "github.com/yuanbenio/universe-go-sdk/types"
	"testing"
)

var (
	pri_key = "50ced2bc6bc71ddfa517121b9df107400c9ba866344567da6aef82fac7824ade"
	pub_key = "02e70aebdffd7b8f7495f35b504f0f3053024e44a87b8c419f7d886659e1475e19"
	content = "原本链是一个分布式的底层数据网络；" +
		"原本链是一个高效的，安全的，易用的，易扩展的，全球性质的，企业级的可信联盟链；" +
		"原本链通过智能合约系统以及数字加密算法，实现了链上数据可持续性交互以及数据传输的安全；" +
		"原本链通过高度抽象的“DTCP协议”与世界上独一无二的“原本DNA”互锁，确保链上数据100%不可篡改；" +
		"原本链通过优化设计后的共识机制和独创的“闪电DNA”算法，已将区块写入速度提高至毫秒级别"
	content_hash = "54ce1d0eb4759bae08f31d00095368b239af91d0dbb51f233092b65788f2a526"
	block_hash   = "4D36473D2FF1FE0772A6C0C55D7911295D8E1E27"
	signature    = "bf52d4d62e58cc280b7dc01d9ab91bc0e2ba9e66c1ff76972c230ad8011fd8af12b70be4f03cbe8a2baf898f0aefe59185cad2eedfcb402505aa299c1acb9e3400"
)

//test result:
//content_hash: 54ce1d0eb4759bae08f31d00095368b239af91d0dbb51f233092b65788f2a526
func TestGenContentHash(t *testing.T) {
	content_hash := app.GenContentHash(content)
	t.Logf("content_hash:%s", content_hash)
}

func GenerateMetadataFromContent() *kts.Metadata {
	//暂时不支持对非文章类型自动补全ContentHash
	md := &kts.Metadata{
		Content:   content,
		BlockHash: block_hash,
		Type:      "article",
		Title:     "原本链go版本sdk测试",
		License: struct {
			Params map[string]string `json:"parameters,omitempty"`
			Type   string            `json:"type,omitempty" binding:"required"`
		}{Type: "cc", Params: map[string]string{
			"y": "4",
			"b": "2",
		}},
	}
	app.FullMetadata(pri_key, md)
	jBs, _ := json.Marshal(md)
	fmt.Println(string(jBs))
	return md
}

//test result:
//{"pubkey":"02e70aebdffd7b8f7495f35b504f0f3053024e44a87b8c419f7d886659e1475e19","block_hash":"4D36473D2FF1FE0772A6C0C55D7911295D8E1E27","signature":"bb559331ef6f1d8fc3fab40cbbe4752dcbf6286061f6615d97281fafbcdad67d16a610784f9338ab3203887ef261c420911c66fadc65028ee89953a7f39551a200","id":"eb59190abed64c2e8fff6ec181fbbae4","category":"原本,DNA,链是,加密算法,数据网络","content_hash":"54ce1d0eb4759bae08f31d00095368b239af91d0dbb51f233092b65788f2a526","type":"article","title":"原本链go版本sdk测试","created":"1521719711","abstract":"原本链是一个分布式的底层数据网络；原本链是一个高效的，安全的，易用的，易扩展的，全球性质的，企业级的可信联盟链；原本链通过智能合约系统以及数字加密算法，实现了链上数据可持续性交互以及数据传输的安全；原本链通过高度抽象的“DTCP协议”与世界上独一无二的“原本DNA”互锁，确保链上数据100%不可篡改；原本链通过优化设计后的共识机制和独创的“闪电DNA”算法，已将区块写入速度提高至毫秒级别","dna":"OTQ1TOMUTDJZZ729GID9K8JVUGBLA1T6AXQ4ZL7K0H75B2DNB","language":"zh-cn","license":{"type":"cc","parameters":{"b":"2","y":"4"}}}
func TestGenerateMetadataFromContent(t *testing.T) {
	GenerateMetadataFromContent()
}

//test result:
//{"pubkey":"02e70aebdffd7b8f7495f35b504f0f3053024e44a87b8c419f7d886659e1475e19","block_hash":"4D36473D2FF1FE0772A6C0C55D7911295D8E1E27","signature":"13173f52bee33add892bdf232f05d5e647fc49dd095eaed1f9de11852047d3eb267e66b4e4115691ff4f2733a2c2636731b826b7acbbdc56628927a55c83533c01","id":"41b74e18aab743369bc4329e8fee5c1f","category":"原本,DNA,链是,加密算法,数据网络","content_hash":"54ce1d0eb4759bae08f31d00095368b239af91d0dbb51f233092b65788f2a526","type":"article","title":"原本链go版本sdk测试","created":"1521719815","abstract":"原本链是一个分布式的底层数据网络；原本链是一个高效的，安全的，易用的，易扩展的，全球性质的，企业级的可信联盟链；原本链通过智能合约系统以及数字加密算法，实现了链上数据可持续性交互以及数据传输的安全；原本链通过高度抽象的“DTCP协议”与世界上独一无二的“原本DNA”互锁，确保链上数据100%不可篡改；原本链通过优化设计后的共识机制和独创的“闪电DNA”算法，已将区块写入速度提高至毫秒级别","dna":"5DOH9KFV6CMSSBXY6MW703YV2Y0XD0V7HC6L45S8KYN2737A9V","language":"zh-cn","license":{"type":"cc","parameters":{"b":"2","y":"4"}}}
//success~ 13173f52bee33add892bdf232f05d5e647fc49dd095eaed1f9de11852047d3eb267e66b4e4115691ff4f2733a2c2636731b826b7acbbdc56628927a55c83533c01
func TestGenMetadataSignature(t *testing.T) {
	md := GenerateMetadataFromContent()
	if sign, err := app.GenMetadataSignature(pri_key, md); err != nil {
		t.Errorf("generator metadata signature error:%s", err.Error())
	} else {
		t.Logf("success~:%s", sign)
	}

}

//test result:
//success~ true
func TestVerifySignature(t *testing.T) {
	md := GenerateMetadataFromContent()
	if b, err := app.VerifySignature(md); err != nil {
		t.Errorf("verify metadata signature error:%s", err.Error())
	} else {
		t.Logf("success~:%s", b)
	}
}

//test result:
//dna: 5DOH9KFV6CMSSBXY6MW703YV2Y0XD0V7HC6L45S8KYN2737A9V
func TestGenerateDNA(t *testing.T) {
	md_sign := "13173f52bee33add892bdf232f05d5e647fc49dd095eaed1f9de11852047d3eb267e66b4e4115691ff4f2733a2c2636731b826b7acbbdc56628927a55c83533c01"
	t.Logf("dna:%s", app.GenerateDNA(md_sign))
}
