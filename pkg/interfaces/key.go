package interfaces

// Key is the interface for all coins
type Key interface {
	Address() (string, error)
	PrivateKeyString() string
	PrivateKeyBytes() []byte
	PublicKeyString() string
	PublicKeyBytes() []byte
}
