package app

import (
	uts "github.com/yuanbenio/universe-go-sdk/utils"
)

// GenPrivKeySecp256k1 : generate secret key by secp256k1
// return:hexadecimal private key ,hexadecimal compressed public key
func GenPrivKeySecp256k1() (string, string) {
	return uts.GenPrivKeySecp256k1()
}

// Sign  : sign function
// return : signature
func Sign(hash, prv []byte) (sig []byte, err error) {
	return uts.Sign(hash, prv)
}
