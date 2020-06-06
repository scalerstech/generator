package utils

import (
	"errors"
	"github.com/schollz/cowyo/encrypt"
)

type EncDec struct {
	u *Utils
}

func (ed *EncDec) Encrypt(toEncrypt, password string) (string, error) {
	return encrypt.EncryptString(toEncrypt, password)
}

func (ed *EncDec) Decrypt(toDecrypt, password string) (string, error) {
	if len(toDecrypt) > 0 {
		return encrypt.DecryptString(toDecrypt, password)
	} else {
		return "", errors.New("EncDec.Decrypt: empty string provided. will not decrypt")
	}
}
