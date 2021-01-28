package main

import (
	"fmt"
	"danilo/7519/jwt"
)

func main() {
	payload := make(map[string]interface{})
	payload["id"] = 123
	secret := "myprivatekey"
	token := jwt.Sign(payload, secret) // returns a string (token)
	fmt.Println(token)
	payloadToken, err := jwt.Verify(token, secret) // returns a map (payload)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(payloadToken)
}
