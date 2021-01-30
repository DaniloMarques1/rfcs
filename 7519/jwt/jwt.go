package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
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

// Will create token following jwt's requirements and return
// the string representing the token and any possible errors
func Sign(payload map[string]interface{}, secret string) (string, error) {
	var jwt Jwt
	// default header values
	jwt.Header.Typ = "JWT"
	jwt.Header.Alg = "HS256" // at least for now, only sha256
	jwt.Payload = payload
	marshedHeader, err := json.Marshal(&jwt.Header)
	if err != nil {
		return "", err
	}
	encodedHeader := base64.RawURLEncoding.EncodeToString([]byte(marshedHeader)) // first part

	jwt.Payload["iat"] = time.Now().Unix()
	marshedPayload, err := json.Marshal(&jwt.Payload)
	if err != nil {
		return "", err
	}
	encodedPayload := base64.RawURLEncoding.EncodeToString([]byte(marshedPayload)) // second part

	jwt.Signature = signPayloadAndHeader(encodedHeader, encodedPayload, secret) // third part
	token := fmt.Sprintf("%s.%s.%s", encodedHeader, encodedPayload, jwt.Signature)

	return token, nil
}

// checks if token is valid, if so returns a map with the payload
func Verify(token, privateKey string) (map[string]interface{}, error) {
	splitedToken, err := getSplittedToken(token)
	if err != nil {
		fmt.Printf("Token error: %v", err)
		return nil, err
	}
	if signature := signPayloadAndHeader(splitedToken[0], splitedToken[1], privateKey); signature != splitedToken[2] {
		return nil, fmt.Errorf("Invalid signature")
	}

	payload, err := base64.RawURLEncoding.DecodeString(splitedToken[1])
	if err != nil {
		fmt.Printf("Error decoding payload %v", err)
		return nil, err
	}

	payloadMap := make(map[string]interface{})
	if err := json.Unmarshal(payload, &payloadMap); err != nil {
		fmt.Printf("Error decoding payload %v", err)
		return nil, err
	}

	return payloadMap, nil
}

// receives a string (token) and returns a slice with len 3
// where slice[0] = header slice[1] = payload and slice[2] = signature
func getSplittedToken(token string) ([]string, error) {
	splittedToken := strings.Split(token, ".")
	if len(splittedToken) != 3 {
		return nil, fmt.Errorf("Invlaid token format\n")
	}
	return splittedToken, nil
}

// will sign the token header and paylaod using the given privatekey
// and the hmac sha256 algorithm, returning the base64url encoded signed string
func signPayloadAndHeader(encodedHeader, encodedPayload, privateKey string) string {
	hash := hmac.New(sha256.New, []byte(privateKey))
	concat := fmt.Sprintf("%s.%s", encodedHeader, encodedPayload)
	hash.Write([]byte(concat))
	return base64.RawURLEncoding.EncodeToString(hash.Sum(nil))
}
