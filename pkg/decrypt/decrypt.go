package decrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"os"

	"turtle/pkg/derive"
)

func Decryptshow(name string, password []byte) error {
	ciphertext, err := read(name)
	if err != nil {
		return fmt.Errorf("%s", "incorrect/non-existing name")
	}
	plaintext, err := decrypt(password, ciphertext)
	if err != nil {
		return fmt.Errorf("%s", "incorrect password")
	}
	fmt.Println(string(plaintext))
	return nil
}

func decrypt(key, data []byte) ([]byte, error) {
	salt, data := data[len(data)-32:], data[:len(data)-32]

	key, _, err := derive.DeriveKey(key, salt)
	if err != nil {
		return nil, err
	}

	blockCipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(blockCipher)
	if err != nil {
		return nil, err
	}

	nonce, ciphertext := data[:gcm.NonceSize()], data[gcm.NonceSize():]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

func read(name string) ([]byte, error) {
	home := os.Getenv("HOME")
	filepath := home + "/turtle-secrets/" + name

	ciphertext, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	return ciphertext, nil
}
