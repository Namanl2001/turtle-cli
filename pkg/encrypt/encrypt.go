package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"os"
	"turtle/pkg/derive"
)

func Encryptsave(name string, password, apikey []byte) error {
	ciphertext, err := encrypt(password, apikey)
	if err != nil {
		return fmt.Errorf("%s", "error in encrypting")
	}
	err = write(name, ciphertext)
	if err != nil {
		return fmt.Errorf("%s, %s", "error in writing to file", err)
	}
	return nil
}

func encrypt(key, data []byte) ([]byte, error) {
	key, salt, err := derive.DeriveKey(key, nil)
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

	nonce := make([]byte, gcm.NonceSize())
	if _, err = rand.Read(nonce); err != nil {
		return nil, err
	}

	ciphertext := gcm.Seal(nonce, nonce, data, nil)

	ciphertext = append(ciphertext, salt...)

	return ciphertext, nil
}

func write(name string, ciphertext []byte) error {
	home := os.Getenv("HOME")
	dirName := home + "/turtle-secrets"
	filepath := dirName + "/" + name

	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		err := os.Mkdir(dirName, 0755)
		if err != nil {
			return fmt.Errorf("error creating directory: %s", err)
		}
	}

	file, _ := os.Stat(filepath)
	if file != nil {
		return fmt.Errorf("%s", "api-key with this name already exists")
	}

	err := os.WriteFile(filepath, ciphertext, 0651)
	if err != nil {
		return err
	}
	return nil
}
