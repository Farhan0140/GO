# SOLID Principles &
# `interface{}` in Go


## Part 1: SOLID Principles (Software Engineering)

### SOLID কী?

**SOLID** হলো **৫টা software design principle** এর সমষ্টি।

এগুলো follow করলে code হয়:

* Clean
* Maintainable
* Scalable
* Bug কম
* Team-friendly

👉 SOLID কোনো programming language না, এটা **code লেখার চিন্তাভাবনার নিয়ম**।

---

## SOLID = ৫টি Principle

| Letter | Full Meaning                    |
| ------ | ------------------------------- |
| S      | Single Responsibility Principle |
| O      | Open / Closed Principle         |
| L      | Liskov Substitution Principle   |
| I      | Interface Segregation Principle |
| D      | Dependency Inversion Principle  |

---

## S — Single Responsibility Principle (SRP)

### অর্থ

> একটা class / struct / module / function এর **একটাই দায়িত্ব** থাকা উচিত।

### কেন দরকার?

* Bug কম হয়
* Change করা সহজ হয়
* Code বুঝতে সুবিধা হয়

### ❌ ভুল উদাহরণ (Go)

```go
type UserService struct {}

func (u UserService) CreateUser() {}
func (u UserService) SendEmail() {}
func (u UserService) SaveToFile() {}
```

এখানে `UserService`:

* User তৈরি করছে
* Email পাঠাচ্ছে
* File-এ save করছে

❌ এক struct = একাধিক দায়িত্ব

### ✅ সঠিক ডিজাইন

```go
type UserService struct {}
type EmailService struct {}
type FileService struct {}
```

📌 **একটা পরিবর্তনের কারণ = একটাই responsibility**

---

## O — Open / Closed Principle (OCP)

### অর্থ

> Code **modify না করে extend করা যাবে**

### ❌ খারাপ উদাহরণ

```go
if userType == "admin" {
	// logic
} else if userType == "guest" {
	// logic
}
```

নতুন role এলে পুরানো code edit করতে হবে ❌

### ✅ ভালো উদাহরণ (Interface ব্যবহার)

```go
type User interface {
	GetRole() string
}
```

নতুন role দরকার হলে → নতুন struct add করো
পুরানো code touch করার দরকার নেই

---

## L — Liskov Substitution Principle (LSP)

### অর্থ

> Parent type যেখানে কাজ করে, child type ও সেখানে কাজ করতে পারবে

### সহজ কথা

* Child কখনো parent-এর behavior ভাঙতে পারবে না

### ❌ ক্লাসিক ভুল উদাহরণ

```text
Bird can Fly
Penguin is Bird (but can't fly)
```

👉 Penguin কে Bird বানালে LSP ভেঙে যায় ❌

---

## I — Interface Segregation Principle (ISP)

### অর্থ

> বড় interface বানিও না
> ছোট ছোট interface বানাও

### ❌ ভুল ডিজাইন

```go
type Worker interface {
	Work()
	Eat()
	Sleep()
}
```

Robot এর Eat/Sleep দরকার নেই ❌

### ✅ ভালো ডিজাইন

```go
type Workable interface {
	Work()
}

type Eatable interface {
	Eat()
}
```

📌 Client যেন অপ্রয়োজনীয় method implement করতে বাধ্য না হয়

---

## D — Dependency Inversion Principle (DIP)

### অর্থ

> High-level module low-level module-এর উপর depend করবে না
> দুজনেই interface-এর উপর depend করবে

### ❌ খারাপ উদাহরণ

```go
type App struct {
	db MySQL
}
```

### ✅ ভালো উদাহরণ

```go
type Database interface {
	Save()
}

type App struct {
	db Database
}
```

📌 এতে করে MySQL → PostgreSQL change করা সহজ হয়

---

## SOLID মনে রাখার এক লাইনের ট্রিক 🧠

> **SOLID = Code যেন মানুষ বুঝতে পারে, শুধু computer না**

---

<br>
<br>
<br>
<br>

# Part 2: `interface{}` in Go

## `interface{}` কী?

```go
interface{}
```

এর মানে:

> **যে কোনো type**

এটা Go এর **empty interface**।

---

## কেন `interface{}` সব type নিতে পারে?

কারণ:

* Interface মানে method set
* Empty interface এ **কোনো method নেই**

👉 তাই সব type-ই empty interface satisfy করে

---

## Basic Example

```go
var x interface{}

x = 10
x = "hello"
x = true
x = 3.14
```

সব valid ✅

---

## Real-life Analogy

```text
interface{} = খালি বাক্স
যেকোনো জিনিস ঢোকানো যায়
```

---

## Function parameter হিসেবে ব্যবহার

```go
func PrintAnything(v interface{}) {
	fmt.Println(v)
}

PrintAnything(10)
PrintAnything("Farhan")
PrintAnything([]int{1,2,3})
```

---

## Problem: Type জানা না থাকলে কী হয়?

```go
var x interface{} = 10

fmt.Println(x + 5) // ❌ compile error
```

কারণ Go জানে না `x` আসলে কোন type

---

## Type Assertion

```go
val, ok := x.(int)

if ok {
	fmt.Println(val + 5)
}
```

---

## Type Switch (Best Practice)

```go
func checkType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("unknown")
	}
}
```

---

## Real-world ব্যবহার কোথায়?

1. JSON decoding

```go
map[string]interface{}
```

2. `fmt.Println`

```go
fmt.Println(a, b, c)
```

3. Database Scan

```go
Scan(&value interface{})
```

---

## ⚠️ `interface{}` বেশি ব্যবহার করলে সমস্যা

* Type safety কমে
* Runtime error বাড়ে
* Code বোঝা কঠিন হয়

📌 Go philosophy:

> **`interface{}` শেষ option হিসেবে ব্যবহার করো**

---

## Go 1.18+ Alternative: Generics

```go
func Print[T any](v T) {
	fmt.Println(v)
}
```

👉 অনেক ক্ষেত্রে `interface{}` এর modern replacement

---

## Final Summary

### SOLID

* Software design করার নিয়ম
* Long-term maintainable code বানায়

### `interface{}`

* Go এর empty interface
* Any type ধরতে পারে
* Power + Danger একসাথে 😄