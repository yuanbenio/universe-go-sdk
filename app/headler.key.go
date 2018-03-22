package app

import (
	uts "universe-go-sdk/utils"
)

// GenPrivKeySecp256k1 生成一对公私密钥
// return:16进制编码对私钥，公钥
func GenPrivKeySecp256k1() (string, string) {
	return uts.GenPrivKeySecp256k1()
}

// Sign  签名函数
// return 签字字节数组
func Sign(hash, prv []byte) (sig []byte, err error) {
	return uts.Sign(hash, prv)
}
