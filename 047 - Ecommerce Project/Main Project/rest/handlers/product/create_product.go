package product

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var newProduct database.Product

	decoder := json.NewDecoder(r.Body)
	error := decoder.Decode(&newProduct)

	if error != nil {
		fmt.Println(error)

		http.Error(w, "Something want wrong", 200)
		return
	}

	Created_Product := database.Store(newProduct)

	util.SendData(w, Created_Product, 201)
}
