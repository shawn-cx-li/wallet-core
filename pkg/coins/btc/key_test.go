package btc

import (
	"testing"
)

const (
	Mnemonic = "abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon about"
)

type TestCase struct {
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

func PrepareTestCases(t *testing.T) []TestCase {
	test := []TestCase{
		{
			"bitcoin mainnet bip44",
			Mnemonic,
			"m/44'/0'/0'/0/0",
			BITCOIN,
			BIP44,
			false,
			"1LqBGSKuX5yYUonjxT5qGfpUsXKYYWeabA",
			"",
			"",
		},
		{
			"bitcoin mainnet bip49",
			Mnemonic,
			"m/49'/0'/0'/0/0",
			BITCOIN,
			BIP49,
			false,
			"37VucYSaXLCAsxYyAPfbSi9eh4iEcbShgf",
			"",
			"",
		},
		{
			"bitcoin testnet bip44",
			Mnemonic,
			"m/44'/1'/0'/0/0",
			BITCOIN_TESTNET,
			BIP44,
			false,
			"mkpZhYtJu2r87Js3pDiWJDmPte2NRZ8bJV",
			"",
			"",
		},
		{
			"bitcoin testnet bip49",
			Mnemonic,
			"m/49'/1'/0'/0/0",
			BITCOIN_TESTNET,
			BIP49,
			false,
			"2Mww8dCYPUpKHofjgcXcBCEGmniw9CoaiD2",
			"",
			"",
		},
		{
			"dash mainnet",
			Mnemonic,
			"m/44'/5'/0'/0/0",
			DASH,
			BIP44,
			false,
			"XoJA8qE3N2Y3jMLEtZ3vcN42qseZ8LvFf5",
			"",
			"",
		},
	}
	return test
}

func TestGenerateAddress(t *testing.T) {
	tests := PrepareTestCases(t)

	for _, c := range tests {
		key, err := NewKey(c.Mnemonic, c.Path, c.Version, c.AddrVersion)
		if err != nil {
			if !c.ExpectErr {
				t.Errorf("unexpected case %s, err: %s", c.Name, err)
			}
			continue
		}

		addr, err := key.Address()
		if err != nil {
			t.Errorf("unexpected case %s, err: %s ", c.Name, err)
		}
		if addr != c.ExpectedAddress {
			t.Errorf("unexpected case %s, address %s, expect %s", c.Name, addr, c.ExpectedAddress)
		}
	}
}
