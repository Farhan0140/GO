package repo

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	IsAdmin   bool   `json:"is_admin"`
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
	users []User
}

func NewUserRepo() UserRepo {
	return &userRepo{}
}

func (u *userRepo) Create(user User) (*User, error) {
	_len := len(u.users)

	if _len == 0 {
		user.ID = 0
		u.users = append(u.users, user)
	} else {
		_len--
		user.ID = u.users[_len].ID + 1
		u.users = append(u.users, user)
	}

	u.users = append(u.users, user)
	return &user, nil
}

func (u *userRepo) Find(email, password string) (*User, error) {
	for _, usr := range u.users {
		if usr.Email == email && usr.Password == password {
			return &usr, nil
		}
	}

	return nil, nil
}