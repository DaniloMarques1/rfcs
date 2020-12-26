package main

import (
	"danilo/4648/base64"
	"fmt"
)

func main() {
	encodedString := base64.EncodeString("this is a cool implementation")
	fmt.Println(encodedString)
	decodedString := base64.DecodeString(encodedString)
	fmt.Println(decodedString)
}
