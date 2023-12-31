package patients

import "CRUD-FINAL/internal/domain"

type IService interface {

	GetAllPatients()([]domain.Patient, error)
	GetByIdentifier(id int)(*domain.Patient, error)
	SavePatient(d *domain.Patient)(int, error)
	DeletePatient(id int)(int, error)
	UpdatePatient(d domain.Patient)(int, error)
}

type Service struct {
	Repository IRepository
}

func (s *Service) GetAllPatients()([]domain.Patient, error){
	patients, err := s.Repository.GetAll()
		if err != nil{
		return nil, err
	}

	return patients, nil
}

func (s *Service) GetByIdentifier(id int)(*domain.Patient, error){
	d, err := s.Repository.GetById(id)
	if err != nil{
		return nil, err
	}

	return d, nil
}

func (s *Service) DeletePatient(id int)(int, error){
	id,err := s.Repository.DeleteById(id);
		if err != nil{
		return 0, err
	}
	return id, nil
}

func (s *Service) SavePatient(p *domain.Patient)(int, error){
  id, err := s.Repository.Save(p)
  if err != nil{
		return 0, err
	}
	return id, nil
}

func (s *Service) UpdatePatient(p domain.Patient)(int, error){
	id, err := s.Repository.Update(p)
	 if err != nil{
		return 0, err
	}
	return id, nil
}
