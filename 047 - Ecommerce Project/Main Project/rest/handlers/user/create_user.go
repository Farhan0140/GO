package user

import (
	"ecommerce/repo"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	IsAdmin   bool   `json:"is_admin"`
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser User

	decoder := json.NewDecoder(r.Body)
	error := decoder.Decode(&newUser)

	if error != nil {
		fmt.Println(error)

		http.Error(w, "Invalid Request Data", http.StatusBadRequest)
		return
	}

	Created_User, err := h.usrRepo.Create(repo.User{
		FirstName: newUser.FirstName,
		LastName: newUser.LastName,
		Email: newUser.Email,
		Password: newUser.Password,
		IsAdmin: newUser.IsAdmin,
	})
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	util.SendData(w, Created_User, http.StatusCreated)
}
