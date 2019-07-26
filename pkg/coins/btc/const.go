package btc

type BlockchainVersion byte
type AddressVersion byte

const (
	ALPHABET = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

	BITCOIN         BlockchainVersion = 0
	BITCOIN_TESTNET BlockchainVersion = 1
	DASH            BlockchainVersion = 2

	BIP44 AddressVersion = 0
	BIP49 AddressVersion = 1
)

var Params = [...]struct {
	Name             string
	Description      string
	PubKeyHashAddrID byte
	ScriptHashAddrID byte
	PrivateKeyID     byte
	HDPrivateKeyID   [4]byte
	HDPublicKeyID    [4]byte
}{
	BITCOIN: {
		Name:             "bitcoin",
		Description:      "bitcoin",
		PubKeyHashAddrID: 0x00,                            // starts with 1
		ScriptHashAddrID: 0x05,                            // starts with 3
		PrivateKeyID:     0x80,                            // starts with 5 (uncompressed) or K (compressed)
		HDPrivateKeyID:   [4]byte{0x04, 0x88, 0xad, 0xe4}, // starts with xprv
		HDPublicKeyID:    [4]byte{0x04, 0x88, 0xb2, 0x1e}, // starts with xpub
	},
	BITCOIN_TESTNET: {
		Name:             "bitcoin_testnet",
		Description:      "bitcoin testnet",
		PubKeyHashAddrID: 0x6f, // starts with m or n
		ScriptHashAddrID: 0xc4, // starts with 2
		PrivateKeyID:     0xef,
		HDPrivateKeyID:   [4]byte{0x04, 0x35, 0x83, 0x94},
		HDPublicKeyID:    [4]byte{0x04, 0x35, 0x87, 0xcf},
	},
	DASH: {
		Name:             "dash",
		Description:      "dash",
		PubKeyHashAddrID: 0x4c,
		ScriptHashAddrID: 0x10,
		PrivateKeyID:     0xcc,
		HDPrivateKeyID:   [4]byte{0x04, 0x88, 0xb2, 0x1e},
		HDPublicKeyID:    [4]byte{0x04, 0x88, 0xad, 0xe4},
	},
}
