package btc

import (
	"testing"
)

const (
	Mnemonic = "abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon about"
)

type testCase struct {
	Name            string
	Mnemonic        string
	Path            string
	Version         BlockchainVersion
	AddrVersion     AddressVersion
	ExpectErr       bool
	ExpectedAddress string
	ExpectedPubKey  string
	ExpectedPrivKey string
}

func PrepareTestCases(t *testing.T) []testCase {
	test := []testCase{
		{
			"bitcoin mainnet bip44",
			Mnemonic,
			"m/44'/0'/0'/0/0",
			BITCOIN,
			BIP44,
			false,
			"1LqBGSKuX5yYUonjxT5qGfpUsXKYYWeabA",
			"03aaeb52dd7494c361049de67cc680e83ebcbbbdbeb13637d92cd845f70308af5e",
			"L4p2b9VAf8k5aUahF1JCJUzZkgNEAqLfq8DDdQiyAprQAKSbu8hf",
		},
		{
			"bitcoin mainnet bip49",
			Mnemonic,
			"m/49'/0'/0'/0/0",
			BITCOIN,
			BIP49,
			false,
			"37VucYSaXLCAsxYyAPfbSi9eh4iEcbShgf",
			"039b3b694b8fc5b5e07fb069c783cac754f5d38c3e08bed1960e31fdb1dda35c24",
			"KyvHbRLNXfXaHuZb3QRaeqA5wovkjg4RuUpFGCxdH5UWc1Foih9o",
		},
		{
			"bitcoin testnet bip44",
			Mnemonic,
			"m/44'/1'/0'/0/0",
			BITCOIN_TESTNET,
			BIP44,
			false,
			"mkpZhYtJu2r87Js3pDiWJDmPte2NRZ8bJV",
			"02a7451395735369f2ecdfc829c0f774e88ef1303dfe5b2f04dbaab30a535dfdd6",
			"cV6NTLu255SZ5iCNkVHezNGDH5qv6CanJpgBPqYgJU13NNKJhRs1",
		},
		{
			"bitcoin testnet bip49",
			Mnemonic,
			"m/49'/1'/0'/0/0",
			BITCOIN_TESTNET,
			BIP49,
			false,
			"2Mww8dCYPUpKHofjgcXcBCEGmniw9CoaiD2",
			"03a1af804ac108a8a51782198c2d034b28bf90c8803f5a53f76276fa69a4eae77f",
			"cULrpoZGXiuC19Uhvykx7NugygA3k86b3hmdCeyvHYQZSxojGyXJ",
		},
		{
			"dash mainnet",
			Mnemonic,
			"m/44'/5'/0'/0/0",
			DASH,
			BIP44,
			false,
			"XoJA8qE3N2Y3jMLEtZ3vcN42qseZ8LvFf5",
			"026fa9a6f213b6ba86447965f6b4821264aaadd7521f049f00db9c43a770ea7405",
			"XGihgbi7c1nVqrjkPSvzJydLVWYW7hTrcXdfSdpFMwi3Xhbabw93",
		},
	}
	return test
}

func TestNewKey(t *testing.T) {
	tests := PrepareTestCases(t)

	for _, c := range tests {
		opts := Opts{c.Version, c.AddrVersion}
		key, err := NewKey(c.Mnemonic, c.Path, opts)
		if err != nil {
			if !c.ExpectErr {
				t.Errorf("unexpected case %s, err: %s", c.Name, err)
			}
			continue
		}

		GetAddressTest(t, &c, key)
		GetPublicKeyTest(t, &c, key)
		GetPrivateKeyTest(t, &c, key)
	}
}

func GetAddressTest(t *testing.T, c *testCase, k *Key) {
	addr, err := k.Address()
	if err != nil {
		t.Errorf("unexpected case %s on Address(), err: %s ", c.Name, err)
	}
	if addr != c.ExpectedAddress {
		t.Errorf("unexpected case %s, address: %s, expect: %s", c.Name, addr, c.ExpectedAddress)
	}
}

func GetPublicKeyTest(t *testing.T, c *testCase, k *Key) {
	pubKey := k.PublicKeyString()

	if pubKey != c.ExpectedPubKey {
		t.Errorf("unexpected case %s, pubKey: %s, expect: %s", c.Name, pubKey, c.ExpectedPubKey)
	}
}

func GetPrivateKeyTest(t *testing.T, c *testCase, k *Key) {
	privKey := k.PrivateKeyString()

	if privKey != c.ExpectedPrivKey {
		t.Errorf("unexpected case %s, privKey: %s, expect: %s", c.Name, privKey, c.ExpectedPrivKey)
	}
}
