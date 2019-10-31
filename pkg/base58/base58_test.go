package base58

import (
	"encoding/hex"
	"testing"
)

type checksumCase struct {
	input    string
	output   string
	alphabet string
}

const btcAlphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

var tests = []checksumCase{
	{
		"e284129cc0922579a535bbf4d1a3b25773090d28c909bc0fed73b5e0222cc372",
		"2im3byRJKgkvJCBka2ZVNGgdGKyG1pignCDv5Pd6ZwSdVNiMvZ",
		btcAlphabet,
	},
}

func TestEncode(t *testing.T) {
	for i := range tests {
		c := tests[i]
		inputBytes, err := hex.DecodeString(c.input)
		// a := make([]byte, 0)
		if err != nil {
			t.Errorf("invalid base58 input %s", c.input)
		}
		o := Encode(inputBytes, c.alphabet)
		if o != c.output {
			t.Errorf("failed to convert checksum %s. expect %s, got %s", c.input, c.output, o)
		}
	}
}

func TestDecode(t *testing.T) {
	for i := range tests {
		c := tests[i]

		outputBytes, err := Decode(c.output, c.alphabet)
		if err != nil {
			t.Errorf("cannot decode base58 output %s", c.output)
		}

		o := hex.EncodeToString(outputBytes[:len(outputBytes)-4])
		if o != c.input {
			t.Errorf("failed to decode %s. expect %s, got %s", c.output, c.input, o)
		}
	}
}
