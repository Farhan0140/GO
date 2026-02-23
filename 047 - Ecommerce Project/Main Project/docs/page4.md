# Tight Coupling in Golang

## 🔒 Tight Coupling কী?

**Tight Coupling** হলো এমন একটি অবস্থা যেখানে একটি struct / module / package অন্য একটি নির্দিষ্ট implementation-এর উপর সরাসরি নির্ভরশীল থাকে। অর্থাৎ, একটিকে পরিবর্তন করলে অন্যটিকেও পরিবর্তন করতে হয়।

Software Engineering–এ coupling বলতে module-গুলোর মধ্যে dependency-র মাত্রাকে বোঝায়। Coupling যত বেশি হবে, system তত কম maintainable ও কম testable হবে।

### সংজ্ঞা

"Coupling is the measure of interdependence between software modules."

(Source: Software Engineering: A Practitioner's Approach – Roger Pressman)

---

## ❌ Tight Coupling এর উদাহরণ (Go)

```go
package main

import "fmt"

type MySQL struct{}

func (m MySQL) Save(data string) {
    fmt.Println("Saving to MySQL:", data)
}

type UserService struct {
    db MySQL
}

func (u UserService) CreateUser(name string) {
    u.db.Save(name)
}
```

### সমস্যা কোথায়?

`UserService` সরাসরি `MySQL` এর উপর নির্ভর করছে।

এখন যদি:

* MySQL বদলে PostgreSQL করতে চাও
* Unit Test করতে চাও
* Mock ব্যবহার করতে চাও

→ তাহলে `UserService` modify করতে হবে।

এটাই Tight Coupling।

---

## 🎯 লক্ষ্য: Loose Coupling

Loose Coupling মানে dependency থাকবে, কিন্তু নির্দিষ্ট implementation-এর উপর নয় — বরং abstraction (interface) এর উপর।

এই ধারণাটি আসে **Dependency Inversion Principle (DIP)** থেকে।

"Depend upon abstractions, not concretions."

(Source: Clean Code – Robert C. Martin)

---

## ✅ Golang এ Tight Coupling কীভাবে Remove করবো?

### Step 1: Interface ব্যবহার করো

```go
type Database interface {
    Save(data string)
}
```

### Step 2: Implementation আলাদা করো

```go
type MySQL struct{}

func (m MySQL) Save(data string) {
    fmt.Println("Saving to MySQL:", data)
}
```

### Step 3: Dependency Inject করো

```go
type UserService struct {
    db Database
}

func NewUserService(db Database) UserService {
    return UserService{db: db}
}

func (u UserService) CreateUser(name string) {
    u.db.Save(name)
}
```

### Step 4: ব্যবহার

```go
func main() {
    mysql := MySQL{}
    service := NewUserService(mysql)

    service.CreateUser("Farhan")
}
```

---

## 🔍 এখন কী পরিবর্তন হলো?

| আগে                        | এখন                                |
| -------------------------- | ---------------------------------- |
| UserService → MySQL        | UserService → Database (interface) |
| Change করলে modify করতে হয় | শুধু implementation বদলালেই হবে    |
| Test করা কঠিন              | Mock দিয়ে সহজে test করা যায়        |

---

## 🧪 Testing এ সুবিধা

```go
type MockDB struct{}

func (m MockDB) Save(data string) {
    fmt.Println("Mock saving:", data)
}
```

এখন test এ:

```go
func main() {
    mock := MockDB{}
    service := NewUserService(mock)

    service.CreateUser("Test User")
}
```

Production code একদম unchanged থাকবে।

---

## 📌 কেন Go তে এটা গুরুত্বপূর্ণ?

Go-তে inheritance নেই, তাই abstraction করার একমাত্র উপায় হলো **interface**।

"Interfaces in Go provide a way to specify the behavior of an object."

(Source: Effective Go – Official Go Documentation)

---

## 🧠 সংক্ষেপে

### Tight Coupling মানে:

* Hard dependency
* Change করলে chain effect
* Test করা কঠিন

### Remove করার উপায়:

1. Interface ব্যবহার করা
2. Constructor দিয়ে dependency inject করা
3. Concrete type এর বদলে abstraction এর উপর depend করা