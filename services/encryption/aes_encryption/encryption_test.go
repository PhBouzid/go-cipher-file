package aes_encryption

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

func TestEncryption_Encrypt(t *testing.T) {
	aesEnc, err := New("Lorem ipsum dolor sit amet, consectetuer adipiscing elit." +
		" Aenean commodo ligula eget dolor. Aenean massa. " +
		"Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus." +
		" Donec quam felis, ultricies nec, pellentesque eu, pretium quis,.")
	assert.NoError(t, err)
	pipeFileReader, pipeFileWriter := io.Pipe()
	pipeEncReader, pipeEncWriter := io.Pipe()

	go func() {
		f, fErr := os.Open("/Users/philippbouzid/books/information_theory_in_network.pdf")
		assert.NoError(t, fErr)
		r := bufio.NewReader(f)
		for {
			buf := make([]byte, 4*1024)
			n, rErr := r.Read(buf)
			buf = buf[:n]
			if n == 0 {
				if rErr == io.EOF {
					break
				}
				if rErr != nil {
					assert.NoError(t, rErr)
					break
				}
			}
			pipeFileWriter.Write(buf)
		}
		err = pipeFileWriter.Close()
		assert.NoError(t, err)
	}()

	aesEnc.Encrypt(pipeFileReader, pipeEncWriter)
	f, err := os.Create("/Users/philippbouzid/books/information_theory_in_network_enc.pdf")
	if err != nil {
		assert.NoError(t, err)
	}

	_, err = io.Copy(f, pipeEncReader)
	if err != nil {
		assert.NoError(t, err)
	}
}
