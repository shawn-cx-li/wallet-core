## How to use

Create a new wallet with a given/new generated mnemonic and derivation path. Output private key and address to console.

1. Build cli `make build-cli`
2. generate mnemonic, 
    ```
    ./main generate-mnemonic
    ```
    sample output:

    ```
    mnemonic: abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon about
    ```
3. create wallet of private key and address for a given blockchain family
    ```
    ./main create --mnemonic="abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon about" --path="m/44'/0'/0'/0/0" --family="btc-bip44"
    ```
    optional config can be loaded from `./config/config.yaml`
    ```
    ./main create --config="./config/config.example.yaml"
    ```
    sample output:
    ```
    mnemonic: abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon about 
    path: m/44'/0'/0'/0/0                        
    privKey: L4p2b9VAf8k5aUahF1JCJUzZkgNEAqLfq8DDdQiyAprQAKSbu8hf 
    addr: 1LqBGSKuX5yYUonjxT5qGfpUsXKYYWeabA 
    ```
4. Use `1LqBGSKuX5yYUonjxT5qGfpUsXKYYWeabA` to receive your Bitcoin. To transfer out fund from the address, import private key `L4p2b9VAf8k5aUahF1JCJUzZkgNEAqLfq8DDdQiyAprQAKSbu8hf` to a third-party app like `Electrum`
