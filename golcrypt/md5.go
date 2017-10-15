package golcrypt

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(content []byte) string {
	md5Hash := md5.New()
	md5Hash.Write(content)
	return hex.EncodeToString(md5Hash.Sum(nil))
}
