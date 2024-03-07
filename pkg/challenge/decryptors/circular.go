package decryptors

import (
	"fmt"
	"strconv"
)

func DecryptCircular(hash, rotatebystr string) string {
	rotateby, err := strconv.Atoi(rotatebystr)
	if err != nil {
		panic(fmt.Sprintf("Error in converting string to integer: %s\n", err.Error()))
	}

	if len(hash) == 0 || rotateby == 0 {
		fmt.Println(hash)
		return hash
	}

	hlength := len(hash)
	rotationidx := hlength - (rotateby % hlength)
	return hash[rotationidx:] + hash[:rotationidx]
}
