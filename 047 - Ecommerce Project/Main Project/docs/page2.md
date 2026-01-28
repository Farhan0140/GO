# Go Middleware

## Part 1: Middleware আসলে কী?

### সহজ ভাষায়

**Middleware = Request আর Handler এর মাঝখানে বসা একটা layer**

```
Client
  ↓
Middleware 1
  ↓
Middleware 2
  ↓
Handler (Business Logic)
  ↓
Response
```

👉 Middleware পারে:

* request আটকাতে
* request modify করতে
* request log করতে
* authentication / authorization check করতে
* CORS handle করতে

📌 Middleware নিজেই response দিতে পারে
📌 আবার চাইলে পরের handler-এ request পাঠাতে পারে

---

## Go তে Middleware এর Standard Shape

```go
func(next http.Handler) http.Handler
```

মানে:

* input: একটা handler
* output: আরেকটা handler (wrapped handler)

---

## Part 2: `middleware` Package (Line by Line Explanation)

### Package declaration

```go
package middleware
```

👉 আলাদা package ব্যবহার করা হয়েছে (clean architecture)

---

### Import

```go
import "net/http"
```

👉 Go এর HTTP server tools ব্যবহার করার জন্য

---

### Middleware Type Definition

```go
type Middleware func(http.Handler) http.Handler
```

👉 Middleware হচ্ছে এমন function:

* যা একটা handler নেয়
* নতুন wrapped handler return করে

এইটাই Go middleware system এর core

---

### Manager Struct

```go
type Manager struct {
	globalMiddlewares []Middleware
}
```

👉 Manager middleware গুলোর controller
👉 `globalMiddlewares` মানে সব route এ apply হবে

---

### Manager Constructor

```go
func NewManager() *Manager {
	return &Manager{
		globalMiddlewares: make([]Middleware, 0),
	}
}
```

👉 নতুন middleware manager তৈরি করে
👉 empty slice initialize করা হয়েছে

---

### Use(): Global Middleware যোগ করা

```go
func (mngr *Manager) Use(middlewares ...Middleware) {
	mngr.globalMiddlewares = append(mngr.globalMiddlewares, middlewares...)
}
```

👉 একাধিক middleware একসাথে যোগ করা যায়
👉 এগুলো সব route এ apply হবে

Example:

```go
manager.Use(middleware.Test, middleware.Logger)
```

---

### With(): Handler কে Middleware দিয়ে Wrap করা

```go
func (mngr *Manager) With(next http.Handler, middlewares ...Middleware) http.Handler {
```

👉 এই function:

* একটা handler নেয়
* route-specific middleware নেয়
* global middleware যুক্ত করে
* final wrapped handler return করে

---

### Step 1: Base handler সেট করা

```go
n := next
```

👉 শুরুতে `n` = আসল handler

---

### Step 2: Route-specific Middleware apply করা

```go
for _, middleware := range middlewares {
	n = middleware(n)
}
```

Flow:

```
Middleware → Handler
```

Order খুব গুরুত্বপূর্ণ

---

### Step 3: Global Middleware apply করা

```go
for _, gblMiddleware := range mngr.globalMiddlewares {
	n = gblMiddleware(n)
}
```

Flow:

```
GlobalMiddleware → RouteMiddleware → Handler
```

---

### Step 4: Final Handler return

```go
return n
```

👉 এখন `n` পুরো middleware chain সহ handler

---

## Part 3: `main.go` — Request Flow Explanation

### Manager তৈরি

```go
manager := middleware.NewManager()
```

---

### Global Middleware Register

```go
manager.Use(middleware.Test, middleware.Logger)
```

👉 সব route এ এই middleware গুলো চলবে

---

### Router তৈরি

```go
mux := http.NewServeMux()
```

---

### Route Example

```go
mux.Handle(
	"GET /products",
	manager.With(
		http.HandlerFunc(handlers.GetProducts),
	),
)
```

Request Flow:

```
Client
 ↓
Test Middleware
 ↓
Logger Middleware
 ↓
GetProducts Handler
```

---

### Server Start

```go
http.ListenAndServe(":8080", util.GlobalRouter(mux))
```

👉 পুরো mux কে GlobalRouter দিয়ে wrap করা হয়েছে

---

## Part 4: `util.GlobalRouter` (Global CORS Middleware)

### Function Signature

```go
func GlobalRouter(mux *http.ServeMux) http.Handler {
```

👉 ServeMux কে middleware হিসেবে wrap করছে

---

### All Request Handler

```go
handleAllRequest := func(w http.ResponseWriter, r *http.Request) {
```

👉 সব request প্রথমে এখানে আসবে

---

### CORS Headers Set করা

```go
w.Header().Set("Access-Control-Allow-Origin", "*")
```

👉 সব origin allow

```go
w.Header().Set("Access-Control-Allow-Method", "GET, POST, PUT, PATCH, DELETE")
```

⚠️ এখানে typo আছে:

* ❌ `Allow-Method`
* ✅ হওয়া উচিত `Access-Control-Allow-Methods`

```go
w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
```

```go
w.Header().Set("Content-Type", "application/json")
```

---

### OPTIONS (Preflight) Handle করা

```go
if r.Method == "OPTIONS" {
	w.WriteHeader(200)
	return
}
```

👉 Browser preflight request এখানেই শেষ
👉 handler পর্যন্ত যাবে না

---

### আসল Request Router এ পাঠানো

```go
mux.ServeHTTP(w, r)
```

---

### Handler Return

```go
return http.HandlerFunc(handleAllRequest)
```

---

## Full Request Flow (Visualization)

```
Browser
 ↓
GlobalRouter (CORS + OPTIONS)
 ↓
ServeMux
 ↓
Global Middleware (Test, Logger)
 ↓
Handler
 ↓
Response
```

---

## Middleware সম্পর্কে Golden Rules 🧠

1. Middleware = handler wrapper
2. Execution order খুব গুরুত্বপূর্ণ
3. Middleware request থামাতে পারে
4. Middleware reusable হওয়া উচিত
5. Global ও Route-specific middleware আলাদা রাখা best practice
