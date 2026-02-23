package product

import (
	"net/http"
	"ecommerce/util"
	"ecommerce/database"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	util.SendData(w, database.List(), 200)
}
