package product

import (
	"ecommerce/repo"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type RequestCreateProduct struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	ImageURL    string  `json:"img"`
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var newProduct RequestCreateProduct

	decoder := json.NewDecoder(r.Body)
	error := decoder.Decode(&newProduct)

	if error != nil {
		fmt.Println(error)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	Created_Product, err := h.productRepo.Create(repo.Product{
		Title: newProduct.Title,
		Description: newProduct.Description,
		Price: newProduct.Price,
		ImageURL: newProduct.ImageURL,
	})
	if err != nil {
		fmt.Println(error)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	util.SendData(w, Created_Product, http.StatusCreated)
}
