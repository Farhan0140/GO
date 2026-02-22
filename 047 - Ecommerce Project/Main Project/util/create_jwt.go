package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {
	ID        int `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	IsAdmin   bool   `json:"is_admin"`
}

func Create_JWT(secret string, data Payload) (string, error) {
	header := Header{ // Assign header
		Alg: "HS256",
		Typ: "JWT",
	}

	headerByteArr, err := json.Marshal(header)
	if err != nil {
		return "", err
	}

	dataByteArr, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	header_encoded_str := base64_URL_Encode(headerByteArr)
	data_encoded_str := base64_URL_Encode(dataByteArr)

	secretByteArr := []byte(secret)

	initial_jwt_part := header_encoded_str + "." + data_encoded_str
	jwt_part_byteArr := []byte(initial_jwt_part)

	// For Signature
	h := hmac.New(sha256.New, secretByteArr)
	h.Write(jwt_part_byteArr)

	signature := h.Sum(nil)
	signatureByteArr := []byte(signature)
	signature_encoded_str := base64_URL_Encode(signatureByteArr)

	jwt := header_encoded_str + "." + data_encoded_str + "." + signature_encoded_str

	return jwt, nil
}


func base64_URL_Encode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}