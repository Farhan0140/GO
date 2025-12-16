# 📘 Go Pointer

## 🔹 1. What is a Pointer?

A **pointer** is a variable that stores the **memory address** of another variable.

```go
var x int = 10
var p *int = &x
```

### Explanation

* `x` stores value `10`
* `&x` gives the memory address of `x`
* `p` stores that address
* `*p` accesses the value at that address (dereferencing)

---

## 🔹 2. Basic Pointer Example

```go
x := 10
p := &x

fmt.Println(x)   // 10
fmt.Println(p)   // memory address (0x...)
fmt.Println(*p)  // 10
```

### Visualization

```
Memory

x        p
┌───┐    ┌──────────┐
│10 │◄───│ address  │
└───┘    └──────────┘
```

---

## 🔹 3. Modify Value Using Pointer

```go
*p = 50
fmt.Println(x) // 50
```

Because `p` points to `x`, modifying `*p` modifies the original variable.

---

## 🔹 4. Pass by Value (Default in Go)

```go
func changeVal(x int) {
	x = 100
}

func main() {
	a := 10
	changeVal(a)
	fmt.Println(a) // 10
}
```

### Explanation

* `a` is copied into `x`
* Changes affect only the copy
* Original value remains unchanged

### Visualization

```
a = 10
x = copy of a = 10
```

---

## 🔹 5. Pass by Reference Using Pointer

```go
func changeVal(x *int) {
	*x = 100
}

func main() {
	a := 10
	changeVal(&a)
	fmt.Println(a) // 100
}
```

### Visualization

```
a ─────┐
       ▼
     *x = 100
```

This allows the function to modify the original variable.

---

## 🔹 6. Pointer with Arrays (Value Type)

Arrays in Go are **value types**, meaning they are copied.

```go
func modify(arr [3]int) {
	arr[0] = 99
}

func main() {
	a := [3]int{1, 2, 3}
	modify(a)
	fmt.Println(a) // [1 2 3]
}
```

The original array is not changed.

---

## 🔹 7. Array Passed Using Pointer

```go
func modify(arr *[3]int) {
	arr[0] = 99
}

func main() {
	a := [3]int{1, 2, 3}
	modify(&a)
	fmt.Println(a) // [99 2 3]
}
```

### Visualization

```
a ─────┐
       ▼
   [99 2 3]
```

---

## 🔹 8. Slice vs Array (Important Concept)

```go
func modify(s []int) {
	s[0] = 99
}

func main() {
	a := []int{1, 2, 3}
	modify(a)
	fmt.Println(a) // [99 2 3]
}
```

Slices internally store a pointer to an array, so passing a slice behaves like reference.

---

## 🔹 9. Pointer with Struct (Pass by Value)

```go
type User struct {
	Name string
}

func update(u User) {
	u.Name = "Changed"
}

func main() {
	user := User{Name: "Farhan"}
	update(user)
	fmt.Println(user.Name) // Farhan
}
```

The struct is copied; original data is unchanged.

---

## 🔹 10. Pointer with Struct (Pass by Reference)

```go
func update(u *User) {
	u.Name = "Changed"
}

func main() {
	user := User{Name: "Farhan"}
	update(&user)
	fmt.Println(user.Name) // Changed
}
```

### Visualization

```
user ─────┐
          ▼
       *User
          │
          └── Name = Changed
```

---

## 🔹 11. Pointer Receiver (Method)

```go
func (u *User) UpdateName(name string) {
	u.Name = name
}

user.UpdateName("Nadim")
```

Go automatically passes `&user` when calling the method.

---

## 🔹 12. new() vs make() vs & Operator

### new()

```go
p := new(int)
*p = 10
```

Allocates memory and returns a pointer.

---

### make() (Only for slice, map, channel)

```go
s := make([]int, 3)
```

Initializes internal structures.

---

### & Operator (Most Common)

```go
x := 10
p := &x
```

---

## ⚠️ Common Mistakes

* Using pointers unnecessarily
* Forgetting to dereference (`*p`)
* Using pointers with slices when not needed

---

## 🧾 Summary Table

| Concept         | Value                 | Pointer |
| --------------- | --------------------- | ------- |
| Copy data       | Yes                   | No      |
| Modify original | ❌                     | ✅       |
| Memory shared   | ❌                     | ✅       |
| Performance     | Slower for large data | Faster  |

---

## 🎯 Interview One-Liners

* Go is always pass-by-value
* Pointers pass memory addresses
* Slices contain pointers internally
* Struct modification requires pointers

---

## 🧠 Final Mental Model

```
Value   → Copy
Pointer → Address → Same Memory
```
