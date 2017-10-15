package golrandom

import (
	cryptoRand "crypto/rand"
	mathRand "math/rand"
	"time"
)

func Token(tokenLength int) string {
	randomFactor := make([]byte, tokenLength*2)
	_, err := cryptoRand.Read(randomFactor)
	if err != nil {
		panic(err)
	}

	mathRand.Seed(time.Now().UnixNano() * int64(randomFactor[0]))

	var letterRunes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_+=")

	token := make([]rune, tokenLength)
	for i := range token {
		token[i] = letterRunes[mathRand.Intn(len(letterRunes))]
	}

	return string(token)
}
