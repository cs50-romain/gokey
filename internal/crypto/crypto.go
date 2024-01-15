package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"crypto/rand"
	"errors"
	"io"
	//"fmt"
	"math/big"

	"golang.org/x/crypto/pbkdf2"
)

func Encrypt(text, key []byte) ([]byte, error) {
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

	//fmt.Printf("Nonce: %x\n", nonce)
	ciphertext := aesgcm.Seal(nil, nonce, text, nil)
	result := append(nonce, ciphertext...)
	//fmt.Printf("Result: %x\n", result)
	return result, nil
}

func Decrypt(ciphertext, key []byte) ([]byte, error) {
	//fmt.Printf("Encrypted Data: %x\n", ciphertext)
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	noncesize := aesgcm.NonceSize()
	if len(ciphertext) < noncesize {
		return nil, errors.New("ciphertext is too short to contain nonce")
	}

	nonce, ciphertext := ciphertext[:noncesize], ciphertext[noncesize:]
	//fmt.Printf("Nonce: %x\n", nonce)
	//fmt.Printf("Ciphertext: %x\n", ciphertext)

	return aesgcm.Open(nil, nonce, ciphertext, nil)
}

func GenerateRandomSalt(length int) ([]byte, error) {
    results := make([]byte, length)
    for i := 0; i < length; i++ {
        salt, err := rand.Int(rand.Reader, big.NewInt(255))
        if err != nil {
            return nil, err
        }
        results[i] = byte(salt.Int64())
    }
    return results, nil
}

func DeriveKey(masterPassword string, salt []byte) []byte {
	// Use PBKDF2 to derive a fixed-length key
	return pbkdf2.Key([]byte(masterPassword), salt, 100000, 32, sha256.New)
}
