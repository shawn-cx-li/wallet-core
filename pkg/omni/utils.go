package main

import (
	"fmt"

	"github.com/btcsuite/btcutil"
	log "github.com/sirupsen/logrus"
)

func paddingHex(h string) (string, error) {
	if len(h) < 16 {
		z := ""
		for i := len(h); i < 16; i++ {
			z += "0"
		}
		return z + h, nil
	} else {
		return "", fmt.Errorf("Invalid hex %s", h)
	}
}

func decimalToHex(f float64) string {
	return fmt.Sprintf("%x", int64(f))
}

func amountToSatoshi(amt float64) (satoshi float64, err error) {
	amount, err := btcutil.NewAmount(amt)
	if err != nil {
		log.Error("failed to convert ", amount, " into BTC Amount")
		return 0, nil
	}

	return amount.ToUnit(btcutil.AmountSatoshi), nil
}

// https://jochen-hoenicke.de/crypto/omni/
// The raw transaction includes two outputs:
// 1. Omni Script: "OP_RETURN 6f6d6e69000000000000001f000000002b752ee0, 0"
// 2. Receiptent with minimum btc: "addr, 0.00000546"
func conductRawTransaction(to, amt string) string {
	// 1. omni script
	protocol := "6f6d6e69" // omni
	function := "00000000" // simple send
	property := "0000001f" // 31 -> Thether

	script := protocol + function + property + amt + ", 0"

	// 2. receiptent
	minBTC := "0.00000546"
	receiptent := to + ", " + minBTC

	return "OP_RETURN" + " " + script + "\n" + receiptent
}
