# CORS & Preflight OPTIONS
---

## 1. CORS কী?

**CORS (Cross-Origin Resource Sharing)** হলো browser-এর একটা security rule।

সহজ ভাষায়:

> এক domain-এর website অন্য domain-এর API call করলে browser আগে check করে — এটা allowed কিনা।

উদাহরণ:

* Frontend: `http://example.com`
* Backend API: `http://api.example.com`

এই দুইটা **origin আলাদা**, তাই browser নিজে থেকেই request block করতে পারে, যদি server permission না দেয়।

---

## 2. OPTIONS (Preflight Request) কী?

### সহজ ভাষায়

Browser যখন দেখে:

* Request টি cross-origin
* এবং request টি **simple request না**

তখন browser আগে server-এর কাছে **permission চাইতে** একটা request পাঠায়।

এই permission-check request টাই হলো:

> **OPTIONS request (Preflight request)**

মানে:

> "আমি কি এই method, এই header, এই origin থেকে request পাঠাতে পারবো?"

---

## 3. কখন Browser OPTIONS (Preflight) পাঠায়?

Browser OPTIONS পাঠায় যদি নিচের যেকোনো একটা সত্য হয়:

### ❌ Request টি simple না হলে

Simple request হতে হলে:

* Method: `GET`, `POST`, `HEAD`
* Header: শুধু standard header
* Content-Type:

  * `text/plain`
  * `multipart/form-data`
  * `application/x-www-form-urlencoded`
* Custom header না থাকতে হবে

### ✅ Preflight হবে যদি:

* Method হয় `PUT`, `DELETE`, `PATCH`
* `Content-Type: application/json`
* Custom header থাকে (`Authorization`, `X-API-KEY`)
* Cookie / token পাঠানো হয় (`credentials: true`)

📌 বাস্তবে প্রায় সব modern API-ই JSON + Authorization ব্যবহার করে → তাই OPTIONS খুব common

---

## 4. Browser থেকে পাঠানো Preflight OPTIONS Request

Browser যেটা পাঠায় সেটা দেখতে এমন:

```http
OPTIONS /api/users HTTP/1.1
Origin: http://example.com
Access-Control-Request-Method: POST
Access-Control-Request-Headers: Content-Type, Authorization
```

এখানে Browser বলছে:

* আমি `POST` method ব্যবহার করতে চাই
* আমি এই header গুলো পাঠাতে চাই
* আমি এই origin থেকে আসছি

---

## 5. Server কী response দিলে browser allow করবে?

Server যদি বলে **"হ্যাঁ, পারো"**, তাহলে browser আসল request পাঠাবে।

Server response-এ অবশ্যই এই header গুলো থাকতে হবে:

```http
HTTP/1.1 204 No Content
Access-Control-Allow-Origin: http://example.com
Access-Control-Allow-Methods: GET, POST, PUT, DELETE, OPTIONS
Access-Control-Allow-Headers: Content-Type, Authorization
Access-Control-Allow-Credentials: true
```

👉 এগুলো না থাকলে browser error দেখাবে:

```
Blocked by CORS policy
```

---

## 6. Go (net/http) এ OPTIONS কীভাবে handle করবে?

### Basic Go example

```go
package main

import (
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	// CORS headers (সব request এর জন্য)
	w.Header().Set("Access-Control-Allow-Origin", "http://example.com")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	// Preflight OPTIONS request হলে এখানেই শেষ
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// আসল request
	w.Write([]byte("Hello from Go API"))
}

func main() {
	http.HandleFunc("/api", handler)
	http.ListenAndServe(":8080", nil)
}
```

📌 সবচেয়ে গুরুত্বপূর্ণ অংশ:

```go
if r.Method == http.MethodOptions {
    w.WriteHeader(http.StatusNoContent)
    return
}
```

এটা না করলে → browser আসল request পাঠাবেই না।

---

## 7. CORS Middleware (Best Practice)

বারবার একই code না লিখে middleware বানানো ভালো।

### Go CORS Middleware

```go
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
```

### Middleware ব্যবহার

```go
mux := http.NewServeMux()
mux.HandleFunc("/api", handler)

http.ListenAndServe(":8080", corsMiddleware(mux))
```

---

## 8. Gin Framework এ CORS handle করা

```go
r := gin.Default()

r.Use(func(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	c.Next()
})
```

---

## 9. Common CORS ভুল (খুব গুরুত্বপূর্ণ)

### ❌ OPTIONS handle না করা

→ Browser API call করবে না

### ❌ `Access-Control-Allow-Origin: *` + `Credentials: true`

→ Invalid (browser block করবে)

### ❌ Header mismatch

Client পাঠায়:

```text
Authorization
```

Server allow করে:

```text
Content-Type
```

→ ❌ Blocked by CORS

---

## 10. Real-life Frontend → Go Backend Example

### Frontend (JavaScript)

```js
fetch("http://localhost:8080/api", {
	method: "POST",
	headers: {
		"Content-Type": "application/json",
		"Authorization": "Bearer token"
	},
	body: JSON.stringify({ name: "Farhan" })
})
```

### Flow:

1. Browser → OPTIONS /api
2. Go server → CORS headers সহ response
3. Browser → POST /api

---

## 11. মনে রাখার জন্য এক লাইনের ট্রিক 🧠

> **OPTIONS request = Browser এর permission check**
> **Go server যদি OPTIONS accept না করে → API dead**