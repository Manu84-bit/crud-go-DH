package dentists

import (
	"CRUD-FINAL/internal/domain"
	"CRUD-FINAL/pkg/clinic"
	"CRUD-FINAL/pkg/web"
	"fmt"
)

type IRepository interface {
	Save(dentist *domain.Dentist) (int, error)
	GetById(id int) (*domain.Dentist, error)
	GetAll()([]domain.Dentist, error)
	Update(d domain.Dentist) (int, error)
	DeleteById(id int) (int, error)
}

type RepositoryImpl struct {
	DentistRepo clinic.SQLDentistInteface
}

func (r *RepositoryImpl) Save(dentist *domain.Dentist) (int, error){
	newDentist, err := r.DentistRepo.Create(dentist)
	if err != nil{
		return 0, web.NewBadRequestApiError("Wrong data")
	}

	return newDentist, nil
}

func (r *RepositoryImpl) GetAll()([]domain.Dentist, error){
	dentists, err := r.DentistRepo.ReadAll()
		if err != nil {
		return nil, web.NewNotFoundException("dentists not found")
	}

	return dentists, nil
}

func (r *RepositoryImpl) GetById(id int) (*domain.Dentist, error){
	dentist, err :=	r.DentistRepo.Read(id)
	if err != nil {
		return nil, web.NewNotFoundException(fmt.Sprintf("dentist id %d not found", id))
	}

	return dentist, nil
}

func (r *RepositoryImpl) DeleteById(id int) (int, error) {
  idDentist, err := r.DentistRepo.Delete(id)
  if err != nil {
	return 0, web.NewNotFoundException(fmt.Sprintf("dentist id %d not exists", idDentist))
  }
  return idDentist, nil
}

func (r *RepositoryImpl) Update(d domain.Dentist) (int, error){
idDentist, err := r.DentistRepo.Update(d)
if err != nil {
	return 0, web.NewNotFoundException(fmt.Sprintf("dentist id %d not exists", idDentist))
  }

  return idDentist, nil
}
