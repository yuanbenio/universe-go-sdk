package app

import (
	"fmt"
	"time"
	kts "project/types"
	"github.com/ethereum/go-xethereum/crypto"
)

//GenContentHash
func GenContentHash(content string) string {
	return kts.Hasher([]byte(content))
}

//GenMetadataSignature
func GenMetadataSignature(private_key string,title string,content_hash string,license *kts.License) string {
	if private_key == nil || license == nil {
		//todo : 抛出错误
		return;
	}
	p, _ := crypto.ToECDSA(private_key)
	hash_string := hex.EncodeToString(crypto.Keccak256(title,content_hash,license.Dumps()))
	return crypto.Sign(hash_string, p)
}

func GenerateDNA(metadataSignature string) string {
	d, _ := hex.DecodeString(metadataSignature)
	digest := base36.EncodeBytes(crypto.Keccak256(d))
	return digest
}

func GenerateMetadataFromContent(private_key string,md *kts.Metadata){
	if md == nil {
		//todo : 抛出错误
		return;
	}
	contentHash = GenContentHash(md.content)

	if md.PubKey == nil {
		//todo: 应该抛出错误
		return;
	}
	md.Signature = GenMetadataSignature(private_key,md.Title,contentHash,md.License)


	//todo :　做一些非空检查 和 字段填充
	md.Category = category
	md.Title = title
	md.Content = content
	md.ContentHash = contentHash
	md.Type = contentType

	md.Created = time.Now().Unix()
	//md.Abstract =
	md.DNA = GenerateDNA()

}



