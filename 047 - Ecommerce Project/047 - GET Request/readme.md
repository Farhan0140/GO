# Go (GET Request) Products API

---

## 1️⃣ `package main`

```go
package main
```

* Go program executable বানাতে `package main` বাধ্যতামূলক
* `main()` function এখান থেকেই entry point
* `go run` করলে execution শুরু হয়

---

## 2️⃣ Import Section

```go
import (
    "encoding/json"
    "fmt"
    "net/http"
)
```

### `encoding/json`

* Go struct → JSON convert
* JSON → Go struct parse
* এখানে response বানাতে ব্যবহৃত

### `fmt`

* Console output / debugging

### `net/http`

* Go built-in HTTP server
* Router, Request, Response, Socket handling

---

## 3️⃣ Product Struct

```go
type Product struct {
    ID          int     `json:"id"`
    Title       string  `json:"title"`
    Description string  `json:"description"`
    Price       float32 `json:"price"`
    ImageURL    string  `json:"img"`
}
```

* Real-world product model
* JSON tag দিয়ে output key customize করা হয়েছে
* Capital field name → exported → encoder access করতে পারে

---

## 4️⃣ Global Products Slice

```go
var products []Product
```

* Global scope
* initially `nil`
* `init()` এ populate করা হয়
* handler থেকে read করা হয়

---

## 5️⃣ `getProducts` Handler

```go
func getProducts(w http.ResponseWriter, r *http.Request) {
```

* `/products` route hit হলে call হয়
* Go runtime আলাদা goroutine এ run করে

### Parameters

* `w` → response লেখার pipe
* `r` → client request info

---

### 5.1️⃣ CORS Header (Cross-Origin Resource Sharing)

```go
w.Header().Set("Access-Control-Allow-Origin", "*")
```

* Browser restriction bypass
* যেকোন origin থেকে request allow

---

### 5.2️⃣ Content-Type Header

```go
w.Header().Set("Content-Type", "application/json")
```

* Client কে জানানো হয় response JSON

---

### 5.3️⃣ HTTP Method Check

```go
if r.Method != "GET" {
    http.Error(w, "POST, PUT, PATCH, DELETE request not allowed", 400)
    return
}
```

* GET ছাড়া অন্য method reject
* `400` = Bad Request

---

### 5.4️⃣ JSON Encode করে Response

```go
encoder := json.NewEncoder(w)
encoder.Encode(products)
```

Flow:

```
products (Go slice)
   ↓
json encoder
   ↓
HTTP response body
```

* Direct stream encode
* Memory efficient

---

## 6️⃣ `main()` Function

```go
func main() {
```

* Program execution এখান থেকে শুরু

---

### 6.1️⃣ ServeMux (Router)

```go
mux := http.NewServeMux()
```

* Go built-in router
* URL → handler mapping

---

### 6.2️⃣ Route Register

```go
mux.HandleFunc("/products", getProducts)
```

* `/products` request → `getProducts()` call

---

### 6.3️⃣ Server Start Log

```go
fmt.Println("Server Running on port: 8080")
```

* Console message only

---

### 6.4️⃣ HTTP Server Start

```go
err := http.ListenAndServe(":8080", mux)
```

এই line সবচেয়ে গুরুত্বপূর্ণ:

Go internally যা করে:

1. TCP socket create
2. Port 8080 bind
3. listen()
4. accept loop
5. epoll দিয়ে FD monitor
6. প্রতি request এ goroutine spawn

* Blocking call

---

### 6.5️⃣ Error Handle

```go
if err != nil {
    fmt.Println("***Error Occurred", err)
}
```

* Port busy / permission issue handle

---

## 7️⃣ `init()` Function

```go
func init() {
```

* `main()` এর আগে auto execute
* initial setup / seed data load

---

## 8️⃣ Product Data Create

```go
prd1 := Product{ ... }
```

* Struct literal
* Stack এ create
* পরে slice এ append

---

## 9️⃣ Products Slice Append

```go
products = append(
    products,
    prd1,
    prd2,
    ...
    prd10,
)
```

* nil slice → allocate
* internal array create
* values copy

Slice internally:

* pointer
* length
* capacity

---

## 🔁 Full Request → Response Flow

```
Browser
  ↓ GET /products
NIC
  ↓
Kernel socket buffer
  ↓
epoll marks FD readable
  ↓
Go netpoller
  ↓
goroutine wakes
  ↓
ServeMux route match
  ↓
getProducts()
  ↓
json.Encode
  ↓
socket send buffer
  ↓
Client receives JSON
```

---

## 🎯 Summary

* `init()` → data load
* `main()` → server start
* ServeMux → routing
* Handler → request processing
* json.Encoder → response
* net/http → epoll + goroutine magic