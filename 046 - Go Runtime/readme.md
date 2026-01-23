# epoll: Linux Kernel I/O Event Notification
## 1️⃣ epoll কী?

👉 **epoll হলো Linux kernel-এর একটি high-performance I/O event notification mechanism**
- Linux -> epoll
- Windows -> IOCP
- macOS / BSD -> kqueue
- Go language -> Runtime netpoll (epoll / kqueue / IOCP internally)

সহজভাবে:

> epoll kernel-কে জিজ্ঞেস করে — "অনেক FD-এর মধ্যে কোনগুলো এখন ready?"

Ready মানে:

* socket readable
* socket writable
* error / close event

---

## 2️⃣ কেন epoll দরকার?

ধরো:

* Server এ 10,000 socket
* সব socket সবসময় active না

❌ একটার পর একটা check করলে:

* CPU waste
* scalability নেই

### পুরনো system call সমস্যা

**select / poll**:

* প্রতিবার সব FD scan
* O(n) complexity
* FD limit

---

## 3️⃣ epoll কীভাবে problem solve করে?

👉 epoll event-driven

Kernel নিজে ready FD track করে
User space শুধু ready FD পায়

Result:

* No scanning
* Low CPU
* High performance

---

## 4️⃣ epoll vs select / poll

| Feature       | select | poll   | epoll |
| ------------- | ------ | ------ | ----- |
| FD scan       | Yes    | Yes    | No    |
| FD limit      | Yes    | No     | No    |
| Performance   | Low    | Medium | High  |
| Kernel notify | No     | No     | Yes   |
| Go uses       | No     | No     | Yes   |

---

## 5️⃣ epoll Kernel Internals

Kernel epoll instance রাখে:

* **Interest List** → Red-Black Tree
* **Ready List** → Linked List

```
epoll instance
 ├─ interest tree
 │    ├─ FD 3
 │    ├─ FD 7
 │    └─ FD 9
 └─ ready list
      ├─ FD 7 (readable)
      └─ FD 9 (writable)
```

---

## 6️⃣ epoll_create()

```c
int epfd = epoll_create1(0);
```

👉 Kernel-এ একটা epoll instance তৈরি করে
👉 Return করে একটা FD (epfd)

* একবার call হয়
* Server lifetime জুড়ে থাকে

---

## 7️⃣ epoll_ctl()

```c
epoll_ctl(epfd, EPOLL_CTL_ADD, fd, &event);
```

👉 Kernel-কে বলে:

* কোন FD
* কোন event (EPOLLIN / EPOLLOUT)

Operations:

* ADD
* MOD
* DEL

Kernel interest tree update হয়

---

## 8️⃣ epoll_wait() — আসল magic

```c
n = epoll_wait(epfd, events, max, timeout);
```

👉 Thread কে sleep করিয়ে দেয়
👉 Busy wait করে না

Timeline:

```
Thread calls epoll_wait
 ↓
Thread sleeps 😴
 ↓
NIC data arrives
 ↓
Kernel marks FD ready
 ↓
FD added to ready list
 ↓
epoll_wait returns
```

---

## 9️⃣ কখন কোন epoll call হয়?

```
Server start:
  epoll_create()

New FD accept:
  epoll_ctl(ADD)

Client sends data:
  Kernel marks FD ready

Runtime waits:
  epoll_wait()
```

---

## 🔟 Go Runtime + epoll

👉 Go নিজে epoll call করে না
👉 Go runtime করে

```
runtime.netpoll()
 └─ epoll_wait()
```

Flow:

```
FD ready
 ↓
epoll_wait returns
 ↓
runtime marks goroutine runnable
 ↓
goroutine scheduled
```

📌 OS thread block হয় না
📌 goroutine block হয়

---

## 1️⃣1️⃣ Level-triggered vs Edge-triggered

* Level-triggered (default)

  * data থাকলে বারবার notify

* Edge-triggered

  * state change হলে notify

Go ব্যবহার করে:
👉 Level-triggered (safe)

---

## 🎯 Interview One-liner

> epoll হলো Linux kernel-এর event notification system যা অনেক FD-এর মধ্যে শুধু ready FD-এর খবর দেয়; epoll_create instance বানায়, epoll_ctl FD register করে, আর epoll_wait thread কে sleep করিয়ে kernel event এ wake করে—এই কারণেই Go server হাজার হাজার connection efficiently handle করতে পারে।

---
</br>
</br>
</br>
</br>
</br>
</br>

# Go Runtime Initialization – Scheduler, epoll, GC


1. Initialize Go Scheduler
2. Go runtime → syscall → kernel → `epoll_create`
3. Setup Garbage Collector (GC)

---

## Overall Program Startup Flow

Go প্রোগ্রাম চালু হলে মোটামুটি নিচের ধাপে কাজ হয়:

```
OS → Go runtime start → runtime initialization → main.main()
```

তোমার উল্লেখ করা তিনটি কাজই ঘটে **runtime initialization phase-এর মধ্যে**, কিন্তু **একই সময়ে নয়** — আলাদা ধাপে।

