package handlers

import (
	"net/http"
	"ecommerce/util"
	"ecommerce/database"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	util.SendData(w, database.List(), 200)
}
