# Singleton Design Pattern in Go

## 1️⃣ Singleton Design Pattern কী?

**Singleton Design Pattern** হলো একটি *Creational Design Pattern* যা নিশ্চিত করে যে—

* একটি struct/class-এর **মাত্র একটি instance** তৈরি হবে
* এবং সেই instance-টিই পুরো application জুড়ে globally accessible থাকবে

এই প্যাটার্নটি মূলত **shared resource management**-এর জন্য ব্যবহার করা হয়।

### কেন দরকার হয়?

যখন এমন কোনো resource থাকে যা একটাই থাকা উচিত:

* Database connection
* Logger
* Config manager
* Cache manager

যেমন: যদি একাধিক database connection object তৈরি হয় অপ্রয়োজনীয়ভাবে, তাহলে resource waste হবে এবং concurrency সমস্যা হতে পারে।

👉 এই ধারণাটি প্রথম জনপ্রিয়ভাবে আসে Design Patterns: Elements of Reusable Object-Oriented Software (Gang of Four – Erich Gamma et al.) বইতে।

---

## 2️⃣ Golang এ Singleton কিভাবে ব্যবহার করতে হয়

Go-তে class নেই, তাই struct + package level variable ব্যবহার করে Singleton implement করা হয়।

### ⚠️ Problem: Concurrent Environment

Go highly concurrent language। তাই multiple goroutine থেকে call হলে race condition হতে পারে।

এজন্য Go-তে standard solution হলো `sync.Once` ব্যবহার করা।

---

## 3️⃣ Proper Singleton Implementation in Go (Thread-Safe)

```go
package singleton

import (
	"fmt"
	"sync"
)

type Database struct {
	Connection string
}

var instance *Database
var once sync.Once

func GetInstance() *Database {
	once.Do(func() {
		fmt.Println("Creating Database Instance...")
		instance = &Database{
			Connection: "Connected",
		}
	})
	return instance
}
```

### 🔎 এখানে কী হচ্ছে?

* `instance` → package level pointer
* `once` → `sync.Once`
* `once.Do()` → নিশ্চিত করে block-এর কোড **শুধুমাত্র একবারই execute হবে**
* যতবার `GetInstance()` call করো, একই instance return করবে

---

### Example Usage

```go
func main() {
	db1 := singleton.GetInstance()
	db2 := singleton.GetInstance()

	if db1 == db2 {
		fmt.Println("Same Instance")
	}
}
```

Output:

```
Creating Database Instance...
Same Instance
```

দেখো — "Creating Database Instance..." একবারই print হচ্ছে।

---

## 4️⃣ `sync.Once` কী?

`sync.Once` হলো Go standard library (`sync` package)-এর একটি primitive যা guarantee করে যে একটি function একবারের বেশি execute হবে না — even in concurrent access.

Official Go Documentation:
[https://pkg.go.dev/sync#Once](https://pkg.go.dev/sync#Once)

---

## 5️⃣ কোথায় Singleton ব্যবহার করা উচিত?

### ✅ Good Use Cases

* Database connection pool
* Logger
* Config loader
* Cache

### ❌ Avoid When

* Global state testing-এ সমস্যা তৈরি করে
* Dependency injection preferable হলে
* High coupling তৈরি করলে

---

## 6️⃣ Advantages & Disadvantages

### ✅ Advantages

* Controlled access
* Memory efficient
* Thread safe (if implemented properly)

### ❌ Disadvantages

* Global state তৈরি করে
* Unit testing কঠিন
* Tight coupling তৈরি করে

---

# 7️⃣ Interview Questions (Basic → Advanced)

### 🟢 Basic

1. Singleton Design Pattern কী?
2. কেন Singleton ব্যবহার করা হয়?
3. Go-তে Singleton কিভাবে implement করো?

---

### 🟡 Intermediate

4. Go-তে Singleton implement করতে `sync.Once` কেন ব্যবহার করা হয়?
5. Naive Singleton implementation-এ concurrency problem কী হতে পারে?
6. Singleton আর Global Variable এর মধ্যে পার্থক্য কী?

---

### 🔴 Advanced

7. Double-checked locking কী? Go-তে প্রয়োজন হয়?
8. Singleton কি Dependency Injection-এর বিকল্প?
9. Singleton testing-এ কী সমস্যা তৈরি করে?
10. Real production environment-এ Singleton misuse-এর example দাও।

---

## 8️⃣ Real-world Insight (Important for Interview)

Production-grade Go project-এ অনেক সময় সরাসরি Singleton না করে:

* Dependency Injection
* Explicit Initialization
* Context-based resource passing

ব্যবহার করা হয়।

কারণ clean architecture-এ hidden dependency avoid করা হয়।

---

## 9️⃣ সংক্ষেপে

```
Singleton =

✔ একটাই instance
✔ global access
✔ controlled creation
✔ sync.Once দিয়ে thread-safe করা হয়
```
