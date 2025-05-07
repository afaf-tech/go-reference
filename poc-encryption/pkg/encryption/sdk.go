package encryption

type Encryptor interface {
	Encrypt(plain []byte) ([]byte, error)
	Decrypt(ciphertext []byte) ([]byte, error)
}
