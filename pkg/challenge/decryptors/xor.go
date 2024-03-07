package decryptors

import "encoding/hex"

func DecryptXOR(hash string) string {
	decoded, _ := hex.DecodeString(hash)
	dhash := decodeXor(decoded, []byte("secret"))
	encoded := hex.EncodeToString([]byte(dhash))

	return encoded
}

func decodeXor(input, key []byte) (output []byte) {
	output = make([]byte, len(input))
	for i := 0; i < len(input); i++ {
		output[i] = input[i] ^ key[i%len(key)]
	}
	return output
}
