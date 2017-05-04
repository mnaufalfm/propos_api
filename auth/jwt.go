package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

func StringToBase64(mes string) string {
	mess := []byte(mes)
	return base64.StdEncoding.EncodeToString(mess)
}

func Base64ToString(bes64 string) string {
	mess, _ := base64.StdEncoding.DecodeString(bes64)
	return string(mess)
}

//message dalam bentuk json
func TokenMaker(message string, key string) string {
	header := `{"alg":"HS256", "typ":"JWT"}`
	a := StringToBase64(header) + "." + StringToBase64(message)
	sign := ComputeHMAC256(a, key)
	token := a + "." + sign
	return token
}

func ComputeHMAC256(mes string, key string) string {
	//fmt.Println("ComputeHMAC256")
	//fmt.Println(mes)
	kun := []byte(key)
	pes := []byte(mes)
	h := hmac.New(sha256.New, kun)
	h.Write(pes)
	return base64.RawStdEncoding.EncodeToString(h.Sum(nil))
}
