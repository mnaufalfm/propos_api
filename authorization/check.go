package jwt

import (
	"crypto/hmac"
	"strings"
)

//untuk melakukan pengecekan keabsahan token
func CheckToken(token string) bool {
	breakToken := strings.Split(token, ".")
	signSend := breakToken[2]
	signReal := TokenMaker(breakToken[0]+"."+breakToken[1], "anggunauranaufalwilliam")
	return hmac.Equal([]byte(signSend), []byte(signReal))
}

