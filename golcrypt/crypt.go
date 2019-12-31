package golcrypt

import "errors"

func Encrypt(blob []byte, key []byte, scheme string) ([]byte, error) {
	if len(key) < 256 {
		key = KeyForAES(key)
	}
	switch scheme {
	case "aes":
		return aesEncrypt(blob, key)
	default:
		return []byte{}, errors.New("unsupported encryption utilized")
	}
}

func Decrypt(blob []byte, key []byte, scheme string) ([]byte, error) {
	if len(key) < 256 {
		key = KeyForAES(key)
	}
	switch scheme {
	case "aes":
		return aesDecrypt(blob, key)
	default:
		return []byte{}, errors.New("unsupported encryption utilized")
	}
}

func aesEncrypt(blob []byte, key []byte) ([]byte, error) {
	block := AESBlock{DataBlob: blob, Key: key}
	err := block.Encrypt()
	return block.Cipher, err
}

func aesDecrypt(blob []byte, key []byte) ([]byte, error) {
	block := AESBlock{Cipher: blob, Key: key}
	err := block.Decrypt()
	return block.DataBlob, err
}
