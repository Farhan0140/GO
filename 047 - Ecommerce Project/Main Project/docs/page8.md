# Domain-Driven Design (DDD) in Go

---

## 1️⃣ Domain-Driven Design (DDD) কী?

**Domain-Driven Design (DDD)** হলো একটি **software design approach** যেখানে software architecture এবং code structure তৈরি করা হয় **business domain** (business problem) বুঝে।

সহজভাবে:

> Domain-Driven Design মানে হলো — software design করার সময় business logic-কেই কেন্দ্র করে architecture তৈরি করা।

এই ধারণাটি জনপ্রিয় করেন Eric Evans তার বই
“Domain-Driven Design: Tackling Complexity in the Heart of Software” (2003)।

---

# 2️⃣ Domain (DDD-তে Domain মানে কী?)

**Domain = যে business problem solve করার জন্য software তৈরি করা হচ্ছে।**

উদাহরণ:

| Application  | Domain                  |
| ------------ | ----------------------- |
| E-commerce   | Product, Order, Payment |
| Banking      | Account, Transaction    |
| Ride sharing | Driver, Ride, Payment   |

DDD-তে software design করা হয় **এই domain model অনুযায়ী।**

---

# 3️⃣ কেন Domain-Driven Design ব্যবহার করা দরকার?

Complex business application-এ code খুব দ্রুত messy হয়ে যায়।

### ❌ Without DDD

সমস্যা:

* Business logic everywhere
* Controller-এ logic
* Database-এ logic
* Code maintain করা কঠিন
* Team communication সমস্যা

---

### ✅ With DDD

DDD ব্যবহার করলে:

1️⃣ Business logic clear হয়
2️⃣ Code structure domain অনুযায়ী হয়
3️⃣ Maintainability বাড়ে
4️⃣ Large team collaboration সহজ হয়
5️⃣ Complex business logic handle করা সহজ হয়

---

# 4️⃣ DDD-এর Core Concepts

## 1️⃣ Entity

Entity হলো এমন object যার **unique identity** আছে।

Example:

```go
type User struct {
	ID    int
	Name  string
	Email string
}
```

এখানে `ID` entity identity।

---

## 2️⃣ Value Object

Value object-এর **identity নেই**, শুধু value গুরুত্বপূর্ণ।

Example:

```go
type Money struct {
	Amount   float64
	Currency string
}
```

---

## 3️⃣ Repository

Repository entity store এবং retrieve করে।

Example:

```go
type UserRepository interface {
	FindByID(id int) (*User, error)
	Save(user *User) error
}
```

---

## 4️⃣ Service (Domain Service)

যে logic কোনো single entity-তে belong করে না।

Example:

```
Order + Payment + Inventory
```

এই logic service-এ থাকে।

---

## 5️⃣ Aggregate

একাধিক entity/value object মিলিয়ে **একটি consistency boundary** তৈরি হয়।

Example:

```
Order
 ├── OrderItem
 └── Payment
```

---

# 5️⃣ DDD Architecture Concept

```text
Application Layer
        ↓
Domain Layer
        ↓
Infrastructure Layer
```

### 1️⃣ Application Layer

* Use cases
* Service orchestration

---

### 2️⃣ Domain Layer

সব business logic এখানে থাকে

যেমন:

* Entities
* Value objects
* Repository interface
* Domain services

---

### 3️⃣ Infrastructure Layer

External systems handle করে

যেমন:

* Database
* HTTP
* Cache
* Messaging

---

# 6️⃣ Golang-এ DDD কিভাবে ব্যবহার করা যায়

ধরি তুমি **mini e-commerce** project বানাচ্ছ।

## Folder Structure

```text
project
│
├── domain
│   ├── product.go
│   └── product_repository.go
│
├── application
│   └── product_service.go
│
├── infrastructure
│   └── product_repository_impl.go
│
├── delivery
│   └── http_handler.go
│
└── main.go
```

---

# 7️⃣ Example Implementation (Simplified)

## Domain Entity

```go
package domain

type Product struct {
	ID    int
	Name  string
	Price float64
}
```

---

## Repository Interface (Domain Layer)

```go
package domain

type ProductRepository interface {
	Save(product *Product) error
	FindByID(id int) (*Product, error)
}
```

---

## Application Service

```go
package application

import "yourapp/domain"

type ProductService struct {
	repo domain.ProductRepository
}

func NewProductService(r domain.ProductRepository) *ProductService {
	return &ProductService{repo: r}
}

func (s *ProductService) CreateProduct(name string, price float64) error {
	product := &domain.Product{
		Name:  name,
		Price: price,
	}

	return s.repo.Save(product)
}
```

---

## Infrastructure Implementation

```go
package infrastructure

import (
	"database/sql"
	"yourapp/domain"
)

type productRepo struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) domain.ProductRepository {
	return &productRepo{db}
}

func (r *productRepo) Save(product *domain.Product) error {
	_, err := r.db.Exec(
		"INSERT INTO products(name, price) VALUES (?, ?)",
		product.Name,
		product.Price,
	)
	return err
}
```

---

# 8️⃣ DDD vs Traditional Layered Architecture

| Feature         | Traditional | DDD             |
| --------------- | ----------- | --------------- |
| Focus           | Database    | Business domain |
| Structure       | Layer based | Domain based    |
| Logic           | Scattered   | Centralized     |
| Maintainability | Medium      | High            |

---

# 9️⃣ Advantages & Disadvantages

## ✅ Advantages

* Complex system handle করা সহজ
* Business logic clean থাকে
* Maintainability high
* Scalability ভালো
* Large team friendly

---

## ❌ Disadvantages

* Learning curve বেশি
* Small project-এ overengineering
* Initial design time বেশি লাগে


# 1️⃣1️⃣ সংক্ষেপে
```
Domain-Driven Design =

✔ Business-centric design
✔ Domain model based architecture
✔ Entities + Value objects + Repository
✔ Large and complex system-এর জন্য best
```
---

### Official Reference

Eric Evans book
[https://domainlanguage.com/ddd/](https://domainlanguage.com/ddd/)

Martin Fowler explanation
[https://martinfowler.com/bliki/DomainDrivenDesign.html](https://martinfowler.com/bliki/DomainDrivenDesign.html)
