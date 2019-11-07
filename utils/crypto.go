package utils

import (
	"crypto/ecdsa"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/yuanbenio/universe-go-sdk/utils/base36"
)

// GenPrivKeySecp256k1
func GenPrivKeySecp256k1() (string, string) {
	p, _ := crypto.GenerateKey()
	return hex.EncodeToString(crypto.FromECDSA(p)), hex.EncodeToString(crypto.CompressPubkey(&p.PublicKey))
}

// GetPubKeyFrom
func GetPubKeyFrom(pubkey string) (*ecdsa.PublicKey, error) {
	if d, err := hex.DecodeString(pubkey); err != nil {
		return nil, err
	} else {
		return crypto.DecompressPubkey(d)
	}
}

// Address
func Address(pubkey string) (string, error) {
	if d, err := hex.DecodeString(pubkey); err != nil {
		return "", err
	} else {
		pub, err := crypto.DecompressPubkey(d)
		if err != nil {
			return "", err
		} else {
			return crypto.PubkeyToAddress(*pub).String(), nil
		}
	}
}

func GetPubKeyFromPri(private_key []byte) (pubKey string) {
	p, _ := crypto.ToECDSA(private_key)
	return hex.EncodeToString(crypto.CompressPubkey(&p.PublicKey))
}

// GetPriKeyFrom
func GetPriKeyFrom(prikey string) (*ecdsa.PrivateKey, error) {
	if d, err := hex.DecodeString(prikey); err != nil {
		return nil, err
	} else {
		return crypto.ToECDSA(d)
	}
}

// Sign
func Sign(hash, prv []byte) (sig []byte, err error) {
	p, err := crypto.ToECDSA(prv)
	if err != nil {
		return nil, err
	}
	return crypto.Sign(hash, p)
}

// VerifySignature
func VerifySignature(pubKey, hash []byte, signature []byte) bool {
	signBs := signature
	if len(signature) == 65 {
		signBs = signature[:len(signature)-1]
	}
	return crypto.VerifySignature(pubKey, hash, signBs) // remove recovery id
}

// Hasher
func Keccak256(data ...[]byte) []byte {
	return crypto.Keccak256(data...)
}

// GenerateDNA
func GenerateDNA(md_sign string) string {
	d, _ := hex.DecodeString(md_sign)
	digest := base36.EncodeBytes(crypto.Keccak256(d))
	return digest
}
