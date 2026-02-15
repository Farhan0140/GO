package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func GetProductById(w http.ResponseWriter, r *http.Request) {
	pId := r.PathValue("id")

	id, err := strconv.Atoi(pId)

	if err != nil {
		http.Error(w, "Enter valid Product id", 400)
		return
	}
	
	product := database.Get(id)
	if product != nil {
		util.SendData(w, product, 200)
		return
	}

	util.SendError(w, 404, "No data found from given id")
}