package repo

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Product struct {
	ID          int     `json:"id" db:"id"` // For customize json
	Title       string  `json:"title" db:"title"`
	Description string  `json:"description" db:"description"`
	Price       float32 `json:"price" db:"price"`
	ImageURL    string  `json:"image_url" db:"image_url"`
}

type ProductRepo interface {
	Create(prd Product) (*Product, error)
	Get(productID int) (*Product, error)
	List() ([]*Product, error)
	Update(prd Product) (*Product, error)
	Delete(productID int) error
}

type productRepo struct {
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) ProductRepo {
	pList := &productRepo{
		db: db,
	}

	return pList
}

func (p *productRepo) Create(prd Product) (*Product, error) {
	query := `
		INSERT INTO products (
			title,
			description,
			price,
			image_url
		) VALUES (
			$1,
			$2,
			$3,
			$4 
		)
		RETURNING id
	`

	row := p.db.QueryRow(query, prd.Title, prd.Description, prd.Price, prd.ImageURL)
	err := row.Scan(&prd.ID)
	if err != nil {
		return nil, err
	}

	return &prd, nil
}

func (p *productRepo) Get(productID int) (*Product, error) {
	var product Product

	query := `
		SELECT 
			id, 
			title,
			description,
			price,
			image_url
		FROM products
		WHERE id = $1
	`

	err := p.db.Get(&product, query, productID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &product, nil
}

func (p *productRepo) List() ([]*Product, error) {
	var productList []*Product

	query := `
		SELECT 
			id,
			title,
			description,
			price,
			image_url
		FROM products
	`

	err := p.db.Select(&productList, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return productList, nil
}

func (p *productRepo) Update(prd Product) (*Product, error) {
	query := `
		UPDATE products
		SET 
			title=$1, 
			description=$2, 
			price=$3, 
			image_url=$4
		WHERE id = $5
	`

	row := p.db.QueryRow(query, prd.Title, prd.Description, prd.Price, prd.ImageURL, prd.ID)
	if err := row.Err(); err != nil {
		return nil, err
	}

	return &prd, nil
}

func (p *productRepo) Delete(productID int) error {
	query := `
		DELETE FROM products
		WHERE id = $1
	`
	_, err := p.db.Exec(query, productID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return err
	}

	return nil
}
