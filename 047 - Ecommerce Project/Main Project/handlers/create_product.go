package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ecommerce/database"
	"ecommerce/util"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct database.Product

	decoder := json.NewDecoder(r.Body)
	error := decoder.Decode(&newProduct)

	if error != nil {
		fmt.Println(error)

		http.Error(w, "Something want wrong", 200)
		return
	}

	newProduct.ID = len(database.Products) + 1
	database.Products = append(database.Products, newProduct)

	util.SendData(w, newProduct, 201)
}
