package btc

type Opts struct {
	version     BlockchainVersion
	addrVersion AddressVersion
}

func NewOpts(v BlockchainVersion, a AddressVersion) Opts {
	return Opts{v, a}
}
