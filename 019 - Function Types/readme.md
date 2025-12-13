1. standard or named function 
1. Anonymous function
1. Function expression or Assign function in variable
1. IIFE - Immediately invoked function expression
1. Variadic function
1. init function [ you can not call this, computer calls this automatically ]
1. Closure - close over
1. Defer function - last in first out
1. Receiver function
1. Higher order function or first class function 
1. Callback function

<br>
<br>


**1) Standard / Named Function**
> যে function-এর নাম থাকে এবং সাধারণভাবে func name() আকারে ডিক্লেয়ার করা হয়।

```
func add(a int, b int) int {
    return a + b
}

```

**2) Anonymous Function**
>যে function-এর কোনো নাম নেই। সাধারণত inline ব্যবহার করা হয়।
```
func() {
    fmt.Println("Hello")
}() // IIFE – Immediately Invoked Function Expression

```


**3) Function Expression / Assign Function to Variable**
>একটি function-কে variable-এ assign করা।
```
square := func(x int) int {
    return x * x
}

fmt.Println(square(5))

```

**4) IIFE – Immediately Invoked Function Expression**
>যে anonymous function ডিক্লেয়ার হওয়ার সাথে সাথে call হয়।
```
func(msg string) {
    fmt.Println(msg)
}("Hello IIFE")

```


**5) Variadic Function**
>যে function-এ variable number of arguments নেওয়া যায়।
- syntax`...type`
```
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

func main() {
    fmt.Println(sum(1, 2, 3, 4))     // 10
    fmt.Println(sum(5))              // 5
    fmt.Println(sum())               // 0

    numbers := []int{10, 20, 30}
    fmt.Println(sum(numbers...))     // 60
}

```


**6) init Function**
>init() হলো একটি special function
>>- program শুরু হওয়ার আগে নিজে নিজে call হয়

>>- manually call করা যায় না

>>- একাধিক init() থাকতে পারে

>>- func init must have no arguments and no return values

>Call order
>>1. package-level variables
>>2. init()
>>3. main()

```
func main () {
	
}

func init() {
	fmt.Println("From init function")
}
```

**7) Closure (Close Over)**
>Closure হলো এমন function যা outer scope-এর variable মনে রাখে।
```
// closure function
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

func main() {
    // Create a new counter closure
    next := counter()

    // Call the closure several times
    fmt.Println(next()) // 1
    fmt.Println(next()) // 2
    fmt.Println(next()) // 3

    // You can make another independent counter
    another := counter()
    fmt.Println(another()) // 1
    fmt.Println(another()) // 2
}
```


**8) Defer Function (LIFO)**
>defer function-কে শেষে execute করার জন্য queue-তে রাখে
>> Last In, First Out (LIFO)
```
func main() {
    defer fmt.Println("World")
    defer fmt.Println("Hello")
}
```


**9) Receiver Function (Method)**
>Receiver function মানে method — যেটা কোনো type-এর সাথে attach থাকে।
```
type User struct {
    name string
}

func (u User) greet() string {
    return "Hello " + u.name
}

func main() {
    user := User{
        name: "Farhan",
    }

    message := user.greet()
    fmt.Println(message)
}
```


**● First-Order Function**
>A first-order function is simply a function where the inputs and outputs are non-function values. It operates on standard types and does not involve other functions as parameters or return values.
```
func add(a int, b int) int {
    return a + b
}

```


**10) Higher order function or first class function (Treated as First class citizen)**
>A higher-order function is a function that takes one or more functions as arguments, returns a function, or both. Go supports this capability because its functions are first-class citizens, meaning functions can be treated like any other value (assigned to variables, passed around, returned, etc.).
>>Rules
    >>1. Parameter --> Function
    >>2. Return --> Function
    >>3. Both

- A first-class citizen (also called a first-class value) is a broader concept. It refers to any entity in a language that can:
    - Be assigned to a variable
    - Be passed as a parameter
    - Be returned from a function
- Functions, integers, structs, slices, and maps are all first-class citizens in Go.
```
● Parameter --> Function

func add (num1 int, num2 int) {
	fmt.Println(num1 + num2)
}

func processOperation (a int, b int, op func(p int, q int)) {
	op(a, b)
}

func main () {
	processOperation(2, 3, add)		// 5
}

```

```
● Both

func add (num1 int, num2 int) {
	fmt.Println(num1 + num2)
}

func processOperation (a int, b int, op func(p int, q int)) func(x int, y int) {
	op(a, b)

    return add
}

func main () {
	processOperation(2, 3, add)		// 5
}
```

```
● Return --> Function

func add(x int, y int) {
    fmt.Println(x + y)
}

func call() func(p int, q int) {
    return add
}

func main() {
    sum := call()

    sum(4, 5) --> 9
}
```


**11) Callback Function**
>যে function-টি অন্য function-এর parameter হিসেবে পাঠানো হয় এবং পরে call করা হয়।
- A higher-order function may use a callback — but not always.
- A callback function requires a higher-order function to be passed into.
- Therefore:
    - All callbacks involve higher-order functions.
    - Not all higher-order functions necessarily involve callbacks, because some higher-order functions simply return functions without calling them.
```
func add (num1 int, num2 int) {
	fmt.Println(num1 + num2)
}

func processOperation (a int, b int, op func(p int, q int)) {
	op(a, b)
}

func main () {
	processOperation(2, 3, add)		// 5
}


```
- `processOperation` is a higher-order function because it accepts a function as a parameter.
- `add` becomes the callback function because:
    - it is passed into another function, and
    - it is executed (called back) inside that function.
