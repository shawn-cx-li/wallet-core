package interfaces

// Key is the interface for all coins
type Key interface {
	Address() (string, error)
	PrivateKeyString() (string, error)
	PrivateKeyBytes() ([]byte, error)
	PublicKeyString() (string, error)
	PublicKeyBytes() []byte
}