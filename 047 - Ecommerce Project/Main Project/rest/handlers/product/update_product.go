package product

import (
	"ecommerce/repo"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Product struct {
	ID          int     `json:"id"` // For customize json
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	ImageURL    string  `json:"img"`
}

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	product_id := r.PathValue("id")
	pId, error := strconv.Atoi(product_id)
	var product Product

	decoder := json.NewDecoder(r.Body)
	error = decoder.Decode(&product)

	if error != nil {
		fmt.Println(error)

		http.Error(w, "Something want wrong", 200)
		return
	}

	product.ID = pId
	updatedProduct, err := h.productRepo.Update(repo.Product{
		ID: product.ID,
		Title: product.Title,
		Description: product.Description,
		Price: product.Price,
		ImageURL: product.ImageURL,
	})
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	util.SendData(w, updatedProduct, 201)
}