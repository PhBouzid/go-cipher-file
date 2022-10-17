package encryption

type Cipher interface {
	Encrypt(data <-chan []byte, output chan<- []byte)
	Decrypt(data <-chan []byte, output chan<- []byte)
}
