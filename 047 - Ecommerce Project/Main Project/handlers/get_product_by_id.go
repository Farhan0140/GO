package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func GetProductById(w http.ResponseWriter, r *http.Request) {
	pId := r.PathValue("productID")

	id, err := strconv.Atoi(pId)

	if err != nil {
		http.Error(w, "Enter valid Product id", 400)
		return
	}

	for _, product := range database.Products {
		if product.ID == id {
			util.SendData(w, product, 200)
			return
		}
	}

	util.SendData(w, "No data found from given id", 404)
}