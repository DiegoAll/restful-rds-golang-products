package repository

import (
	"context"
	"database/sql"
	"restful-rds-golang-products/models"
)

// InternalTrafficRepository es una interfaz para operaciones de tráfico interno (productos).
type ProductRepositoryInterface interface {
	InsertProduct(ctx context.Context, record *models.Product) error
	// Aquí se pueden añadir más métodos para productos, como Get, Update, Delete, GetAll
	GetAllProducts(ctx context.Context) ([]models.Product, error)
}

// ProductsRepository implementa InternalTrafficRepository para PostgreSQL.
type ProductsRepository struct {
	DB *sql.DB
}

// NewPostgresInternalTrafficRepository crea una nueva instancia de ProductsRepository.
func NewPostgresProductRepository(db *sql.DB) *ProductsRepository {
	return &ProductsRepository{DB: db}
}

// InsertProduct inserta un nuevo producto en la base de datos.
func (r *ProductsRepository) InsertProduct(ctx context.Context, product *models.Product) error {
	query := `
        INSERT INTO products (name, description, price, created_at, updated_at)
        VALUES ($1, $2, $3, NOW(), NOW())
        RETURNING id` // Usamos RETURNING id para obtener el ID generado

	err := r.DB.QueryRowContext(
		ctx,
		query,
		product.Name,
		product.Description,
		product.Price,
	).Scan(&product.Id) // Escaneamos el ID generado de vuelta al modelo
	return err
}

// GetAllProducts obtiene todos los productos de la base de datos.
func (r *ProductsRepository) GetAllProducts(ctx context.Context) ([]models.Product, error) {
	query := `
        SELECT id, name, description, price, created_at, updated_at
        FROM products`

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
