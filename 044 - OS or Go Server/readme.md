# File Descriptor, Process & Socket with Go Server

## 1️⃣ File Descriptor (FD) কী?

👉 **File Descriptor হলো একটি ছোট integer number** যেটা kernel ব্যবহার করে কোনো **open resource** কে identify করার জন্য।

📌 Resource হতে পারে:

* File
* Socket
* Pipe
* Terminal
* Network connection

> **FD = kernel-এর দেওয়া handle**

---

## File Descriptor কোথায় থাকে?

❌ FD ফাইলের ভেতরে থাকে না
❌ Program-এর variable না

✅ FD থাকে **kernel-এর ভিতরে**, আর process শুধু একটা **number** জানে।

```
Process
  └─ fd = 3   ─────▶ Kernel table entry
```

---

## 2️⃣ Process এর সাথে File Descriptor এর সম্পর্ক

### প্রতিটা process এর কি আলাদা FD থাকে?

👉 **হ্যাঁ — প্রতিটা process এর আলাদা File Descriptor Table থাকে** ✅

```
Process A                  Process B
FD Table                   FD Table
0 → stdin                  0 → stdin
1 → stdout                 1 → stdout
2 → stderr                 2 → stderr
3 → file.txt               3 → socket
```

📌 FD number একই হলেও resource আলাদা হতে পারে।

---

## Kernel Side FD Table Structure

```
Process
 └─ FD Table
      ├─ 0 → stdin
      ├─ 1 → stdout
      ├─ 2 → stderr
      ├─ 3 → File Object
```

FD entry point করে:

```
FD → File Object → inode / Socket
```

---

## Default File Descriptors

| FD | Meaning |
| -- | ------- |
| 0  | stdin   |
| 1  | stdout  |
| 2  | stderr  |

সব process এর ক্ষেত্রেই এই তিনটা থাকে।

---

## 3️⃣ File Descriptor vs File Object (Important)

* **FD** → per-process
* **File Object** → kernel-level, shareable

```
Process A FD 3 ─┐
                ├─ File Object ──▶ inode
Process B FD 5 ─┘
```

📌 `fork()` করলে FD share হয়
📌 `dup()` করলে FD duplicate হয়

---

## 4️⃣ Socket কী?

👉 **Socket হলো network communication-এর জন্য special file-like object**।

Kernel-এর দৃষ্টিতে:

> **Socket = file**

তাই socket-ও FD দিয়ে identify হয়।

---

## Socket কী represent করে?

একটা socket এর মধ্যে থাকে:

* Local IP : Port
* Remote IP : Port
* Protocol (TCP / UDP)
* Connection state

```
Socket
 ├─ Local IP:Port
 ├─ Remote IP:Port
 └─ State (LISTEN, ESTABLISHED)
```

---

## 5️⃣ Socket তৈরি হওয়ার High-Level Flow

```
socket()
  ↓
bind()
  ↓
listen()
  ↓
accept()
```

Kernel যা করে:

* Socket create করে
* File Descriptor দেয়

```
fd = 3  → listening socket
fd = 4  → client connection socket
```

---

## 6️⃣ Go Server এর সাথে Socket এর সম্পর্ক

👉 **Go server মানে আসলে socket-based server**।

Go code:

```go
ln, _ := net.Listen("tcp", ":8080")
```

ভেতরে যা ঘটে (simplified):

```
Go runtime
  ↓
syscall socket()
  ↓
Kernel creates socket
  ↓
Kernel returns FD
```

---

## 7️⃣ Go Server Runtime Flow (Visualization)

```
Client
  ↓ TCP SYN
Kernel
  ↓ creates connection socket
  ↓ assigns FD (e.g. 7)
Go runtime
  ↓ wraps FD into net.Conn
  ↓ passes to goroutine
```

📌 প্রতিটা client connection:

* আলাদা socket
* আলাদা FD
* আলাদা goroutine (সাধারণত)

---

## 8️⃣ Go Server: Goroutine + Socket + FD

Typical Go server code:

```go
for {
    conn, _ := ln.Accept()
    go handle(conn)
}
```

Visualization:

```
OS Process
 ├─ FD 3 → Listening socket
 ├─ FD 4 → Client A socket
 ├─ FD 5 → Client B socket
 │
 ├─ Goroutine A → FD 4
 └─ Goroutine B → FD 5
```

---

## 9️⃣ Blocking I/O হলে কী হয়?

Go runtime ব্যবহার করে:

* Network poller
* Non-blocking socket

যাতে:

* এক goroutine block হলেও
* OS thread block না হয়

📌 Backend এ epoll / kqueue / IOCP ব্যবহার হয়।

---

## 🔟 Big Picture (সব একসাথে)

```
Client
   ↓
Socket
   ↓
File Descriptor (FD)
   ↓
Process FD Table
   ↓
Go Runtime
   ↓
Goroutine
```

---

## 🎯 One-Line Master Explanation (Interview Level)

> **File Descriptor হলো kernel-managed integer handle যা process-এর FD table এ থাকে; প্রতিটা process-এর আলাদা FD table থাকে; socket হলো network-এর জন্য special file object যা FD দিয়ে access হয়; আর Go server এই socket FD গুলোকে goroutine দিয়ে handle করে।**

---


