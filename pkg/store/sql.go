package store

import (
	"CRUD-GO/internal/domain"
	"database/sql"
)

type SqlStoreImp struct {
	DB *sql.DB
}

func (s *SqlStoreImp) Read(id int)(*domain.Product, error) {
	var product domain.Product
	query:= "SELECT * FROM products WHERE id = ?"
	row := s.DB.QueryRow(query, id)
	err:= row.Scan(
		&product.Id, 
		&product.Name,
		&product.Quantity, 
		&product.CodeValue,
		&product.IsPublished, 
		&product.Expiration,
		&product.Price,
		)
	if err != nil {
		return nil, err
	}

	return &product, err
}