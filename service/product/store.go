package product

import (
	"database/sql"
	"fmt"

	"github.com/arturfil/go_repository_hex/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
    return &Store{db: db}
}

func (s *Store) GetProducts() ([]*types.Product, error) {
    rows, err := s.db.Query("SELECT * FROM products")
    if err != nil {
        return nil, err
    }

    var products []*types.Product
    for rows.Next() {
        product, err := scanRowIntoProduct(rows)
        if err != nil {
            fmt.Println("error!", err)
            return nil, err
        }
        products = append(products, product)
    }

    return products, nil
}

func scanRowIntoProduct(rows *sql.Rows) (*types.Product, error) {
    product := &types.Product{}

    err := rows.Scan(
        &product.ID,
        &product.Name,
        &product.Descrtipion,
        &product.Image,
        &product.Price,
        &product.Quantity,
        &product.CreatedAt,
        &product.UpdatedAt,
    )
    if err != nil {
        return nil, err
    }

    return product, nil
}
