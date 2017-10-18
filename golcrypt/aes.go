package golcrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

type DataBlob []byte
type Cipher []byte
type Key []byte

type AESBlock struct {
	DataBlob DataBlob
	Cipher   Cipher
	Key      Key
}

func (aesBlock *AESBlock) Encrypt() error {
	crypt, err := aes.NewCipher(aesBlock.Key)
	if err != nil {
		return err
	}

	galoisCounterMode, err := cipher.NewGCM(crypt)
	if err != nil {
		return err
	}

	nonce := make([]byte, galoisCounterMode.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return err
	}

	aesBlock.Cipher = galoisCounterMode.Seal(nonce, nonce, aesBlock.DataBlob, nil)
	return nil
}

func (aesBlock *AESBlock) Decrypt() error {
	crypt, err := aes.NewCipher(aesBlock.Key)
	if err != nil {
		return err
	}

	galoisCounterMode, err := cipher.NewGCM(crypt)
	if err != nil {
		return err
	}

	nonceSize := galoisCounterMode.NonceSize()
	if len(aesBlock.Cipher) < nonceSize {
		return errors.New("ciphertext too short")
	}

	var nonce []byte
	nonce, aesBlock.Cipher = aesBlock.Cipher[:nonceSize], aesBlock.Cipher[nonceSize:]
	aesBlock.DataBlob, err = galoisCounterMode.Open(nil, nonce, aesBlock.Cipher, nil)
	return err
}
