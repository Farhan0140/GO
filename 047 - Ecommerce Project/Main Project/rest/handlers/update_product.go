package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	product_id := r.PathValue("id")
	pId, error := strconv.Atoi(product_id)
	var product database.Product

	decoder := json.NewDecoder(r.Body)
	error = decoder.Decode(&product)

	if error != nil {
		fmt.Println(error)

		http.Error(w, "Something want wrong", 200)
		return
	}

	product.ID = pId
	updatedProduct := database.Update(product)

	util.SendData(w, updatedProduct, 201)
}