package product

import (
	"net/http"
	"ecommerce/util"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {

	lst, err := h.productRepo.List()
	if err != nil {
		println(err)

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	util.SendData(w, lst, http.StatusOK)
}
