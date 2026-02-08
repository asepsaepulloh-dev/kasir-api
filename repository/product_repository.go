package repository

import (
	"database/sql"
	"errors"
	"kasir-api/model"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (repo *ProductRepository) GetAll(name string) ([]model.Product, error) {
	query := `SELECT p.id, p.name, p.price, p.stock, c.id, c.name, c.description 
			  FROM products p 
			  LEFT JOIN categories c ON p.category_id = c.id`

	args := []interface{}{}

	if name != "" {
		query += " WHERE p.name ILIKE $1"
		args = append(args, "%"+name+"%")
	}

	rows, err := repo.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]model.Product, 0)

	for rows.Next() {
		var p model.Product
		var categoryID sql.NullInt64
		var categoryName sql.NullString
		var categoryDescription sql.NullString

		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &categoryID, &categoryName, &categoryDescription)
		if err != nil {
			return nil, err
		}

		if categoryID.Valid {
			p.Category = &model.Category{
				ID:          int(categoryID.Int64),
				Name:        categoryName.String,
				Description: categoryDescription.String,
			}
		}

		products = append(products, p)
	}

	return products, nil
}

func (repo *ProductRepository) Create(product *model.Product) error {
	var categoryID *int

	// If category name is provided, look up the category ID
	if product.CategoryName != "" {
		var id int
		categoryQuery := "SELECT id FROM categories WHERE name = $1"
		err := repo.db.QueryRow(categoryQuery, product.CategoryName).Scan(&id)
		if err == sql.ErrNoRows {
			return errors.New("category not found")
		}
		if err != nil {
			return err
		}
		categoryID = &id
	}

	// Insert product with category_id
	query := "INSERT INTO products (name, price, stock, category_id) VALUES ($1, $2, $3, $4) RETURNING id"
	err := repo.db.QueryRow(query, product.Name, product.Price, product.Stock, categoryID).Scan(&product.ID)
	return err
}

// GetByID - ambil produk by ID
func (repo *ProductRepository) GetByID(id int) (*model.Product, error) {
	query := `SELECT p.id, p.name, p.price, p.stock, c.id, c.name, c.description 
			  FROM products p 
			  LEFT JOIN categories c ON p.category_id = c.id 
			  WHERE p.id = $1`

	var p model.Product
	var categoryID sql.NullInt64
	var categoryName sql.NullString
	var categoryDescription sql.NullString

	err := repo.db.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &categoryID, &categoryName, &categoryDescription)
	if err == sql.ErrNoRows {
		return nil, errors.New("produk tidak ditemukan")
	}
	if err != nil {
		return nil, err
	}

	if categoryID.Valid {
		p.Category = &model.Category{
			ID:          int(categoryID.Int64),
			Name:        categoryName.String,
			Description: categoryDescription.String,
		}
	}

	return &p, nil
}

func (repo *ProductRepository) Update(product *model.Product) error {
	var categoryID *int

	// If category name is provided, look up the category ID
	if product.CategoryName != "" {
		var id int
		categoryQuery := "SELECT id FROM categories WHERE name = $1"
		err := repo.db.QueryRow(categoryQuery, product.CategoryName).Scan(&id)
		if err == sql.ErrNoRows {
			return errors.New("category not found")
		}
		if err != nil {
			return err
		}
		categoryID = &id
	}

	// Update product with category_id
	query := "UPDATE products SET name = $1, price = $2, stock = $3, category_id = $4 WHERE id = $5"
	result, err := repo.db.Exec(query, product.Name, product.Price, product.Stock, categoryID, product.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("produk tidak ditemukan")
	}

	return nil
}

func (repo *ProductRepository) Delete(id int) error {
	query := "DELETE FROM products WHERE id = $1"
	result, err := repo.db.Exec(query, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("produk tidak ditemukan")
	}

	return err
}
