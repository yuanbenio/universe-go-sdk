package app

import (
	"fmt"
	kts "project/types"
	"github.com/ethereum/go-xethereum/crypto"
)

//GenContentHash
func GenContentHash(content string) string {
	return kts.Hasher([]byte(content))
}

//SignForMetadata
func SignForContent(title string,) string {
	
}

func GenerateMetadataFromContent(pubKey string,content string,title string,contentType string,source string,extra,data,license *kes.License){
	contentHash = GenContentHash(content)
	md := &kts.Metadata{}

	md.DNA = GenerateDNA()

}

func GenerateDNA(metadataSignature string) string {
	d, _ := hex.DecodeString(metadataSignature)
	digest := base36.EncodeBytes(crypto.Keccak256(d))
	return digest
}



