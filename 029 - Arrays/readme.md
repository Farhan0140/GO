# 📘 Go Array

## 🔹 What is an Array in Go?

An **array** in Go is a collection of **fixed-size**, **same-type** elements stored in **contiguous memory**.

✔ Size is fixed at compile time

✔ Elements are indexed starting from 0

✔ Arrays are **value types** (copied on assignment)

---

## 1️⃣ Zero-Value Array

```go
var arr [5]int
```

All elements are initialized with the **zero value** of the type.

### Visualization

```
Index:  0   1   2   3   4
Value:  0   0   0   0   0
```

---

## 2️⃣ Assigning Value to a Specific Index

```go
arr[2] = 99
```

### Visualization

```
Index:  0   1   2   3   4
Value:  0   0  99   0   0
```

Array index access is **O(1)** and very fast.

---

## 3️⃣ Full Array Initialization

```go
arr1 := [5]int{1, 2, 3, 4, 5}
```

### Visualization

```
Index:  0   1   2   3   4
Value:  1   2   3   4   5
```

---

## 4️⃣ Partial / Indexed Initialization

```go
arr2 := [10]int{1: 99, 5: 100, 7: 3}
```

Only specified indices are initialized; others get zero values.

### Visualization

```
Index:  0   1   2   3   4   5   6   7   8   9
Value:  0  99   0   0   0 100   0   3   0   0
```

---

## 5️⃣ Length of an Array

```go
len(arr1) // 5
len(arr2) // 10
```

`len()` is a **compile-time constant** for arrays and runs in **O(1)** time.

---

## 6️⃣ Iterating Using Traditional `for` Loop

```go
for i := 0; i < len(arr2); i++ {
	fmt.Println(arr2[i])
}
```

Best when you need full control over indices.

---

## 7️⃣ Iterating Using `range`

```go
for idx, val := range arr1 {
	fmt.Printf("index [%d] -> value [%d]\n", idx, val)
}
```

### Range Iteration Visualization

```
idx = 0 → val = 1
idx = 1 → val = 2
idx = 2 → val = 3
idx = 3 → val = 4
idx = 4 → val = 5
```

---

## ⚠️ Important Concept: Arrays are Value Types

```go
a := [3]int{1, 2, 3}
b := a
b[0] = 99
```

Result:

```
a → [1 2 3]
b → [99 2 3]
```

Changes in `b` do NOT affect `a`.

---

## 🧠 Memory Visualization

```
arr1 (stack memory)
┌────┬────┬────┬────┬────┐
│ 1  │ 2  │ 3  │ 4  │ 5  │
└────┴────┴────┴────┴────┘
```

Arrays use **contiguous memory**, enabling fast access.

---

## 🧾 Summary Table

| Feature     | Go Array   |
| ----------- | ---------- |
| Size        | Fixed      |
| Type        | Value      |
| Zero Value  | Yes        |
| Memory      | Contiguous |
| Performance | Very Fast  |
| Flexibility | ❌          |

---

## 🚫 When NOT to Use Arrays?

* When size is dynamic
* When frequent insert/delete operations are needed

👉 Use **slices** instead in most real-world Go programs.

