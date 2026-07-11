# WaitGroup in Go

# Previous Scenario

ধরুন আপনার Dashboard API-তে ৪টি service call করতে হবে।

1. User Service
2. Order Service
3. Notification Service
4. Analytics Service

প্রতিটি API call শেষ হতে প্রায়:

```text id="mtvkql"
8 seconds
```

সময় লাগে।

---

# Without Goroutine

```go id="pk9yol"
user := getUser()
orders := getOrders()
notifications := getNotifications()
analytics := getAnalytics()
```

Total time:

```text id="b7xk2r"
8 + 8 + 8 + 8
= 32 seconds
```

---

# With Goroutine

```go id="k7v19j"
go getUser()
go getOrders()
go getNotifications()
go getAnalytics()
```

Total time:

```text id="08zn4z"
≈ 8 seconds
```

---

# Problem

Handler জানে না:

```text id="4kn9r0"
সব goroutine শেষ হয়েছে কিনা।
```

Handler response পাঠিয়ে দিতে পারে goroutine শেষ হওয়ার আগেই।

---

# Solution: WaitGroup

WaitGroup মূলত একটি counter।

Main goroutine অপেক্ষা করে যতক্ষণ না counter শূন্য (0) হয়।

---

# Basic Example

```go id="d4u98m"
var wg sync.WaitGroup

wg.Add(4)

go func() {
	defer wg.Done()
	getUser()
}()

go func() {
	defer wg.Done()
	getOrders()
}()

go func() {
	defer wg.Done()
	getNotifications()
}()

go func() {
	defer wg.Done()
	getAnalytics()
}()

wg.Wait()
```

এরপর response return করা যাবে।

---

# How WaitGroup Works

WaitGroup-এর ভেতরে মূলত একটি counter থাকে।

```text id="z6sn65"
Counter = 0
```

---

## Step 1

```go id="hbyejd"
wg.Add(4)
```

Counter:

```text id="2iqqof"
Counter = 4
```

মানে:

```text id="vt4xmu"
আমার ৪টি goroutine finish হওয়া পর্যন্ত অপেক্ষা করতে হবে।
```

---

## Step 2

প্রথম goroutine শেষ:

```go id="pbhgtg"
wg.Done()
```

Counter:

```text id="yslqzv"
Counter = 3
```

---

দ্বিতীয় goroutine:

```text id="uvlwwf"
Counter = 2
```

---

তৃতীয় goroutine:

```text id="xok3tx"
Counter = 1
```

---

চতুর্থ goroutine:

```text id="e3yzon"
Counter = 0
```

---

এখন:

```go id="ly7g0o"
wg.Wait()
```

unblock হয়ে যাবে।

---

# Visualization

```text id="20o7lm"
wg.Add(4)

Counter = 4

Done()
Counter = 3

Done()
Counter = 2

Done()
Counter = 1

Done()
Counter = 0

wg.Wait() returns
```

---

# Internal Mechanism

WaitGroup internally:

* Counter রাখে।
* কত goroutine waiting করছে তা track করে।
* Counter শূন্য হলে waiting goroutine-গুলোকে wake up করে।

Conceptually:

```text id="t4mzsg"
state:
    counter
    waiters
    semaphore
```

---

# How Wait() Works

```go id="qqvtbg"
wg.Wait()
```

যদি:

```text id="x4fxvt"
counter > 0
```

তাহলে:

```text id="upyrtq"
main goroutine sleep করবে।
```

---

যখন:

```text id="zv0e95"
counter == 0
```

তখন:

```text id="8f7vku"
main goroutine wake up হবে।
```

---

# Add()

```go id="hylrhv"
wg.Add(n)
```

মানে:

```text id="b1d30f"
counter += n
```

---

উদাহরণ:

```go id="6yag3m"
wg.Add(4)
```

```text id="mcv87m"
counter = counter + 4
```

---

# Done()

```go id="9ppv1r"
wg.Done()
```

আসলে:

```go id="vvq0ed"
wg.Add(-1)
```

অর্থাৎ:

```text id="c4bl05"
counter--
```

---

# Wait()

```go id="48m4a4"
wg.Wait()
```

মানে:

```text id="yj4iv8"
while counter > 0
    sleep
```

