package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"fmt"
	"net/http"
	"strconv"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")
	pId, error := strconv.Atoi(productId)
	if error != nil {
		fmt.Println(error)

		http.Error(w, "Something want wrong", 200)
		return
	}

	database.Delete(pId)
	util.SendData(w, "Product Successfully Deleted", 200)
}