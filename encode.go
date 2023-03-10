package sencoding

import (
	"encoding/base64"
	"strings"
)

func Encode(str string) string {

	var tempres string
	for _, i := range str {
		tempres += string(i ^ 4/2)
	}

	base := base64.StdEncoding.EncodeToString([]byte(tempres))
	rev := Reverse(base)

	var res string
	for _, i := range rev {
		res += string(i) + GenerateSym(1)
	}

	res = strings.ReplaceAll(res, "=", "")

	return res

}
