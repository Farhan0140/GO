package repo

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int    `json:"id" db:"id"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"password" db:"password"`
	IsAdmin   bool   `json:"is_admin" db:"is_admin"`
}

type UserRepo interface {
	Create(user User) (*User, error)
	Find(email, password string) (*User, error)
	// Get(userID int) (*User, error)
	// List()	(*[]User, error)
	// Update(user User) (*User, error)
	// Delete(userID int) error
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) Create(user User) (*User, error) {
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, err
	}

	user.Password = string(hashedPassword)

	query := `
		INSERT INTO users (
			first_name, 
			last_name, 
			email, 
			password, 
			is_admin
		)
		VALUES (
			:first_name,
			:last_name,
			:email,
			:password,
			:is_admin
		)
		RETURNING id
	`

	var userId int
	rows, err := u.db.NamedQuery(query, user)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(&userId)
	}

	user.ID = userId

	fmt.Println("New User Id ", userId)
	return &user, err
}

func (u *userRepo) Find(email, password string) (*User, error) {
	var user User
	query := `
		SELECT id, first_name, last_name, email, password, is_admin
		FROM users
		WHERE email = $1
		LIMIT 1
	`

	err := u.db.Get(&user, query, email)
	if err != nil {
		fmt.Println(err)
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}

	return &user, nil
}
