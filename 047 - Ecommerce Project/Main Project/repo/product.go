package repo

type Product struct {
	ID          int     `json:"id"` // For customize json
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	ImageURL    string  `json:"img"`
}

type ProductRepo interface {
	Create(prd Product) (*Product, error)
	Get(productID int) (*Product, error)
	List()	(*[]Product, error)
	Update(prd Product) (*Product, error)
	Delete(productID int) error
}

type productRepo struct {
	products []Product
}

func NewProductRepo() ProductRepo {
	pList := &productRepo{}

	GenerateInitialProducts(pList)
	return pList
}

func (p *productRepo) Create(prd Product) (*Product, error) {
	_len := len(p.products)

	if _len == 0 {
		prd.ID = 0
		p.products = append(p.products, prd)
	} else {
		_len--
		prd.ID = p.products[_len].ID + 1
		p.products = append(p.products, prd)
	}
	return &prd, nil
}

func (p *productRepo) Get(productID int) (*Product, error) {
	for _, prd := range p.products {
		if prd.ID == productID {
			return &prd, nil
		}
	}

	return nil, nil
}

func (p *productRepo) List() (*[]Product, error) {
	return &p.products, nil
}

func (p *productRepo) Update(prd Product) (*Product, error) {
	var new_product Product
	for idx, product := range p.products {
		if product.ID == prd.ID {
			p.products[idx] = prd
			new_product = p.products[idx]
		}
	}

	return &new_product, nil
}

func (p *productRepo) Delete(productID int) error {
	var temp []Product

	for _, product := range p.products {
		if product.ID != productID {
			temp = append(temp, product)
		}
	}

	p.products = temp
	return nil
}

func GenerateInitialProducts(p *productRepo) {
	prd1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "Orange is Orange",
		Price:       160.34,
		ImageURL:    "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQVc1ggU1JM6immtNCh08GzQoO3pdcxBtesN8p8T8CqvkZFj3pj",
	}

	prd2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "Apple is Green",
		Price:       100.00,
		ImageURL:    "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR5j8eIjlCdEweGdYheC_xL0DtYPhajuq-sVA&s",
	}

	prd3 := Product{
		ID:          3,
		Title:       "Banana",
		Description: "Banana is Yellow",
		Price:       60.50,
		ImageURL:    "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQc9bH2z92pEalMYHCHnK3ii7Y6WLMNv8IgE2RJ-na3rBhvXb0dbjQQbrT7ODSaL45f--wj6nN8pHrI3rSAlcFX3dlnNzz9lG_aU6iPoBw&s=10",
	}

	prd4 := Product{
		ID:          4,
		Title:       "Mango",
		Description: "Mango is Sweet",
		Price:       180.75,
		ImageURL:    "https://encrypted-tbn1.gstatic.com/images?q=tbn:ANd9GcTUb33Kqs_oPuUvWhURRtEEYDd7K3zpjvOmGuFwt0quwKCFyMS3",
	}

	prd5 := Product{
		ID:          5,
		Title:       "Pineapple",
		Description: "Pineapple is Tropical",
		Price:       220.00,
		ImageURL:    "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQpineapple",
	}

	prd6 := Product{
		ID:          6,
		Title:       "Grapes",
		Description: "Grapes are Purple",
		Price:       140.25,
		ImageURL:    "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQgrapes",
	}

	prd7 := Product{
		ID:          7,
		Title:       "Watermelon",
		Description: "Watermelon is Juicy",
		Price:       300.00,
		ImageURL:    "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQwatermelon",
	}

	prd8 := Product{
		ID:          8,
		Title:       "Papaya",
		Description: "Papaya is Healthy",
		Price:       120.90,
		ImageURL:    "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQpapaya",
	}

	prd9 := Product{
		ID:          9,
		Title:       "Strawberry",
		Description: "Strawberry is Red",
		Price:       250.40,
		ImageURL:    "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQstrawberry",
	}

	prd10 := Product{
		ID:          10,
		Title:       "Guava",
		Description: "Guava is Fresh",
		Price:       90.00,
		ImageURL:    "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQguava",
	}

	p.products = append(
		p.products,
		prd1,
		prd2,
		prd3,
		prd4,
		prd5,
		prd6,
		prd7,
		prd8,
		prd9,
		prd10,
	)
}
