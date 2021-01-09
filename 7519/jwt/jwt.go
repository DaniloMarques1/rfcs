package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type Header struct {
	Typ string `json:"typ"`
	Alg string `json:"alg"`
}

type Jwt struct {
	Header    Header
	Payload   map[string]interface{}
	Signature string
}

//  Will create token following jwt's requirements
func Sign(payload map[string]interface{}, secret string) string {
	// TODO add exp time claim
	var jwt Jwt
	// default header values
	jwt.Header.Typ = "JWT"
	jwt.Header.Alg = "HS256" // at least for now, only sha256
	jwt.Payload = payload
	marshedHeader, err := json.Marshal(&jwt.Header)
	if err != nil {
		fmt.Printf("Error %v \n", err)
		os.Exit(-1)
	}
	encodedHeader := base64.RawURLEncoding.EncodeToString([]byte(marshedHeader)) // first part

	jwt.Payload["iat"] = time.Now().Unix()
	marshedPayload, err := json.Marshal(&jwt.Payload)
	if err != nil {
		fmt.Printf("Error %v \n", err)
		os.Exit(-1)
	}
	encodedPayload := base64.RawURLEncoding.EncodeToString([]byte(marshedPayload)) // second part

	jwt.Signature = signPayloadAndHeader(encodedHeader, encodedPayload, secret) // third part
	token := fmt.Sprintf("%s.%s.%s", encodedHeader, encodedPayload, jwt.Signature)
	return token
}

// checks if token is valid, if so returns a map with the payload
func Verify(token, privateKey string) (map[string]interface{}, error) {
	splitedToken, err := getSplittedToken(token)
	if err != nil {
		fmt.Printf("Token error: %v", err)
                return nil, err
        }
	payload, err := base64.RawURLEncoding.DecodeString(splitedToken[1])
	if err != nil {
		fmt.Printf("Error decoding payload %v", err)
		return nil, err
	}
	payloadMap := make(map[string]interface{})
	err = json.Unmarshal(payload, &payloadMap)
	if err != nil {
		fmt.Printf("Error decoding payload %v", err)
		return nil, err
	}

	return payloadMap, nil
}

func getSplittedToken(token string) ([]string, error) {
	splittedToken := strings.Split(token, ".")
	if len(splittedToken) != 3 {
		return nil, fmt.Errorf("Invlaid token\n")
	}
	return splittedToken, nil
}

// will sign the token header and paylaod using the given privatekey
// and the hmac sha256 algorithm, returning the base64url encoded signed string
func signPayloadAndHeader(encodedHeader, encodedPayload, privateKey string) string {
	//privateKey = base64.RawURLEncoding.EncodeToString([]byte(privateKey))
	h := hmac.New(sha256.New, []byte(privateKey))
	concat := fmt.Sprintf("%s.%s", encodedHeader, encodedPayload)
	h.Write([]byte(concat))
	return base64.RawURLEncoding.EncodeToString(h.Sum(nil))
}

