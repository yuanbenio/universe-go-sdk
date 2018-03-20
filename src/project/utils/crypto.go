package utils

import (
	"crypto/ecdsa"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/primasio/go-base36"
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
		return crypto.ToECDSAPub(d), nil
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
	p, _ := crypto.ToECDSA(prv)
	return crypto.Sign(hash, p)
}

// VerifySignature
func VerifySignature(hash []byte, signature []byte) bool {
	pubkey, _ := crypto.Ecrecover(hash, signature)
	return crypto.VerifySignature(pubkey, hash, signature)
}

// Hasher
func Hasher(data ...[]byte) string {
	return hex.EncodeToString(crypto.Keccak256(data...))
}

//GenerateDNA
func GenerateDNA(metadataSignature string) string {
	d, _ := hex.DecodeString(metadataSignature)
	digest := base36.EncodeBytes(crypto.Keccak256(d))
	return digest
}
