package main

import (
	"encoding/hex"
	"fmt"
	"universe-go-sdk/app"
)

//test result:
//private_key:50ced2bc6bc71ddfa517121b9df107400c9ba866344567da6aef82fac7824ade
//public_key:02e70aebdffd7b8f7495f35b504f0f3053024e44a87b8c419f7d886659e1475e19
func GenPrivKeySecp256k1_test() {
	pri, pub := app.GenPrivKeySecp256k1()
	fmt.Println(fmt.Sprintf("private_key:%s\n public_key:%s", pri, pub))
}

//test result:
// sign success, signature： 9c8dec698554008b1be8204f616240ab0ab71a79cec743153abe3f394e06364d5ef027766ae17f72a56036cc05b065174c742703b6f094751e6f55d7de55476100
func Sign_test() {
	content := "原本链是一个分布式的底层数据网络；" +
		"原本链是一个高效的，安全的，易用的，易扩展的，全球性质的，企业级的可信联盟链；" +
		"原本链通过智能合约系统以及数字加密算法，实现了链上数据可持续性交互以及数据传输的安全；" +
		"原本链通过高度抽象的“DTCP协议”与世界上独一无二的“原本DNA”互锁，确保链上数据100%不可篡改；" +
		"原本链通过优化设计后的共识机制和独创的“闪电DNA”算法，已将区块写入速度提高至毫秒级别"
	pri_key := "50ced2bc6bc71ddfa517121b9df107400c9ba866344567da6aef82fac7824ade"
	priBs, _ := hex.DecodeString(pri_key)
	conHash, _ := hex.DecodeString(app.GenContentHash(content))
	if signBs, err := app.Sign(conHash, priBs); err != nil {
		fmt.Println("sign error", err.Error())
	} else {
		fmt.Println("sign success, signature：", hex.EncodeToString(signBs))
	}
}

func main() {
	GenPrivKeySecp256k1_test()
}
