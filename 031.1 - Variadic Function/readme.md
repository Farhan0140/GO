# Variadic Functions
## What is a Variadic Function?

A **variadic function** is a function that can accept a **variable number of
arguments** of the **same type**.

In simple terms:
> You can call the function with 0, 1, or many arguments.

---

## One-Line Definition

> A variadic function allows passing **zero or more values** using `...`.

---

## Syntax

```go
func functionName(param ...Type) {
    // param behaves like a slice
}
```

### Rules
- `...` is called the **variadic operator**
- Inside the function, `param` is a **slice**
- Only **one variadic parameter** is allowed
- The variadic parameter **must be the last parameter**

---

## Basic Example

```go
func Sum(nums ...int) int {
    total := 0
    for _, v := range nums {
        total += v
    }
    return total
}
```

### Usage

```go
fmt.Println(Sum(1, 2, 3))   // 6
fmt.Println(Sum(10, 20))    // 30
fmt.Println(Sum())          // 0
```

---

## What Happens Internally?

When you write:

```go
Sum(1, 2, 3)
```

The compiler converts it to:

```go
Sum([]int{1, 2, 3}...)
```

So internally:
- A slice is created
- Slice header (pointer, length, capacity) is passed
- The function receives a **slice**

---

## Memory Visualization

```
nums
 ├─ pointer → [1, 2, 3]
 ├─ len = 3
 └─ cap = 3
```

---

## Zero Arguments Case

```go
Sum()
```

Internally:

```
nums = []int{}
len = 0
cap = 0
```

This is an **empty slice**, not `nil`.

---

## Passing a Slice to a Variadic Function

```go
numbers := []int{4, 5, 6}
fmt.Println(Sum(numbers...))
```

- `numbers...` expands the slice
- Each element is passed as a separate argument

---

## Variadic with Other Parameters

```go
func PrintInfo(name string, scores ...int) {
    fmt.Println("Name:", name)
    fmt.Println("Scores:", scores)
}
```

---

## append() Is a Variadic Function

```go
func append(slice []T, elems ...T) []T
```

---

## Pass by Value or Reference

Slice header is copied, underlying array is shared.

```go
func Modify(nums ...int) {
    nums[0] = 99
}
```