উদাহরণ কোড:
```
package main

import (
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I'm Nadim. I'm 24 years old....")
}

func main () {
	mux := http.NewServeMux()	// Router

	mux.HandleFunc("/hello", HelloHandler)	// Route 
	mux.HandleFunc("/about", AboutHandler)	// Route 

	fmt.Println("Server Running on port: 3000")

	err := http.ListenAndServe(":3000", mux)

	if err != nil {
		fmt.Println("***Error Occurred", err)
	}
}
```
</br> </br>

# Go HTTP Server Request Flow — Kernel থেকে Handler পর্যন্ত (Bangla)

```go
http.ListenAndServe(":3000", mux)
```

---

## 🧠 STEP 0: Server start হওয়ার সময় কী ঘটে

```go
http.ListenAndServe(":3000", mux)
```

এই লাইনে Go runtime kernel-এর সাথে কথা বলে নিচের system call গুলো করে:

```
socket()
bind()
listen()
```

Kernel যা করে:

* একটি **listening socket** তৈরি করে
* ধরে নেওয়া যাক:

```
Listening Socket FD = 3
```

📌 এই FD 3 kernel-এর কাছে registered থাকে এবং server এখন client request-এর জন্য অপেক্ষা করে।

### Diagram

```
Go Process
 └─ FD 3 → Listening Socket (Kernel)
```

---

## 🌍 STEP 1: Client `/about` এ request পাঠায়

Browser-এ user লিখলো:

```
http://localhost:3000/about
```

Browser যা তৈরি করলো:

```
GET /about HTTP/1.1
Host: localhost:3000
```

এই request:

```
Client Browser
   ↓
TCP Packet
   ↓
Network
```

---

## 💳 STEP 2: Server-এর NIC request গ্রহণ করে

Server machine-এর **NIC (Network Interface Card)**:

```
NIC
 └─ RX Buffer (receive buffer)
```

NIC যা করে:

* TCP packet RX buffer-এ রাখে
* CPU-কে **hardware interrupt** দেয়

📌 Data আগেই binary থাকে, এখানে কোনো conversion হয় না।

### Diagram

```
Network
   ↓
NIC RX Buffer
```

---

## 🧠 STEP 3: Kernel data socket receive buffer-এ রাখে

Kernel interrupt পেয়ে দেখে:

* Port = 3000
* Listening socket FD = 3

তারপর flow হয়:

```
NIC RX Buffer
   ↓ (DMA)
Kernel Socket Receive Buffer (FD 3)
```

Kernel এখন বলে:

> "FD 3 readable হয়ে গেছে"

### Diagram

```
Kernel
 └─ FD 3
     └─ Receive Buffer (data ready)
```

---

## 🧵 STEP 4: Go runtime epoll দিয়ে এটা detect করে

Go runtime আগে থেকেই:

```
epoll_wait()
```

Kernel জানায়:

```
FD 3 readable
```

Go runtime তখন:

```
accept()
```

Kernel:

* নতুন client connection socket তৈরি করে
* নতুন FD দেয়

```
Client Connection FD = 7
```

### Diagram

```
Kernel
 ├─ FD 3 → Listening Socket
 └─ FD 7 → Client Socket
```

---

## 🧵 STEP 5: নতুন goroutine তৈরি হয়

Go runtime:

```
go c.serve(conn)
```

এখানে:

* FD 7 কে `net.Conn` দিয়ে wrap করা হয়
* একটি নতুন goroutine তৈরি হয়

```
OS Thread
 └─ Goroutine
      └─ net.Conn (FD 7)
```

📌 পুরনো goroutine আবার next request-এর জন্য অপেক্ষা করতে থাকে।

---

## 🚦 STEP 6: Goroutine HTTP parse করে route match করে

নতুন goroutine:

```
FD 7 থেকে read()
```

Kernel socket buffer থেকে data user space-এ দেয়।

Go HTTP server:

* HTTP request parse করে
* URL পায়: `/about`

ServeMux routing:

```
/hello → HelloHandler
/about → AboutHandler ✅
```

### Diagram

```
Goroutine
   ↓
http.Server
   ↓
ServeMux
   ↓
AboutHandler()
```

---

## 📝 STEP 7: Handler response লেখে

```go
fmt.Fprintln(w, "I'm Nadim. I'm 24 years old...")
```

Flow:

```
User Space
   ↓
Kernel Socket SEND Buffer (FD 7)
   ↓
NIC TX Buffer
```

Kernel data network-এ পাঠায়।

---

## 🌍 STEP 8: Client response পায়

Client side flow:

```
NIC
 ↓
Kernel TCP Stack
 ↓
Browser
 ↓
HTML Render
```

User browser-এ output দেখে।

---

## 🧠 Complete Big Picture Diagram

```
Client Browser
     ↓
Network
     ↓
Server NIC (RX)
     ↓
Kernel Socket Buffer
     ↓
FD 3 readable
     ↓
epoll_wait (Go runtime)
     ↓
accept()
     ↓
FD 7 (client socket)
     ↓
goroutine
     ↓
ServeMux
     ↓
Handler
     ↓
write()
     ↓
Kernel Send Buffer
     ↓
NIC TX
     ↓
Client Browser
```

---

## 🎯 One-Line Final Explanation (Interview Level)

> **Client request NIC দিয়ে kernel socket buffer-এ আসে; kernel listening socket FD readable করে; Go runtime epoll দিয়ে detect করে accept করে নতুন FD পায়; সেই FD দিয়ে goroutine HTTP request parse করে ServeMux দিয়ে handler চালায়; response আবার socket send buffer হয়ে NIC দিয়ে client-এ যায়।**

---
