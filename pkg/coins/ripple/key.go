package ripple

import (
	"crypto/ecdsa"
	"encoding/hex"

	ethCrypto "github.com/ethereum/go-ethereum/crypto"
	// log "github.com/sirupsen/logrus"

	"github.com/shawn-cx-li/wallet-core/pkg/base58"
	"github.com/shawn-cx-li/wallet-core/pkg/crypto"
	"github.com/shawn-cx-li/wallet-core/pkg/utils"
)

// Key is the ripple key bag
type Key struct {
	*ecdsa.PrivateKey
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

	return &Key{privKey}, nil
}

func (k *Key) Address() (string, error) {
	pubKeyByte := k.PublicKeyBytes()

	accountID := utils.Sha256RipeMD160(pubKeyByte)
	h, err := newHash(accountID, RIPPLE_ACCOUNT_ID)
	if err != nil {
		return "", err
	}

	return base58.Encode(h, ALPHABET), nil
}

func (k *Key) PrivateKeyString() string {
	return hex.EncodeToString(k.privateKey())
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
