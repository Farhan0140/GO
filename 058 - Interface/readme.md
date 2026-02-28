# Go Interface – Complete Explanation (Bangla)

---

## 1️⃣ What is an Interface in GoLang?

Go-তে **Interface** হলো একটা type যা বলে দেয় — *কোন method গুলো থাকতে হবে*, কিন্তু **কিভাবে implement হবে সেটা বলে না**।

মানে Interface শুধু behavior define করে, implementation না।

### 🔹 Simple Example

```go
package main

import "fmt"

// 1. Interface
type Animal interface {
	Sound() string
}

// 2. Struct
type Dog struct{}

type Cat struct{}

// 3. Method implementation
func (d Dog) Sound() string {
	return "Woof"
}

func (c Cat) Sound() string {
	return "Meow"
}

func main() {
	var a Animal

	a = Dog{}
	fmt.Println(a.Sound()) // Woof

	a = Cat{}
	fmt.Println(a.Sound()) // Meow
}
```

### 🔎 এখানে কী হলো?

* `Animal` interface বলছে → `Sound()` method থাকতে হবে।
* `Dog` আর `Cat` সেই method implement করেছে।
* তাই `Dog` এবং `Cat` — দুটোই `Animal` হয়ে গেছে।

⚡ Go-তে `implements` keyword লাগে না।
যে struct method fulfill করবে, সে automatically interface implement করবে।

---

## 2️⃣ Why do I have to use Interface?

Interface ব্যবহার করার প্রধান কারণগুলো:

### ✅ 1. Flexibility

একই interface দিয়ে অনেক ধরনের object handle করা যায়।

উপরের উদাহরণে:

* `Dog`
* `Cat`
* ভবিষ্যতে `Cow`

সবই `Animal` হিসেবে কাজ করবে।

---

### ✅ 2. Loose Coupling (Clean Architecture)

ধরো তুমি backend project বানাচ্ছো

```go
type Database interface {
	Save(data string) error
}
```

এখন:

* MySQL
* PostgreSQL
* MongoDB

সবই `Database` interface implement করতে পারে।

তাহলে তোমার business logic database এর উপর depend করবে না —
Depend করবে শুধু `Database` interface এর উপর।

এটা professional Go backend এ খুব গুরুত্বপূর্ণ।

---

### ✅ 3. Testing সহজ হয় (Mocking)

Testing করার সময় fake implementation বানানো যায়।

---

## 3️⃣ Relation between Abstraction and Interface

### 🔹 Abstraction কী?

Abstraction মানে:
👉 User কে শুধু প্রয়োজনীয় জিনিস দেখানো
👉 ভিতরের implementation লুকিয়ে রাখা

### 🔹 Interface কিভাবে Abstraction দেয়?

Interface:

* শুধু method signature দেখায়
* ভিতরের code লুকিয়ে রাখে

Example:

```go
type Payment interface {
	Pay(amount float64) error
}
```

User জানে:
👉 `Pay()` method আছে
❌ কিন্তু কিভাবে টাকা কাটছে সেটা জানে না

এটাই abstraction।

---

## 4️⃣ Why Interface is called Pure Abstraction?

Go-তে interface এর ভিতরে:

❌ কোনো variable থাকে না
❌ কোনো implementation থাকে না
✔ শুধু method signature থাকে

Example:

```go
type Shape interface {
	Area() float64
}
```

এখানে:

* কোনো field নেই
* কোনো body নেই

তাই একে বলে **Pure Abstraction**।

Java-তে interface-এ default method থাকতে পারে
কিন্তু Go-তে interface একদম clean — pure behavior contract।

---

## 5️⃣ How to connect multiple interface with single struct?

একটা struct একাধিক interface implement করতে পারে।

### 🔹 Example

```go
package main

import "fmt"

type Reader interface {
	Read() string
}

type Writer interface {
	Write(data string)
}

type File struct{}

// Reader implement
func (f File) Read() string {
	return "Reading file..."
}

// Writer implement
func (f File) Write(data string) {
	fmt.Println("Writing:", data)
}

func main() {
	var r Reader
	var w Writer

	f := File{}

	r = f
	w = f

	fmt.Println(r.Read())
	w.Write("Hello")
}
```

### 🔎 কী হলো এখানে?

* `File` struct
* `Read()` method আছে → তাই `Reader`
* `Write()` method আছে → তাই `Writer`

অর্থাৎ এক struct → multiple interface implement করতে পারে।

---

## 🔥 Bonus: Interface Composition (Advanced)

```go
type ReadWriter interface {
	Reader
	Writer
}
```

এখানে `ReadWriter` = `Reader + Writer`

---

# 🧠 Final Summary

| বিষয়               | ব্যাখ্যা                                     |
| ------------------ | -------------------------------------------- |
| Interface          | শুধু method contract                         |
| Purpose            | Flexibility + Loose coupling                 |
| Abstraction        | Implementation লুকানো                        |
| Pure abstraction   | শুধু method signature                        |
| Multiple interface | এক struct অনেক interface implement করতে পারে |
