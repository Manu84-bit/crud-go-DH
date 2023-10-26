package product

import "CRUD-GO/internal/domain"

type IService interface {
	GetByIdentifier(id int)(*domain.Product, error)
}

type Service struct {
	Repository IRepository
}
func (s *Service) GetByIdentifier(id int)(*domain.Product, error){
	prod, err := s.Repository.GetById(id)
	if err != nil{
		return nil, err
	}

	return prod, nil
}