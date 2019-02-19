package golcrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
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

func KeyForAES(key []byte) []byte {
	hasher := md5.New()
	hasher.Write(key)
	return []byte(hex.EncodeToString(hasher.Sum(nil)))
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
