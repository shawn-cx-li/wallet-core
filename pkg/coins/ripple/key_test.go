package ripple

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
	ExpectErr       bool
	ExpectedAddress string
	ExpectedPubKey  string
	ExpectedPrivKey string
}

func prepareTestCases(t *testing.T) []testCase {
	test := []testCase{
		{
			"ripple mainnet bip44",
			Mnemonic,
			"m/44'/144'/0'/0/0",
			false,
			"rHsMGQEkVNJmpGWs8XUBoTBiAAbwxZN5v3",
			"031d68bc1a142e6766b2bdfb006ccfe135ef2e0e2e94abb5cf5c9ab6104776fbae",
			"90802a50aa84efb6cdb225f17c27616ea94048c179142fecf03f4712a07ea7a4",
		},
	}
	return test
}

func TestNewKey(t *testing.T) {
	tests := prepareTestCases(t)

	for _, c := range tests {
		key, err := NewKey(c.Mnemonic, c.Path, Opts{})
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
