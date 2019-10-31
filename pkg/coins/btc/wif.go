package btc

import (
	"errors"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	base58 "github.com/shawn-cx-li/wallet-core/pkg/base58"
)

// WifString creates the Wallet Import Format string encoding of a WIF structure.
// See DecodeWIF for a detailed breakdown of the format and requirements of
// a valid WIF string.
func (k *Key) WifString() string {
	// Precalculate size.  Maximum number of bytes before base58 encoding
	// is one byte for the network, 32 bytes of private key, possibly one
	// extra byte if the pubkey is to be compressed, and finally four
	// bytes of checksum.
	encodeLen := 1 + btcec.PrivKeyBytesLen + 4 + 1
	// if w.CompressPubKey {
	// 	encodeLen++
	// }

	a := make([]byte, 0, encodeLen)
	a = append(a, Params[k.opts.version].PrivateKeyID)
	// Pad and append bytes manually, instead of using Serialize, to
	// avoid another call to make.
	a = paddedAppend(btcec.PrivKeyBytesLen, a, k.PrivateKey.D.Bytes())
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

type WIF struct {
	// PrivKey is the private key being imported or exported.
	PrivKey *btcec.PrivateKey

	// CompressPubKey specifies whether the address controlled by the
	// imported or exported private key was created by hashing a
	// compressed (33-byte) serialized public key, rather than an
	// uncompressed (65-byte) one.
	CompressPubKey bool

	// netID is the bitcoin network identifier byte used when
	// WIF encoding the private key.
	netID byte
}

const compressMagic byte = 0x01

func NewWIF(privKey *btcec.PrivateKey, net *chaincfg.Params, compress bool) (*WIF, error) {
	if net == nil {
		return nil, errors.New("no network")
	}
	return &WIF{privKey, compress, net.PrivateKeyID}, nil
}

// String creates the Wallet Import Format string encoding of a WIF structure.
// See DecodeWIF for a detailed breakdown of the format and requirements of
// a valid WIF string.
func (w *WIF) String() string {
	// Precalculate size.  Maximum number of bytes before base58 encoding
	// is one byte for the network, 32 bytes of private key, possibly one
	// extra byte if the pubkey is to be compressed, and finally four
	// bytes of checksum.
	encodeLen := 1 + btcec.PrivKeyBytesLen + 4
	if w.CompressPubKey {
		encodeLen++
	}

	a := make([]byte, 0, encodeLen)
	a = append(a, w.netID)
	// Pad and append bytes manually, instead of using Serialize, to
	// avoid another call to make.
	a = paddedAppend(btcec.PrivKeyBytesLen, a, w.PrivKey.D.Bytes())
	if w.CompressPubKey {
		a = append(a, compressMagic)
	}
	cksum := chainhash.DoubleHashB(a)[:4]
	a = append(a, cksum...)
	return base58.Encode(a, ALPHABET)
}
