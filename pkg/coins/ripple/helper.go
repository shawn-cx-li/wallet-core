package ripple

import (
	"fmt"
)

func newHash(b []byte, version HashVersion) ([]byte, error) {
	n := hashTypes[version].Payload
	if len(b) > n {
		return nil, fmt.Errorf("Hash is wrong size, expected: %d got: %d", n, len(b))
	}

	return append([]byte{byte(RIPPLE_ACCOUNT_ID)}, b...), nil
}
