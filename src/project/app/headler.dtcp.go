package app

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/satori/go.uuid"
	"github.com/yanyiwu/gojieba"
	kts "project/types"
	uts "project/utils"
	"strings"
	"time"
)

//GenContentHash
func GenContentHash(content string) string {
	return uts.Hasher([]byte(content))
}

//GenMetadataSignature
func GenMetadataSignature(private_key string, md *kts.Metadata) (string, error) {
	if private_key == "" || md == nil {
		return "", errors.New("there must be a private key and license")
	}
	prvBs, _ := hex.DecodeString(private_key)
	//j,_ := json.Marshal(md)

	h := crypto.Keccak256(md.DumpsRmSignSort())
	if signBs, err := uts.Sign(h, prvBs); err != nil {
		return "", err
	} else {
		return hex.EncodeToString(signBs), nil
	}
}

func VerifySignature(md *kts.Metadata) (bool, error) {
	if md == nil || md.PubKey == "" {
		return false, errors.New("public key is empty or metadata is nil")
	}
	//conBs, _ := hex.DecodeString(md.ContentHash)
	//h := crypto.Keccak256([]byte(md.Title), conBs, md.DumpsLicense())
	h := crypto.Keccak256(md.DumpsRmSignSort())
	if signBs, err := hex.DecodeString(md.Signature); err != nil {
		return false, err
	} else {
		d1, _ := hex.DecodeString(md.PubKey)
		return crypto.VerifySignature(d1, h, signBs[:len(signBs)-1]), nil // remove recovery id
	}
}

//GenerateDNA
func GenerateDNA(md_sign string,block_hash string) string {
	return uts.GenerateDNA(md_sign,block_hash)
}

//GenerateMetadataFromContent
func GenerateMetadataFromContent(private_key string, md *kts.Metadata) (err error) {
	if md == nil {
		return errors.New("metadata is nil")
	}
	if md.Content == "" {
		return errors.New("metadata content is empty")
	}
	if md.BlockHash == "" {
		return errors.New("block hash is empty")
	}
	if md.ContentHash == "" {
		contentHash := GenContentHash(md.Content)
		md.ContentHash = contentHash
	}

	if md.PubKey == "" {
		priBs, _ := hex.DecodeString(private_key)
		md.PubKey = uts.GetPubKeyFromPri(priBs)
	}
	if md.Title == "" {
		md.Title = md.Type
	}
	if private_key == "" {
		return errors.New("there must be a private key")
	}

	if err != nil {
		return err
	}

	if md.ID == "" {
		md.ID = strings.Replace(uuid.Must(uuid.NewV4()).String(), "-", "", -1)
	}

	if md.Type == "" {
		md.Type = "article"
	}

	if md.Language == "" {
		md.Language = "zh-cn"
	}

	md.Created = fmt.Sprintf("%d", time.Now().Unix())

	switch md.Type {

	case "article":

		if md.Abstract == "" {
			_s := strings.Split(md.Content, "")
			if len(_s) > 200 {
				md.Abstract = strings.Join(_s[:200], "")
			} else {
				md.Abstract = strings.Join(_s, "")
			}
		}

		x := gojieba.NewJieba()
		defer x.Free()

		_j := x.Extract(md.Content, 5)
		if md.Category != "" {
			_j = append(_j, md.Category)
		}
		md.Category = strings.Join(_j, ",")
	case "image", "video", "audio":
		//todo : 添加图片的处理
		if md.ContentHash == "" {
			return errors.New("there must be a contentHash if the content type is image、video or audio")
		}

	default:
		return errors.New("content type is nonsupport")
	}
	signature, err := GenMetadataSignature(private_key, md)
	if md.Signature == "" {
		md.Signature = signature
	}
	if md.DNA == "" {
		md.DNA = GenerateDNA(signature,md.BlockHash)
	}

	return nil

}
