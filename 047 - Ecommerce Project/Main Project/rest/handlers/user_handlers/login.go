package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type requestLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var reqLogin requestLogin

	decoder := json.NewDecoder(r.Body)
	error := decoder.Decode(&reqLogin)

	if error != nil {
		fmt.Println(error)

		http.Error(w, "Invalid Request Data", http.StatusBadRequest)
		return
	}

	usr := database.Find(reqLogin.Email, reqLogin.Password)
	if usr == nil {
		http.Error(w, "Invalid credentials", http.StatusBadRequest)
	}

	util.SendData(w, usr, http.StatusCreated)
}
