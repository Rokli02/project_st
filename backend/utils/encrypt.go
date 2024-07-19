package utils

import (
	"crypto/sha256"
	"hash"
)

type Encrypter struct {
	secret    string
	newHasher func() hash.Hash
}

func NewEncrypter(secret string) *Encrypter {
	e := &Encrypter{
		secret:    secret,
		newHasher: sha256.New,
	}

	return e
}

func (e *Encrypter) Encrypt(text string) string {

	return text
}

func (e *Encrypter) Decrypt(scramble string) string {

	return scramble
}

func (e *Encrypter) Hash(text string) string {
	hasher := e.newHasher()
	hasher.Write([]byte(text))
	return string(hasher.Sum([]byte(e.secret)))
}
