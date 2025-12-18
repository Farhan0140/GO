# 📘 Go Slice
## 🔹 1. What is a Slice?

A **slice** in Go is a **dynamic, flexible view over an array**.

Important facts:

* A slice does **not own data**
* A slice **references an underlying array**
* A slice is lightweight and cheap to copy

> **Array owns data, Slice views data**

---

## 🔹 2. Slice Declaration (Nil Slice)

```go
var s []int
fmt.Println(s)        // []
fmt.Println(s == nil) // true
```

### Explanation

* Slice is declared but not initialized
* No underlying array exists
* This is called a **nil slice**

---

## 🔹 3. Slice Literal Initialization

```go
s := []int{1, 2, 3}
fmt.Println(s) // [1 2 3]
```

### Visualization

```
Slice Header
┌──────────────┐
│ ptr ─────┐   │
│ len = 3  │   │
│ cap = 3  │   │
└──────────────┘
             ▼
Underlying Array
┌───┬───┬───┐
│ 1 │ 2 │ 3 │
└───┴───┴───┘
```

---

## 🔹 4. Slice Internal Structure (VERY IMPORTANT)

Internally, Go slice is represented as:

```go
type slice struct {
    ptr *T
    len int
    cap int
}
```

Meaning:

* `ptr` → pointer to the first element of the array
* `len` → number of accessible elements
* `cap` → total capacity from pointer position

---

## 🔹 5. len vs cap

```go
fmt.Println(len(s)) // usable elements
fmt.Println(cap(s)) // allocated capacity
```

* `len` controls iteration
* `cap` controls growth using `append`

---

## 🔹 6. make() Function (Most Used)

```go
s := make([]int, 3, 5)
```

Meaning:

* Length = 3
* Capacity = 5

### Visualization

```
Underlying Array (size 5)
┌───┬───┬───┬───┬───┐
│ 0 │ 0 │ 0 │ _ │ _ │
└───┴───┴───┴───┴───┘
        ↑
      len=3
```

---

## 🔹 7. Creating Slice from Array

```go
arr := [5]int{10, 20, 30, 40, 50}
s := arr[1:4]
```

### Visualization

```
Array
┌───┬───┬───┬───┬───┐
│10 │20 │30 │40 │50 │
└───┴───┴───┴───┴───┘
      ▲       ▲
      1       4

Slice → [20 30 40]
```

The slice and array share the **same memory**.

---

## 🔹 8. Slice Modification Effect

```go
s[0] = 99
fmt.Println(arr) // [10 99 30 40 50]
```

Modifying slice modifies the array.

---

## 🔹 9. Slice Passed to Function

```go
func modify(s []int) {
    s[0] = 100
}
```

Explanation:

* Slice header is copied (ptr, len, cap)
* Underlying array is shared
* Original data is modified

> Go is **always pass-by-value**

---

## 🔹 10. append() — The Heart of Slice

```go
s := []int{1,2,3}
s = append(s, 4)
```

Two cases:

### Case 1: Capacity Available

* Same array is reused

### Case 2: Capacity Exceeded

* New array allocated
* Old data copied
* Slice pointer updated

---

## 🔹 11. append() Reallocation Example

```go
s := make([]int, 0, 2)
s = append(s, 1)
s = append(s, 2)
s = append(s, 3)
```

### Visualization

```
Old Array (cap=2)
[1 2]

New Array (cap=4)
[1 2 3 _]
```

---

## 🔹 12. Slice Sharing Trap (INTERVIEW FAVORITE)

```go
a := []int{1,2,3,4}
b := a[:2]
b = append(b, 99)
```

### Result

```
a = [1 2 99 4]
b = [1 2 99]
```

Why?

* Same underlying array
* append overwrote element

---

## 🔹 13. Avoid Slice Sharing Bugs

```go
b := append([]int{}, a[:2]...)
```

Forces a new array allocation.

---

## 🔹 14. Nil Slice vs Empty Slice

```go
var s1 []int   // nil slice
s2 := []int{}  // empty slice
```

| Property | Nil Slice | Empty Slice |
| -------- | --------- | ----------- |
| len      | 0         | 0           |
| cap      | 0         | 0           |
| == nil   | true      | false       |

Important for JSON & APIs.

---

## 🔹 15. Slice of Struct

```go
type User struct {
    Name string
}

users := []User{{"Farhan"}, {"Nadim"}}
```

Structs are stored inside the underlying array.

---

## 🔹 16. Pointer vs Slice

❌ Avoid this:

```go
func f(s *[]int)
```

✅ Use this:

```go
func f(s []int)
```

