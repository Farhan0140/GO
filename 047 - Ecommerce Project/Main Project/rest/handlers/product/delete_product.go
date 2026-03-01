package product

import (
	"ecommerce/util"
	"fmt"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")
	pId, error := strconv.Atoi(productId)
	if error != nil {
		fmt.Println(error)

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	error = h.productRepo.Delete(pId)
	if error != nil {
		fmt.Println(error)

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	util.SendData(w, "Product Successfully Deleted", 200)
}