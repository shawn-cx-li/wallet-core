package eth

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"

	ethCrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/shawn-cx-li/wallet-core/pkg/crypto"
	"github.com/shawn-cx-li/wallet-core/pkg/interfaces"
)

type Key struct {
	*ecdsa.PrivateKey
	opts Opts
}

func NewKey(mnemonic, path string, opts Opts) (interfaces.Key, error) {
	return newKey(mnemonic, path, opts)
}

func newKey(mnemonic, path string, opts Opts) (*Key, error) {
	seed, err := crypto.RecoverSeed(mnemonic, "")
	if err != nil {
		return nil, err
	}
	privKey, _, err := crypto.GetPrivateKey(seed, path)
	if err != nil {
		return nil, err
	}
	if ValidateOpts(opts) != nil {
		return nil, ValidateOpts(opts)
	}

	return &Key{
		privKey,
		opts,
	}, nil
}

// ImportKey converts a secret string to a Key
func ImportKey(privKey string, opts Opts) (interfaces.Key, error) {
	return importKey(privKey, opts)
}

func importKey(privKey string, opts Opts) (*Key, error) {
	if hasHexPrefix(privKey) {
		privKey = privKey[2:]
	}
	if !isHex(privKey) {
		return nil, fmt.Errorf("invalid private key %s", privKey)
	}
	if ValidateOpts(opts) != nil {
		return nil, ValidateOpts(opts)
	}
	key, err := crypto.StringToECDSA(privKey)
	if err != nil {
		return nil, err
	}

	return &Key{
		key,
		opts,
	}, nil
}

func (k *Key) Address() (string, error) {
	return ethCrypto.PubkeyToAddress(k.PublicKey).String(), nil
}

// PrivateKeyString returns the Wallet Import Format (WIF)
func (k *Key) PrivateKeyString() string {
	return "0x" + hex.EncodeToString(k.privateKey())
}

func (k *Key) PrivateKeyBytes() []byte {
	return k.privateKey()
}

func (k *Key) PublicKeyString() string {
	return "0x" + hex.EncodeToString(k.publicKey())
}

func (k *Key) PublicKeyBytes() []byte {
	return k.publicKey()
}

func (k *Key) privateKey() []byte {
	return ethCrypto.FromECDSA(k.PrivateKey)
}

func (k *Key) publicKey() []byte {
	return ethCrypto.CompressPubkey(&k.PublicKey)
}
