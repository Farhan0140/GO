# Pagination | Query Parameters | Offset

## What is Pagination?

Pagination হলো অনেকগুলো data একসাথে না পাঠিয়ে ছোট ছোট অংশে (pages) ভাগ করে পাঠানোর পদ্ধতি।

উদাহরণ:

ধরুন আপনার database-এ **১,০০,০০০ users** আছে। যদি API একবারে সব users পাঠায়:

```http
GET /users
```

তাহলে:

* Response অনেক বড় হবে।
* Memory বেশি লাগবে।
* Database slow হয়ে যাবে।
* Network traffic বেড়ে যাবে।

তাই Pagination ব্যবহার করা হয়।

---

# What are Query Parameters?

URL-এর `?` এর পরে যেগুলো আসে সেগুলোকে Query Parameters বলে।

উদাহরণ:

```http
GET /users?page=1&limit=10
```

এখানে:

| Parameter | Meaning                  |
| --------- | ------------------------ |
| page      | কোন page চান             |
| limit     | প্রতি page-এ কত data চান |

---

# Pagination Example

Database Table:

| id | name   |
| -- | ------ |
| 1  | Farhan |
| 2  | Rahim  |
| 3  | Karim  |
| 4  | Hasan  |
| 5  | Nadim  |
| 6  | Sakib  |
| 7  | Tamim  |
| 8  | Rifat  |
| 9  | Jahid  |
| 10 | Rakib  |
| 11 | Nafis  |
| 12 | Rony   |

---

Request:

```http
GET /users?page=1&limit=5
```

Response:

```text
1 2 3 4 5
```

---

Request:

```http
GET /users?page=2&limit=5
```

Response:

```text
6 7 8 9 10
```

---

Request:

```http
GET /users?page=3&limit=5
```

Response:

```text
11 12
```

---

# What is Offset?

Offset মানে হলো কতগুলো row skip করতে হবে।

Formula:

```text
offset = (page - 1) * limit
```

---

## Example: page = 1

```text
offset = (1 - 1) * 10
       = 0
```

---

## Example: page = 2

```text
offset = (2 - 1) * 10
       = 10
```

---

## Example: page = 3

```text
offset = (3 - 1) * 10
       = 20
```

---

# SQL Example

```sql
SELECT *
FROM users
LIMIT 10
OFFSET 20;
```

মানে:

```text
প্রথম ২০টা row skip করো,
পরের ১০টা row দাও।
```

---

# Repository Layer Example (Go)

ধরি:

```go
type User struct {
	ID   int
	Name string
}
```

Repository:

```go
func (r *UserRepo) GetUsers(
	ctx context.Context,
	limit int,
	offset int,
) ([]User, error) {

	query := `
		SELECT id, name
		FROM users
		LIMIT $1
		OFFSET $2
	`

	rows, err := r.db.QueryContext(
		ctx,
		query,
		limit,
		offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []User{}

	for rows.Next() {
		var u User

		err := rows.Scan(
			&u.ID,
			&u.Name,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil
}
```

---

# Repository for Total Count

Pagination metadata পাঠানোর জন্য total row count দরকার।

```go
func (r *UserRepo) CountUsers(
	ctx context.Context,
) (int, error) {

	var total int

	query := `
		SELECT COUNT(*)
		FROM users
	`

	err := r.db.QueryRowContext(
		ctx,
		query,
	).Scan(&total)

	return total, err
}
```

---

# Handler Example

```go
func (h *Handler) GetUsers(
	w http.ResponseWriter,
	r *http.Request,
) {
	page := 1
	limit := 10

	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}

	if limitStr != "" {
		limit, _ = strconv.Atoi(limitStr)
	}

	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	users, err := h.repo.GetUsers(
		r.Context(),
		limit,
		offset,
	)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	total, err := h.repo.CountUsers(
		r.Context(),
	)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	resp := map[string]any{
		"page":  page,
		"limit": limit,
		"total": total,
		"data":  users,
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(resp)
}
```

