package product

import (
	"CRUD-GO/internal/domain"
	"CRUD-GO/pkg/store"
	"CRUD-GO/pkg/web"
	"fmt"
)

type IRepository interface {
	GetById(id int) (*domain.Product, error)
}

type RepositoryImpl struct {
	Store store.StoreInteface 
}

func (r *RepositoryImpl) GetById(id int) (*domain.Product, error){
	product, err :=	r.Store.Read(id)
	if err != nil {
		return nil, web.NewNotFoundException(fmt.Sprintf("product id %d not found", id))
	}

	return product, nil
}