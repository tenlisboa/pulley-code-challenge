package decryptors

import (
	"encoding/base64"
	"fmt"
)

func DecryptHex(hash string) string {
	dhash, err := base64.StdEncoding.DecodeString(hash)
	if err != nil {
		panic(fmt.Sprintf("Error in decoding base64: %s", err.Error()))
	}

	return string(dhash)
}
