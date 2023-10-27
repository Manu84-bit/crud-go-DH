package appointments

import (
	"CRUD-FINAL/internal/domain"
)

type IService interface {
	GetByIdentifier(id int)(*domain.AppointmentDTO, error)
	GetAllAppointments()([]domain.AppointmentDTO, error)
	SaveAppointment(d *domain.Appointment)(int, error)
	DeleteAppointment(id int)(int, error)
	UpdateAppointment(a domain.Appointment)(int, error)
}

type Service struct {
	Repository IRepository
}
func (s *Service) GetByIdentifier(id int)(*domain.AppointmentDTO, error){
	a, err := s.Repository.GetByIdentifier(id)
	if err != nil{
		return nil, err
	}

	return a, nil
}

func (s *Service) GetAllAppointments()([]domain.AppointmentDTO, error){
	dentists, err := s.Repository.GetAllAppointments()
		if err != nil{
		return nil, err
	}

	return dentists, nil
}

func (s *Service) SaveAppointment(a *domain.Appointment)(int, error){
  id, err := s.Repository.SaveAppointment(a)
  if err != nil{
		return 0, err
	}
	return id, nil
}

func (s *Service) DeleteAppointment(id int)(int, error){
	id,err := s.Repository.DeleteById(id)
		if err != nil{
		return 0, err
	}
	return id, nil
}

func (s *Service) UpdateAppointment(a domain.Appointment)(int, error){
	id, err := s.Repository.UpdateAppointment(a)
	 if err != nil{
		return 0, err
	}
	return id, nil
}
