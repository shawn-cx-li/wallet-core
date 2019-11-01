package eth

import (
	"crypto/ecdsa"
	"encoding/hex"

	ethCrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/shawn-cx-li/wallet-core/pkg/crypto"
)

type Key struct {
	*ecdsa.PrivateKey
	mnemonic string
	path     string
	opts     Opts
}

type Opts struct{}

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
	return ethCrypto.PubkeyToAddress(k.PublicKey).String(), nil
}

// PrivateKeyString returns the Wallet Import Format (WIF)
func (k *Key) PrivateKeyString() (string, error) {
	return "0x" + hex.EncodeToString(k.getPrivateKey()), nil
}

func (k *Key) PrivateKeyBytes() ([]byte, error) {
	return k.getPrivateKey(), nil
}

func (k *Key) PublicKeyString() (string, error) {
	return "0x" + hex.EncodeToString(k.getPublicKey()), nil
}

func (k *Key) PublicKeyBytes() []byte {
	return k.getPublicKey()
}

func (k *Key) getPrivateKey() []byte {
	return ethCrypto.FromECDSA(k.PrivateKey)
}

func (k *Key) getPublicKey() []byte {
	return ethCrypto.CompressPubkey(&k.PublicKey)
}
