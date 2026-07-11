# Why Goroutine and Go Channels Matter in API Calling

## Scenario

ধরুন আপনার একটি API endpoint আছে:

```http
GET /dashboard
```

এই API call করার সময় আপনাকে ৪টি আলাদা service call করতে হবে:

1. User Service
2. Order Service
3. Notification Service
4. Analytics Service

প্রতিটি service response দিতে প্রায় **8 seconds** সময় নেয়।

---

# Without Goroutine (Sequential Execution)

Code:

```go
user := getUser()
orders := getOrders()
notifications := getNotifications()
analytics := getAnalytics()
```

Execution:

```text
getUser()           -> 8 sec
getOrders()         -> 8 sec
getNotifications()  -> 8 sec
getAnalytics()      -> 8 sec
--------------------------------
Total               -> 32 sec
```

প্রায়:

```text
32.34 seconds
```

কারণ:

* প্রথম function শেষ না হলে দ্বিতীয় function শুরু হবে না।
* দ্বিতীয় function শেষ না হলে তৃতীয় function শুরু হবে না।
* সবগুলো কাজ একটার পর একটা (Sequentially) চলবে।

---

# Timeline

```text
0s --------8s--------16s--------24s--------32s

User Service
████████

Order Service
        ████████

Notification Service
                ████████

Analytics Service
                        ████████
```

---

# Problem

User API call করছে।

কিন্তু response পেতে:

```text
32.34 seconds
```

অপেক্ষা করতে হচ্ছে।

এত বেশি latency production system-এর জন্য খারাপ।

---

# Solution: Goroutine

Go-এর সবচেয়ে powerful feature হলো:

```text
Goroutine
```

এগুলো lightweight threads।

একই সময়ে অনেকগুলো কাজ চালাতে পারে।

---

# Using Goroutines

```go
go getUser()
go getOrders()
go getNotifications()
go getAnalytics()
```

এখন চারটি function একসাথে execute হবে।

---

# Timeline

```text
0s -------------------------------- 8s

User Service
████████

Order Service
████████

Notification Service
████████

Analytics Service
████████
```

সবগুলো service একই সময়ে run করছে।

Total:

```text
≈ 8 seconds
```

---

# Performance Improvement

Without Goroutine:

```text
32.34 sec
```

With Goroutine:

```text
8 sec
```

Improvement:

```text
24+ seconds saved
```

অর্থাৎ:

```text
4x faster response time
```

---

# Why This Happens?

কারণ API call গুলো বেশিরভাগ সময়:

* Database I/O
* Network I/O
* HTTP requests
* Microservice communication

এগুলোর জন্য wait করে।

CPU কাজ খুব কম করে।

এই waiting time-এর মধ্যে অন্য কাজগুলো execute করা যায়।

Goroutine ঠিক এই কাজটাই করে।

---

# Real Example

Without Goroutine:

```go
func GetDashboard() {
	user := getUser()
	orders := getOrders()
	notifications := getNotifications()
	analytics := getAnalytics()

	_ = user
	_ = orders
	_ = notifications
	_ = analytics
}
```

---

# Using Goroutine

```go
func GetDashboard() {
	go getUser()
	go getOrders()
	go getNotifications()
	go getAnalytics()
}
```

কিন্তু এখানে একটা সমস্যা আছে।

Main function return করে ফেলতে পারে।

Goroutine শেষ হওয়ার আগেই handler response পাঠিয়ে দিতে পারে।

তাই আমাদের synchronization দরকার।

---

# This is why Channels Matter

Channel হলো goroutine-এর মধ্যে communication mechanism।

এক goroutine অন্য goroutine-কে data পাঠাতে পারে।

---

# Example

```go
userCh := make(chan User)
```

এক goroutine:

```go
go func() {
	userCh <- getUser()
}()
```

অন্য goroutine:

```go
user := <-userCh
```

---

# Channel Diagram

```text
Goroutine A
      |
      | send data
      ↓
    Channel
      |
      | receive data
      ↓
Goroutine B
```

---

# Real API Example

```go
func GetDashboard() Dashboard {

	userCh := make(chan User)
	orderCh := make(chan []Order)
	notiCh := make(chan []Notification)
	analyticsCh := make(chan Analytics)

	go func() {
		userCh <- getUser()
	}()

	go func() {
		orderCh <- getOrders()
	}()

	go func() {
		notiCh <- getNotifications()
	}()

	go func() {
		analyticsCh <- getAnalytics()
	}()

	user := <-userCh
	orders := <-orderCh
	notifications := <-notiCh
	analytics := <-analyticsCh

	return Dashboard{
		User:          user,
		Orders:        orders,
		Notifications: notifications,
		Analytics:     analytics,
	}
}
```

---

# Execution Flow

```text
Request
   |
   |
   +------ Goroutine 1 ------ User Service
   |
   +------ Goroutine 2 ------ Order Service
   |
   +------ Goroutine 3 ------ Notification Service
   |
   +------ Goroutine 4 ------ Analytics Service
   |
   |
Wait on Channels
   |
Collect Results
   |
Return JSON Response
```

---

# Timeline

```text
0s -------------------------------- 8s

Request Started
     |
     +---- User Service
     |
     +---- Order Service
     |
     +---- Notification Service
     |
     +---- Analytics Service
     |
Receive from Channels
     |
Return Response
```

Total:

```text
≈ 8 seconds
```

---

# Why Channels Matter in API Calling

Channels provide:

### 1. Synchronization

Handler জানে কখন সব goroutine শেষ হয়েছে।

---

### 2. Data Communication

এক goroutine-এর result অন্য জায়গায় safely পাঠানো যায়।

---

### 3. Prevent Race Conditions

Shared variable ব্যবহার না করে data pass করা যায়।

---

### 4. Clean Concurrency

Code readable এবং maintainable হয়।

---

# Production Use Cases

### Calling Multiple Microservices

```text
User Service
Order Service
Payment Service
Notification Service
```

---

### Multiple Database Queries

```text
PostgreSQL
Redis
ElasticSearch
```

---

### External APIs

```text
Google API
Stripe API
AWS API
```

---

### Dashboard APIs

```text
Statistics
Recent Orders
Notifications
User Profile
```

সবগুলো একসাথে fetch করা যায়।

---

# Why Goroutine Matters

Because:

```text
Sequential:
8 + 8 + 8 + 8
= 32 seconds
```

```text
Concurrent:
max(8,8,8,8)
= 8 seconds
```

---

# Why Channels Matter

Because:

```text
Goroutines run concurrently
Channels collect results safely
Handler waits for completion
Then returns response
```

---

# Final Takeaway

Without Goroutines:

```text
API Response Time ≈ 32.34 sec
```

With Goroutines + Channels:

```text
API Response Time ≈ 8 sec
```

Savings:

```text
≈ 24 seconds
```

This is why Goroutines and Channels are extremely important in modern Go APIs, especially in:

* Microservices
* Aggregation APIs
* Dashboard APIs
* External API calls
* High-performance backend systems
