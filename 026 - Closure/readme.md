```
package main

import "fmt"

const a = 10
var p = 100


func Outer() func() {
	money := 100
	age := 30

	fmt.Println("Age: ", age)

	show := func () {
		money = money + a + p
		fmt.Println("Money: ", money)
	}

	return show
}

func Call() {
	inc1 := Outer()
	inc1()
	inc1()

	inc2 := Outer()
	inc2()
	inc2()
}

func main () {
	Call()
}

func init() {
	fmt.Println("--- From init ---")
}
```

<br>

# 🔧 1. Compilation Phase
The Go compiler (gc) performs several key steps:

## ✅ Lexing & Parsing
Your source becomes an AST (`Abstract Syntax Tree`).

## ✅ Type Checking & Constant Folding
- `const a = 10` → treated as a **compile-time constant**.
    - In generated code, `a` is inlined (replaced with literal `10`).
    - So `money + a + p` becomes `money + 10 + p` at compile time.
- Global `var p = 100` → stored in the initialized data section (`.data`) of the binary.

## ✅ Escape Analysis (Critical!)
Go analyzes whether variables escape the stack frame:

- age := 30 → only used in fmt.Println → stays on stack → freed when Outer() returns.
- money := 100 → captured by closure show → escapes to heap.

>You can verify this with:
>> `go build -gcflags="-m"`
>
>It will print:
>> `./main.go:13:2: moved to heap: money`

## ✅ Code Generation
- Closures are compiled as structs containing:
    - A pointer to the function code
    - Pointers to captured variables (here: `&money`)
- The binary now contains:
    - `.text`: all function machine code
    - `.data`: global `p = 100`
    - `.rodata`: string literals like `"Age: "`, and constant `a` (as immediate value)

<br>

| Code Element | Memory Segment | Why |
|--------------|----------------|-----|
| `const a = 10` | Code segment (or read-only data) | Constants are embedded directly into instructions or stored in `.rodata`. Immutable. |
| `var p = 100` | Data segment (initialized) | Global variable with known value at compile time → goes to `.data`. |
| `func Outer, Call, main, init` | Code segment | All executable instructions live here. Read-only. |
| `init()` function | Registered to run before `main()` | Part of Go's runtime initialization. |

<br>

# ⚡ 2. Execution Phase

Go runtime starts → runs `init()` → `main()` → `Call()`.

**Step 0: Runtime Initialization**
- Go runtime starts (sets up goroutines, GC, memory allocator).
- Runs all init() functions in dependency order → prints `--- From init ---`

**Step 1: init() runs**
```
--- From init ---
```

**Step 2: `Call()` is invoked**

Inside `Call()`:
```
inc1 := Outer()  // ← First call to Outer()
```

→ `Outer()` executes:

- **Stack frame** for `Outer()` is created.
- Creates local variables: `money = 100`, `age = 30` → live on stack (for this call).
- Prints: `Age: 30`
- Defines anonymous function `show` that captures `money`, `a`, and `p`.
- Returns `show` → but note: `money` would normally die when `Outer()` returns…

…but because `show` closes over `money`, Go promotes `money` to the heap!

>💡 Closure Rule: If a local variable is referenced by a returned closure, it’s allocated on the heap, not the stack.

-` money = 100` → escaped, so:
    - Runtime allocates `money` on heap (e.g., at address `0x10010`)
    - Stack holds a **pointer** to that heap location (but you don’t see it—Go hides this).

Then, closure `show` is created as:
```
type closure_struct struct {
    money_ptr *int  // points to 0x10010
}
```
- `show’s` code knows to use `*money_ptr + 10 + p`
- `Outer()` returns this closure → stack frame destroyed, but heap `money` lives on.

<br>

**-> Call inc1() Twice**

Each time:

- Load global `p` from **.data segment** (address known at link time).
- Load constant `10` as immediate value.
- Dereference `money_ptr` → read/update heap value.
- Print result.
After two calls: heap memory at `0x10010` holds `320`.
```
inc1() → Money: 100 + 10 + 100 = 210
inc1() → Money: 210 + 10 + 100 = 320
```

**Call `Outer()` Again → `inc2`**

- New **stack frame**
- New heap allocation for `money` (e.g., `0x10020`)
- New closure pointing to `0x10020`
- Independent of `inc1`!
```
inc2() → 210
inc2() → 320
```

