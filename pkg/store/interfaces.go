package store

import "CRUD-GO/internal/domain"

type StoreInteface interface {
	Read(id int) (*domain.Product, error)
}