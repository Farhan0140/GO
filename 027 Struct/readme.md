# 📘 Go Struct 


## 🧱 Struct Definition

```go
type Person struct {
    Name   string
    Age    int
    Salary float32
    Dept   string
}
```

### 🔍 Explanation

Each `Person` instance contains:

* `Name`   → Person name
* `Age`    → Person age
* `Salary` → Monthly salary
* `Dept`   → Department name

---

## 🛠️ Instantiation Types

### ✅ Type‑1: Zero‑Value Instantiation

```go
var person1 Person
```

📌 All fields are initialized with **zero values**:

```
Name   = ""
Age    = 0
Salary = 0.0
Dept   = ""
```

Later, values are assigned manually:

```go
person1.Name = "Nadim"
person1.Age = 23
person1.Salary = 300.00
person1.Dept = "EEE"
```

🧠 Memory: Typically allocated on the **stack** (if it does not escape).

---

### ✅ Type‑2: Instantiation with Initialization (Composite Literal)

```go
person2 := Person{
    Name:   "Farhan",
    Age:    23,
    Salary: 100.00,
    Dept:   "CSE",
}
```

📌 Clean, readable, and the **most commonly used** style in Go projects.

---

### ✅ Type‑3: Declare First, Initialize Later

```go
var person3 Person
person3 = Person{
    Name:   "Sadi",
    Age:    27,
    Salary: 100000.50,
    Dept:   "CSE",
}
```

📌 Useful when declaration and initialization happen in different scopes.

---

## 🖨️ Struct Printing Methods

### 🔹 Output‑1: `fmt.Println()`

```go
fmt.Println(person1)
```

**Output:**

```
{Nadim 23 300 EEE}
```

📌 Default formatting, field names are hidden.

---

### 🔹 Output‑2: `fmt.Printf("%v")`

```go
fmt.Printf("%v", person2)
```

**Output:**

```
{Farhan 23 100 CSE}
```

📌 Same as `Println`, but more formatting control.

---

### 🔹 Output‑3: `fmt.Printf("%+v")` ⭐ (Recommended)

```go
fmt.Printf("%+v\n", person3)
```

**Output:**

```
{Name:Sadi Age:27 Salary:100000.5 Dept:CSE}
```

📌 Includes **field names** — best for debugging.

---

### 🔹 Output‑4: `fmt.Printf("%#v")`

```go
fmt.Printf("%#v\n", person3)
```

**Output:**

```
main.Person{Name:"Sadi", Age:27, Salary:100000.5, Dept:"CSE"}
```

📌 Shows **Go‑syntax representation** — useful for deep debugging and copy‑paste.

---

### 🔹 Output‑5: Custom / User‑Friendly Print

```go
fmt.Printf(
    "Name: %s\nAge: %d\nSalary: %.2f\nDepartment: %s",
    person1.Name,
    person1.Age,
    person1.Salary,
    person1.Dept,
)
```

**Output:**

```
Name: Nadim
Age: 23
Salary: 300.00
Department: EEE
```

📌 Best for **user‑facing output**.

---

## 🧠 Visualization (Conceptual)

```
Person (struct)
│
├── Name   → "Nadim"
├── Age    → 23
├── Salary → 300.00
└── Dept   → "EEE"
```

Each variable (`person1`, `person2`, `person3`) is a **separate instance** of the same struct layout.


![image alt]()

---

## ✅ Key Takeaways

* `var p Person` → zero‑value instantiation
* `Person{}` → composite literal (most common)
* `%+v` is the **most practical format** for struct debugging
* `%#v` helps understand Go internals
* Field‑wise printing is ideal for real users


