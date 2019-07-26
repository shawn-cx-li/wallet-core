package crypto

import (
	"math/big"
	"testing"
)

const (
	password = "abcdefg"
	path1    = "m/44'/60'/0'/0/0"
	path2    = "m/44'/60'/0'/0/1"
)

type Transaction struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount big.Int `json:"amount"`
}

func TestMnemonicGenerateAndRecover(t *testing.T) {
	m, s, err := GetNewSeed(password)
	if err != nil {
		t.Errorf("Failed to create new seed, err: %s", err)
	}

	rs, err := RecoverSeed(m, password)
	if err != nil {
		t.Errorf("Failed to recover seed, err: %s", err)
	}

	if s != rs {
		t.Errorf("Seed %s does not match the recovered seed %s", s, rs)
	}
}

func TestSignAndVerify(t *testing.T) {
	_, s, err := GetNewSeed(password)
	if err != nil {
		t.Errorf("Failed to create new seed, err: %s", err)
	}

	privKey1, _, err := GetPrivateKey(s, path1)
	if err != nil {
		t.Errorf("Failed to get private key 1, err: %s", err)
	}

	privKey2, _, err := GetPrivateKey(s, path2)
	if err != nil {
		t.Errorf("Failed to create new seed, err: %s", err)
	}

	tx := &Transaction{
		From:   "abc",
		To:     "efg",
		Amount: *big.NewInt(1000),
	}

	txHash, err := ToHash(tx)
	if err != nil {
		t.Errorf("Failed to hash message, err: %v", err)
	}

	sig, err := Sign(privKey1, txHash)
	if err != nil {
		t.Errorf("Failed to sign message, err: %s", err)
	}

	valid1, err := Verify(&privKey1.PublicKey, sig, txHash)
	if err != nil {
		t.Errorf("Failed to verity message, err: %s", err)
	}

	valid2, err := Verify(&privKey2.PublicKey, sig, txHash)
	if err != nil {
		t.Errorf("Failed to verity message, err: %s", err)
	}

	if valid1 != true {
		t.Errorf("Failed in valid result 1")
	}

	if valid2 != false {
		t.Errorf("Failed in valid result 2")
	}
}
