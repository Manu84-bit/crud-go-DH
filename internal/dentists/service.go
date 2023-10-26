package dentists

import "CRUD-FINAL/internal/domain"

type IService interface {
	GetByIdentifier(id int)(*domain.Dentist, error)
	GetAllDentists()([]domain.Dentist, error)
	SaveDentist(d *domain.Dentist)(int, error)
	DeleteDentist(id int)(int, error)
	UpdateDentist(d domain.Dentist)(int, error)
}

type Service struct {
	Repository IRepository
}

func (s *Service) GetAllDentists()([]domain.Dentist, error){
	dentists, err := s.Repository.GetAll()
		if err != nil{
		return nil, err
	}

	return dentists, nil
}
func (s *Service) GetByIdentifier(id int)(*domain.Dentist, error){
	d, err := s.Repository.GetById(id)
	if err != nil{
		return nil, err
	}

	return d, nil
}

func (s *Service) DeleteDentist(id int)(int, error){
	id,err := s.Repository.DeleteById(id);
		if err != nil{
		return 0, err
	}
	return id, nil
}

func (s *Service) SaveDentist(d *domain.Dentist)(int, error){
  id, err := s.Repository.Save(d)
  if err != nil{
		return 0, err
	}
	return id, nil
}

func (s *Service) UpdateDentist(d domain.Dentist)(int, error){
	id, err := s.Repository.Update(d)
	 if err != nil{
		return 0, err
	}
	return id, nil
}



