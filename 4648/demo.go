package main

import (
	"fmt"

	"danilo/4648/base64"
)

func main() {
	url := "https://duckduckgo.com/?t=ffab&q=base64+rfc&atb=v240-1&ia=answer"
	encodedString := base64.UrlEncodeString(string(url))
	fmt.Println(encodedString)
	decodedString := base64.UrlDecodeString(encodedString)
	fmt.Println(decodedString)
	fmt.Println()
	str := "Normal string to be encoded"
	encodedString = base64.StdEncodeString(str)
	fmt.Println(encodedString)
	decodedString = base64.StdDecodeString(encodedString)
	fmt.Println(decodedString)
}
