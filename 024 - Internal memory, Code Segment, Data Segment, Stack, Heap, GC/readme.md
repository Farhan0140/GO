**1️⃣ Internal Memory**
>Internal Memory বলতে প্রোগ্রাম চলার সময় RAM-এর ভেতরে যেই জায়গাগুলো ব্যবহার হয়, সেগুলোর মোট ধারণা
- Code Segment
- Data Segment
- Stack
- Heap
<!-- <img> -->
---

**2️⃣ Code Segment (Text Segment)**

- এখানে থাকে তোমার লেখা Go কোডের compiled machine instructions
- এটা read-only (পরিবর্তন করা যায় না)
>
```
func add(a int, b int) int {
    return a + b
}

```
- এই add ফাংশনের actual machine code থাকবে Code Segment-এ

---

**3️⃣ Data Segment**
>এখানে থাকে Global & Package-level variable।
- Data Segment দুই ভাগে ভাগ করা যায়:
    - Initialized Data  `var x int = 10`
    - Uninitialized Data (BSS)  `var y int`

- প্রোগ্রাম শুরুতেই মেমোরি allocate হয়
- পুরো প্রোগ্রাম চলা পর্যন্ত থাকে


---

**4️⃣ Stack**
- Function call সম্পর্কিত data এখানে থাকে
- খুব দ্রুত কাজ করে (Fast)
- LIFO (Last In First Out)
- Stack এ কী থাকে?
    - Local variables
    - Function parameters
    - Return address

### What is a Stack Frame?
>A stack frame is the memory allocated on the stack for a single function call, containing its local data and execution context.
- A stack frame is a block of memory on the stack that is created every time a function is called.
- It stores everything that function needs to run and is destroyed when the function returns.

- stack frame holds:
    - Function parameters
    - Local variables
    - Return address (where to go back after the function finishes)
    - Saved registers / bookkeeping data (compiler/runtime use)

```
func sum(a int, b int) int {
    c := a + b
    return c
}

func main() {
    x := sum(2, 3)
}

```
### Stack behavior

```
When main() starts:

[ Stack Frame: main ]
- x
```

```
When sum(2,3) is called:

[ Stack Frame: sum ]
- a = 2
- b = 3
- c = 5
- return address
--------------------
[ Stack Frame: main ]
- x

```

```
When sum() returns:

[ Stack Frame: main ]
- x = 5

```
The `sum stack frame` is destroyed automatically.
- Key properties of a stack frame   
    - Created on function call
    - Destroyed on function return
    - Very fast (no GC involved)
    - LIFO order (Last In, First Out)
    - Exists per goroutine in Go

### Stack frame in Go (important notes)
- Go uses **goroutine stacks**
- Initial stack is small (~2KB)
- Stack **grows & shrinks automatically**
- Compiler uses **escape analysis** to decide whether a variable stays in the stack frame or moves to heap

```
func f() *int {
    x := 10   // cannot stay in stack frame
    return &x
}

x escapes → allocated on heap, not stack frame
```


---



**5️⃣ Heap**
- যখন data এর lifetime function শেষ হওয়ার পরও দরকার, তখন Heap ব্যবহার হয়
- Dynamic memory allocation
```
func createNumber() *int {
    x := 10
    return &x
}

```
এখানে x Heap এ যাবে কারণ function শেষ হলেও x দরকার


---

**6️⃣ GC (Garbage Collector)**
- GC হলো Go-এর স্বয়ংক্রিয় পরিষ্কারক
- Heap থেকে অপ্রয়োজনীয় memory মুছে ফেলে
- GC কীভাবে কাজ করে?
    - কোন object আর ব্যবহার হচ্ছে না → detect করে
    - সেই memory free করে দেয়
```
func test() {
    x := new(int)
    x = nil
}

```
x আর reference নেই → GC পরে memory free করবে