---

# Request Example

```http
GET /users?page=2&limit=5
```

---

# JSON Response দেখতে কেমন হয়?

```json
{
  "page": 2,
  "limit": 5,
  "total": 12,
  "data": [
    {
      "id": 6,
      "name": "Sakib"
    },
    {
      "id": 7,
      "name": "Tamim"
    },
    {
      "id": 8,
      "name": "Rifat"
    },
    {
      "id": 9,
      "name": "Jahid"
    },
    {
      "id": 10,
      "name": "Rakib"
    }
  ]
}
```

---

# Pagination JSON-এর Elements

| Field | Meaning                  |
| ----- | ------------------------ |
| page  | বর্তমান page             |
| limit | প্রতি page-এ item সংখ্যা |
| total | database-এ মোট item      |
| data  | বর্তমান page-এর data     |

---

# Better Pagination Response (Production API)

Production API-তে সাধারণত আরও metadata পাঠানো হয়।

```json
{
  "page": 2,
  "limit": 5,
  "total": 12,
  "total_pages": 3,
  "has_next": true,
  "has_prev": true,
  "next_page": 3,
  "prev_page": 1,
  "data": [
    {
      "id": 6,
      "name": "Sakib"
    },
    {
      "id": 7,
      "name": "Tamim"
    }
  ]
}
```

---

# Meaning of Each Field

| Field       | Meaning            |
| ----------- | ------------------ |
| page        | Current page       |
| limit       | Items per page     |
| total       | Total rows         |
| total_pages | মোট কত page        |
| has_next    | পরের page আছে কিনা |
| has_prev    | আগের page আছে কিনা |
| next_page   | পরের page number   |
| prev_page   | আগের page number   |
| data        | Actual data        |

---

# Calculate Total Pages

```go
totalPages := int(math.Ceil(
	float64(total) /
	float64(limit),
))
```

---

# Production Response Struct

```go
type PaginationResponse struct {
	Page       int         `json:"page"`
	Limit      int         `json:"limit"`
	Total      int         `json:"total"`
	TotalPages int         `json:"total_pages"`
	HasNext    bool        `json:"has_next"`
	HasPrev    bool        `json:"has_prev"`
	NextPage   int         `json:"next_page,omitempty"`
	PrevPage   int         `json:"prev_page,omitempty"`
	Data        interface{} `json:"data"`
}
```

---

# Request → Response Flow

```text
Client
   ↓
GET /users?page=2&limit=5
   ↓
Handler
   ↓
offset=(2-1)*5=5
   ↓
Repository
   ↓
SELECT * FROM users
LIMIT 5 OFFSET 5
   ↓
Database
   ↓
JSON Response
```

---

# Complete Flow Explanation

1. Client URL-এ `page` এবং `limit` পাঠায়।
2. Handler Query Parameters থেকে `page` এবং `limit` read করে।
3. Handler offset calculate করে:

```text
offset = (page - 1) * limit
```

4. Repository `LIMIT` এবং `OFFSET` ব্যবহার করে data fetch করে।
5. Repository total row count বের করে।
6. Handler pagination metadata তৈরি করে।
7. JSON response client-কে return করে।

---

# Summary

Pagination API-তে সবচেয়ে common fields:

| Field       | Description                 |
| ----------- | --------------------------- |
| page        | Current page number         |
| limit       | Number of items per page    |
| total       | Total rows in database      |
| total_pages | Total number of pages       |
| has_next    | Next page exists or not     |
| has_prev    | Previous page exists or not |
| next_page   | Next page number            |
| prev_page   | Previous page number        |
| data        | Actual data of current page |

---

# Final Formula

```text
offset = (page - 1) * limit

total_page = total_item / limit
```

এটাই হলো Go backend-এ সবচেয়ে common **Offset-based Pagination Implementation**।
