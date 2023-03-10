package main

import (
	"crypto/rand"
	"math/big"
)

func Reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func Random(min, max int64) int {
	bg := big.NewInt(max - min)

	n, err := rand.Int(rand.Reader, bg)
	if err != nil {
		panic(err)
	}
	return int(n.Int64() + min)
}

func GenerateSym(length int) string {
	letterRunes := []rune(")(0-")
	mixedRunes := []rune(")(0-")
	b := make([]rune, length)
	b[0] = letterRunes[Random(0, int64(len(mixedRunes)-1))]
	for i := range b {
		if i != 0 {
			b[i] = mixedRunes[Random(0, int64(len(mixedRunes)-1))]
		}
	}

	return string(b)
}
