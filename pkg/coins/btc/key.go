package btc

import (
	"crypto/ecdsa"
	"fmt"

	ethCrypto "github.com/ethereum/go-ethereum/crypto"

	"github.com/shawn-cx-li/wallet-core/pkg/crypto"
	"github.com/shawn-cx-li/wallet-core/pkg/key"
)

type Key struct {
	*ecdsa.PrivateKey
	version     BlockchainVersion
	addrVersion AddressVersion
	mnemonic    string
	path        string
}

func NewKey(mnemonic, path string, version BlockchainVersion, addrVersion AddressVersion) (key.Key, error) {
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
		version,
		addrVersion,
		mnemonic,
		path,
	}, nil
}

func (k *Key) Address() (string, error) {
	switch k.addrVersion {
	case BIP44:
		return k.newAddressPubKeyHash()
	case BIP49:
		return k.newAddressScriptHash()
	default:
		return "", fmt.Errorf("unexpected address version")
	}
}

func (k *Key) PrivateKeyString() (string, error) { return "", nil }
func (k *Key) PrivateKeyBytes() ([]byte, error)  { return nil, nil }
func (k *Key) PublicKeyString() (string, error)  { return "", nil }

func (k *Key) PublicKeyBytes() []byte {
	pubKey := k.PublicKey
	return ethCrypto.CompressPubkey(&pubKey)
}
