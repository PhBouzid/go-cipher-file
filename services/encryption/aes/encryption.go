package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

type Encryption struct {
	Passphrase string
	Cipher     cipher.AEAD
}

const AES256 = 32

func New(passphrase string) (*Encryption, error) {

	passphraseBytes := []byte(passphrase)

	if len(passphraseBytes) < AES256 {
		var padd0x00 byte = 0
		padBytes := bytes.Repeat([]byte{padd0x00}, AES256-len(passphraseBytes))
		passphraseBytes = append(passphraseBytes, padBytes...)
	}

	passphraseBytes = passphraseBytes[:AES256]

	cryptoBlockCipher, err := aes.NewCipher(passphraseBytes)
	if err != nil {
		return nil, err
	}

	aesCipher, err := cipher.NewGCM(cryptoBlockCipher)
	if err != nil {
		return nil, err
	}

	encryption := &Encryption{
		Passphrase: passphrase,
		Cipher:     aesCipher,
	}

	return encryption, nil
}

func (e *Encryption) Encrypt(data <-chan []byte, output chan<- []byte) {

}

func (e *Encryption) Decrypt(data <-chan []byte, output chan<- []byte) {

}
