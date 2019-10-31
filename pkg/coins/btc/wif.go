package btc

import (
	base58 "github.com/shawn-cx-li/wallet-core/pkg/base58"
)

const compressMagic byte = 0x01

// WifString creates the Wallet Import Format string encoding of a WIF structure.
// See DecodeWIF for a detailed breakdown of the format and requirements of
// a valid WIF string.
func (k *Key) WifString() string {
	// Precalculate size.  Maximum number of bytes before base58 encoding
	// is one byte for the network, 32 bytes of private key, possibly one
	// extra byte if the pubkey is to be compressed, and finally four
	// bytes of checksum.
	encodeLen := 1 + Params[k.opts.version].PrivKeyBytesLen + 4 + 1
	// if w.CompressPubKey {
	// 	encodeLen++
	// }

	a := make([]byte, 0, encodeLen)
	a = append(a, Params[k.opts.version].PrivateKeyID)
	// Pad and append bytes manually, instead of using Serialize, to
	// avoid another call to make.
	a = paddedAppend(Params[k.opts.version].PrivKeyBytesLen, a, k.PrivateKey.D.Bytes())
	// if w.CompressPubKey {
	a = append(a, compressMagic)
	// }
	// cksum := utils.DoubleSha256(a)[:4]
	// a = append(a, cksum...)
	return base58.Encode(a, ALPHABET)
}

// paddedAppend appends the src byte slice to dst, returning the new slice.
// If the length of the source is smaller than the passed size, leading zero
// bytes are appended to the dst slice before appending src.
func paddedAppend(size uint, dst, src []byte) []byte {
	for i := 0; i < int(size)-len(src); i++ {
		dst = append(dst, 0)
	}
	return append(dst, src...)
}
