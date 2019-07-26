package btc

import (
	"testing"
)

type TestCase {
	Mnemonic string
	Path string
	Version BlockchainVersion
	AddrVersion AddressVersion
	ExpectErr bool
	ExpectedAddress string
	ExpectedPubKey string
	ExpectedPrivKey string
}

func PrepareTestCases() []TestCase {
	
}

func TestGenerateAddress(t *testing.T) {

}
