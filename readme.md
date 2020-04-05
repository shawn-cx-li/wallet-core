### Description
This project is aimed to provide a universal HD wallet derivation. Use a single mnemonic to generate private/public key pair and address for various Blockchains.

### Examples
Examples can be found in [examples](https://github.com/shawn-cx-li/wallet-core/blob/master/examples/demo.go)

### Command Line Tool
Find doc of cli in [cli](https://github.com/shawn-cx-li/wallet-core/blob/master/cli/readme.md)

### Sample Result
To derive private key, public key and address from mnemonic `abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon about` for bitcoin BIP44 specifics

```
key, _ := btc.NewKey(mnemonic, btcBIP44Path, btc.NewOpts(btc.BITCOIN, btc.BIP44))
privKey := key.PrivateKeyString()
pubKey := key.PublicKeyString()
addr, _ := key.Address()
log.Info(privKey)   // L4p2b9VAf8k5aUahF1JCJUzZkgNEAqLfq8DDdQiyAprQAKSbu8hf
log.Info(pubKey)    // 03aaeb52dd7494c361049de67cc680e83ebcbbbdbeb13637d92cd845f70308af5e
log.Info(addr)      // 1LqBGSKuX5yYUonjxT5qGfpUsXKYYWeabA
```
