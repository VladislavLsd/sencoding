package se

import (
	"encoding/base64"
	"strings"
)

func Decode(str string) string {

	var rmvd = str
	for _, i := range ")(0-" {
		rmvd += strings.ReplaceAll(rmvd, string(i), "")
	}

	decoded, _ := base64.StdEncoding.DecodeString(Reverse(rmvd))

	var tempres string
	for _, i := range string(decoded) {
		tempres += string(i ^ 4/2) //bitwise
	}

	return string(tempres)
}
