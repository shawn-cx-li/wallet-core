package base58

import (
	"math/big"

	"github.com/shawn-cx-li/wallet-core/pkg/utils"
)

var bigRadix = big.NewInt(58)
var bigZero = big.NewInt(0)

func Base58Encode(b []byte, alphabet string) string {
	checksum := utils.DoubleSha256(b)
	b = append(b, checksum[0:4]...)
	x := new(big.Int)
	x.SetBytes(b)

	answer := make([]byte, 0)
	for x.Cmp(bigZero) > 0 {
		mod := new(big.Int)
		x.DivMod(x, bigRadix, mod)
		answer = append(answer, alphabet[mod.Int64()])
	}

	// leading zero bytes
	for _, i := range b {
		if i != 0 {
			break
		}
		answer = append(answer, alphabet[0])
	}

	// reverse
	alen := len(answer)
	for i := 0; i < alen/2; i++ {
		answer[i], answer[alen-1-i] = answer[alen-1-i], answer[i]
	}

	return string(answer)
}
