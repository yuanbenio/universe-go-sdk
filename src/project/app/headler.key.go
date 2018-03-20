package app

import (
	uts "project/utils"
)

// GenPrivKeySecp256k1 生成一对公私密钥
// return:私钥，公钥
func GenPrivKeySecp256k1() (string, string) {
	return uts.GenPrivKeySecp256k1()
}
func Sign(hash, prv []byte) (sig []byte, err error) {
	return uts.Sign(hash, prv)
}
