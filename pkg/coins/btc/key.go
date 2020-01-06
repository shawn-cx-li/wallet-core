package btc

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
	mnemonic string
	path     string
	opts     Opts
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

	return &Key{
		privKey,
		mnemonic,
		path,
		opts,
	}, nil
}

func (k *Key) Address() (string, error) {
	switch k.opts.addrVersion {
	case BIP44:
		return k.newAddressPubKeyHash(k.PublicKeyBytes())
	case BIP49:
		return k.newAddressScriptHash(k.PublicKeyBytes())
	default:
		return "", fmt.Errorf("unexpected address version")
	}
}

// PrivateKeyString returns the Wallet Import Format (WIF)
func (k *Key) PrivateKeyString() string {
	return k.WifString()
}

func (k *Key) PrivateKeyBytes() []byte {
	return k.privateKey()
}

func (k *Key) PublicKeyString() string {
	return hex.EncodeToString(k.publicKey())
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
