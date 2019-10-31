package btc

import (
	"github.com/shawn-cx-li/wallet-core/pkg/utils"

	base58 "github.com/shawn-cx-li/wallet-core/pkg/base58"
)

func (k *Key) newAddressPubKeyHash(pubKeyByte []byte) (string, error) {

	accountID := utils.Sha256RipeMD160(pubKeyByte)
	h := append([]byte{Params[k.opts.version].PubKeyHashAddrID}, accountID...)

	return base58.Encode(h, ALPHABET), nil
}

func (k *Key) newAddressScriptHash(pubKeyByte []byte) (string, error) {

	accountID := utils.Sha256RipeMD160(pubKeyByte)

	scriptSig := append([]byte{0x00, 0x14}, accountID...)
	keyHash := utils.Sha256RipeMD160(scriptSig)

	h := append([]byte{Params[k.opts.version].ScriptHashAddrID}, keyHash...)

	return base58.Encode(h, ALPHABET), nil
}
