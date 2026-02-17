# JWT Creation in Go (HS256)
---

# JWT Standard Structure

A JWT token follows this format:

```
Base64Url(Header) + "." + Base64Url(Payload) + "." + Base64Url(Signature)
```

JWT specification reference (RFC 7519):
[https://datatracker.ietf.org/doc/html/rfc7519](https://datatracker.ietf.org/doc/html/rfc7519)

Go crypto documentation:
[https://pkg.go.dev/crypto/hmac](https://pkg.go.dev/crypto/hmac)

Base64 documentation:
[https://pkg.go.dev/encoding/base64](https://pkg.go.dev/encoding/base64)

---

# 1️⃣ Imports

```go
import (
    "crypto/hmac"
    "crypto/sha256"
    "encoding/base64"
    "encoding/json"
)
```

Explanation:

* `crypto/hmac` → Used to generate HMAC signatures
* `crypto/sha256` → SHA256 hashing algorithm
* `encoding/base64` → For Base64 URL encoding
* `encoding/json` → Converts structs to JSON

JWT requires:

* Header → JSON
* Payload → JSON
* Signature → HMAC SHA256

---

# 2️⃣ Header Struct

```go
type Header struct {
    Alg string `json:"alg"`
    Typ string `json:"typ"`
}
```

Purpose:
The JWT header specifies:

* Which algorithm is used (`HS256`)
* Token type (`JWT`)

Example JSON:

```json
{
  "alg": "HS256",
  "typ": "JWT"
}
```

---

# 3️⃣ Payload Struct

```go
type Payload struct {
    ID        int    `json:"id"`
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
    Email     string `json:"email"`
    Password  string `json:"password"`
    IsAdmin   bool   `json:"is_admin"`
}
```

Purpose:
Stores user-related data inside the token.

Example JSON:

```json
{
  "id": 1,
  "first_name": "Farhan",
  "last_name": "Nadim",
  "email": "farhan@gmail.com",
  "password": "123456",
  "is_admin": true
}
```

⚠️ Security Warning:
Passwords should NEVER be stored inside JWT payload.
JWT payload is encoded, NOT encrypted.

---

# 4️⃣ Create_JWT Function

```go
func Create_JWT(secret string, data Payload) (string, error)
```

Parameters:

* `secret` → Secret key used to generate signature
* `data` → Payload struct

Returns:

* JWT string
* error (if any occurs)

---

# 5️⃣ Assign Header

```go
header := Header{
    Alg: "HS256",
    Typ: "JWT",
}
```

Algorithm fixed as:

* HS256 → HMAC + SHA256

---

# 6️⃣ Convert Header to JSON

```go
headerByteArr, err := json.Marshal(header)
```

Example output:

```
{"alg":"HS256","typ":"JWT"}
```

Returned as a byte slice.

---

# 7️⃣ Convert Payload to JSON

```go
dataByteArr, err := json.Marshal(data)
```

Example output:

```
{"id":1,"first_name":"Farhan",...}
```

---

# 8️⃣ Base64 URL Encoding

```go
header_encoded_str := base64_URL_Encode(headerByteArr)
data_encoded_str := base64_URL_Encode(dataByteArr)
```

Example Encoded Header:

```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9
```

Example Encoded Payload:

```
eyJpZCI6MSwiZmlyc3RfbmFtZSI6IkZhcmhhbiIs...
```

---

# 9️⃣ Create Initial JWT Part

```go
initial_jwt_part := header_encoded_str + "." + data_encoded_str
```

Now it becomes:

```
header.payload
```

Example:

```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZmly...
```

---

# 🔟 Convert Secret to Byte Slice

```go
secretByteArr := []byte(secret)
```

If secret = "mysecret"
It becomes a byte array.

---

# 1️⃣1️⃣ Generate Signature

```go
h := hmac.New(sha256.New, secretByteArr)
h.Write(jwt_part_byteArr)
signature := h.Sum(nil)
```

This performs:

```
HMAC_SHA256(
   key = secret,
   message = header.payload
)
```

Purpose:

* Ensures token integrity
* Prevents tampering
* Without secret key, valid signature cannot be generated

---

# 1️⃣2️⃣ Encode Signature

```go
signature_encoded_str := base64_URL_Encode(signatureByteArr)
```

Signature is Base64 URL encoded.

---

# 1️⃣3️⃣ Create Final JWT

```go
jwt := header_encoded_str + "." + data_encoded_str + "." + signature_encoded_str
```

Final Structure:

```
header.payload.signature
```

Example:

```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9
.
eyJpZCI6MSwiZmlyc3RfbmFtZSI6IkZhcmhhbiIs...
.
XyJhdfkjhdsfkjhsdfkjhsdf
```

---

# base64_URL_Encode Function

```go
func base64_URL_Encode(data []byte) string {
    return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
```

Why URL Encoding?

Standard Base64 contains:

```
+  /  =
```

These characters can cause issues in URLs.

JWT uses:

* URL-safe Base64
* No padding

---

# Complete Flow Summary

Given:

```go
secret := "mysecret"
```

User payload:

```go
Payload{
  ID: 1,
  FirstName: "Farhan",
  ...
}
```

Steps:

1. Convert Header → JSON
2. Convert Payload → JSON
3. Base64 URL encode both
4. Create `header.payload`
5. Generate HMAC SHA256 signature
6. Encode signature
7. Combine into final JWT

---

# 🔐 Important Security Notes

❌ Do NOT store passwords in JWT
❌ JWT is NOT encrypted
✔ JWT is signed

JWT provides:

* Integrity
* Authenticity (if secret is secure)

JWT does NOT provide:

* Confidentiality

<br>
<br>
<br>

# **Minimizing Manual JWT Creation in Go Using Library (HS256)**

Manually building JWT in production systems is not recommended because:

* Signature validation mistakes are possible
* Algorithm confusion attack risks exist
* Expiration (`exp`), issued at (`iat`) may be forgotten
* Security maintenance becomes difficult

This document explains how to minimize manual JWT code using a proper Go library.

---

# Recommended Go JWT Library

Most widely used and actively maintained library:

`github.com/golang-jwt/jwt/v5`

GitHub:
[https://github.com/golang-jwt/jwt](https://github.com/golang-jwt/jwt)

Go Documentation:
[https://pkg.go.dev/github.com/golang-jwt/jwt/v5](https://pkg.go.dev/github.com/golang-jwt/jwt/v5)

This is the community-maintained fork. The old `dgrijalva/jwt-go` package is deprecated.

---

# Step 1: Install

```bash
go get github.com/golang-jwt/jwt/v5
```

---

# Step 2: Minimal JWT Creation (HS256)

Your previous manual 60+ lines of JWT code can now be reduced to around 10–15 lines.

```go
package util

import (
    "time"

    "github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
    ID        int    `json:"id"`
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
    Email     string `json:"email"`
    IsAdmin   bool   `json:"is_admin"`
    jwt.RegisteredClaims
}

func CreateJWT(secret string, data CustomClaims) (string, error) {

    // Add expiration
    data.RegisteredClaims = jwt.RegisteredClaims{
        ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
        IssuedAt:  jwt.NewNumericDate(time.Now()),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)

    signedToken, err := token.SignedString([]byte(secret))
    if err != nil {
        return "", err
    }

    return signedToken, nil
}
```

---

# What Was Removed Compared to Manual Implementation?

In the manual version you had to:

* json.Marshal header
* json.Marshal payload
* Base64 encode header
* Base64 encode payload
* Concatenate header.payload
* Create HMAC using hmac.New
* Use sha256 hashing
* Encode signature
* Combine everything manually

In the library version:

```
jwt.NewWithClaims()
token.SignedString()
```

Everything is handled internally.

---

# **JWT Verification (Now Very Simple)**

Manual verification is complex and error-prone.

Using the library:

```go
func VerifyJWT(secret, tokenString string) (*CustomClaims, error) {

    token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{},
        func(token *jwt.Token) (interface{}, error) {
            return []byte(secret), nil
        })

    if err != nil {
        return nil, err
    }

    claims, ok := token.Claims.(*CustomClaims)
    if !ok || !token.Valid {
        return nil, err
    }

    return claims, nil
}
```

The library automatically validates:

* Signature
* Expiration
* Token structure
* Token validity

---

# Why Using a Library Is Better

1. RFC-compliant implementation
2. Handles edge cases
3. Built-in expiration support
4. Standard claims support
5. Receives future security updates

---

# Important Security Improvement

In your manual code, you included:

```
Password string `json:"password"`
```

Passwords must NEVER be stored inside JWT payload.

JWT is NOT encrypted.
JWT is signed only.

Production payload should contain:

* User ID
* Role
* Permissions

---

# Production-Level Minimal Version

Cleaner production-ready claims structure:

```go
type Claims struct {
    UserID  int  `json:"user_id"`
    IsAdmin bool `json:"is_admin"`
    jwt.RegisteredClaims
}
```

---

# Manual vs Library Comparison

| Feature       | Manual          | Library  |
| ------------- | --------------- | -------- |
| Signature     | Manual HMAC     | Built-in |
| Expiration    | Not implemented | Built-in |
| Validation    | Manual          | Built-in |
| Security Risk | High            | Low      |
| Code Size     | Large           | Small    |

---

# Final Recommendation

For production systems:

✔ Always use `golang-jwt/jwt`
❌ Never manually build JWT unless it is strictly for learning purposes

