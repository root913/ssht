package hash

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"io"
	"io/ioutil"
	"log"
	"os"
)

type Hash struct{}

func (h *Hash) createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func (h *Hash) encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(h.createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

func (h *Hash) decrypt(data []byte, passphrase string) []byte {
	key := []byte(h.createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}

func (h *Hash) EncryptFile(filename string, data []byte, passphrase string) {
	f, _ := os.Create(filename)
	defer f.Close()
	f.Write(h.encrypt(data, passphrase))
}

func (h *Hash) DecryptFile(filename string, passphrase string) ([]byte, error) {
	if _, err := os.Stat(filename); err != nil {
		return nil, err
	}
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return h.decrypt(data, passphrase), nil
}

func NewHash() *Hash {
	return &Hash{}
}
