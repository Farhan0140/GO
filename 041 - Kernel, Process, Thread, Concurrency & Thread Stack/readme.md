# Kernel, Process, Thread, Concurrency & Thread Stack

## 1️⃣ Kernel কী?

👉 **Kernel হলো Operating System-এর হৃদয় (core)** ❤️

সহজ ভাষায়:

> **Kernel হলো সেই প্রোগ্রাম যেটা সবসময় CPU-তে চলে এবং পুরো কম্পিউটারকে কন্ট্রোল করে।**

---

## Kernel কী কী কাজ করে?

Kernel-এর প্রধান দায়িত্বগুলো:

* 🧠 CPU control
* 🧵 Process ও Thread manage করা
* 🧠 Memory management
* 💽 Disk / I/O control
* 🔐 Security & isolation

📌 User program **কখনোই সরাসরি hardware access করতে পারে না**—সবকিছু kernel-এর মাধ্যমে হয়।

---

## Kernel কোথায় থাকে?

* Kernel থাকে **RAM-এ**
* CPU-র **privileged mode (kernel mode)** এ execute হয়
* User program চলে **user mode** এ

```
CPU
├─ Kernel Mode   ← Kernel
└─ User Mode    ← Application / Program
```

---

## 2️⃣ Kernel কেন Process ও Thread track করে?

কারণ:

> **CPU একটাই (logical), কিন্তু কাজ অনেক**

একই সময়ে:

* Browser চলছে
* Music player চলছে
* Code editor চলছে
* Background services চলছে

👉 কে কখন CPU পাবে, কে অপেক্ষা করবে—এই সিদ্ধান্ত নিতে হয়
➡️ **এই কাজটা kernel করে**

---

## Kernel কী কী তথ্য track করে?

### 🔹 Process Control Block (PCB)

* Process ID (PID)
* State (running, waiting, ready)
* Memory information
* Open files
* Parent process

### 🔹 Thread Control Block (TCB)

* Thread ID
* Stack pointer
* Program counter
* Register state

📌 এই তথ্য ছাড়া kernel **context switch করতে পারে না**।

---

## 3️⃣ Concurrency কী এবং kernel কেন execute করে?

### Concurrency মানে কী?

👉 **Concurrency মানে:**

> অনেক কাজ *একই সময়ে হচ্ছে বলে মনে হওয়া*

এমনকি single CPU হলেও:

```
Time
│ Task A ──┐
│          ├─ interleaving
│ Task B ──┘
```

---

### Kernel কেন concurrency চালায়?

কারণ:

* I/O খুব slow (disk, network)
* CPU অনেক fast
* CPU idle থাকলে resource waste হয়

উদাহরণ:

```
Process A → waiting for disk
Process B → uses CPU
```

👉 Kernel smart ভাবে CPU switch করে

📌 একে বলে **Context Switching**

---

## 4️⃣ Parallelism কী? Kernel-এর ভূমিকা

👉 **Parallelism মানে:**

> একাধিক CPU/Core এ সত্যিকারের একসাথে কাজ

```
Core 1 → Thread A
Core 2 → Thread B
```

Kernel যা করে:

* Thread গুলোকে different core-এ assign করে
* Load balancing করে
* Core affinity manage করে

---

## 5️⃣ OS নিজে কেন Concurrency বা Parallelism execute করতে পারে না?

কারণ:

> **OS = Kernel + tools + UI + utilities**

কিন্তু:

* OS নিজে hardware-এ execute হয় না
* CPU সরাসরি বুঝে শুধু **kernel instructions**

📌 তাই:

* Scheduling
* Context switching
* Core assignment

👉 এই সবকিছু **শুধু kernel করে**

---

## 6️⃣ Thread কী?

👉 **Thread হলো process-এর ভেতরের সবচেয়ে ছোট execution unit**

```
Process
├─ Thread 1
├─ Thread 2
└─ Thread 3
```

Thread গুলো share করে:

* Code
* Heap
* Open files

কিন্তু ❗ **Stack আলাদা হয়**

---

## 7️⃣ Stack কী এবং কী কাজে লাগে?

Stack ব্যবহার হয়:

* Function call
* Local variables
* Function arguments
* Return address

---

## 8️⃣ কেন প্রতিটা Thread-এর আলাদা Stack দরকার?

এই প্রশ্নটা খুব গুরুত্বপূর্ণ 🔥

### যদি stack shared হতো?

ধরি:

```
Thread A → func1()
Thread B → func2()
```

Shared stack হলে:

* Local variable overwrite হতো
* Return address নষ্ট হতো
* Program crash করতো 💥

---

### Separate Stack মানে কী?

```
Thread A Stack
┌────────────┐
│ func1 vars │
└────────────┘

Thread B Stack
┌────────────┐
│ func2 vars │
└────────────┘
```

✔ Data corruption হয় না
✔ Thread independently run করতে পারে
✔ Context switching fast হয়

---

## 9️⃣ Kernel Context Switch করতে Stack কেন দরকার?

Context switch মানে:

```
Thread A → pause
Thread B → resume
```

Kernel যা করে:

* Thread A এর **stack pointer save করে**
* Thread B এর **stack pointer load করে**

📌 Stack না থাকলে thread resume করা অসম্ভব

---

## 🔟 Concurrency vs Parallelism (Bangla Table)

| বিষয়      | Concurrency      | Parallelism         |
| --------- | ---------------- | ------------------- |
| CPU       | 1 হলেও চলে       | Multiple core দরকার |
| Execution | Interleaving     | Truly simultaneous  |
| Control   | Kernel scheduler | Kernel scheduler    |
| Goal      | CPU utilization  | Speed               |

---

## 🎯 One-Line Master Explanation (Interview Level)

> **Kernel CPU control করে বলে process ও thread track করে; concurrency ও parallelism kernel চালায় কারণ শুধু kernel-ই hardware access পায়; আর প্রতিটা thread-এর আলাদা stack লাগে কারণ function execution state isolate না করলে program crash করবে।**
