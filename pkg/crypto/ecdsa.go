package crypto

import (
	"crypto/ecdsa"
	"encoding/hex"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/ethereum/go-ethereum/accounts"
	bip39 "github.com/tyler-smith/go-bip39"
)

// GetNewMnemonic Generates a random mnemonic
// A mnemonic can be used to create/recover seed
func GetNewMnemonic() (mnemonic string, err error) {
	entropy, err := bip39.NewEntropy(128)
	if err != nil {
		return "", err
	}

	mnemonic, err = bip39.NewMnemonic(entropy)
	return
}

// GetNewSeed generates a random mnemonic and the deterministic seed with specified password
func GetNewSeed(pwd string) (mnemonic, seed string, err error) {
	mnemonic, err = GetNewMnemonic()
	if err != nil {
		return
	}

	seedBytes := bip39.NewSeed(mnemonic, pwd)
	seed = hex.EncodeToString(seedBytes)
	return
}

// RecoverSeed recovers the deterministic seed with specified mnemonic and password
func RecoverSeed(mnemonic, pwd string) (seed string, err error) {
	seedBytes := bip39.NewSeed(mnemonic, pwd)
	seed = hex.EncodeToString(seedBytes)
	return
}

// GetPrivateKey derives ECDSA private key with seed and path
// path e.g. "m/44'/60'/0'/0/1"
func GetPrivateKey(seed string, path string) (*ecdsa.PrivateKey, string, error) {
	// Generate a new master node using the seed.
	seedBytes, err := hex.DecodeString(seed)
	if err != nil {
		return nil, "", err
	}
	masterKey, err := hdkeychain.NewMaster(seedBytes, &chaincfg.MainNetParams)
	if err != nil {
		return nil, "", err
	}

	// Derive the key path
	keyPath, err := accounts.ParseDerivationPath(path)
	if err != nil {
		return nil, "", err
	}

	// Get key by iterating the key path
	for _, n := range keyPath {
		masterKey, err = masterKey.Child(n)
		if err != nil {
			return nil, "", err
		}
	}

	privateKey, err := masterKey.ECPrivKey()
	if err != nil {
		return nil, "", err
	}

	privateKeyECDSA := privateKey.ToECDSA()

	return privateKeyECDSA, path, nil
}
