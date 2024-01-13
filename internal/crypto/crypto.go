package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

func Encrypt(text, key []byte) ([]byte, error) {
	desiredKeySize := 32
	if len(key) < desiredKeySize {
		// Pad the key with zeros or any other strategy.
		for i := len(key); i < desiredKeySize; i++ {
			key = append(key, 0)
		}
	} else if len(key) > desiredKeySize {
		// Truncate the key.
		key = key[:desiredKeySize]
	}

	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, aesgcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return aesgcm.Seal(nonce, nonce, text, nil), nil
}

func Decrypt(ciphertext, key []byte) ([]byte, error) {
	desiredKeySize := 32
	if len(key) < desiredKeySize {
		// Pad the key with zeros or any other strategy.
		for i := len(key); i < desiredKeySize; i++ {
			key = append(key, 0)
		}
	} else if len(key) > desiredKeySize {
		// Truncate the key.
		key = key[:desiredKeySize]
	}

	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	noncesize := aesgcm.NonceSize()
	nonce, ciphertext := ciphertext[:noncesize], ciphertext[noncesize:]
	return aesgcm.Open(nil, nonce, ciphertext, nil)
}
