package challenge

import (
	"strings"

	"github.com/tenlisboa/pulley-challenge/pkg/challenge/decryptors"
)

type Challenge struct {
	Challenger       string            `json:"challenger"`
	EncryptedPath    string            `json:"encrypted_path"`
	EncryptionMethod encryptionMethods `json:"encryption_method"`
	ExpiresIn        string            `json:"expires_in"`
	Hint             string            `json:"hint"`
	Instructions     string            `json:"instructions"`
	Level            int32             `json:"level"`
}

type encryptionMethods string

const (
	Base64    encryptionMethods = "encoded as base64"
	NonHex                      = "inserted some non-hex characters"
	Circular                    = "circularly rotated left by "
	XOR                         = "encrypted with XOR"
	Scrambled                   = "scrambled! original positions as base64 encoded messagepack: "
)

func Decode(em encryptionMethods, hash string) string {
	if strings.Contains(string(em), string(Base64)) {
		return decryptors.DecryptHex(hash)
	}

	if strings.Contains(string(em), string(NonHex)) {
		return decryptors.FilterNonHex(hash)	
	}

	if strings.Contains(string(em), string(Circular)) {
		rotatebystr, _ := strings.CutPrefix(string(em), string(Circular))
		return decryptors.DecryptCircular(hash, rotatebystr)
	}

	if strings.Contains(string(em), string(XOR)) {
		return decryptors.DecryptXOR(hash)		
	}

	if strings.Contains(string(em), string(Scrambled)) {
		messagepack, _ := strings.CutPrefix(string(em), string(Scrambled))
		return decryptors.DecryptMsgPack(hash, messagepack)
	}

	return hash
}

