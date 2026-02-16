package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	var s string
	s = "Aaab"

	// String to Byte
	byte_array := []byte(s)

	// fmt.Println(s)
	// fmt.Println(byte_array)

	// Encode byte to string
	enc := base64.URLEncoding
	enc = enc.WithPadding(base64.NoPadding)
	Base64str := enc.EncodeToString(byte_array)
	// fmt.Println(Base64str)

	// Decode string to byte
	decodedByte, error := enc.DecodeString(Base64str)
	if error != nil {
		fmt.Println("Something want wrong: \n", error)
		return
	}
	fmt.Println(decodedByte)



	// Sha-256 (Secure Hash Algorithm)
	data := []byte("130")
	hash := sha256.Sum256(data)
	fmt.Println("SHA-256: ", hash)	// SHA-256:  [202 216 175 133 254 207 73 119 169 37 151 135 180 84 193 248 131 204 11 21 56 96 128 188 54 192 178 86 120 236 92 86]

	hashToString := hex.EncodeToString(hash[:])
	fmt.Println(hashToString)	// cad8af85fecf4977a9259787b454c1f883cc0b15386080bc36c0b25678ec5c56


	// HMAC (Hash-based Message Authentication Code)

	secret := []byte("-My-Secret-")
	message := []byte("My Name is Farhan Nadim")

	h := hmac.New(sha256.New, secret)
	h.Write(message)

	text := h.Sum(nil)
	fmt.Println("HMAC-SHA-256: ", text)
	// fmt.Println("HMAC-SHA-256: ", hex.EncodeToString(text[:]))
}
