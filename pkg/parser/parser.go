package parser

import (
	"encoding/json"
	"fmt"
)

type SomeStruct interface{}

func DecodeTo[Type SomeStruct](body []byte) Type {
	var data Type
	if err := json.Unmarshal(body, &data); err != nil {
		panic(fmt.Sprintf("Error decoding json: %s", err.Error()))
	}

	return data
}
