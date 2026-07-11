# Hey Load Testing Tool

## What is `hey`?

`hey` হলো Go ভাষায় তৈরি একটি **HTTP load testing tool**। এটি ব্যবহার করে আপনি একটি API বা ওয়েবসাইটে অনেকগুলো request পাঠিয়ে দেখতে পারেন:

* Server কতগুলো request per second (RPS) handle করতে পারে।
* Response latency (response time) কত।
* Server কত concurrent user handle করতে পারে।
* কোনো performance bottleneck আছে কিনা।

Repository: https://github.com/rakyll/hey

---

# Installation

```bash
go install github.com/rakyll/hey@latest
```

অথবা

```bash
git clone https://github.com/rakyll/hey.git
cd hey
go build
```

---

# Basic Syntax

```bash
hey [options] <url>
```

উদাহরণ:

```bash
hey https://example.com
```

ডিফল্টভাবে:

* `-n 200` → মোট 200 টি request পাঠাবে।
* `-c 50` → একসাথে 50 টি concurrent worker ব্যবহার করবে।

---

# Available Options

```text
-n  Number of requests to run. Default is 200.
-c  Number of workers to run concurrently. Default is 50.
-q  Rate limit, in queries per second (QPS) per worker.
-z  Duration of application to send requests.
-o  Output type.

-m  HTTP method.
-H  Custom HTTP header.
-t  Timeout for each request.
-A  HTTP Accept header.
-d  HTTP request body.
-D  HTTP request body from file.
-T  Content-type.
-a  Basic authentication.
-x  HTTP Proxy address.
-h2 Enable HTTP/2.

-host HTTP Host header.

-disable-compression
-disable-keepalive
-disable-redirects
-cpus
```

---

# Most Used Commands (বাংলা ব্যাখ্যা)

# 1. `-n` (Number of Requests)

কতগুলো request পাঠাতে চান।

```bash
hey -n 1000 http://localhost:8000
```

অর্থ:

* মোট 1000 request পাঠাবে।

---

# 2. `-c` (Concurrency)

একসাথে কতগুলো request চলবে।

```bash
hey -n 1000 -c 100 http://localhost:8000
```

অর্থ:

* মোট request = 1000
* একসাথে 100 request চলবে।

এটি অনেকটা:

> "একই সময়ে 100 জন user website ব্যবহার করছে।"

আরেকটি উদাহরণ:

```bash
hey -n 10000 -c 500 http://localhost:8000
```

মানে:

* মোট 10000 request
* একসাথে 500 request।

---

# `-n` এবং `-c` এর পার্থক্য

```bash
hey -n 1000 -c 10
```

Execution:

```text
Batch 1 → 10 requests
Batch 2 → 10 requests
Batch 3 → 10 requests
...
```

মোট:

```text
1000 requests
```

---

# 3. `-q` (Rate Limit)

প্রতি worker প্রতি সেকেন্ডে কত request পাঠাবে।

```bash
hey -n 1000 -c 10 -q 5 http://localhost:8000
```

মানে:

* 10 worker
* প্রতিটি worker 5 request/sec পাঠাবে।

মোট throughput:

```text
10 × 5 = 50 requests/sec
```

---

# 4. `-z` (Duration)

কতক্ষণ test চলবে।

```bash
hey -z 30s http://localhost:8000
```

মানে:

* 30 সেকেন্ড request পাঠাবে।
* `-n` ignore করবে।

আরও উদাহরণ:

```bash
hey -z 5m http://localhost:8000
```

মানে:

* 5 মিনিট load test চলবে।

---

# 5. `-m` (HTTP Method)

Request method নির্ধারণ করে।

GET:

```bash
hey -m GET http://localhost:8000/users
```

POST:

```bash
hey -m POST http://localhost:8000/users
```

DELETE:

```bash
hey -m DELETE http://localhost:8000/users/1
```

Supported Methods:

* GET
* POST
* PUT
* DELETE
* HEAD
* OPTIONS

---

# 6. `-d` (Request Body)

POST বা PUT request এ body পাঠাতে।

```bash
hey -m POST \
-d '{"name":"Farhan"}' \
http://localhost:8000/users
```

---

# 7. `-D` (Body from File)

File থেকে request body পাঠাতে।

`data.json`

```json
{
  "name": "Farhan",
  "age": 25
}
```

Run:

```bash
hey -m POST -D data.json http://localhost:8000/users
```

---

# 8. `-T` (Content-Type)

```bash
hey -m POST \
-T application/json \
-d '{"name":"Farhan"}' \
http://localhost:8000/users
```

Header:

```http
Content-Type: application/json
```

ডিফল্ট Content-Type:

```text
text/html
```

---

# 9. `-H` (Custom Header)

Header যোগ করতে।

