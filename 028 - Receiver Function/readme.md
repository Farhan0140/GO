# 📘 Go Receiver Function (Method)


## 🧠 What is a Receiver Function?

A **receiver function** (also called a **method**) is a function that is **attached to a specific type**.

> In Go, when a function has a receiver, it becomes a behavior of that type.

> A receiver function is a function that is bound to a type and can be called using dot (`.`) notation, making it a method of that type.


---

## 🔹 Normal Function vs Receiver Function

### ❌ Normal Function

```go
func Print_User_Details(usr User) {
	// usr is just a parameter
}
```

Call style:

```go
Print_User_Details(user1)
```

📌 The function does NOT belong to `User`.

---

### ✅ Receiver Function (Method)

```go
func (usr User) User_Details() {
	// usr is the receiver
}
```

Call style:

```go
user2.User_Details()
```

📌 The function **belongs to `User`**.

---

## 🧩 Receiver Syntax Breakdown

```go
func (usr User) User_Details() {}
```

| Part           | Meaning                                  |
| -------------- | ---------------------------------------- |
| `usr`          | Receiver variable (like `this` / `self`) |
| `User`         | Receiver type                            |
| `User_Details` | Method name                              |

---

## 🔁 How Method Call Works (Visualization)

### Struct Instance in Memory

```
user2
│
├── Name  → "Nadim"
├── Age   → 19
└── Phone → "01912345678"
```

### Method Binding to Type

```
User Type
│
├── User_Details()
└── Call(money int)
```

### Method Call Flow

```
user2.User_Details()
   │
   └── usr = copy of user2
         ├─ usr.Name
         ├─ usr.Age
         └─ usr.Phone
```

---

## ⚠️ Important: Value Receiver Behavior

Your methods use **value receivers**:

```go
func (usr User) User_Details()
```

✔ A **copy** of `User` is passed
✔ Original struct is NOT modified
✔ Best for read-only operations

---

## 🔑 Pointer Receiver (Advanced Concept)

To modify the original struct, use a pointer receiver:

```go
func (usr *User) UpdateAge(newAge int) {
	usr.Age = newAge
}
```

Call:

```go
user2.UpdateAge(25)
```

### Pointer Visualization

```
user2 ─────┐
           ▼
        *User
          │
          └── Age updated
```

---

## 📌 Why Receiver Functions Matter

* Enables **OOP-style design**
* Keeps related logic close to data
* Required to implement **interfaces**
* Improves code readability
* Widely used in real-world Go projects

---

## 🧾 Summary Table

| Feature           | Normal Function | Receiver Function |
| ----------------- | --------------- | ----------------- |
| Attached to type  | ❌               | ✅                 |
| Call syntax       | `func(x)`       | `x.func()`        |
| OOP style         | ❌               | ✅                 |
| Interface support | ❌               | ✅                 |
