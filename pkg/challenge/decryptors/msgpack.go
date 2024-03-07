package decryptors

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/johejo/msgpb64"
)

func DecryptMsgPack(hash, msgpack string) string {
	var original []int
	err := msgpb64.NewDecoder(base64.StdEncoding, strings.NewReader(msgpack)).Decode(&original)

	if err != nil {
		panic(fmt.Sprintf("Error in unmarshalling messagepack: %s", err.Error()))
	}

	rhash := []rune(hash)
	nhash := make([]rune, len(hash))
	for ap, op := range original {
		nhash[op] = rhash[ap]
	}

	return string(nhash)
}
