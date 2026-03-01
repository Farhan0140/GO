package database

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	IsAdmin   bool   `json:"is_admin"`
}

var users []User

func (u User) Store() User {
	// if u.ID != 0 {
	// 	return u
	// }

	_len := len(users)

	if _len == 0 {
		u.ID = 0
		users = append(users, u)
	} else {
		_len--
		u.ID = users[_len].ID + 1
		users = append(users, u)
	}

	users = append(users, u)
	return u
}

func Find(email, password string) *User {
	for _, usr := range users {
		if usr.Email == email && usr.Password == password {
			return &usr
		}
	}

	return nil
}
