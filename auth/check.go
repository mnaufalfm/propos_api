package jwt

import (
	"crypto/hmac"
	"strings"
)

//untuk melakukan pengecekan keabsahan token. Return true jika token valid
func CheckToken(token string) bool {
	breakToken := strings.Split(token, ".")
	//fmt.Println("MasukToken")
	if len(breakToken) < 2 {
		//fmt.Println("Maksimal")
		return false
	}
	if breakToken[0] == "" || breakToken[1] == "" || breakToken[2] == "" {
		//fmt.Println("Lemah")
		return false
	}
	signSend := breakToken[2]
	signReal := ComputeHMAC256(breakToken[0]+"."+breakToken[1], "anggunauranaufalwilliam")
	//fmt.Println(signSend)
	//fmt.Println(signReal)
	return hmac.Equal([]byte(signSend), []byte(signReal))
}
