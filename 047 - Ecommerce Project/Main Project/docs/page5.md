# Dependency and Dependency Injection in Golang

## 🔹 Dependency কী?

**Dependency** মানে হলো—একটি struct, function, বা module অন্য একটি struct, function, বা service-এর উপর নির্ভর করছে।

সহজ ভাষায়:
যদি `A` কাজ করতে `B` লাগে, তাহলে `B` হলো `A`-র dependency।

### উদাহরণ (Real-world analogy)

ধরো, একটি `OrderService` অর্ডার সংরক্ষণ করতে `Database` ব্যবহার করছে।
এখানে `Database` হলো `OrderService`-এর dependency।

---

## 🔹 Dependency Injection (DI) কী?

**Dependency Injection** হলো এমন একটি design pattern যেখানে dependency-টা struct-এর ভিতরে নিজে তৈরি না করে, বাইরে থেকে দিয়ে দেওয়া হয়।

অর্থাৎ:

> “নিজে বানাবো না, বাইরে থেকে নিবো।”

এটা loosely coupled architecture তৈরি করতে সাহায্য করে।

---

# ❌ Example 1: Dependency Injection ছাড়া (Tightly Coupled)

```go
package main

import "fmt"

type MySQL struct{}

func (m MySQL) Save(data string) {
	fmt.Println("Saved to MySQL:", data)
}

type OrderService struct{}

func (o OrderService) CreateOrder() {
	db := MySQL{} // ❌ নিজেই dependency তৈরি করছে
	db.Save("New Order")
}

func main() {
	service := OrderService{}
	service.CreateOrder()
}
```

### 🔴 সমস্যা কী?

* `OrderService` সরাসরি `MySQL` এর উপর নির্ভর করছে।
* পরে যদি `Postgres` ব্যবহার করতে চাও → পুরো কোড বদলাতে হবে।
* Unit testing কঠিন হবে।

এটা tightly coupled design।

---

# ✅ Example 2: Dependency Injection ব্যবহার করে (Loosely Coupled)

### Step 1: Interface তৈরি করি

```go
type Database interface {
	Save(data string)
}
```

### Step 2: Implementation তৈরি করি

```go
type MySQL struct{}

func (m MySQL) Save(data string) {
	fmt.Println("Saved to MySQL:", data)
}
```

### Step 3: Dependency Inject করি

```go
type OrderService struct {
	db Database // interface ব্যবহার করছি
}

func NewOrderService(database Database) OrderService {
	return OrderService{db: database}
}

func (o OrderService) CreateOrder() {
	o.db.Save("New Order")
}
```

### Step 4: main এ inject করি

```go
func main() {
	mysql := MySQL{}
	service := NewOrderService(mysql) // ✅ Inject করছি
	service.CreateOrder()
}
```

---

## 🔹 এখানে কী হলো?

* `OrderService` জানেই না কোন database ব্যবহার হচ্ছে।
* শুধু জানে—একটা `Database` interface লাগবে।
* তাই চাইলে এখন সহজে `Postgres`, `MongoDB` ইত্যাদি যোগ করা যাবে।

---

# 🔥 কেন Dependency Injection গুরুত্বপূর্ণ?

### 1️⃣ Loose Coupling

Component গুলো একে অপরের উপর কম নির্ভরশীল হয়।

### 2️⃣ Easy Testing

Mock database ব্যবহার করা যায়।

```go
type MockDB struct{}

func (m MockDB) Save(data string) {
	fmt.Println("Mock Save:", data)
}
```

Testing এ:

```go
mock := MockDB{}
service := NewOrderService(mock)
```

### 3️⃣ Scalable Architecture

Large project-এ maintain করা সহজ।

---

# 🔹 Go-তে Dependency Injection কিভাবে করা হয়?

Go-তে সাধারণত ৩ভাবে DI করা হয়:

### 1️⃣ Constructor Injection (সবচেয়ে common)

```go
func NewService(dep Dependency) *Service
```

### 2️⃣ Field Injection

Struct তৈরি করার পর field assign করা।

### 3️⃣ Method Injection

Method parameter হিসেবে dependency পাঠানো।

---

# 🔥 গুরুত্বপূর্ণ কথা

Go-তে built-in DI container নেই (যেমন Java-তে Spring আছে)।
তবে জনপ্রিয় library আছে:

* Google তৈরি করা Wire
* Uber এর Fx
* Uber এর Dig

তবে Go-তে সাধারণত manual DI-ই বেশি ব্যবহার করা হয় (constructor দিয়ে)।

---

# 🎯 সংক্ষেপে

| বিষয়                 | অর্থ                                              |
| -------------------- | ------------------------------------------------- |
| Dependency           | যে জিনিস ছাড়া অন্য struct কাজ করতে পারে না        |
| Dependency Injection | সেই dependency বাইরে থেকে inject করা              |
| লাভ                  | Loose coupling, testable code, clean architecture |
