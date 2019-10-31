package main

import (
	"github.com/shawn-cx-li/wallet-core/pkg/coins/bch"
)

func main() {
	// demo_btc()
	// demo_ripple()
	bchD()
}

func bchD() {
	bch.GenerateAddress("abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon about", "m/44'/145'/0'/0/0")
}
