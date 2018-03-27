package main

import (
	"encoding/hex"
	_ "encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"universe-go-sdk/app"
	kts "universe-go-sdk/types"
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
func GenContentHash_test() {
	content_hash := app.GenContentHash(content)
	fmt.Println("content_hash:", content_hash)
}

//test result:
//{"pubkey":"02e70aebdffd7b8f7495f35b504f0f3053024e44a87b8c419f7d886659e1475e19","block_hash":"4D36473D2FF1FE0772A6C0C55D7911295D8E1E27","signature":"bb559331ef6f1d8fc3fab40cbbe4752dcbf6286061f6615d97281fafbcdad67d16a610784f9338ab3203887ef261c420911c66fadc65028ee89953a7f39551a200","id":"eb59190abed64c2e8fff6ec181fbbae4","category":"原本,DNA,链是,加密算法,数据网络","content_hash":"54ce1d0eb4759bae08f31d00095368b239af91d0dbb51f233092b65788f2a526","type":"article","title":"原本链go版本sdk测试","created":"1521719711","abstract":"原本链是一个分布式的底层数据网络；原本链是一个高效的，安全的，易用的，易扩展的，全球性质的，企业级的可信联盟链；原本链通过智能合约系统以及数字加密算法，实现了链上数据可持续性交互以及数据传输的安全；原本链通过高度抽象的“DTCP协议”与世界上独一无二的“原本DNA”互锁，确保链上数据100%不可篡改；原本链通过优化设计后的共识机制和独创的“闪电DNA”算法，已将区块写入速度提高至毫秒级别","dna":"OTQ1TOMUTDJZZ729GID9K8JVUGBLA1T6AXQ4ZL7K0H75B2DNB","language":"zh-cn","license":{"type":"cc","parameters":{"b":"2","y":"4"}}}
func GenerateMetadataFromContent_test() *kts.Metadata {
	//暂时不支持对非文章类型自动补全ContentHash
	md := &kts.Metadata{
		Content:   content,
		BlockHash: block_hash,
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
	app.FullMetadata(pri_key, md)
	jBs, _ := json.Marshal(md)
	fmt.Println(string(jBs))
	return md
}

//test result:
//{"pubkey":"02e70aebdffd7b8f7495f35b504f0f3053024e44a87b8c419f7d886659e1475e19","block_hash":"4D36473D2FF1FE0772A6C0C55D7911295D8E1E27","signature":"13173f52bee33add892bdf232f05d5e647fc49dd095eaed1f9de11852047d3eb267e66b4e4115691ff4f2733a2c2636731b826b7acbbdc56628927a55c83533c01","id":"41b74e18aab743369bc4329e8fee5c1f","category":"原本,DNA,链是,加密算法,数据网络","content_hash":"54ce1d0eb4759bae08f31d00095368b239af91d0dbb51f233092b65788f2a526","type":"article","title":"原本链go版本sdk测试","created":"1521719815","abstract":"原本链是一个分布式的底层数据网络；原本链是一个高效的，安全的，易用的，易扩展的，全球性质的，企业级的可信联盟链；原本链通过智能合约系统以及数字加密算法，实现了链上数据可持续性交互以及数据传输的安全；原本链通过高度抽象的“DTCP协议”与世界上独一无二的“原本DNA”互锁，确保链上数据100%不可篡改；原本链通过优化设计后的共识机制和独创的“闪电DNA”算法，已将区块写入速度提高至毫秒级别","dna":"5DOH9KFV6CMSSBXY6MW703YV2Y0XD0V7HC6L45S8KYN2737A9V","language":"zh-cn","license":{"type":"cc","parameters":{"b":"2","y":"4"}}}
//success~ 13173f52bee33add892bdf232f05d5e647fc49dd095eaed1f9de11852047d3eb267e66b4e4115691ff4f2733a2c2636731b826b7acbbdc56628927a55c83533c01
func GenMetadataSignature_test(pri_key string, md *kts.Metadata) {
	if sign, err := app.GenMetadataSignature(pri_key, md); err != nil {
		fmt.Println("generator metadata signature error", err.Error())
		panic(err)
	} else {
		fmt.Println("success~", sign)
	}

}

//test result:
//success~ true
func VerifySignature_test(md *kts.Metadata) {
	if b, err := app.VerifySignature(md); err != nil {
		fmt.Println("verify metadata signature error", err.Error())
		panic(err)
	} else {
		fmt.Println("success~", b)
	}
}

//test result:
//dna: 5DOH9KFV6CMSSBXY6MW703YV2Y0XD0V7HC6L45S8KYN2737A9V
func GenerateDNA_test() {
	md_sign := "13173f52bee33add892bdf232f05d5e647fc49dd095eaed1f9de11852047d3eb267e66b4e4115691ff4f2733a2c2636731b826b7acbbdc56628927a55c83533c01"
	fmt.Println("dna:", app.GenerateDNA(md_sign))
}

func main() {
	//GenerateDNA_test()

	//	{"pubkey":"","title":"原本链java版本sdk测试","type":""}

	//s := "9eda9afe54859080783c288fab3bdd3e78dda8878b33359a7e1ef0d4818e1ce0"
	//pri := "3c4dbee4485557edce3c8878be34373c1a41d955f38d977cfba373642983ce4c"
	//p,_ := hex.DecodeString(pri)
	////s1,_ := hex.DecodeString(s)
	//s1 := utils.Hasher([]byte(s))
	//s3,_ := hex.DecodeString(s1)
	//s2,_ := utils.Sign(s3,p)
	//
	//fmt.Println("sign:",hex.EncodeToString(s2))

	md := &kts.Metadata{
		Abstract:    "原本链是一个分布式的底层数据网络；原本链是一个高效的，安全的，易用的，易扩展的，全球性质的，企业级的可信联盟链；原本链通过智能合约系统以及数字加密算法，实现了链上数据可持续性交互以及数据传输的安全；原本链通过高度抽象的“DTCP协议”与世界上独一无二的“原本DNA”互锁，确保链上数据100%不可篡改；原本链通过优化设计后的共识机制和独创的“闪电DNA”算法，已将区块写入速度提高至毫秒级别",
		BlockHash:   "4D36473D2FF1FE0772A6C0C55D7911295D8E1E27",
		Category:    "原本,数据,DNA,安全,区块",
		ContentHash: "54ce1d0eb4759bae08f31d00095368b239af91d0dbb51f233092b65788f2a526",
		Created:     "1522067969129",
		ID:          "4c9b26e165344cf391822cb4c221e8b5",
		Language:    "zh-cn",
		License: struct {
			Type       string            `json:"type,omitempty" binding:"required"`
			Parameters map[string]string `json:"parameters,omitempty" binding:"required"`
		}{Type: "cc", Parameters: map[string]string{
			"y": "4",
			"b": "2",
		}},
		PubKey: "03d75b59a801f6db4bbb501ff8b88743902aa83a3e54237edcd532716fd27dea77",
		Title:  "原本链java版本sdk测试",
		Type:   "article",
		//Signature:"ffd1515581f7962444291faf67f27f3ef13b9401f52ba076f2cd5b25f88341a923492d08b82cf7b6801c5d8e840bc771d960e74bde50c79387e151f8f86079b601",
		Signature: "ffd1515581f7962444291faf67f27f3ef13b9401f52ba076f2cd5b25f88341a923492d08b82cf7b6801c5d8e840bc771d960e74bde50c79387e151f8f86079b601",
		DNA:       "65SO0BNCXRLWIEIKVFSZAAL8WG2964P91N3S29T8HS3YP1RQ67",
	}

	fmt.Println("signature from md:", md.Signature)

	sign, _ := app.GenMetadataSignature("3c4dbee4485557edce3c8878be34373c1a41d955f38d977cfba373642983ce4c", md)
	fmt.Println("gen signature from md:", sign)

	s, _ := hex.DecodeString(sign)
	pub, _ := crypto.Ecrecover(crypto.Keccak256(md.DumpsRmSignSort()), s)
	fmt.Println("pubkey :", hex.EncodeToString(pub))

	v, _ := app.VerifySignature(md)
	fmt.Println("verify :", v)

	dna := app.GenerateDNA(md.Signature)
	fmt.Println("dna from md sign:", dna)

	dna = app.GenerateDNA(sign)
	fmt.Println("dna from gen sign :", dna)

	dna = app.GenerateDNA("ffd1515581f7962444291faf67f27f3ef13b9401f52ba076f2cd5b25f88341a923492d08b82cf7b6801c5d8e840bc771d960e74bde50c79387e151f8f86079b600")
	fmt.Println("dna from java sign :", dna)

	fmt.Println("dna :", md.DNA)

	//res := app.Metadata_save("http://119.23.22.129:9000", "", md)
	//if res.Code == "error" {
	//	fmt.Println("metadata post error : ", res.Msg)
	//} else {
	//	fmt.Println("success~ dna: ", res.Data.Dna)
	//}

	//hash := crypto.Keccak256(md.DumpsRmSignSort())
	//d,_ := hex.DecodeString(md.Signature)
	//pub,_ := crypto.Ecrecover(hash,d)
	//p,_ := crypto.DecompressPubkey(pub)
	//fmt.Println(hex.EncodeToString(crypto.FromECDSAPub(p)))
	//
	//
	//
	//r,_ := app.VerifySignature(md)
	//fmt.Println(r)

}
