package main

func main() {
	conf, err := getConfig()
	if err != nil {
		panic(err)
	}

	wallet := NewWallet(conf.Wallet.Mnemonic, conf.Wallet.Path)
	wallet.printKeys(conf.Family)
}