Slices already contain pointers internally.

---

## 🔹 17. Complete Memory Model

```
Slice Header
┌──────────────┐
│ ptr ─────┐   │
│ len      │   │
│ cap      │   │
└──────────────┘
             ▼
Underlying Array
┌─────────────────┐
│ actual elements │
└─────────────────┘
```

---

## 🧾 Array vs Slice Comparison

| Feature      | Array     | Slice       |
| ------------ | --------- | ----------- |
| Size         | Fixed     | Dynamic     |
| Copy         | Full copy | Header copy |
| Pass to func | Expensive | Cheap       |
| Usage        | Rare      | Everywhere  |

---

## 🎯 Interview One-Liners

* Slice is a descriptor over an array
* Slice contains pointer, length, capacity
* append may reallocate memory
* Slice behaves like reference but is pass-by-value

---

## 🧠 Golden Rule

```
Array owns data
Slice views data
```

<br><br><br>

# 📘 Go Slice Append & Sharing

## 🔹 Code

```go
var x []int

x = append(x, 1) // [1], len:1, cap:1
x = append(x, 2) // [1,2], len:2, cap:2
x = append(x, 3) // [1,2,3], len:3, cap:4

y := x // y = [1,2,3], len:3, cap:4

x = append(x, 4) // [1,2,3,4], len:4, cap:4
y = append(y, 5) // [1,2,3,5], len:4, cap:4

x[0] = 99

fmt.Println(x)
fmt.Println(y)
```

---

## 🔹 Step 1: Slice Declaration

```go
var x []int
```

* `x` is a **nil slice**
* No underlying array exists
* `len = 0`, `cap = 0`

```
x
┌──────────────┐
│ ptr = nil    │
│ len = 0      │
│ cap = 0      │
└──────────────┘
```

---

## 🔹 Step 2: `x = append(x, 1)`

* No array exists → Go allocates a new array
* Capacity becomes `1`

```
Underlying Array A (cap=1)
┌───┐
│ 1 │
└───┘

x → ptr → A
len = 1, cap = 1
```

---

## 🔹 Step 3: `x = append(x, 2)`

* Capacity full (`cap=1`)
* Go allocates a new array
* Existing data copied

```
Underlying Array B (cap=2)
┌───┬───┐
│ 1 │ 2 │
└───┴───┘

x → ptr → B
len = 2, cap = 2
```

---

## 🔹 Step 4: `x = append(x, 3)`

* Capacity full again
* Go **doubles capacity** for performance

```
Underlying Array C (cap=4)
┌───┬───┬───┬───┐
│ 1 │ 2 │ 3 │ _ │
└───┴───┴───┴───┘

x → ptr → C
len = 3, cap = 4
```

---

## 🔹 Step 5: `y := x`

⚠️ **Critical Rule**

* Only the slice header is copied
* Underlying array is NOT copied

```
x ──┐
    ├──► Underlying Array C
y ──┘

Both:
len = 3, cap = 4
```

---

## 🔹 Step 6: `x = append(x, 4)`

* Capacity available
* No new array allocation
* Element written at index `3`

```
Underlying Array C
┌───┬───┬───┬───┐
│ 1 │ 2 │ 3 │ 4 │
└───┴───┴───┴───┘

x: len = 4, cap = 4
y: len = 3, cap = 4
```

`y` cannot see `4` yet because its length is still `3`.

---

## 🔹 Step 7: `y = append(y, 5)`

* Capacity still available
* Append happens at index `len(y) = 3`
* This **overwrites** the existing value

```
Underlying Array C
┌───┬───┬───┬───┐
│ 1 │ 2 │ 3 │ 5 │
└───┴───┴───┴───┘

x: [1 2 3 5]
y: [1 2 3 5]
```

⚠️ Value `4` is permanently lost.

---

## 🔹 Step 8: `x[0] = 99`

* Direct modification of underlying array
* Both slices observe the change

```
Underlying Array C
┌────┬───┬───┬───┐
│ 99 │ 2 │ 3 │ 5 │
└────┴───┴───┴───┘
```

---

## 🔹 Final Output

```go
[99 2 3 5]
[99 2 3 5]
```

### ❓ Why `x` is NOT `[99 1 2 3 4]`

* `y = append(y, 5)` overwrote index `3`
* That index previously contained `4`

---

### ❓ Why `y` is NOT `[1 2 3 5]`

* `x[0] = 99` modified shared memory
* Both slices share the same array

---

## 🔹 Why Capacity Grows (2× Rule)

### Reason: Performance Optimization

* Frequent reallocation is expensive
* Doubling capacity makes append **amortized O(1)**

