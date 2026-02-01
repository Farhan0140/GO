## 1️⃣ `time.Now()`

### 👉 কী করে

বর্তমান **লোকাল সময়** রিটার্ন করে।

### 🧠 কখন ব্যবহার করি

* লগ টাইমস্ট্যাম্প
* বর্তমান সময় দেখাতে
* কোনো কাজ কখন হয়েছে সেটা ধরতে

### 🧪 কোড

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
}
```

### 📤 আউটপুট (উদাহরণ)

```
2026-02-01 15:42:10.123456 +0600 +06
```

---

## 2️⃣ `time.Date()`

### 👉 কী করে

নিজে হাতে নির্দিষ্ট তারিখ ও সময় বানায়।

### 🧠 কখন ব্যবহার করি

* ফিক্সড তারিখ দরকার হলে
* জন্মতারিখ, ইভেন্ট টাইম
* টাইম তুলনা করতে

### 🧪 কোড

```go
t := time.Date(2026, 2, 1, 10, 30, 0, 0, time.Local)
fmt.Println(t)
```

### 📤 আউটপুট

```
2026-02-01 10:30:00 +0600 +06
```

---

## 3️⃣ `time.Sleep()`

### 👉 কী করে

প্রোগ্রামকে নির্দিষ্ট সময়ের জন্য থামিয়ে রাখে।

### 🧠 কখন ব্যবহার করি

* Delay দিতে
* Retry logic
* Hardware / API wait

### 🧪 কোড

```go
fmt.Println("Start")
time.Sleep(2 * time.Second)
fmt.Println("End")
```

### 📤 আউটপুট

```
Start
(২ সেকেন্ড পরে)
End
```

---

## 4️⃣ `time.Since()`

### 👉 কী করে

কোনো সময় থেকে এখন পর্যন্ত কত সময় গেছে তা বলে।

### 🧠 কখন ব্যবহার করি

* Execution time মাপতে
* Performance check
* Timeout logic

### 🧪 কোড

```go
start := time.Now()
time.Sleep(1 * time.Second)
fmt.Println(time.Since(start))
```

### 📤 আউটপুট

```
1.0002345s
```

---

## 5️⃣ `time.Until()`

### 👉 কী করে

এখন থেকে ভবিষ্যতের কোনো সময় পর্যন্ত কত সময় বাকি আছে।

### 🧠 কখন ব্যবহার করি

* Countdown
* Event reminder
* Scheduler

### 🧪 কোড

```go
future := time.Now().Add(10 * time.Minute)
fmt.Println(time.Until(future))
```

### 📤 আউটপুট

```
9m59.999s
```

---

## 6️⃣ `time.Add()`

### 👉 কী করে

কোনো সময়ের সাথে সময় যোগ বা বিয়োগ করে।

### 🧠 কখন ব্যবহার করি

* Expiry time
* Token validity
* Deadline set

### 🧪 কোড

```go
now := time.Now()
after := now.Add(24 * time.Hour)
fmt.Println(after)
```

### 📤 আউটপুট

```
2026-02-02 15:42:10 +0600 +06
```

---

## 7️⃣ `time.Format()`

### 👉 কী করে

সময়কে সুন্দর স্ট্রিং আকারে দেখায়।

### 🧠 কখন ব্যবহার করি

* UI / API response
* Report
* Log

### 🧪 কোড

```go
now := time.Now()
fmt.Println(now.Format("2006-01-02 15:04:05"))
```

### 📤 আউটপুট

```
2026-02-01 15:42:10
```

🧠 **মনে রাখার ট্রিক**
👉 Go-তে format layout সবসময়
`2006-01-02 15:04:05`

---

## 8️⃣ `time.Parse()`

### 👉 কী করে

স্ট্রিং থেকে time বানায়।

### 🧠 কখন ব্যবহার করি

* User input time
* API date parse
* Database date

### 🧪 কোড

```go
t, _ := time.Parse("2006-01-02", "2026-02-01")
fmt.Println(t)
```

### 📤 আউটপুট

```
2026-02-01 00:00:00 +0000 UTC
```

---

## 9️⃣ `time.NewTicker()`

### 👉 কী করে

নির্দিষ্ট সময় পরপর কাজ চালায়।

### 🧠 কখন ব্যবহার করি

* Sensor reading
* Auto refresh
* Background task

### 🧪 কোড

```go
ticker := time.NewTicker(1 * time.Second)

for t := range ticker.C {
	fmt.Println("Tick at", t)
}
```

### 📤 আউটপুট

```
Tick at 15:42:11
Tick at 15:42:12
Tick at 15:42:13
```

---

## 🔟 `time.After()`

### 👉 কী করে

নির্দিষ্ট সময় পরে একবার সিগন্যাল দেয়।

### 🧠 কখন ব্যবহার করি

* Timeout
* Auto stop

### 🧪 কোড

```go
<-time.After(3 * time.Second)
fmt.Println("Time up!")
```

### 📤 আউটপুট

```
(৩ সেকেন্ড পরে)
Time up!
```

---

## 🔥 সংক্ষেপে টেবিল

| Function   | ব্যবহার     |
| ---------- | ----------- |
| `Now()`    | বর্তমান সময় |
| `Sleep()`  | Delay       |
| `Add()`    | সময় যোগ     |
| `Since()`  | সময় মাপা    |
| `Format()` | সময় দেখানো  |
| `Parse()`  | সময় পড়া     |
| `Ticker()` | বারবার কাজ  |
| `After()`  | Timeout     |

