package crypto

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"math/big"
)

type ecdsaSignature struct {
	R, S *big.Int
}

func ToHash(i interface{}) ([]byte, error) {
	b, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}

	h := sha256.New()
	h.Write(b)

	return h.Sum(nil), nil
}

func encodeSignature(sig *ecdsaSignature) string {
	cmb := append(sig.R.Bytes(), sig.S.Bytes()...)
	return base64.StdEncoding.EncodeToString(cmb)
}

func decodeSignature(signature string) (sig *ecdsaSignature, err error) {
	sigBytes, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return
	}

	keySize := len(sigBytes) / 2
	r := big.NewInt(0).SetBytes(sigBytes[:keySize])
	s := big.NewInt(0).SetBytes(sigBytes[keySize:])

	sig = &ecdsaSignature{
		R: r,
		S: s,
	}
	return
}

func Sign(privKey *ecdsa.PrivateKey, hash []byte) (signature string, err error) {
	r, s, err := ecdsa.Sign(rand.Reader, privKey, hash)
	if err != nil {
		return
	}

	sig := &ecdsaSignature{
		R: r,
		S: s,
	}

	return encodeSignature(sig), nil
}

func Verify(pubKey *ecdsa.PublicKey, signature string, hash []byte) (valid bool, err error) {
	sig, err := decodeSignature(signature)
	if err != nil {
		return
	}

	valid = ecdsa.Verify(pubKey, hash, sig.R, sig.S)
	return
}