### Simplified Growth Rule

| Current Capacity | New Capacity |
| ---------------- | ------------ |
| < 1024           | ×2           |
| ≥ 1024           | +25%         |

> Exact behavior may vary by Go version

---

## 🔹 Role of Underlying Array

* Stores actual data
* Slice only describes a window over it
* Multiple slices can share it
* append may reuse or replace it

---

## 🔹 How to Avoid This Bug

```go
y := append([]int{}, x...)
```

* Forces a new array allocation
* Prevents memory sharing

---

## 🧠 Golden Rule

```
Slice header is copied
Underlying array is shared
```

---

## 🎯 Interview One-Liners

* Slice append may overwrite shared data
* Capacity controls memory reuse
* Slice is pass-by-value but reference-like
* Most slice bugs come from shared arrays


<br><br><br>

# 📘 Go Slice Sharing, Append & Capacity
## 🔹 Given Code

```go
func Change_Slice( p []int ) []int {
    p[0] = 100
    p = append(p, 99)

    return p
}

func main () {
    x := []int{1,2,3,4,5}
    x = append(x, 6)
    x = append(x, 7)

    a := x[4:]

    y := Change_Slice(a)

    fmt.Println(x)
    fmt.Println(y)

    fmt.Println(x[0:8])
}
```

---

## 🔹 Step 1: Create slice `x`

```go
x := []int{1,2,3,4,5}
x = append(x, 6)
x = append(x, 7)
```

### Explanation

* Slice literal creates an underlying array
* Appends extend the same array
* Final state:

```
x = [1 2 3 4 5 6 7]
len(x) = 7
cap(x) = 8   // important
```

### Visualization

```
Underlying Array A (cap=8)
index:  0 1 2 3 4  5  6  7
value: [1 2 3 4 5  6  7  _]
```

---

## 🔹 Step 2: Create slice `a` from `x`

```go
a := x[4:]
```

### What this means

* `a` starts from index `4` of `x`
* It shares the **same underlying array**

```
a = [5 6 7]
len(a) = 3
cap(a) = 4   // from index 4 to end
```

### Visualization

```
Underlying Array A
[1 2 3 4 | 5 6 7 _]
            ▲
            a starts here
```

---

## 🔹 Step 3: Call `Change_Slice(a)`

```go
y := Change_Slice(a)
```

### Important Rule

> Go is **pass‑by‑value**

* Slice header `(ptr, len, cap)` is copied
* Underlying array is **shared**

---

## 🔹 Step 4: Inside `Change_Slice` — `p[0] = 100`

```go
p[0] = 100
```

### What does this modify?

```
p[0] → a[0] → x[4]
```

### Memory after modification

```
Underlying Array A
[1 2 3 4 | 100 6 7 _]
```

⚠️ `x` is already modified at this point.

---

## 🔹 Step 5: Inside `Change_Slice` — `append(p, 99)`

```go
p = append(p, 99)
```

### Capacity Check

```
len(p) = 3
cap(p) = 4
```

* Capacity available
* **No new array allocated**
* Append writes at index `7`

### Memory after append

```
Underlying Array A
[1 2 3 4 | 100 6 7 99]
```

```
p (returned as y)
len = 4
cap = 4
```

---

## 🔹 Step 6: Function Returns

```go
return p
```

* `y` now points to the same underlying array
* `a`, `x`, `y` all share memory

---

## 🔹 Final State of All Slices

### x

```
x = [1 2 3 4 100 6 7]
len = 7
cap = 8
```

### y

```
y = [100 6 7 99]
```

### a

```
a = [100 6 7]
```

---

## 🔹 Why `fmt.Println(x[0:8])` Works

```go
fmt.Println(x[0:8])
```

### Rule

```
0 ≤ low ≤ high ≤ cap(slice)
```

* `len(x) = 7`
* `cap(x) = 8`

✔️ `0:8` is valid because slicing uses **capacity**, not length.

### Result

```
[1 2 3 4 100 6 7 99]
```

---

## 🔹 Final Output

```text
[1 2 3 4 100 6 7]
[100 6 7 99]
[1 2 3 4 100 6 7 99]
```

---

## 🧠 Core Rules Reinforced

1. Slice header is copied
2. Underlying array is shared
3. append reuses array if capacity allows
4. Modifying one slice affects all sharers
5. Slicing upper bound uses capacity, not length

---

## 🎯 Interview One‑Liners

* Slice is pass‑by‑value but reference‑like
* append may or may not allocate new memory
* Capacity controls memory reuse
* Most slice bugs come from shared arrays
