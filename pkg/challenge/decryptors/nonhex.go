package decryptors

import "strings"

var hexchars = map[rune]bool{
	'0': true,
	'1': true,
	'2': true,
	'3': true,
	'4': true,
	'5': true,
	'6': true,
	'7': true,
	'8': true,
	'9': true,
	'a': true,
	'b': true,
	'c': true,
	'd': true,
	'e': true,
	'f': true,
}

func FilterNonHex(hash string) string {
	s := strings.Map(func(r rune) rune {
		if !hexchars[r] {
			return '_'
		}

		return r
	}, hash)

	return strings.ReplaceAll(s, "_", "")
}
