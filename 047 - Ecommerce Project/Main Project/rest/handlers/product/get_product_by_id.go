package product

import (
	"ecommerce/util"
	"net/http"
	"strconv"
)

func (h *Handler) GetProductById(w http.ResponseWriter, r *http.Request) {
	pId := r.PathValue("id")

	id, err := strconv.Atoi(pId)

	if err != nil {
		http.Error(w, "Enter valid Product id", 400)
		return
	}
	
	product, err := h.productRepo.Get(id)
	if product != nil {
		util.SendData(w, product, 200)
		return
	}
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	util.SendError(w, 404, "No data found from given id")
}