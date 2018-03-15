package app

import (
	"fmt"
	kts "project/types"
	"github.com/ethereum/go-xethereum/crypto"
)

// GenPrivKeySecp256k1 生成一对公私密钥
// return:私钥，公钥
func GenPrivKeySecp256k1 () (string,string){
	return kts.GenPrivKeySecp256k1()
}
func Sign(hash, prv []byte) (sig []byte, err error) {
	return kts.Sign(hash,prv)
}