So output:
```
--- From init ---
Age: 30
Money: 210
Money: 320
Age: 30
Money: 210
Money: 320
```

| Data / Code | Segment | Notes |
|-------------|---------|-------|
| `const a = 10` | Code / `.rodata` | Immutable, baked into binary |
| `var p = 100` | Data segment (`.data`) | Global, initialized |
| Function code (`main`, etc.) | Code segment | Executable instructions |
| `age` in `Outer()` | Stack | Not captured → dies on return |
| `money` in `Outer()` | Heap | Escaped to heap due to closure capture |
| Closure function objects | Heap | Function + captured environment (pointer to `money`) |
| `inc1`, `inc2` (variables) | Stack (in `Call`) | Local pointers to heap-allocated closures |

<br>

# 🧱 3. Memory Layout During Execution

```
┌──────────────────────────────┐
│        CODE SEGMENT          │ ← Read-only
│  - main(), Call(), Outer()   │
│  - show() (anonymous func)   │
│  - Machine instructions      │
└──────────────────────────────┘

┌──────────────────────────────┐
│        DATA SEGMENT          │ ← Read/Write, initialized
│  p = 100                     │ ← Global variable
└──────────────────────────────┘

┌──────────────────────────────┐
│        HEAP                  │ ← Dynamic, GC-managed
│  0x10010: money = 320        │ ← Captured by inc1's closure
│  0x10020: money = 320        │ ← Captured by inc2's closure
│  ...                         │
└──────────────────────────────┘

┌──────────────────────────────┐
│        STACK (Call())        │ ← Grows downward
│  inc2: closure → 0x20020     │ ← Points to heap closure object
│  inc1: closure → 0x20010     │
│  ...                         │
└──────────────────────────────┘

Each closure object (0x20010, 0x20020) contains:
   - func pointer → show() in code segment
   - env: { money_ptr → 0x10010 or 0x10020 }
```

<br>

# ♻️ Garbage Collection (GC)

- Go’s **GC** tracks heap allocations.
- Each `money` (heap) is referenced by its closure (`inc1` or `inc2`).
- When `Call()` ends, `inc1` and `inc2` go out of scope → closures become unreachable → **GC can reclaim** both closure and its `money`.


```
Closure মানে হচ্ছে পূর্বের কিছু মনে রাখা, like পুনর্জন্ম(Logically পুনর্জন্ম এর সাথে Closure সম্পর্ক নাই) এর মত,
ধরি ফাহিম আর সীমান্ত, একটা বোরখা পড়া মেয়ে কে খুব ভালোবাসে, সীমান্ত চায় না ফাহিম তাকে ভালোবাসুক তাই
সে ফাহিম কে খুন করে, ফাহিম মারা যাওয়ার সময় একটা প্রতিশ্রুতি রেখে যায়:
	“যে আমাকে হত্যা করেছে, তার নাম আমি ভুলব না।”
-----
ফাহিমের শরীর শেষ হয়ে গেল (মৃত্যু = outer function শেষ)।
কিন্তু তার আত্মা(Heap memory) শুধু একটাই জিনিস ধরে রাখল →
👉 সীমান্তই তাকে হত্যা করেছে
-----

এখন ফাহিম এর পুনর্জন্ম হয় হিংস্র কুকুর হয়ে, এখন সে আগের জীবন এর সব কিছু ভুলে গেলেও তাকে যে সীমান্ত খুন
করেছে এইটা তার আত্মা(Heap memory) মনে রেখেছে প্রতিশোধ নেওয়ার জন্য, কিন্তু যখন সে সীমান্তকে দেখে— 
	➡️ তার ভিতরের সেই সংরক্ষিত স্মৃতি সক্রিয় হয়
	➡️ সে বুঝতে পারে: “এই লোকটাই আমাকে শেষ করেছিল”, 
তাই সে সীমান্ত কে যেয়ে কামড়ে মেরে ফেলে, এই যে এই যে—
	- শরীর শেষ
	- জীবন শেষ
	- সময় পার হয়ে গেছে
	- তবুও নির্দিষ্ট একটা তথ্য মনে আছে
👉 এটাই Closure

Closure মানে: জীবন/ফাংশন শেষ, কিন্তু কিছু স্মৃতি জীবিত।

```