---

## 1️⃣ Initialize Go Scheduler

### 👉 কখন হয়?

* **প্রোগ্রাম শুরু হওয়ার একদম শুরুতেই**
* `main()` ফাংশন কল হওয়ার **আগে**

### এই ধাপে কী ঘটে?

Go runtime নিচের abstraction গুলো initialize করে:

* **M:P:G**
* **M (OS Thread)**
* **P (Processor / Logical CPU)**
* **G (Goroutine)**

এছাড়াও:

* `GOMAXPROCS` অনুযায়ী কতগুলো `P` লাগবে তা নির্ধারণ করে
* প্রথম OS thread (`M0`) তৈরি হয়
* Scheduler চালু হয়

### সহজভাবে বলা যায়:

> “Go এখন goroutine চালানোর জন্য সম্পূর্ণ প্রস্তুত”

📌 **Scheduler initialize না হলে Go কোড চলতেই পারবে না** — তাই এটা সবসময় হয়।

---

## 2️⃣ Go runtime → syscall → kernel → `epoll_create`

### 👉 কখন হয়?

* Scheduler initialize হওয়ার **পরে**
* **শুধু Linux-এ**
* এবং সবচেয়ে গুরুত্বপূর্ণ: **Lazy / On-demand ভাবে**

### গুরুত্বপূর্ণ তথ্য

Go runtime **প্রোগ্রাম শুরুতেই `epoll_create` কল করে না**।

### তাহলে `epoll_create` কখন কল হয়?

যখন Go runtime বুঝে যে:

* Network I/O লাগবে
* অথবা timer / netpoll দরকার

তখন এই flow হয়:

```
runtime → syscall → kernel → epoll_create
```

### উদাহরণ

এই কোডে `epoll_create` হবেই:

```go
net.Listen("tcp", ":8080")
```

এই কোডে `epoll_create` লাগবে না:

```go
fmt.Println("Hello")
```

📌 তাই বলা যায়:

> **`epoll_create` = On-demand (Lazy initialization)**

---

## 3️⃣ Setup Garbage Collector (GC)

### 👉 কখন হয়?

* Scheduler initialize হওয়ার **পরই**
* কিন্তু **GC তখনই run করে না**

### GC setup বলতে কী বোঝায়?

এই ধাপে Go runtime:

* Heap structure তৈরি করে
* GC metadata initialize করে
* Write barrier enable করে
* GC controller প্রস্তুত করে

### GC কখন actually run করে?

GC চলা শুরু করে তখনই, যখন:

* Heap allocation বাড়ে
* Allocation threshold cross করে
* অথবা `runtime.GC()` manually কল করা হয়

📌 অর্থাৎ:

> **GC engine প্রস্তুত থাকে**, কিন্তু **প্রয়োজন না হলে চলে না**

---

## Runtime Initialization Timeline (সহজ ভিজ্যুয়াল)

```
Program start
   ↓
Initialize Go runtime
   ↓
1️⃣ Initialize Scheduler (G / M / P)
   ↓
3️⃣ Setup GC (GC ready, not running)
   ↓
main.main() call
   ↓
2️⃣ epoll_create (only if network/timer needed)
```

---

## Summary Table

| কাজ                  | কখন হয়        | সবসময় হয়? |
| -------------------- | ------------- | --------- |
| Scheduler initialize | Program start | ✅ Yes     |
| GC setup             | Program start | ✅ Yes     |
| `epoll_create`       | On-demand     | ❌ No      |

---

## One-line Summary 🧠

> **Scheduler এবং GC setup হয় প্রোগ্রাম শুরুতেই, কিন্তু `epoll_create` হয় তখনই যখন Go runtime-এর I/O বা netpoll দরকার হয়।**

---

## Notes

* এই ব্যাখ্যা Linux-based Go runtime-এর জন্য প্রযোজ্য
* Windows বা BSD-তে netpoll mechanism আলাদা হতে পারে
* Go runtime-এর অনেক অংশ lazy initialization follow করে performance optimize করার জন্য

---
</br>
</br>
</br>
</br>
</br>
</br>


# Go Scheduler Internals — Interview‑Ready & Production‑Level Guide
---

## 1. The M–P–G Model (Core of Go Scheduling)

Go uses a **three‑entity scheduling model** to efficiently multiplex goroutines onto OS threads.

---

### 🔹 G — Goroutine

* Lightweight execution unit created using the `go` keyword
* Starts with a **small stack (2–4 KB)**
* Stack **grows and shrinks dynamically**
* Very cheap compared to OS threads
* Applications can easily have **thousands to millions** of goroutines

📌 **Interview Tip:** Goroutines are *not threads*; they are scheduled by the Go runtime.

---

### 🔹 M — Machine (OS Thread)

* Represents an **actual OS thread**
* Managed by the kernel scheduler
* Executes Go code **only when it holds a P**
* Can block when:

  * A syscall is executed
  * A `cgo` call is made
  * Blocking I/O occurs

