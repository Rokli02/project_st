package utils

type Encrypter struct {
	Secret string
}

func (e *Encrypter) Encrypt(text string) string {

	return text
}

func (e *Encrypter) Decrypt(scramble string) string {

	return scramble
}
