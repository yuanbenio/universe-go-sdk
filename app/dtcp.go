package app

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	kts "github.com/yuanbenio/universe-go-sdk/types"
	uts "github.com/yuanbenio/universe-go-sdk/utils"
	"strings"
	"time"
)

var (
	ErrPrivateKeyNil = errors.New("private key is nil")
	ErrPublicKeyNil  = errors.New("public key is nil")
	ErrBlockInfoNil  = errors.New("block information is nil")
	ErrContentNil    = errors.New("content and contentHash can't be nil at the same time")
	ErrTypeNil       = errors.New("type is nil")
	ErrTypeInvalid   = errors.New("invalid type")
	ErrExtendDataNil = errors.New("extends data is nil")
	ErrTitleNil      = errors.New("title is nil")
	ErrCategoryNil   = errors.New("category is nil")
	ErrSubKeysNil    = errors.New("subKeys is nil")
)

const (
	DefaultLanguage = "zh-CN"
)

//GenContentHash: generate content_hash
//params: content
//return: hexadecimal keccak(content)
func GenContentHash(content string) string {
	return uts.Hasher([]byte(content))
}

//GenMetadataSignature: calculation metadata.Signature
//params: hexadecimal privateKey and metadata
//return: hexadecimal signature
func GenMetadataSignature(privateKey string, md *kts.Metadata) (string, error) {
	if privateKey == "" {
		return "", ErrPrivateKeyNil
	}
	if md == nil {
		return "", ErrMDNil
	}
	prvBs, err := hex.DecodeString(privateKey)
	if err != nil {
		return "", err
	}
	h := crypto.Keccak256(md.DumpsRmSignSort())
	if signBs, err := uts.Sign(h, prvBs); err != nil {
		return "", err
	} else {
		return hex.EncodeToString(signBs), nil
	}
}

//VerifySignature: verify metadata`s signature
//params: metadata
//result: verify result
func VerifySignature(md *kts.Metadata) (bool, error) {
	if md == nil {
		return false, ErrMDNil
	}
	if md.PubKey == "" {
		return false, ErrPublicKeyNil
	}
	h := crypto.Keccak256(md.DumpsRmSignSort())
	if signBs, err := hex.DecodeString(md.Signature); err != nil {
		return false, err
	} else {
		d1, err := hex.DecodeString(md.PubKey)
		if err != nil {
			return false, err
		}
		return crypto.VerifySignature(d1, h, signBs[:len(signBs)-1]), nil // remove recovery id
	}
}

//GenerateDNA: generate metadata`s lightning dna
//params: metadata`s signature
//result: base36 decimal string
func GenerateDNA(mdSign string) string {
	return uts.GenerateDNA(mdSign)
}

//FullMetadata: full the metadata
//params: hexadecimal private key and metadata
//result: full metadata
func FullMetadata(privateKey string, md *kts.Metadata) (err error) {
	if md == nil {
		return ErrMDNil
	}
	if md.BlockHash == "" || md.BlockHeight == "" {
		return ErrBlockInfoNil
	}
	if md.License.Type == "" || (md.License.Type != "none" && md.License.Params == nil) {
		return ErrLicenseNil
	}
	if md.ContentHash == "" {
		if md.Content == "" {
			return ErrContentNil
		}
		contentHash := GenContentHash(md.Content)
		md.ContentHash = contentHash
	}

	if md.PubKey == "" {
		priBs, _ := hex.DecodeString(privateKey)
		md.PubKey = uts.GetPubKeyFromPri(priBs)
	}
	if privateKey == "" {
		return ErrPrivateKeyNil
	}

	if md.Type == "" {
		return ErrTypeNil
	}

	if md.Language == "" {
		md.Language = DefaultLanguage
	}

	if md.Created == "" {
		md.Created = fmt.Sprintf("%d", time.Now().Unix())
	}

	switch md.Type {
	case kts.PRIVATE.Value(), kts.CUSTOM.Value():
		//pass
	case kts.ARTICLE.Value():
		if md.Abstract == "" && md.Content != "" {
			_s := strings.Split(md.Content, "")
			if len(_s) > 200 {
				md.Abstract = strings.Join(_s[:200], "")
			} else {
				md.Abstract = strings.Join(_s, "")
			}
		}
	case kts.IMAGE.Value(), kts.VIDEO.Value(), kts.AUDIO.Value():
		if md.Data == nil {
			return ErrExtendDataNil
		}
	default:
		return ErrTypeInvalid
	}
	if md.Type != kts.PRIVATE.Value() {
		if md.Title == "" {
			return ErrTitleNil
		}
		if md.Category == "" {
			return ErrCategoryNil
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
	md.Content = ""
	return nil

}

//GenRegisterAccountReq
//params: hexadecimal private key and metadata
//result: full metadata
func GenRegisterAccountReq(privateKey string, subKeys []string) (*kts.RegisterAccountReq, error) {
	if subKeys == nil || len(subKeys) == 0 {
		return nil, ErrSubKeysNil
	}
	if privateKey == "" {
		return nil, ErrPrivateKeyNil
	}
	req := &kts.RegisterAccountReq{
		Subkeys: subKeys,
	}
	priKeyBs, err := hex.DecodeString(privateKey)
	if err != nil {
		return nil, err
	} else {
		req.Pubkey = uts.GetPubKeyFromPri(priKeyBs)
	}

	_d, err := json.Marshal(subKeys)
	if err != nil {
		return nil, err
	}
	h := crypto.Keccak256(_d)
	if signBs, err := uts.Sign(h, priKeyBs); err != nil {
		return nil, err
	} else {
		req.Signature = hex.EncodeToString(signBs)
	}
	return req, nil

}
