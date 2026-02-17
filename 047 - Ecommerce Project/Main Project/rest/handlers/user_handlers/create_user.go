package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser database.User

	decoder := json.NewDecoder(r.Body)
	error := decoder.Decode(&newUser)

	if error != nil {
		fmt.Println(error)

		http.Error(w, "Invalid Request Data", http.StatusBadRequest)
		return
	}

	Created_User := newUser.Store()

	util.SendData(w, Created_User, http.StatusCreated)
}