package jwt

import (
	"encoding/base64"
	"encoding/json"
	"testing"
)

/*
func TestMain(m *testing.M) {
}
*/

func TestGetSplittedToken(t *testing.T) {
	token := "header.payload.signature"
	splittedToken, err := getSplittedToken(token)
	if err != nil {
		t.Errorf("Error getting the splitted token %v\n", err)
	}

	if splittedToken[0] != "header" {
		t.Errorf("Wrong header in token %v\n", err)
	}

	if splittedToken[1] != "payload" {
		t.Errorf("Wrong payload in token %v\n", err)
	}

	if splittedToken[2] != "signature" {
		t.Errorf("Wrong signature in token %v\n", err)
	}
}

func TestGetSplittedTokenFail(t *testing.T) {
	token := "headerpayloadsignature"
	_, err := getSplittedToken(token)
	if err == nil {
		t.Errorf("A wrong token format should yeld a error\n")
	}
}

func TestSignPayloadAndHeader(t *testing.T) {
	expect := "aGLWGAlPb6y_LnEZrfHBvxrMZqVZlO-YUzGlVUv_1Wo"
	privatekey := "privatekey"
	result := signPayloadAndHeader("header", "payload", privatekey)
	if result != expect {
		t.Errorf("Error signing header and payload with private key using the HMACSHA256\n")
	}
}

func TestSign(t *testing.T) {
	payload := make(map[string]interface{})
	payload["name"] = "Danilo"
	privateKey := "privateKey"
	token, err := Sign(payload, privateKey)
	if err != nil {
		t.Errorf("Error signing token %v\n", err)
	}
	splittedToken, err := getSplittedToken(token)
	if err != nil {
		t.Errorf("Error splitting token %v\n", err)
	}

	headerBytes, err := base64.RawURLEncoding.DecodeString(splittedToken[0])
	if err != nil {
		t.Errorf("Error decoding header %v\n", err)
	}
	var header map[string]interface{}
	if err := json.Unmarshal(headerBytes, &header); err != nil {
		t.Errorf("Error parsing header %v\n", err)
	}

	if header["typ"] != "JWT" || header["alg"] != "HS256" {
		t.Errorf("Wrong header returned")
	}

	payloadBytes, err := base64.RawURLEncoding.DecodeString(splittedToken[1])
	if err != nil {
		t.Errorf("Error decoding Payload %v\n", err)
	}
	var payloadReturned map[string]interface{}
	if err := json.Unmarshal(payloadBytes, &payloadReturned); err != nil {
		t.Errorf("Error parsing payload %v\n", err)
	}

	if payloadReturned["name"] != payload["name"] {
		t.Errorf("Wrong payload returned")
	}
}

func TestVerify(t *testing.T) {
	privateKey := "privateKey"
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwibmFtZSI6IkRhbmlsbyIsImlhdCI6MTUxNjIzOTAyMn0.G0Hbu7ep_JiCtPuq4t9FvbVNJGKbwZV_A5m2b08EmhY"
	payload, err := Verify(token, privateKey)
	if err != nil {
		t.Errorf("Error verifying the token %v\n", err)
	}

	if payload["id"] != float64(1) {
		t.Errorf("Wrong id returned in the payload\n")
	}

	if payload["name"] != "Danilo" {
		t.Errorf("Wrong name returned in the payload\n")
	}

}

func TestVerifyError(t *testing.T) {
	privateKey := "wrong" // wrong private key
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwibmFtZSI6IkRhbmlsbyIsImlhdCI6MTUxNjIzOTAyMn0.G0Hbu7ep_JiCtPuq4t9FvbVNJGKbwZV_A5m2b08EmhY"
	_, err := Verify(token, privateKey)
	if err == nil {
		t.Errorf("Token should not be validate because of wrong private key\n")
	}
}
