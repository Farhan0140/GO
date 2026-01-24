# 🛒 Go POST Request
---

## 📦 Package Declaration
---

## 📥 Import Section

```go
import (
    "encoding/json"
    "fmt"
    "net/http"
)
```

### Explanation (one by one)

#### `encoding/json`

* Converts **Go structs ↔ JSON**
* Used to:

  * Read JSON from request body
  * Send JSON in response

❌ Without this:

* You cannot decode incoming JSON
* You cannot send proper JSON responses

---

#### `net/http`

* Core HTTP library in Go
* Used to:

  * Create server
  * Handle routes
  * Handle requests & responses

❌ Without this:

* No server
* No API

---

## 🧱 Product Struct

```go
type Product struct {
    ID          int     `json:"id"`
    Title       string  `json:"title"`
    Description string  `json:"description"`
    Price       float32 `json:"price"`
    ImageURL    string  `json:"img"`
}
```

### Why this struct exists

* Represents a **product model**
* Used to store and send product data

### JSON Tags (VERY IMPORTANT)

Example:

```go
ID int `json:"id"`
```

Why needed:

* Converts Go field names to JSON keys

If NOT used:

```json
{
  "ID": 1,
  "Title": "Apple"
}
```

With tags:

```json
{
  "id": 1,
  "title": "Apple"
}
```

Frontend **expects lowercase JSON keys**, so tags are necessary.

---

## 🗃️ Global Products Slice

```go
var products []Product
```

### Why this is needed

* Stores all products in memory
* Acts like a temporary database

### If not used

* Data will be lost
* No place to store products

⚠️ Note:

* Data resets when server restarts
* This is NOT persistent storage

---

## ➕ Create Product Handler

```go
func createProduct(w http.ResponseWriter, r *http.Request) {
```

### What this function does

* Handles **POST /create-product**
* Receives JSON
* Saves product
* Sends JSON response

---

## 🌍 CORS HEADERS (VERY IMPORTANT)

```go
w.Header().Set("Access-Control-Allow-Origin", "*")
w.Header().Set("Access-Control-Allow-Method", "POST")
w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
```

### 🔥 THIS PART FIXES CORS ERROR 🔥

#### What is CORS?

* Browser security rule
* Blocks frontend requests to different origin

Example problem:

* Frontend: `http://localhost:3000`
* Backend: `http://localhost:8080`

Browser blocks request ❌

---

### Explanation (line by line)

#### `Access-Control-Allow-Origin: *`

* Allows requests from **any frontend**

❌ Without this:

* Browser throws **CORS error**
* API works in Postman but fails in browser

---

#### `Access-Control-Allow-Method: POST`

* Allows POST requests

❌ Without this:

* Browser blocks POST method

---

#### `Access-Control-Allow-Headers: Content-Type`

* Allows JSON content-type

❌ Without this:

* Browser blocks JSON body

---

## 🧪 Preflight OPTIONS Request

```go
if r.Method == "OPTIONS" {
    w.WriteHeader(200)
    return
}
```

### Why this exists

* Browser sends OPTIONS request before POST
* Checks permission first

❌ Without this:

* Browser never sends POST
* CORS error occurs

---

## 🚦 Method Validation

```go
if r.Method != "POST" {
    http.Error(w, "Please give me correct JSON", 400)
    return
}
```

### Why this is needed

* Protects API from wrong methods

❌ Without this:

* GET, PUT, DELETE may break logic

---

## 📥 Decode JSON Body

```go
var newProduct Product

decoder := json.NewDecoder(r.Body)
error := decoder.Decode(&newProduct)
```

### What this does

* Reads JSON from request body
* Converts JSON → Go struct

❌ Without decoding

* No data from frontend

---

## ❌ Error Handling

```go
if error != nil {
    http.Error(w, "Something want wrong", 200)
    return
}
```

### Why needed

* Handles invalid JSON

❌ Without this

* Server may crash

---

## 🆔 Assign Product ID

```go
newProduct.ID = len(products) + 1
```

### Why needed

* Generates unique ID

❌ Without this

* Products have ID = 0

---

## 💾 Save Product

```go
products = append(products, newProduct)
```

### What happens

* Adds product to memory

---

## 📤 Send Response

```go
w.WriteHeader(201)
encoder := json.NewEncoder(w)
encoder.Encode(newProduct)
```

### Why

* Sends created product back
* Status `201 Created`

---

## 🚀 Main Function

```go
func main() {
```

### Router

```go
mux := http.NewServeMux()
```

* Acts as router

---

### Routes

```go
mux.HandleFunc("/products", getProducts)
mux.HandleFunc("/create-product", createProduct)
```

* Connects URL → function

---

### Start Server

```go
http.ListenAndServe(":8080", mux)
```

* Starts server on port 8080

❌ Without this

* Server never starts

---

## ⚙️ init() Function

```go
func init() {
```

### Why init()

* Runs **before main()**
* Used to preload data

❌ Without this

* Product list empty