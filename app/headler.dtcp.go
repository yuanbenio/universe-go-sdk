package app

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"strings"
	"time"
	kts "github.com/yuanbenio/universe-go-sdk/types"
	uts "github.com/yuanbenio/universe-go-sdk/utils"
)

// GenContentHash : generate content_hash
// params:content
// return:hexadecimal keccak(content)
func GenContentHash(content string) string {
	return uts.Hasher([]byte(content))
}

// GenMetadataSignature : calculation metadata.Signature
// params:hexadecimal privateKey and metadata
// return: hexadecimal signature
func GenMetadataSignature(privateKey string, md *kts.Metadata) (string, error) {
	if privateKey == "" || md == nil {
		return "", errors.New("there must be a private key and license")
	}
	prvBs, _ := hex.DecodeString(privateKey)

	h := crypto.Keccak256(md.DumpsRmSignSort())
	if signBs, err := uts.Sign(h, prvBs); err != nil {
		return "", err
	} else {
		return hex.EncodeToString(signBs), nil
	}
}

// VerifySignature : verify metadata`s signature
// params : metadata
// result : verify result
func VerifySignature(md *kts.Metadata) (bool, error) {
	if md == nil || md.PubKey == "" {
		return false, errors.New("public key is empty or metadata is nil")
	}
	h := crypto.Keccak256(md.DumpsRmSignSort())
	if signBs, err := hex.DecodeString(md.Signature); err != nil {
		return false, err
	} else {
		d1, _ := hex.DecodeString(md.PubKey)
		return crypto.VerifySignature(d1, h, signBs[:len(signBs)-1]), nil // remove recovery id
	}
}

// GenerateDNA : generate metadata`s lightning dna
// params : metadata`s signature
// result : base36 decimal string
func GenerateDNA(mdSign string) string {
	return uts.GenerateDNA(mdSign)
}

// FullMetadata : full the metadata
// params : hexadecimal private key and metadata
// result : full metadata
func FullMetadata(privateKey string, md *kts.Metadata) (err error) {
	if md == nil {
		return errors.New("metadata is nil")
	}
	if md.BlockHash == "" || md.BlockHeight == ""{
		return errors.New("block hash or block height is empty")
	}
	if md.License.Type == "" || (md.License.Type !="none" && md.License.Parameters == nil) {
		return errors.New("license is nil")
	}
	if md.ContentHash == "" {
		if md.Content == "" {
			return errors.New("content and contentHash can't be empty at the same time")
		}
		contentHash := GenContentHash(md.Content)
		md.ContentHash = contentHash
	}

	if md.PubKey == "" {
		priBs, _ := hex.DecodeString(privateKey)
		md.PubKey = uts.GetPubKeyFromPri(priBs)
	}
	if privateKey == "" {
		return errors.New("there must be a private key")
	}

	if md.Type == "" {
		return errors.New("type is empty")
	}

	if md.Language == "" {
		md.Language = "zh-CN"
	}

	md.Created = fmt.Sprintf("%d", time.Now().Unix())

	switch md.Type {
	case kts.PRIVATE.Value(),kts.CUSTOM.Value():
		//pass
	case kts.ARTICLE.Value():
		if md.Abstract == "" && md.Content != ""{
			_s := strings.Split(md.Content, "")
			if len(_s) > 200 {
				md.Abstract = strings.Join(_s[:200], "")
			} else {
				md.Abstract = strings.Join(_s, "")
			}
		}
	case kts.IMAGE.Value(), kts.VIDEO.Value(), kts.AUDIO.Value():
		if md.Data == nil {
			return errors.New("Please add extends  data!")
		}
	default:
		return errors.New("content type is nonsupport")
	}
	if md.Type != kts.PRIVATE.Value() {
		if md.Title == "" {
			return errors.New("title is empty")
		}
		if md.Category == "" {
			return errors.New("category can't be  empty !")
		}
	}

	signature, err := GenMetadataSignature(privateKey, md)
	if err != nil {
		return err
	}
	if md.Signature == "" {
		md.Signature = signature
	}
	if md.DNA == "" {
		md.DNA = GenerateDNA(signature)
	}
	return nil

}
