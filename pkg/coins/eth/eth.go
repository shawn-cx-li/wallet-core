package eth

import (
	"encoding/hex"

	ethCrypto "github.com/ethereum/go-ethereum/crypto"
	log "github.com/sirupsen/logrus"

	"github.com/shawn-cx-li/wallet-core/pkg/crypto"
)

type Wallet struct {
	Address string
	PrivKey string
	PubKey  string
}

func GenerateAddress(mnemonic, path string) (addr string, err error) {
	seed, err := crypto.RecoverSeed(mnemonic, "")
	if err != nil {
		return
	}

	privKey, _, err := crypto.GetPrivateKey(seed, path)
	if err != nil {
		return
	}

	address := ethCrypto.PubkeyToAddress(privKey.PublicKey)
	addr = address.String()

	log.Info(address.String())
	log.Info(hex.EncodeToString(address[:]))
	return
}
