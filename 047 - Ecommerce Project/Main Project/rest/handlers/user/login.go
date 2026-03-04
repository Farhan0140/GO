package user

import (
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type requestLogin struct {
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var reqLogin requestLogin

	decoder := json.NewDecoder(r.Body)
	error := decoder.Decode(&reqLogin)

	if error != nil {
		fmt.Println(error)

		http.Error(w, "Invalid Request Data", http.StatusBadRequest)
		return
	}

	usr, err := h.usrRepo.Find(reqLogin.Email, reqLogin.Password) // database.Find()
	if usr == nil {
		http.Error(w, "Invalid credentials", http.StatusBadRequest)
		return
	}
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	access_token, err := util.Create_JWT(h.cnf.SecretKey, util.Payload{
		ID: usr.ID,
		FirstName: usr.FirstName,
		LastName: usr.LastName,
		Email: usr.Email,
		IsAdmin: usr.IsAdmin,
	})
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	util.SendData(w, access_token, http.StatusCreated)
}