---

# Timeline

```text id="t58k7e"
Request Started
      |
      +---- Goroutine 1
      |
      +---- Goroutine 2
      |
      +---- Goroutine 3
      |
      +---- Goroutine 4
      |
      |
      wg.Wait()
      |
All Done
      |
Return Response
```

---

# Common Panic #1

```go id="z7f4mw"
wg.Add(1)

go func() {
    panic("boom")
    wg.Done()
}()
```

সমস্যা:

```text id="j5o0gy"
Done() কখনো execute হবে না।
```

Counter:

```text id="8f4r7i"
1
```

Wait:

```text id="xk0x4z"
Forever
```

Deadlock।

---

# Solution

সবসময়:

```go id="v5chye"
defer wg.Done()
```

---

Correct:

```go id="0dh1wl"
go func() {
	defer wg.Done()

	panic("boom")
}()
```

panic হলেও:

```text id="7ukfnp"
Done() execute হবে।
```

---

# Common Panic #2

Negative WaitGroup Counter

```go id="of3q8s"
wg.Add(1)

wg.Done()
wg.Done()
```

Result:

```text id="25q8dm"
panic:
sync: negative WaitGroup counter
```

কারণ:

```text id="esd6th"
Counter = -1
```

যা invalid।

---

# Common Panic #3

```go id="0u6mgg"
go func() {
	wg.Add(1)
	defer wg.Done()
}()
```

কখনো কখনো:

```text id="7x7s13"
panic:
sync: WaitGroup misuse
```

কারণ:

`Add()` goroutine start হওয়ার পরে call হয়েছে।

---

# Correct Way

```go id="x0ylqk"
wg.Add(1)

go func() {
	defer wg.Done()
}()
```

---

# Best Practice

```go id="vf95s4"
wg.Add(n)

for i := 0; i < n; i++ {
	go func() {
		defer wg.Done()
	}()
}

wg.Wait()
```

---

# Dashboard API Example

```go id="iy0j6j"
var wg sync.WaitGroup

wg.Add(4)

go func() {
	defer wg.Done()
	user = getUser()
}()

go func() {
	defer wg.Done()
	orders = getOrders()
}()

go func() {
	defer wg.Done()
	notifications = getNotifications()
}()

go func() {
	defer wg.Done()
	analytics = getAnalytics()
}()

wg.Wait()
```

Total time:

```text id="zn4g7l"
≈ 8 seconds
```

---

# WaitGroup vs Channel

| Feature            | WaitGroup | Channel |
| ------------------ | --------- | ------- |
| Wait for goroutine | ✅         | ✅       |
| Pass data          | ❌         | ✅       |
| Synchronization    | ✅         | ✅       |
| Communication      | ❌         | ✅       |
| Collect results    | ❌         | ✅       |
| Signal completion  | ✅         | ✅       |

---

# When to Use WaitGroup

যখন শুধু:

```text id="3dq2sh"
goroutine শেষ হওয়া পর্যন্ত wait করতে হবে।
```

---

Example:

```go id="cvaxw3"
sendEmail()
generatePDF()
writeLog()
```

---

# When to Use Channel

যখন:

```text id="rjlwm1"
goroutine থেকে result ফেরত লাগবে।
```

---

Example:

```go id="swd7ia"
user := getUser()
orders := getOrders()
```

result collect করতে হবে।

---

# Production Pattern

অনেক সময় দুটো একসাথে ব্যবহার করা হয়।

```go id="lkgbmh"
var wg sync.WaitGroup
resultCh := make(chan User)
```

```go id="k0od7i"
go func() {
	defer wg.Done()
	resultCh <- getUser()
}()
```

```go id="n6c6j4"
wg.Wait()
close(resultCh)
```

---

# Final Takeaway

### WaitGroup

```text id="w4pn9w"
"Tell me when all goroutines are finished."
```

### Channel

```text id="5mpjck"
"Give me the result produced by goroutines."
```

### Together

```text id="jzmbcr"
WaitGroup → completion synchronization
Channel → data communication
```

Go concurrent programming-এ এই দুটো প্রায় সব production-grade API, microservice, worker pool, fan-out/fan-in এবং dashboard aggregation pattern-এর foundation।
