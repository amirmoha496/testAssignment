package main

import (
	"crypto/sha512"
	b64 "encoding/base64"
)

//CryptoManager class provides methids for Crypto Operations
type CryptoManager struct {
}

//SHA2Hash It computes the SHA256 hash for the given input and returns a Base 64 encoded hashed value for input
//data -> Input for which hash to be computed
func (cm CryptoManager) SHA2Hash(data string) string {
	sha2 := sha512.Sum512([]byte(data))
	enc := b64.StdEncoding.EncodeToString(sha2[:])
	return enc
}
