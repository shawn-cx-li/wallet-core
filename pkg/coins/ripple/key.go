package ripple

import (
	"crypto/ecdsa"

	ethCrypto "github.com/ethereum/go-ethereum/crypto"
	// log "github.com/sirupsen/logrus"

	"github.com/shawn-cx-li/wallet-core/pkg/base58"
	"github.com/shawn-cx-li/wallet-core/pkg/crypto"
	"github.com/shawn-cx-li/wallet-core/pkg/key"
	"github.com/shawn-cx-li/wallet-core/pkg/utils"
)

type Key struct {
	*ecdsa.PrivateKey
}

func NewRippleKey(mnemonic, path string) (key.Key, error) {
	seed, err := crypto.RecoverSeed(mnemonic, "")
	if err != nil {
		return nil, err
	}
	privKey, _, err := crypto.GetPrivateKey(seed, path)
	if err != nil {
		return nil, err
	}

	return &Key{privKey}, nil
}

func (k *Key) Address() (string, error) {
	pubKeyByte := k.PublicKeyBytes()

	accountID := utils.Sha256RipeMD160(pubKeyByte)
	h, err := newHash(accountID, RIPPLE_ACCOUNT_ID)
	if err != nil {
		return "", err
	}

	return base58.Base58Encode(h, ALPHABET), nil
}

func (k *Key) PrivateKeyString() (string, error) { return "", nil }
func (k *Key) PrivateKeyBytes() ([]byte, error)  { return nil, nil }
func (k *Key) PublicKeyString() (string, error)  { return "", nil }

func (k *Key) PublicKeyBytes() []byte {
	pubKey := k.PublicKey
	return ethCrypto.CompressPubkey(&pubKey)
}
