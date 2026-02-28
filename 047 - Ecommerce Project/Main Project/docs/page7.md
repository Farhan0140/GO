# Repository Design Pattern in Go

---

## 1️⃣ Repository Design Pattern কী?

**Repository Design Pattern** হলো একটি architectural pattern যা **business logic layer** এবং **data access layer**–এর মাঝে একটি abstraction layer তৈরি করে।

সহজভাবে বললে:

> Repository হলো এমন একটি layer যা database access logic কে encapsulate করে এবং application-কে clean interface দেয়।

এই প্যাটার্নটি মূলত Domain-Driven Design (DDD) থেকে এসেছে, যা popular হয় “Domain-Driven Design: Tackling Complexity in the Heart of Software” (Eric Evans, 2003) বইয়ের মাধ্যমে।

---

## 2️⃣ কেন Repository ব্যবহার করা দরকার?

ধরো তুমি সরাসরি handler/controller থেকে SQL query লিখছো। তাহলে কী সমস্যা হবে?

### ❌ Without Repository

* Business logic + SQL mixed হয়ে যাবে
* Code duplication হবে
* Testing কঠিন হবে
* Database change করলে সব জায়গায় modify করতে হবে

---

### ✅ With Repository

Repository ব্যবহার করলে:

1. **Separation of Concerns** বজায় থাকে
2. Database change (MySQL → PostgreSQL → MongoDB) সহজ হয়
3. Unit testing সহজ হয় (mock repository করা যায়)
4. Clean Architecture follow করা যায়
5. Loose coupling পাওয়া যায়

---

## 3️⃣ Conceptual Architecture

```
Handler / Service
        ↓
    Repository (Interface)
        ↓
  Database Implementation
```

এখানে business layer জানে না database কীভাবে কাজ করছে।
সে শুধু repository interface জানে।

---

# 4️⃣ Golang এ Repository কিভাবে ব্যবহার করতে হয়

ধরি তুমি একটি mini e-commerce project করছো।
আমরা `Product` entity নিয়ে example করবো।

---

## Step 1️⃣: Model

```go
package model

type Product struct {
	ID    int
	Name  string
	Price float64
}
```

---

## Step 2️⃣: Repository Interface

```go
package repository

import "yourapp/model"

type ProductRepository interface {
	Create(product model.Product) (model.Product, error)
	FindByID(id int) (model.Product, error)
	FindAll() ([]model.Product, error)
	Delete(id int) error
}
```

এখানে আমরা abstraction define করছি।
এখনো database implementation লিখিনি।

---

## Step 3️⃣: Database Implementation (MySQL example)

```go
package repository

import (
	"database/sql"
	"yourapp/model"
)

type productRepo struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepo{db: db}
}

func (r *productRepo) Create(product model.Product) (model.Product, error) {
	query := "INSERT INTO products(name, price) VALUES (?, ?)"
	result, err := r.db.Exec(query, product.Name, product.Price)
	if err != nil {
		return model.Product{}, err
	}

	id, _ := result.LastInsertId()
	product.ID = int(id)
	return product, nil
}
```

এভাবে বাকি method গুলো implement করা হবে।

---

## Step 4️⃣: Service Layer Use

```go
type ProductService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(name string, price float64) error {
	product := model.Product{
		Name:  name,
		Price: price,
	}
	_, err := s.repo.Create(product)
	return err
}
```

এখানে service জানেই না MySQL, PostgreSQL নাকি MongoDB ব্যবহার হচ্ছে।
এটাই Repository Pattern এর power।

---

# 5️⃣ Real Database Change Scenario

ধরো আজ MySQL ব্যবহার করছো।
আগামীকাল PostgreSQL ব্যবহার করতে চাও।

তুমি শুধু নতুন implementation লিখবে:

```go
type postgresProductRepo struct { ... }
```

Business logic untouched থাকবে।

---

# 6️⃣ Advantages & Disadvantages

## ✅ Advantages

* Loose coupling
* Easy testing (mock repository)
* Clean architecture friendly
* Database independence
* Maintainability high

## ❌ Disadvantages

* Extra abstraction layer
* Small project হলে overengineering হতে পারে
* Boilerplate code বেশি হয়

---

# 7️⃣ Interview Questions (Basic → Advanced)

## 🟢 Basic

1. Repository Pattern কী?
2. Repository Pattern আর DAO কি একই?
3. কেন Repository ব্যবহার করা হয়?

---

## 🟡 Intermediate

4. Repository Pattern কিভাবে loose coupling তৈরি করে?
5. Go-তে interface কেন গুরুত্বপূর্ণ Repository pattern-এ?
6. Repository এবং Service layer-এর মধ্যে পার্থক্য কী?

---

## 🔴 Advanced

7. Repository pattern কি microservice architecture-এ দরকার?
8. Repository vs Unit of Work pattern পার্থক্য কী?
9. Repository pattern overuse করলে কী সমস্যা হতে পারে?
10. Clean Architecture-এ Repository কোথায় বসে?

---

# 8️⃣ Repository vs Direct DB Access

| বিষয়            | Direct DB Access | Repository |
| --------------- | ---------------- | ---------- |
| Coupling        | High             | Low        |
| Testing         | Hard             | Easy       |
| Maintainability | Low              | High       |
| DB Change       | Difficult        | Easy       |

---

# 9️⃣ সংক্ষেপে

```
Repository Pattern =

✔ Database logic encapsulation
✔ Business logic clean রাখা
✔ Interface-based abstraction
✔ Easy testing
✔ Database independent design
```

