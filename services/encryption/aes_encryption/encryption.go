package aes_encryption

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"hash"
	"io"
)

type Encryption struct {
	Passphrase string
	Block      cipher.Block
	Stream     cipher.Stream
	IV         []byte
	Mac        hash.Hash
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

	var iv [aes.BlockSize]byte

	aesCipherStream := cipher.NewOFB(cryptoBlockCipher, iv[:])
	if err != nil {
		return nil, err
	}

	encryption := &Encryption{
		Passphrase: passphrase,
		Block:      cryptoBlockCipher,
		Stream:     aesCipherStream,
	}

	return encryption, nil
}

func (e *Encryption) Encrypt(in *io.PipeReader, out *io.PipeWriter) {
	encWriter := &cipher.StreamWriter{
		S: e.Stream,
		W: out,
	}

	_, err := io.Copy(encWriter, in)
	if err != nil {
		panic(err)
	}
}

func (e *Encryption) Decrypt(in *io.PipeReader, out *io.PipeWriter) {
	decReader := &cipher.StreamReader{
		S: e.Stream,
		R: in,
	}

	_, err := io.Copy(out, decReader)
	if err != nil {
		panic(err)
	}
}
