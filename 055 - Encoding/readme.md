# Encoding, Hashing, HMAC, and JWT Authentication in Go

## 1️⃣ Base64 — Encoding (NOT Security)

### What Base64 is

* **Base64 is encoding**, not encryption, not hashing.
* Purpose: **convert binary data into readable ASCII text**
* Used when data must be sent over:

  * HTTP headers
  * JSON
  * URLs
  * Emails

➡️ Anyone can **decode Base64 back to original data**

---

### How it works

* Takes **3 bytes (24 bits)** → splits into **4 groups of 6 bits**
* Maps each 6-bit group to a character set:

```
A–Z a–z 0–9 + /
```

---

### Example (Go)

```go
data := []byte("Aaab")

encoded := base64.StdEncoding.EncodeToString(data)
fmt.Println(encoded)   // QWFhYg==

decoded, _ := base64.StdEncoding.DecodeString(encoded)
fmt.Println(string(decoded)) // Aaab
```

---

### When to use Base64

✅ Encoding binary → text
❌ **Never** for passwords or authentication

---

## 2️⃣ SHA (Secure Hash Algorithm) — One-way Hashing

### What hashing means

* One-way function
* Input → fixed-length output
* Cannot be reversed
* Same input → same output

Used for:

* Password hashing
* Data integrity
* Digital signatures

---

## SHA-1 (❌ Broken)

* Output: **160 bits (20 bytes)**
* Vulnerable to **collision attacks**
* ❌ **Never use**

```go
// DO NOT USE
sha1.Sum(data)
```

---

## SHA-256 (✅ Secure)

* Output: **256 bits (32 bytes)**
* Strong, widely used
* Used in:

  * JWT
  * TLS
  * Blockchain

```go
data := []byte("130")
hash := sha256.Sum256(data)

fmt.Println(hex.EncodeToString(hash[:]))
// cad8af85fecf4977a9259787b454c1f883cc0b15386080bc36c0b25678ec5c56
```

---

## SHA-512 (✅ Stronger, Bigger)

* Output: **512 bits (64 bytes)**
* More secure but:

  * Bigger size
  * Slightly slower
* Used in high-security environments

```go
hash := sha512.Sum512(data)
```

---

## SHA Comparison

| Algorithm | Output  | Secure? | Use Case        |
| --------- | ------- | ------- | --------------- |
| SHA-1     | 160 bit | ❌ No    | Deprecated      |
| SHA-256   | 256 bit | ✅ Yes   | JWT, APIs       |
| SHA-512   | 512 bit | ✅ Yes   | Banking, crypto |

---

## 3️⃣ HMAC — **(Hash-Based Message Authentication Code)**

### Problem with plain SHA

If we hash:

```
SHA256("message")
```

Anyone can do the same.

---

### Solution: HMAC

HMAC = **Hash + Secret Key**

```
HMAC(secret, message)
```

✔ Ensures:

* Integrity (message not modified)
* Authenticity (sender knows secret)

---

## 4️⃣ HMAC-SHA-256 (🔥 Most Important for JWT)

### What it is

* HMAC using SHA-256
* Requires **secret key**
* Output cannot be forged without secret

---

### Example (Go)

```go
secret := []byte("-My-Secret-")
message := []byte("My Name is Farhan Nadim")

h := hmac.New(sha256.New, secret)
h.Write(message)

mac := h.Sum(nil)
fmt.Println(hex.EncodeToString(mac))
```

---

### What’s happening internally

1. Message is padded
2. Secret key is mixed using XOR
3. SHA-256 runs twice
4. Output is fixed 256-bit hash

➡️ If **message OR secret changes → output changes**

---

## 5️⃣ Encoding vs Hashing vs HMAC (CRITICAL)

| Feature     | Base64 | SHA-256           | HMAC-SHA-256     |
| ----------- | ------ | ----------------- | ---------------- |
| Reversible  | ✅ Yes  | ❌ No              | ❌ No             |
| Uses Secret | ❌ No   | ❌ No              | ✅ Yes            |
| Security    | ❌ None | ⚠️ Integrity only | ✅ Authentication |
| JWT Usage   | ❌      | ❌                 | ✅                |

---

# 6️⃣ BEST for JWT Authentication ✅

### ✅ Answer: **HMAC-SHA-256 (HS256)**

---

## Why HMAC-SHA-256 is best for JWT

JWT has **3 parts**:

```
header.payload.signature
```

---

### JWT Signature Generation

```text
signature = HMAC-SHA256(
  base64(header) + "." + base64(payload),
  secret
)
```

---

### Why this works

✔ Payload can be read
✔ Payload **cannot be modified**
✔ Only server knows the secret
✔ Fast and scalable
✔ No database lookup required

---

## JWT Verification Flow

1. Client sends JWT
2. Server recomputes HMAC-SHA-256 using secret
3. Compare signatures
4. If match → token is valid

---

## JWT Algorithms Comparison

| Algorithm | Type         | Use Case                   |
| --------- | ------------ | -------------------------- |
| HS256     | HMAC-SHA-256 | Best for APIs              |
| RS256     | RSA          | Public/private key systems |
| ES256     | ECDSA        | High-security, complex     |

➡️ **For most Go + REST APIs → HS256 is BEST**

---

## Final Recommendation 🔥

✔ Use **Base64** → encoding JWT parts
✔ Use **SHA-256** → internal hashing
✔ Use **HMAC-SHA-256** → JWT signing
❌ Never use SHA-1
❌ Never store passwords using plain SHA