```bash
hey -H "Authorization: Bearer TOKEN" \
http://localhost:8000
```

একাধিক Header:

```bash
hey \
-H "Authorization: Bearer TOKEN" \
-H "X-API-Key: 12345" \
http://localhost:8000
```

---

# 10. `-A` (Accept Header)

```bash
hey -A application/json http://localhost:8000
```

Header:

```http
Accept: application/json
```

---

# 11. `-a` (Basic Authentication)

```bash
hey -a admin:password http://localhost:8000
```

Equivalent Header:

```http
Authorization: Basic xxxxx
```

---

# 12. `-t` (Timeout)

```bash
hey -t 60 http://localhost:8000
```

মানে:

* প্রতিটি request সর্বোচ্চ 60 সেকেন্ড অপেক্ষা করবে।

```bash
hey -t 0 http://localhost:8000
```

মানে:

* Infinite timeout।

---

# 13. `-h2` (HTTP/2)

```bash
hey -h2 https://example.com
```

HTTP/2 ব্যবহার করে request পাঠাবে।

---

# 14. `-disable-keepalive`

Default:

```text
Request1 ----\
Request2 ----- Same TCP Connection
Request3 ----/
```

Command:

```bash
hey -disable-keepalive http://localhost:8000
```

এখন:

```text
Request1 -> New TCP
Request2 -> New TCP
Request3 -> New TCP
```

এটি server-এর TCP connection handling test করার জন্য খুব useful।

---

# 15. `-disable-compression`

```bash
hey -disable-compression http://localhost:8000
```

Server response gzip না করে raw data পাঠাবে।

---

# 16. `-disable-redirects`

```bash
hey -disable-redirects http://localhost:8000
```

302 বা অন্য redirect automatically follow করবে না।

---

# 17. `-cpus`

```bash
hey -cpus 4 -n 10000 -c 500 http://localhost:8000
```

`hey` নিজে কত CPU core ব্যবহার করবে তা নির্ধারণ করে।

Default:

```text
Current machine's CPU cores.
```

---

# 18. `-o` (Output Type)

Default:

```bash
hey http://localhost:8000
```

Summary output দেখাবে।

CSV output:

```bash
hey -o csv http://localhost:8000
```

সব metrics CSV format-এ export হবে।

---

# 19. `-x` (HTTP Proxy)

```bash
hey -x localhost:8080 http://localhost:8000
```

Request proxy server এর মাধ্যমে পাঠানো হবে।

---

# 20. `-host`

Custom Host header সেট করার জন্য।

```bash
hey -host api.example.com http://localhost:8000
```

---

# Real World Examples

## Test GET API

```bash
hey -n 10000 -c 100 http://localhost:8000/api/users
```

---

## Test POST API

```bash
hey \
-n 10000 \
-c 100 \
-m POST \
-T application/json \
-d '{"name":"Farhan"}' \
http://localhost:8000/api/users
```

---

## 30 Seconds Stress Test

```bash
hey -z 30s -c 200 http://localhost:8000
```

---

## Authenticated API Test

```bash
hey \
-n 10000 \
-c 100 \
-H "Authorization: Bearer TOKEN" \
http://localhost:8000/api/profile
```

---

# Most Common Backend Interview Command

```bash
hey -n 10000 -c 100 http://localhost:8000
```

এটি ব্যবহার করে দেখা হয়:

* Requests Per Second (RPS)
* Average Latency
* Fastest Response
* Slowest Response
* Error Rate
* Throughput

এগুলোর মাধ্যমে বোঝা যায় আপনার Go, Django, Node.js বা অন্য কোনো backend server কতটা load handle করতে পারে।

---

# Summary

সবচেয়ে বেশি ব্যবহৃত options:

| Option                 | কাজ                      |
| ---------------------- | ------------------------ |
| `-n`                   | মোট request সংখ্যা       |
| `-c`                   | Concurrent worker সংখ্যা |
| `-q`                   | Rate limit               |
| `-z`                   | Duration ভিত্তিক test    |
| `-m`                   | HTTP Method              |
| `-d`                   | Request body             |
| `-D`                   | File থেকে body           |
| `-T`                   | Content-Type             |
| `-H`                   | Custom Header            |
| `-A`                   | Accept Header            |
| `-a`                   | Basic Authentication     |
| `-t`                   | Timeout                  |
| `-h2`                  | HTTP/2                   |
| `-disable-keepalive`   | Keep-alive বন্ধ          |
| `-disable-compression` | Compression বন্ধ         |
| `-disable-redirects`   | Redirect follow বন্ধ     |
| `-cpus`                | CPU cores                |
| `-o csv`               | CSV output               |
| `-x`                   | Proxy                    |
| `-host`                | Custom Host header       |
