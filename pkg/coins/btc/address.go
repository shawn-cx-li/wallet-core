package btc

import (
	"github.com/shawn-cx-li/wallet-core/pkg/utils"

	base58x "github.com/shawn-cx-li/wallet-core/pkg/base58"
)

func (k *Key) newAddressPubKeyHash() (string, error) {
	pubKeyByte := k.PublicKeyBytes()

	accountID := utils.Sha256RipeMD160(pubKeyByte)
	h := append([]byte{Params[k.version].PubKeyHashAddrID}, accountID...)

	return base58x.Base58Encode(h, ALPHABET), nil
}

func (k *Key) newAddressScriptHash() (string, error) {
	pubKeyByte := k.PublicKeyBytes()

	accountID := utils.Sha256RipeMD160(pubKeyByte)

	scriptSig := append([]byte{0x00, 0x14}, accountID...)
	keyHash := utils.Sha256RipeMD160(scriptSig)

	h := append([]byte{Params[k.version].ScriptHashAddrID}, keyHash...)

	return base58x.Base58Encode(h, ALPHABET), nil
}
