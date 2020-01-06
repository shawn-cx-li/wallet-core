package eth

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
			"bitcoin mainnet bip44",
			Mnemonic,
			"m/44'/60'/0'/0/0",
			false,
			"0x9858EfFD232B4033E47d90003D41EC34EcaEda94",
			"0x0237b0bb7a8288d38ed49a524b5dc98cff3eb5ca824c9f9dc0dfdb3d9cd600f299",
			"0x1ab42cc412b618bdea3a599e3c9bae199ebf030895b039e9db1e30dafb12b727",
		},
	}
	return test
}

func TestNewKey(t *testing.T) {
	tests := prepareTestCases(t)

	for _, c := range tests {
		opts := Opts{}
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
		ImportKeyTest(t, &c, key)
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

func ImportKeyTest(t *testing.T, c *testCase, k *Key) {
	importedKey, err := ImportKey(k.PrivateKeyString(), Opts{})
	if err != nil {
		t.Errorf("unable to import key from %s", k.PrivateKeyString())
	}

	if importedKey.PrivateKeyString() != k.PrivateKeyString() {
		t.Errorf("unexpected private key from imported secret, expect %s, got %s", k.PrivateKeyString(), importedKey.PrivateKeyString())
	}

	if importedKey.PublicKeyString() != k.PublicKeyString() {
		t.Errorf("unexpected public key from imported secret, expect %s, got %s", k.PublicKeyString(), importedKey.PublicKeyString())
	}
}
