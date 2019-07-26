package utils

import (
	"crypto/sha256"

	"golang.org/x/crypto/ripemd160"
)

func DoubleSha256(b []byte) []byte {
	hasher := sha256.New()
	hasher.Write(b)
	sha := hasher.Sum(nil)
	hasher.Reset()
	hasher.Write(sha)
	return hasher.Sum(nil)
}

func Sha256RipeMD160(b []byte) []byte {
	ripe := ripemd160.New()
	sha := sha256.New()
	sha.Write(b)
	ripe.Write(sha.Sum(nil))
	return ripe.Sum(nil)
}
