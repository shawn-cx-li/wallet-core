package btc

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"

	ethCrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/shawn-cx-li/wallet-core/pkg/crypto"
)

type Key struct {
	*ecdsa.PrivateKey
	mnemonic string
	path     string
	opts     Opts
}

func NewKey(mnemonic, path string, opts Opts) (*Key, error) {
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

func (k *Key) getPrivateKey() []byte {
	return ethCrypto.FromECDSA(k.PrivateKey)
}

// PrivateKeyString returns the Wallet Import Format (WIF)
func (k *Key) PrivateKeyString() (string, error) {

	return k.WifString(), nil
}
func (k *Key) PrivateKeyBytes() ([]byte, error) {
	return k.getPrivateKey(), nil
}

func (k *Key) getPublicKey() []byte {
	pubKey := k.PublicKey
	return ethCrypto.CompressPubkey(&pubKey)
}

func (k *Key) PublicKeyString() (string, error) {
	return hex.EncodeToString(k.getPublicKey()), nil
}

func (k *Key) PublicKeyBytes() []byte {
	return k.getPublicKey()
}