📌 If an M blocks, the Go runtime may **create or wake another M** so execution continues.

---

### 🔹 P — Processor (Scheduler Token)

* A **logical processor**, not a CPU core
* Represents **permission to execute Go code**
* Owns run queues and scheduler state
* Without a P, an M **cannot run Go code**

📌 **Key Insight:** P is the most important concept — it controls parallelism.

---

### 🔑 The Golden Rule (Very Important)

> **At any time:**
> `Running Goroutines ≤ Number of P ≤ Number of M`

📌 This invariant is often asked directly in interviews.

---

## 2. How Many P Exist? (GOMAXPROCS)

The number of P is controlled by:

```go
GOMAXPROCS
```

### Default Behavior

* `GOMAXPROCS = number of CPU cores`

### Examples

| CPU Cores | GOMAXPROCS | Number of P |
| --------- | ---------- | ----------- |
| 4         | default    | 4           |
| 4         | 2          | 2           |

📌 **Interview Tip:** Changing `GOMAXPROCS` limits *parallel execution*, not concurrency.

---

## 3. How Many M Exist?

* Minimum: **1 (M0)**
* Maximum: **Dynamic (no hard limit)**

### New M Is Created When:

* A goroutine blocks in a syscall
* A `cgo` call blocks
* A goroutine waits on network I/O

📌 Therefore, in practice:

> **Usually: M ≥ P**

---

## 4. Run Queues — Where Goroutines Wait
* slot of local run-queue: 256

Each P owns **exactly one Local Run Queue**.

```
P0 → Local Run Queue
P1 → Local Run Queue
P2 → Local Run Queue
P3 → Local Run Queue
```

Additionally, the runtime maintains a:

```
Global Run Queue (GRQ)
```

---

## 5. Local vs Global Run Queue

| Queue            | Owned By     | Characteristics | Purpose                   |
| ---------------- | ------------ | --------------- | ------------------------- |
| Local Run Queue  | Individual P | Lock‑free, fast | Primary scheduling        |
| Global Run Queue | Runtime      | Locked, slower  | Load balancing & fairness |

### Scheduling Priority Order

1. Local Run Queue (fastest)
2. Global Run Queue
3. Work stealing from another P

📌 **Interview Tip:** Local queues exist for performance; global queue exists for fairness.

---

## 6. Local Run Queue — Internal Structure

The Local Run Queue is implemented as:

* A **fixed‑size circular (ring) buffer**
* Size = **256 slots** (runtime constant)
* Each slot holds a pointer to a `*G`

```
[ G ][ G ][ G ][   ][   ] ...
  ↑           ↑
 head        tail
```

### Why This Matters

* No heap allocation
* No locks
* Cache‑friendly
* Constant‑time operations

---

## 7. Enqueue & Dequeue Logic

### Enqueue (goroutine becomes runnable)

```
tail++
runq[tail % 256] = G
```

### Dequeue (scheduler picks next goroutine)

```
G = runq[head % 256]
head++
```

📌 **Key Point:** Modulo arithmetic enables the circular behavior.

---

## 8. Local Run Queue = Ring Buffer (Confirmed)

Yes — the Local Run Queue **is a ring buffer**.

### Reasons Go Uses a Ring Buffer

* Extremely fast modulo operations
* Predictable memory access
* Avoids allocation and GC pressure
* Lock‑free design

### Runtime Definition

```go
type runq struct {
    head uint32
    tail uint32
    q    [256]*g
}
```

📌 This is a textbook circular queue implementation.

---

## 9. Global Run Queue — Why It Exists

The Global Run Queue:

* Uses a linked list
* Is protected by a lock
* Is intentionally slower

### Used When:

* Local Run Queue is full
* Goroutines are created after syscall returns
* Scheduler enforces fairness across Ps

📌 **Design Principle:** Global queue is a fallback, not the fast path.

---

## 10. Real Scheduling Flow (Step‑by‑Step)

```
Does M have a P?
 ├─ No → Park M
 └─ Yes
     ↓
Does P have runnable G in Local Run Queue?
 ├─ Yes → Pop G → Execute
 └─ No
     ↓
Is Global Run Queue non‑empty?
 ├─ Yes → Take batch → Execute
 └─ No
     ↓
Steal ~half of another P's run queue
```

---

## 11. One‑Line Mental Model (Must Remember)

> **P owns a ring‑buffer run queue**
> **M borrows P to execute G**
> **Global queue exists only for fairness and load balancing**

---

## 12. Ultra‑Important Interview Takeaways

* P ≠ OS thread
* P controls **parallelism**
* Goroutines are scheduled onto Ms via Ps
* Run queues belong to Ps, not Ms
* Local Run Queue = ring buffer
* Global Run Queue is slow by design

---

## 13. Common Interview Questions This Covers

* How does Go schedule goroutines?
* Why does Go need P if it already has threads?
* What happens when a goroutine blocks?
* Difference between local and global run queues
* Why is work stealing needed?
* How does GOMAXPROCS affect performance?


---
